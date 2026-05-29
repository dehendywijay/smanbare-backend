package repository

import (
	"fmt"
	"gin-app/internal/models"

	"gorm.io/gorm"
)

type AlumniRepository struct {
	GormDB *gorm.DB
}

func NewAlumniRepository(db *gorm.DB) *AlumniRepository {
	return &AlumniRepository{
		GormDB: db,
	}
}

func (r *AlumniRepository) GetAllAlumni() (*[]models.Alumni, error) {
	var alumni *[]models.Alumni
	err := r.GormDB.Find(&alumni).Error
	if err != nil {
		return nil, fmt.Errorf("Gagal mengambil data alumni: %v", err)
	}
	return alumni, nil
}

func (r *AlumniRepository) CreateAlumni(alumni models.Alumni) error {
	err := r.GormDB.Create(&alumni).Error
	if err != nil {
		return fmt.Errorf("Gagal menyimpan data alumni: %v", err)
	}
	return nil
}

func (r *AlumniRepository) UpdateAlumni(id string, updatedAlumni *models.Alumni) error {
	var alumni models.Alumni
	err := r.GormDB.First(&alumni, id).Error
	if err != nil {
		return fmt.Errorf("Gagal mengambil data alumni: %v", err)
	}

	err = r.GormDB.Model(&alumni).Updates(updatedAlumni).Error
	if err != nil {
		return fmt.Errorf("Gagal memperbarui data alumni: %v", err)
	}
	return nil
}

func (r *AlumniRepository) DeleteAlumni(id string) error {
	var alumni models.Alumni
	if err := r.GormDB.First(&alumni, id).Error; err != nil {
		return fmt.Errorf("Gagal mengambil data alumni: %v", err)
	}

	err := r.GormDB.Delete(&alumni).Error
	if err != nil {
		return fmt.Errorf("Gagal menghapus data alumni: %v", err)
	}
	return nil
}

func (r *AlumniRepository) GetFotoAlumni(id string) (string, error) {
	var alumni models.Alumni
	err := r.GormDB.Select("foto").First(&alumni, id).Error
	if err != nil {
		return "", fmt.Errorf("Gagal mengambil foto alumni: %v", err)
	}
	return alumni.Foto, nil
}
