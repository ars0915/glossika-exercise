package usecase

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/ars0915/glossika-exercise/config"
	"github.com/ars0915/glossika-exercise/constant"
)

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

type Claims struct {
	UserID uint `json:"user_id"`
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
