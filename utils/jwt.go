package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		[]byte(secretKey),
	}
}

// 创建一个token
func (j *JWT) CreateToken(data string, timeout int64) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:  "TimeToken",
		Subject: data,
	}
	if timeout > 0 {
		claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(timeout)).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(token string) (string, error) {
	var err error
	var claims jwt.StandardClaims
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	err = claims.Valid()
	if err != nil {
		return "", err
	} else {
		return claims.Subject, err
	}
}

func MakeTimedToken(data string, secretKey string, timeout int64) (string, error) {
	j := NewJWT(secretKey)
	return j.CreateToken(data, timeout)
}

func ParseTimedToken(secretKey, token string) (string, error) {
	j := NewJWT(secretKey)
	return j.ParseToken(token)
}
