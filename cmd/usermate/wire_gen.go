// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"usermate/internal/biz"
	"usermate/internal/conf"
	"usermate/internal/data"
	"usermate/internal/server"
	"usermate/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	userMateRepo :=data.NewUserMateRepo(dataData,logger)
	userMateUsecase := biz.NewUserMateUsecase(userMateRepo,logger)
	userMateService:= service.NewUserMateService(userMateUsecase,logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService,userMateService ,logger)
	httpServer := server.NewHTTPServer(confServer, greeterService,userMateService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
