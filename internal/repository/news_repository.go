package repository

import (
	"fmt"
	"gin-app/internal/models"

	"gorm.io/gorm"
)

type NewsRepository struct {
	GormDB *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{
		GormDB: db,
	}
}

func (r *NewsRepository) CreateNews(data models.News) error {
	err := r.GormDB.Create(&data).Error
	if err != nil {
		return fmt.Errorf("Gagal membuat news: %w", err)
	}
	return nil
}

func (r *NewsRepository) GetNews() ([]models.News, error) {
	var news []models.News
	err := r.GormDB.Order("created_at DESC").Find(&news).Error
	if err != nil {
		return nil, fmt.Errorf("Gagal mendapatkan news: %w", err)
	}
	return news, nil
}

func (r *NewsRepository) GetNewsByID(slug string) (models.News, error) {
	var news models.News
	err := r.GormDB.Where("slug = ?", slug).First(&news).Error
	if err != nil {
		return models.News{}, fmt.Errorf("Gagal mendapatkan news by ID: %w", err)
	}
	return news, nil
}

func (r *NewsRepository) UpdateNews(slug string, data models.News) error {
	var news models.News
	err := r.GormDB.Where("slug = ?", slug).First(&news).Error
	if err != nil {
		return fmt.Errorf("Gagal mendapatkan news by ID: %w", err)
	}
	err = r.GormDB.Model(&news).Updates(data).Error
	if err != nil {
		return fmt.Errorf("Gagal mengupdate news: %w", err)
	}
	return nil
}

func (r *NewsRepository) DeleteNews(slug string) error {
	err := r.GormDB.Where("slug = ?", slug).Delete(&models.News{}).Error
	if err != nil {
		return fmt.Errorf("Gagal menghapus news: %w", err)
	}

	return nil
}

func (r *NewsRepository) GetFotoNews(slug string) (string, error) {
	var news models.News
	err := r.GormDB.Select("thumbnail").Where("slug = ?", slug).First(&news).Error
	if err != nil {
		return "", err
	}
	return news.Thumbnail, nil
}
