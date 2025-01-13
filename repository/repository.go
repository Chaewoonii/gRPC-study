package repository

import (
	"rpc-server/config"
	"rpc-server/gRPC/client"
	auth "rpc-server/gRPC/proto"
)

type Repository struct {
	cfg *config.Config

	gRPCClient *client.GRPCClient // Auth Data 생성요청 처리를 위해 gRPCCLient 받기: app.go 수정
}

func NewRepository(cfg *config.Config, gRPCClient *client.GRPCClient) (*Repository, error) {
	r := &Repository{cfg: cfg, gRPCClient: gRPCClient} // Auth Data 생성요청 처리를 위해 gRPCCLient 받기

	return r, nil
}

func (r Repository) CreateAuth(name string) (*auth.AuthData, error) {
	return r.gRPCClient.CreateAuth(name)
}
