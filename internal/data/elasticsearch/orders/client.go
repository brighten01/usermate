package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/olivere/elastic/v7"
	"usermate/internal/common"
	"usermate/internal/conf"
	"usermate/pkg/utils"
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
		Address:  esConf.OrdersEs.Host,
		Username: esConf.OrdersEs.Username,
		Password: esConf.OrdersEs.Password,
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

// 订单搜索
func (es *Client) SearchOrderDataWithOptions(ctx context.Context, opts utils.SearchOptions) ([]utils.OrderList, error) {
	query := elastic.NewBoolQuery()

	if opts.Query != "" {
		query.Must(elastic.NewQueryStringQuery(opts.Query))
	}

	if !opts.StartTime.IsZero() {
		query.Filter(elastic.NewRangeQuery("start_time").Gte(opts.StartTime))
	}

	if !opts.EndTime.IsZero() {
		query.Filter(elastic.NewRangeQuery("end_time").Lte(opts.EndTime))
	}

	if opts.Status != 0 {
		query.Filter(elastic.NewTermQuery("status", opts.Status))
	}

	if opts.MinAmount > 0 {
		query.Filter(elastic.NewRangeQuery("amount").Gte(opts.MinAmount))
	}

	if opts.MaxAmount > 0 {
		query.Filter(elastic.NewRangeQuery("amount").Lte(opts.MaxAmount))
	}

	if opts.UserMateID != 0 {
		query.Filter(elastic.NewTermQuery("user_mate_id", opts.UserMateID))
	}

	search := es.client.Search().
		Index(common.OrderIndex).
		Query(query)

	if opts.OrderBy != "" {
		search.Sort(opts.OrderBy, opts.OrderDesc)
	}

	if opts.Page < common.DefaultPage {
		opts.Page = common.DefaultPage
	}
	if opts.Size < common.DefaultPageSize {
		opts.Size = common.DefaultPageSize
	}

	searchResult, err := search.
		From((opts.Page - 1) * opts.Size).
		Size(opts.Size).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", common.ErrDataSearch, err)
	}

	var orders []utils.OrderList
	for _, hit := range searchResult.Hits.Hits {
		var order utils.OrderList
		if err := json.Unmarshal(hit.Source, &order); err != nil {
			es.log.WithContext(ctx).Errorf("解析订单数据失败: %v", err)
			continue
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// 搜索订单处理数据
func (es *Client) SearchOrderData(ctx context.Context, query string, page, size int) ([]utils.OrderList, error) {
	opts := utils.SearchOptions{
		Query: query,
		Page:  page,
		Size:  size,
	}
	orders, err := es.SearchOrderDataWithOptions(ctx, opts)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
