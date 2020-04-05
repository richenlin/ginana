// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"ginana/internal/config"
	"ginana/internal/db"
	"ginana/internal/server/http"
	"ginana/internal/service"
	"ginana/internal/service/user"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := db.NewDB(configConfig)
	if err != nil {
		return nil, nil, err
	}
	iUser, err := user.New(gormDB, configConfig)
	if err != nil {
		return nil, nil, err
	}
	serviceService, err := service.New(gormDB, iUser)
	if err != nil {
		return nil, nil, err
	}
	syncedEnforcer, err := db.NewCasbin(iUser, configConfig)
	if err != nil {
		return nil, nil, err
	}
	engine, err := http.NewGin(serviceService, syncedEnforcer)
	if err != nil {
		return nil, nil, err
	}
	server, err := http.NewHttpServer(engine, configConfig)
	if err != nil {
		return nil, nil, err
	}
	app, cleanup, err := NewApp(serviceService, server)
	if err != nil {
		return nil, nil, err
	}
	return app, func() {
		cleanup()
	}, nil
}

// wire.go:

var initProvider = wire.NewSet(config.NewConfig, db.NewDB)

var serviceProvider = wire.NewSet(user.New)
