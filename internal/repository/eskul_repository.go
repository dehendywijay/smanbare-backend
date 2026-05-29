package repository

import (
	"fmt"

	"gin-app/internal/models"

	"gorm.io/gorm"
)

type EskulRepository struct {
	GormDB *gorm.DB
}

func NewEskulRepository(db *gorm.DB) *EskulRepository {
	return &EskulRepository{
		GormDB: db,
	}
}

func (r *EskulRepository) CreateEskul(eskul models.Eskul)  error {
	err := r.GormDB.Create(&eskul).Error
	if err != nil {
		return fmt.Errorf("Gagal menyimpan data eskul: %v", err)
	}
	return nil
}

func (r *EskulRepository) GetEskul() (*[]models.Eskul, error) {
	var eskul *[]models.Eskul
	err := r.GormDB.Find(&eskul).Error
	if err != nil {
		return nil, fmt.Errorf("Gagal mengambil data eskul: %v", err)
	}
	return eskul, nil
}

func (r *EskulRepository) GetEskulByID(slug string) (models.Eskul, error) {
	var eskul models.Eskul
	err := r.GormDB.Where("slug = ?", slug).First(&eskul).Error	
	if err != nil {
		return eskul, fmt.Errorf("Gagal mengambil data eskul: %v", err)
	}
	return eskul, nil
}

func (r *EskulRepository) EditEskul(slug string, updatedEskul models.Eskul) error {
	var eskul models.Eskul
	err := r.GormDB.Where("slug = ?", slug).First(&eskul).Error
	if err != nil {
		return fmt.Errorf("Gagal mengambil data eskul: %v", err)
	}
	err = r.GormDB.Model(&eskul).Updates(updatedEskul).Error
	if err != nil {
		return  fmt.Errorf("Gagal memperbarui data eskul: %v", err)
	}
	return nil
}

func (r *EskulRepository) DeleteEskul(slug string) error {
	var eskul models.Eskul
	err := r.GormDB.Where("slug = ?", slug).First(&eskul).Error
	if err != nil {
		return fmt.Errorf("Gagal mengambil data eskul: %v", err)
	}
	err= r.GormDB.Delete(&eskul).Error
	if err != nil {
		return fmt.Errorf("Gagal menghapus data eskul: %v", err)
	}
	return nil
}

func (r *EskulRepository) GetFotoEskul(slug string) (string, error) {
	var eskul models.Eskul
	err := r.GormDB.Select("foto").Where("slug = ?", slug).First(&eskul).Error
	if err != nil {
		return "", fmt.Errorf("Gagal mengambil foto eskul: %v", err)
	}
	return eskul.Foto, nil
}