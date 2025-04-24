package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"usermate/internal/conf"
	"usermate/pkg/utils"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserMateRepo)

type Data struct {
	rdb          *redis.Client
	kafkaConf    *conf.Kafka
	elasticConf  *conf.ElasticSearch
	db           *gorm.DB
	kafka        *kafka.Writer
	searchClient *utils.ESClient
	writer       *kafka.Writer
}

// NewData .
func NewData(c *conf.Data, kafkaconf *conf.Kafka, elasticConf *conf.ElasticSearch, logger log.Logger) (*Data, func(), error) {
	var ctx context.Context

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: c.Database.Driver,
		DSN:        c.Database.Source,
	}), &gorm.Config{})

	if err != nil {
		log.NewHelper(logger).Errorf("init database failed err  %v ", err)
		return nil, nil, err
	}
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaconf.Host),
		Topic:    kafkaconf.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaconf.Host},
		Topic:   kafkaconf.Topic,
		//GroupID:  kafkaconf.Group, // 消费者组ID
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	defer func() {
		err := writer.Close()
		if err != nil {
			log.NewHelper(logger).Errorf("writer close err %v", err)
		}
	}()

	defer func() {
		err := reader.Close()
		if err != nil {
			log.NewHelper(logger).Errorf("reader close err %v", err)
		}
	}()

	log.NewHelper(logger).Infof("init kafka client success")
	// 初始化elaticsearch
	searchClient, err := utils.NewOrderESClient(elasticConf, logger)
	if err != nil {
		log.NewHelper(logger).Errorf("elasticsearch 初始化失败:", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	//收发消息
	go func() {
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.NewHelper(logger).Errorf("消费失败:%v", err)
				break
			}
			log.NewHelper(logger).Infof("收到消息: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s\n",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			err = searchClient.UpsertOrderData(ctx, msg.Value, elasticConf)
			if err != nil {
				log.NewHelper(logger).Errorf("订单数据写入ES失败: %v", err)
			}
		}

	}()

	//初始化完成
	data := &Data{
		db:           db,
		rdb:          rdb,
		kafkaConf:    kafkaconf,
		writer:       writer,
		searchClient: searchClient,
	}
	return data, cleanup, nil
}
