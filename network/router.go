package network

import (
	"github.com/gin-gonic/gin"
	"rpc-server/config"
	"rpc-server/service"
)

type Network struct {
	cfg *config.Config

	// 네트워크는 전송받은 요청에 대해 서비스에게 보내기 때문에 서비스 객체를 가지고 있어야 함
	service *service.Service

	engin *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	r := &Network{cfg: cfg, service: service, engin: gin.New()}

	return r, nil
}

func (n *Network) StartServer() {
	n.engin.Run(":8080")
}
