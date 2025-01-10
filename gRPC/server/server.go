package server

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"rpc-server/config"
	"rpc-server/gRPC/paseto"
	auth "rpc-server/gRPC/proto"
)

type GRPCServer struct {
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]auth.AuthData
}

func NewGRPCServer(cfg *config.Config) error {
	if lis, err := net.listen("tcp", cfg.GRPC.URL); err != nil {
		return err
	} else {

		server := grpc.NewServer([]grpc.ServerOption{}...)

		auth.RegisterAuthServiceServer(grpc.NewServer(server, &GRPCServer{
			pasetoMaker:    paseto.NewPasetoMaker(cfg),
			tokenVerifyMap: make(map[string]*auth.AuthData),
		})

		// gRPC를 서버로 등록한다
		reflection.Register(server)

		//if err = server.Serve(lis); err != nil {
		//	panic(err)
		//}
		//fmt.Println("여기를 오지 않습니다.")
		// 따라서, 새로운 쓰레드를 형성해서 서버를 연다.
		go func() {
			log.Println("Start GRPC Server")
			if err = server.Serve(lis); err != nil {
				panic(err)
			}
		}()
	}

	return nil
}

