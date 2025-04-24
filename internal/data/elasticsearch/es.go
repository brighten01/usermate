package elasticsearch

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"usermate/internal/conf"
	"usermate/internal/data/elasticsearch/mates"
	"usermate/internal/data/elasticsearch/orders"
)

type ES struct {
	OrderClient    *orders.Client
	UserMateClient *mates.Client
}

var ProviderSet = wire.NewSet(NewES)

func NewES(conf *conf.ElasticSearch, log log.Logger) (*ES, error) {
	order_client, err := orders.NewESClient(log, conf)
	usermate_client, err := mates.NewESClient(log, conf)
	if err != nil {
		return nil, err
	}
	return &ES{
		OrderClient:    order_client,
		UserMateClient: usermate_client,
	}, err
}
