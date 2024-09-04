package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/ars0915/glossika-exercise/config"
	"github.com/ars0915/glossika-exercise/constant"
)

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func GenerateJWT(userID uint) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(config.Conf.JWT.ExpireDuration).Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    constant.ServiceName, // 设置发行人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.Conf.JWT.Secret)
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return config.Conf.JWT.Secret, nil
	})
	return token, claims, err
}
