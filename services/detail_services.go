package services

import (
	"gin-app/config"
	"gin-app/models"
)

func CreateKepalaSekolah(data models.KepalaSekolah) (models.KepalaSekolah, error) {
	err := config.DB.Create(&data).Error
	return data, err
}

func EditKepalaSekolah(id string, data models.KepalaSekolah) (models.KepalaSekolah, error) {
	var kepalaSekolah models.KepalaSekolah
	err := config.DB.Where("id = ?", id).First(&kepalaSekolah).Error
	if err != nil {
		return kepalaSekolah, err
	}
	err = config.DB.Model(&kepalaSekolah).Updates(data).Error
	return kepalaSekolah, err
}