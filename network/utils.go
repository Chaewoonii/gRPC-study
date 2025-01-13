package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 미들웨어: API가 실행되기 전 특정 함수를 먼저 태우는(?) 방식

// auth 에 대해 검증 하고 가져오는 코드
func (n *Network) verifyLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Bearer token을 가져온다.
		t := getAuthToken(c)

		if t == "" { // 토큰 값이 없는 경우
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
		} else {
			if _, err := n.gRPCClient.VerifyAuth(t); err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort() // 토큰 검증, 실패한 경우
			} else { // 토큰 검증, 성공한 경우
				c.Next()
			}
		}
	}
}

// Bearer token 을 가져오는 함수
func getAuthToken(c *gin.Context) string {
	var token string

	// 토큰 헤더의 Authorization 값을 가져옴: Bearer ~~~ 형태
	authToken := c.Request.Header.Get("Authorization")

	// 토큰 앞의 "Bearer " 부분을 삭제, 실질적인 토큰 값만을 가져옴.
	authSided := strings.Split(authToken, " ")
	// 비어 있는 토큰에 값을 할당
	if len(authSided) > 1 {
		token = authSided[1]
	}
	return token
}
