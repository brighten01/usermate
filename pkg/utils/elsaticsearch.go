package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"usermate/internal/common"
	"usermate/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/olivere/elastic/v7"
	"github.com/patrickmn/go-cache"
)

// OrderList 订单列表结构
type OrderList struct {
	ID                  int64     `json:"id"`
	OrderID             string    `json:"order_id"`
	UID                 int64     `json:"uid"`
	UserMateID          int64     `json:"user_mate_id"`
	ServiceCategory     string    `json:"service_category"`
	StartTime           time.Time `json:"start_time"`
	EndTime             time.Time `json:"end_time"`
	Status              int8      `json:"status"`
	Amount              float64   `json:"amount"`
	Discount            float64   `json:"discount"`
	Avatar              string    `json:"avatar"`
	LinkURL             string    `json:"link_url"`
	IsOrderAfter        int8      `json:"is_order_after"`
	Gender              int8      `json:"gender"`
	Level               int8      `json:"level"`
	Duration            int8      `json:"duration"`
	ServiceCategoryID   int       `json:"service_category_id"`
	ServiceCategoryName string    `json:"service_category_name"`
	Wechat              string    `json:"wechat"`
	Note                string    `json:"note"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Payment             int64     `json:"payment"`
	Nickname            string    `json:"nickname"`
}

// Validate 验证订单数据
func (o *OrderList) Validate() error {
	if o.OrderID == "" {
		return errors.New("order_id is required")
	}
	if o.UID == 0 {
		return errors.New("uid is required")
	}
	if o.UserMateID == 0 {
		return errors.New("user_mate_id is required")
	}
	if o.Amount < 0 {
		return errors.New("amount cannot be negative")
	}
	return nil
}

// ToJSON 将订单数据转换为JSON
func (o *OrderList) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}

// FromJSON 从JSON解析订单数据
func (o *OrderList) FromJSON(data []byte) error {
	return json.Unmarshal(data, o)
}

// ESConfig Elasticsearch配置
type ESConfig struct {
	Address  string
	Username string
	Password string
}

// ESClient Elasticsearch客户端
type ESClient struct {
	client *elastic.Client
	log    *log.Helper
	cache  *cache.Cache
}

// NewESClient 创建新的ES客户端
func NewOrderESClient(esconf *conf.ElasticSearch, logger log.Logger) (*ESClient, error) {
	cfg := &ESConfig{
		Address:  esconf.OrdersEs.Host,
		Username: esconf.OrdersEs.Username,
		Password: esconf.OrdersEs.Password,
	}

	if cfg == nil || cfg.Address == "" {
		return nil, errors.New("es 初始化失败")
	}

	options := []elastic.ClientOptionFunc{
		elastic.SetURL(cfg.Address),
		elastic.SetSniff(false),
	}

	if cfg.Username != "" && cfg.Password != "" {
		options = append(options, elastic.SetBasicAuth(cfg.Username, cfg.Password))
	}

	client, err := elastic.NewClient(options...)
	if err != nil {
		return nil, common.ErrClientCreation
	}

	return &ESClient{
		client: client,
		log:    log.NewHelper(logger),
		cache:  cache.New(common.DefaultCacheExpiration, common.DefaultCacheCleanupInterval),
	}, nil
}

// WriteOrderData 写入订单数据
func (es *ESClient) WriteOrderData(ctx context.Context, orderInfo []byte) error {
	var orderList *OrderList
	if err := json.Unmarshal(orderInfo, &orderList); err != nil {
		return fmt.Errorf("%w: %v", common.ErrDataUnmarshal, err)
	}

	result, err := es.client.Index().
		Index(common.OrderIndex).
		BodyJson(orderList).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("%w: %v", common.ErrDataWrite, err)
	}

	es.log.WithContext(ctx).Infof("订单数据写入成功: %v", result.Id)
	return nil
}

// SearchOptions 搜索选项
type SearchOptions struct {
	Query      string
	Page       int
	Size       int
	StartTime  time.Time
	EndTime    time.Time
	Status     int8
	MinAmount  float64
	MaxAmount  float64
	UserMateID int64
	OrderBy    string
	OrderDesc  bool
}

// BatchWriteResult 批量写入结果
type BatchWriteResult struct {
	Success int
	Failed  int
	Errors  []error
}

// SearchOrderDataWithOptions 使用选项搜索订单数据

// BatchWriteOrderData 批量写入订单数据
func (es *ESClient) BatchWriteOrderData(ctx context.Context, orders []*OrderList) (*BatchWriteResult, error) {
	if len(orders) == 0 {
		return &BatchWriteResult{}, nil
	}

	bulk := es.client.Bulk()
	for _, order := range orders {
		if err := order.Validate(); err != nil {
			return &BatchWriteResult{
				Failed: 1,
				Errors: []error{fmt.Errorf("订单验证失败: %w", err)},
			}, nil
		}

		req := elastic.NewBulkIndexRequest().
			Index(common.OrderIndex).
			Doc(order)
		bulk.Add(req)
	}

	result, err := bulk.Do(ctx)
	if err != nil {
		return &BatchWriteResult{
			Failed: len(orders),
			Errors: []error{err},
		}, nil
	}

	var failed int
	var errors []error
	for _, item := range result.Items {
		for _, v := range item {
			if v.Error != nil {
				failed++
				errors = append(errors, fmt.Errorf("批量写入失败: %v", v.Error))
			}
		}
	}

	return &BatchWriteResult{
		Success: len(orders) - failed,
		Failed:  failed,
		Errors:  errors,
	}, nil
}

// Close 关闭ES客户端
func (es *ESClient) Close() {
	if es.client != nil {
		es.client.Stop()
	}
}

type OrderRequest struct {
	UserMateId string
	OrderId    string
}

// UpsertOrderData 更新或插入订单数据
func (es *ESClient) UpsertOrderData(ctx context.Context, data []byte, esconf *conf.ElasticSearch) error {
	var order OrderList
	if err := json.Unmarshal(data, &order); err != nil {
		es.log.WithContext(ctx).Infof("%w: %v", common.ErrDataWrite, err)
		return fmt.Errorf("%w: %v", common.ErrDataUnmarshal, err)
	}

	_, err := es.client.Index().
		Index(esconf.OrdersEs.Index).
		Id(order.OrderID).
		BodyJson(order).
		Do(context.Background())
	if err != nil {
		es.log.WithContext(ctx).Infof("%w: %v", common.ErrDataWrite, err)
		return fmt.Errorf("%w: %v", common.ErrDataWrite, err)
	}

	return nil
}
