package paseto

import (
	"github.com/o1egl/paseto"
	"rpc-server/config"
)

// Paseto: JWT 보다 가벼움, 보다 안전한 암호화 알고리즘 제공
type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

// 새로운 토큰을 만드는 함수
func (m *PasetoMaker) CreateNewToken() (string, error) {
	return "", nil
}

// 토큰을 받아 검증하는 함수
func (m *PasetoMaker) VerifyToken(token string) error {
	return nil
}
