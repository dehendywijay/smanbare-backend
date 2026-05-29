package services

import (
	"errors"
	"gin-app/config"
	"gin-app/dto"
	"gin-app/models"
	"gin-app/utility" 
)

func LoginAdmin(username, password string) (dto.AdminDTO, error) {
	var admin models.Admin

	err := config.DB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return dto.AdminDTO{}, errors.New("Data tidak ada")
	}

	if !utility.CheckPassword(admin.Password, password) {
		return dto.AdminDTO{}, errors.New("invalid password")
	}

	accessToken, err := utility.GenerateAccessToken(admin.ID)
	if err != nil {
		return dto.AdminDTO{}, err
	}

	refreshToken, err := utility.GenerateRefreshToken(admin.ID)
	if err != nil {
		return dto.AdminDTO{}, err
	}

	return dto.AdminDTO{
		Username:     admin.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

