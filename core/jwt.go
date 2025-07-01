package core

import (
	"blog-admin/global"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId uint   `json:"user_id"`
	Phone  string `json:"phone"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uint, Phone string) (string, error) {
	claims := &Claims{
		UserId: userId,
		Phone:  Phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.JWTSecret))
}

// 解析token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return global.JWTSecret, nil
	})
	if err != nil || !token.Valid {

		return nil, err
	}

	return claims, err
}
