package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"rpc-server/config"
	"rpc-server/gRPC/paseto"
	auth "rpc-server/gRPC/proto"
	"time"
)

type GRPCServer struct {
	auth.AuthServiceServer
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]*auth.AuthData
}

func NewGRPCServer(cfg *config.Config) error {
	fmt.Println("Start GRPC Server")
	fmt.Println(cfg.GRPC.URL) //config.toml에 GRPC URL 값을 넣어줌.
	if lis, err := net.Listen("tcp", cfg.GRPC.URL); err != nil {
		return err
	} else {

		server := grpc.NewServer([]grpc.ServerOption{}...)

		auth.RegisterAuthServiceServer(server, &GRPCServer{
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

// auth_grpc.pb.go의 AuthServiceServer 인터페이스를 사용하기 위한 내부 함수 구현
// CreateAuth(context.Context, *CreateTokenReq) (*CreateTokenRes, error)
// VerifyAuth(context.Context, *VerifyTokenReq) (*VerifyTokenRes, error)

func (s *GRPCServer) CreateAuth(_ context.Context, req *auth.CreateTokenReq) (*auth.CreateTokenRes, error) {
	data := req.Auth
	token := data.Token
	s.tokenVerifyMap[token] = data

	return &auth.CreateTokenRes{Auth: data}, nil
}

func (s *GRPCServer) VerifyAuth(_ context.Context, req *auth.VerifyTokenReq) (*auth.VerifyTokenRes, error) {
	token := req.Token
	res := &auth.VerifyTokenRes{V: &auth.Verify{
		Auth: nil,
	}}

	log.Println("token", token)

	if authData, ok := s.tokenVerifyMap[token]; !ok { //토큰 검증 실패(서버에서 관리되지 않는 토큰임)
		res.V.Status = auth.ResponseType_FAILED
		return res, errors.New("token not existed at server")
	} else if err := s.pasetoMaker.VerifyToken(token); err != nil {
		log.Println("err", err.Error())
		return nil, errors.New("failed to verify token")
	} else if authData.ExpireDate < time.Now().Unix() { // 토큰 만료 시 삭제
		delete(s.tokenVerifyMap, token)
		res.V.Status = auth.ResponseType_EXPIRED_DATE
		return res, errors.New("token Expired")
	} else {
		res.V.Status = auth.ResponseType_SUCCESS // 토큰 검증 성공
		return res, nil
	}
}

func (s *GRPCServer) apply(options *interface{}) {
	//TODO implement me
	panic("implement me")
}
