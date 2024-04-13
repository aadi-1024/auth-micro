package jwtUtil

import (
	"github.com/aadi-1024/auth-micro/pkg/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JwtConfig struct {
	jwtSecret string
	expiry    time.Duration
}

func NewJwtConfig(exp time.Duration) *JwtConfig {
	return &JwtConfig{
		jwtSecret: os.Getenv("JWT_SECRET"),
		expiry:    exp,
	}
}

func (j *JwtConfig) GenerateToken(uid int) (string, error) {
	clms := models.Claims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiry)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clms)
	signedString, err := token.SignedString(j.jwtSecret)
	if err != nil {
		return "", nil
	}
	return signedString, nil
}
