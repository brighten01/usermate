//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"usermate/internal/biz"
	"usermate/internal/conf"
	"usermate/internal/data"
	"usermate/internal/data/elasticsearch"
	"usermate/internal/server"
	"usermate/internal/service"
)

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, kafka *conf.Kafka, elastic *conf.ElasticSearch, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, elasticsearch.ProviderSet, newApp))
}
