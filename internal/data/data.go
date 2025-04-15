package data

import (
	"usermate/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserMateRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	rdb *redis.Client
	db  *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
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
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	data := &Data{
		db:  db,
		rdb: rdb,
	}
	return data, cleanup, nil
}
