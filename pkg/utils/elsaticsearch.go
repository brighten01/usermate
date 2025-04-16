package utils

import (
	"context"
	"encoding/json"
	"time"
	"usermate/internal/conf"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/olivere/elastic/v7"
)

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

// 写入订单数据
func UpsertOrderData(client *elastic.Client, orderInfo []byte, logger log.Logger, esconf *conf.ElasticSearch) error {
	var orderList *OrderList

	err := json.Unmarshal(orderInfo, &orderList)
	if err != nil {
		log.NewHelper(logger).Infof("订单数据解析失败: %v", err)
	}
	result, err := client.Update().
		Index(esconf.OrderIndex).
		Id("id").
		Doc(orderList).
		DocAsUpsert(true).
		Do(context.Background())
	if err != nil {
		log.NewHelper(logger).Infof("订单数据写入失败: %v", err)
		return err
	}

	log.NewHelper(logger).Infof("订单数据写入成功: %v", result.Id)
	return nil
}

type OrderRequest struct {
	UserMateId string
	OrderId    string
}

// todo 根据查找逻辑更新文档
func SearchOrderData(client *elastic.Client, esconf *conf.ElasticSearch, orderRequest *OrderRequest) ([]*OrderList, error) {
	var orders []*OrderList
	ctx := context.Background()
	query := elastic.NewBoolQuery()

	if orderRequest.UserMateId != "" {
		query = query.Must(elastic.NewTermsQuery("user_mate_id", orderRequest.UserMateId))
	}
	if orderRequest.OrderId != "" {
		query = query.Must(elastic.NewTermsQuery("order_id", orderRequest.OrderId))
	}
	result, err := client.Search().Index(esconf.OrderIndex).Query(query).Sort("id", true).From(0).Size(10).Pretty(true).Do(ctx)
	if err != nil {
		return orders, err
	}

	if result.Hits.TotalHits.Value > 0 {
		for _, hit := range result.Hits.Hits {
			var order OrderList
			err := json.Unmarshal(hit.Source, &order)
			if err != nil {
				//todo log
			}
			orders = append(orders, &order)
		}
	}
	return orders, nil
}
