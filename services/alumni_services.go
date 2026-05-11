package services

import (
	"gin-app/config"
	"gin-app/models"
)

func GetAllAlumni() ([]models.Alumni, error) {
	var alumni []models.Alumni
	err := config.DB.Find(&alumni).Error
	return alumni, err
}

func CreateAlumni(alumni models.Alumni) (models.Alumni, error) {
	result := config.DB.Create(&alumni)
	return alumni, result.Error
}

func UpdateAlumni(id string, updatedAlumni models.Alumni) (models.Alumni, error)  {
	var alumni models.Alumni
	err := config.DB.First(&alumni, id).Error
	if err != nil {
		return alumni, err
	}

	err = config.DB.Model(&alumni).Updates(updatedAlumni).Error
	return alumni, err
}

func DeleteAlumni(id string) error {
	var alumni models.Alumni
	if err := config.DB.First(&alumni, id).Error; err != nil {
		return err
	}
	return config.DB.Delete(&alumni).Error
}

func GetFotoAlumni(id string) (string, error) {
	var alumni models.Alumni
	err := config.DB.Select("foto").First(&alumni, id).Error
	if err != nil {
		return "", err
	}
	return alumni.Foto, nil
}