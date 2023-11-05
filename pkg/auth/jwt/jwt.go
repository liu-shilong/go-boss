package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("golang-jwt")

type JwtClaims struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// 生成JWT
func GenerateToken(username string, password string) (string, error) {
	claims := JwtClaims{
		username,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 定义过期时间
			Issuer:    "go-boss",                                          // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// 解析token
func ParseToken(tokenString string) (*JwtClaims, error) {
	var mc = new(JwtClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
