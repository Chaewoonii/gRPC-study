package service

import (
	"rpc-server/config"
	"rpc-server/repository"
)

type Service struct {
	cfg *config.Config

	// service는 repository와 통신
	repository *repository.Repository
}

func NewService(cfg *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{cfg: cfg, repository: repository}

	return r, nil
}
