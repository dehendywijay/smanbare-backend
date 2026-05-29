package services

import (
	"fmt"
	"gin-app/internal/dto"
	"gin-app/internal/models"
	"gin-app/internal/repository"
	"gin-app/pkg/hash"
	jwttoken "gin-app/pkg/jwt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func getRefreshSecret() []byte {
	secret := os.Getenv("REFRESH_SECRET_KEY")
	if secret == "" {
		return []byte("REFRESH_SECRET_KEY")
	}
	return []byte(secret)
}

type AuthService struct {
	r *repository.AuthRepository
}

func NewAuthService(r *repository.AuthRepository) *AuthService {
	return &AuthService{
		r: r,
	}
}

func (s *AuthService) LoginAdmin(username, password string) (dto.AdminDTO, error) {
	var admin models.Admin

	admin, err := s.r.LoginAdmin(username, password)
	if err != nil {
		return dto.AdminDTO{}, fmt.Errorf("gagal login: %s", err)
	}

	if !hash.CheckPassword(admin.Password, password) {
		return dto.AdminDTO{}, fmt.Errorf("gagal login: password salah")
	}

	accessToken, err := jwttoken.GenerateAccessToken(admin.ID)
	if err != nil {
		return dto.AdminDTO{}, err
	}

	refreshToken, err := jwttoken.GenerateRefreshToken(admin.ID)
	if err != nil {
		return dto.AdminDTO{}, err
	}

	return dto.AdminDTO{
		Username:     admin.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return getRefreshSecret(), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("refresh token tidak valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != "refresh" {
		return "", fmt.Errorf("refresh token tidak valid")
	}

	userIDFloat := claims["user_id"].(float64)
	userID := uint(userIDFloat)

	newAccessToken, err := jwttoken.GenerateAccessToken(userID)
	if err != nil {
		return "", fmt.Errorf("gagal membuat access token baru")
	}

	return newAccessToken, nil
}
