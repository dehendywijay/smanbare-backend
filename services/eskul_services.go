package services

import (
	"gin-app/config"
	"gin-app/models"
)

func CreateEskul(eskul models.Eskul) (models.Eskul, error) {
	result := config.DB.Create(&eskul)
	return eskul, result.Error
}

func GetEskul() ([]models.Eskul, error) {
	var eskul []models.Eskul
	result := config.DB.Find(&eskul)
	return eskul, result.Error
}

func GetEskulByID(slug string) (models.Eskul, error) {
	var eskul models.Eskul
	result := config.DB.Where("slug = ?", slug).First(&eskul)
	return eskul, result.Error
}

func EditEskul(slug string, updatedEskul models.Eskul) (models.Eskul, error) {
	var eskul models.Eskul
	err := config.DB.Where("slug = ?", slug).First(&eskul).Error
	if err != nil {
		return eskul, err
	}
	err = config.DB.Model(&eskul).Updates(updatedEskul).Error
	return eskul, err
}

func DeleteEskul(slug string) error {
	var eskul models.Eskul
	err := config.DB.Where("slug = ?", slug).First(&eskul).Error
	if err != nil {
		return err
	}
	return config.DB.Delete(&eskul).Error
}

func GetFotoEskul(slug string) (string, error) {
	var eskul models.Eskul
	err := config.DB.Select("foto").Where("slug = ?", slug).First(&eskul).Error
	if err != nil {
		return "", err
	}
	return eskul.Foto, nil
}