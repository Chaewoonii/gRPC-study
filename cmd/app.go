package cmd

import (
	"rpc-server/config"
	"rpc-server/network"
	"rpc-server/repository"
	"rpc-server/service"
)

type App struct {
	cfg *config.Config //포인터 타입

	service    *service.Service
	repository *repository.Repository
	network    *network.Network
}

func NewApp(cfg *config.Config) {
	a := &App{cfg: cfg}

	var err error

	if a.repository, err = repository.NewRepository(cfg); err != nil {
		panic(err)
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
	} else if a.network, err = network.NewNetwork(cfg, a.service); err != nil {
		panic(err)
	} else {
		a.network.StartServer()
	}
}