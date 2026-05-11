package services

import (
	"gin-app/config"
	"gin-app/models"
)

func CreateGuru(guru models.Guru) (models.Guru, error) {
	result := config.DB.Create(&guru)
	return guru, result.Error
}

func GetGuru() ([]models.Guru, error) {
	var guru []models.Guru
	result := config.DB.Order("nama asc").Find(&guru)
	return guru, result.Error
}

func EditGuru(id string, updatedGuru models.Guru) (models.Guru, error) {
	var guru models.Guru
	err := config.DB.First(&guru, id).Error
	if err != nil {
		return guru, err
	}
	err = config.DB.Model(&guru).Updates(updatedGuru).Error
	return guru, err
}

func DeleteGuru(id string) error {
	var guru models.Guru
	err := config.DB.First(&guru, id).Error
	if err != nil {
		return err
	}
	return config.DB.Delete(&guru).Error
}

func GetFotoGuru(id string) (string, error) {
	var guru models.Guru
	err := config.DB.Select("foto").First(&guru, id).Error
	if err != nil {
		return "", err
	}
	return guru.Foto, nil
}


func EditKepala(id string, updatedKepalaSekolah models.KepalaSekolah) (models.KepalaSekolah, error) {
	var kepalaSekolah models.KepalaSekolah
	err := config.DB.First(&kepalaSekolah, id).Error
	if err != nil {
		return kepalaSekolah, err
	}
	err = config.DB.Model(&kepalaSekolah).Updates(updatedKepalaSekolah).Error
	return kepalaSekolah, err
}

func GetFotoKepala(id string) (string, error) {
	var kepalaSekolah models.KepalaSekolah
	err := config.DB.Select("foto").First(&kepalaSekolah, id).Error
	if err != nil {
		return "", err
	}
	return kepalaSekolah.Foto, nil
}

func CreateKepala(kepalaSekolah models.KepalaSekolah) (models.KepalaSekolah, error) {
	result := config.DB.Create(&kepalaSekolah)
	return kepalaSekolah, result.Error
}

func GetKepalaByID(id string) (models.KepalaSekolah, error) {
	var kepalaSekolah models.KepalaSekolah
	err := config.DB.First(&kepalaSekolah, id).Error
	return kepalaSekolah, err
}