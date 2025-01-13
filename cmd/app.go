package cmd

import (
	"rpc-server/config"
	"rpc-server/gRPC/client"
	"rpc-server/network"
	"rpc-server/repository"
	"rpc-server/service"
)

type App struct {
	cfg *config.Config //포인터 타입

	gRPCClient *client.GRPCClient // grpc 클라이언트
	service    *service.Service
	repository *repository.Repository
	network    *network.Network
}

func NewApp(cfg *config.Config) {
	a := &App{cfg: cfg}

	var err error

	client.NewGRPCClient(cfg) // grpc 클라이언트

	if a.gRPCClient, err = client.NewGRPCClient(cfg); err != nil { // grpc 연결 확인
		panic(err)
	} else if a.repository, err = repository.NewRepository(cfg, a.gRPCClient); err != nil {
		panic(err)
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
	} else if a.network, err = network.NewNetwork(cfg, a.service, a.gRPCClient); err != nil { // 네트워크가 gRPC 클라이언트도 가지고 있어야 함
		panic(err)
	} else {
		a.network.StartServer()
	}
}
