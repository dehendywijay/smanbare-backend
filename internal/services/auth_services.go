package services

import (
	"fmt"
	"gin-app/internal/dto"
	"gin-app/internal/repository"
	"gin-app/pkg/hash"
	jwttoken "gin-app/pkg/jwt"

)

type AuthService struct {
	r *repository.AuthRepository
	RepoToken *repository.TokenRepository
	Token *jwttoken.JWTService

}

func NewAuthService(r *repository.AuthRepository, token *jwttoken.JWTService, repoToken *repository.TokenRepository) *AuthService {
	return &AuthService{
		r: r,
		Token: token,
		RepoToken: repoToken,
	}
}

func (s *AuthService) LoginAdmin(username, password string) (dto.AdminDTO, error) {

	admin, err := s.r.LoginAdmin(username, password)
	if err != nil {
		return dto.AdminDTO{}, fmt.Errorf("gagal login: %s", err)
	}

	if !hash.CheckPassword(admin.Password, password) {
		return dto.AdminDTO{}, fmt.Errorf("gagal login: password salah")
	}

	accessToken, err := s.Token.GenerateAccessToken(1)
	if err != nil {
		return dto.AdminDTO{}, err
	}

	refreshToken, err := s.Token.GenerateRefreshToken()
	if err != nil {
		return dto.AdminDTO{}, err
	}

	err = s.RepoToken.SetToken(refreshToken)
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
	_, err := s.RepoToken.GetToken(refreshToken)
	if err != nil {
		return "", err
	}

	newAccessToken, err := s.Token.GenerateAccessToken(1)
	if err != nil {
		return "", fmt.Errorf("gagal membuat access token baru")
	}

	return newAccessToken, nil
}
