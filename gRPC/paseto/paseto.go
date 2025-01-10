package paseto

import (
	"crypto/rand"
	"github.com/o1egl/paseto"
	"rpc-server/config"
	auth "rpc-server/gRPC/proto"
)

// Paseto: JWT 보다 가벼움, 보다 안전한 암호화 알고리즘 제공
type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg *config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

// 새로운 토큰을 만드는 함수
func (m *PasetoMaker) CreateNewToken(auth *auth.AuthData) (string, error) {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	return m.Pt.Encrypt(m.Key, auth, randomBytes)
}

// 토큰을 받아 검증하는 함수
func (m *PasetoMaker) VerifyToken(token string) error {
	var auth *auth.AuthData
	return m.Pt.Decrypt(token, m.Key, auth, nil)
}
