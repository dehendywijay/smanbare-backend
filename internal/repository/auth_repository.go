package repository

import (
	"fmt"
	"gin-app/internal/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	GormDB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		GormDB: db,
	}
}

func (r *AuthRepository) LoginAdmin(username, password string) (models.Admin, error) {
	var admin models.Admin

	err := r.GormDB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return models.Admin{}, fmt.Errorf("Data tidak ada: %s", err)
	}

	return admin, nil
}

