package mates

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/olivere/elastic/v7"
	"usermate/internal/common"
	"usermate/internal/conf"
)

type Client struct {
	log    *log.Helper
	client *elastic.Client
	index  string
}

type ESConfig struct {
	Address  string
	Username string
	Password string
}

func NewESClient(logger log.Logger, esConf *conf.ElasticSearch) (*Client, error) {
	cfg := &ESConfig{
		Address:  esConf.Host,
		Username: esConf.Username,
		Password: esConf.Password,
	}
	if cfg == nil || cfg.Address == "" {
		return nil, common.ErrInvalidConfig
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
		return nil, err
	}

	return &Client{
		client: client,
		log:    log.NewHelper(logger),
	}, nil
}

func (c *Client) SearchMates(ctx context.Context, req SearchMateDTO) (*elastic.SearchResult, error) {

	return &elastic.SearchResult{}, nil
}

func (c *Client) AddUserMates(ctx context.Context, dto AddMateDTO) error {

	return nil
}
