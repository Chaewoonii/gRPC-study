package paseto

import (
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
func (m *PasetoMaker) CreateNewToken(auth auth.AuthData) (string, error) {
	// CreateNewToken에는 footer 값이 들어가는데, VerifyToken에서는 footer 값이 없음: 검증 에러 -> footer 값 삭제
	//randomBytes := make([]byte, 16)
	//rand.Read(randomBytes)
	//return m.Pt.Encrypt(m.Key, auth, randomBytes)
	return m.Pt.Encrypt(m.Key, auth, nil)
}

// 토큰을 받아 검증하는 함수
func (m *PasetoMaker) VerifyToken(token string) error {
	var auth auth.AuthData
	return m.Pt.Decrypt(token, m.Key, &auth, nil)
}
