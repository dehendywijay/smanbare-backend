package jwttoken

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getAccessSecret() []byte {
	secret := os.Getenv("ACCESS_SECRET_KEY")
	if secret == "" {
		return []byte("ACCESS_SECRET_KEY")
	}
	return []byte(secret)
}

func getRefreshSecret() []byte {
	secret := os.Getenv("REFRESH_SECRET_KEY")
	if secret == "" {
		return []byte("REFRESH_SECRET_KEY")
	}
	return []byte(secret)
}
func GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    "access",
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getAccessSecret())
}

func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    "refresh",
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getRefreshSecret())
}
