package services

import (
	"gin-app/config"
	"gin-app/models"

)


func CreateNews(data models.News) (models.News, error) {
	err := config.DB.Create(&data).Error
	return data, err
}

func GetNews() ([]models.News, error) {
	var news []models.News
	err := config.DB.Order("created_at DESC").Find(&news).Error
	return news, err
}

func GetNewsByID(slug string) (models.News, error) {
	var news models.News

	err := config.DB.Where("slug = ?", slug).First(&news).Error
	
	return news, err
}

func UpdateNews(slug string, data models.News) (models.News, error) {
	var news models.News
	err := config.DB.Where("slug = ?", slug).First(&news).Error
	if err != nil {
		return news, err
	}
	err = config.DB.Model(&news).Updates(data).Error
	return news, err
}

func DeleteNews(slug string) (string, error) {
	var news models.News

	err := config.DB.Where("slug = ?", slug).First(&news).Error
	if err != nil {
		return "", err
	}

	thumbnail := news.Thumbnail

	err = config.DB.Delete(&news).Error
	if err != nil {
		return "", err
	}

	return thumbnail, nil
}

func GetFotoNews(slug string) (string, error) {
	var news models.News
	err := config.DB.Select("thumbnail").Where("slug = ?", slug).First(&news).Error
	if err != nil {
		return "", err
	}
	return news.Thumbnail, nil
}