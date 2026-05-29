package repository

import (
	"fmt"
	"gin-app/internal/models"

	"gorm.io/gorm"
)

type GuruRepository struct {
	GormDB *gorm.DB
}

func NewGuruRepository(db *gorm.DB) *GuruRepository {
	return &GuruRepository{
		GormDB: db,
	}
}

func (r *GuruRepository) CreateGuru(guru models.Guru) error {
	err := r.GormDB.Create(&guru).Error
	if err != nil {
		return fmt.Errorf("Gagal menyimpan data guru: %v", err)
	}
	return nil
}

func (r *GuruRepository) GetGuru() ([]models.Guru, error) {
	var guru []models.Guru
	err := r.GormDB.Order("nama asc").Find(&guru).Error
	if err!= nil {
		return nil, fmt.Errorf("Gagal mengambil data guru: %v", err)
	}
	return guru, nil
}

func (r *GuruRepository) EditGuru(id string, updatedGuru models.Guru) error {
	var guru models.Guru
	err := r.GormDB.First(&guru, id).Error
	if err != nil {
		return fmt.Errorf("gagal menemukan data guru: %v", err)
	}
	err = r.GormDB.Model(&guru).Updates(updatedGuru).Error
	if err != nil {
		return fmt.Errorf("gagal mengupdate data guru: %v", err)
	}
	return nil
}

func (r *GuruRepository) DeleteGuru(id string) error {
	var guru models.Guru
	err := r.GormDB.First(&guru, id).Error
	if err != nil {
		return fmt.Errorf("gagal menemukan data guru: %v", err)
	}

	err = r.GormDB.Delete(&guru).Error
	if err != nil {
		return fmt.Errorf("gagal menghapus data guru: %v", err)
	}
	return nil
}

func (r *GuruRepository) GetFotoGuru(id string) (string, error) {
	var guru models.Guru
	err := r.GormDB.Select("foto").First(&guru, id).Error
	if err != nil {
		return "", fmt.Errorf("gagal mengambil foto guru: %v", err)
	}
	return guru.Foto, nil
}

func (r *GuruRepository) EditKepala(id string, updatedKepalaSekolah models.KepalaSekolah) error {
	var kepalaSekolah models.KepalaSekolah
	err := r.GormDB.First(&kepalaSekolah, id).Error
	if err != nil {
		return fmt.Errorf("gagal menemukan data kepala sekolah: %v", err)
	}
	err = r.GormDB.Model(&kepalaSekolah).Updates(updatedKepalaSekolah).Error
	if err != nil {
		return fmt.Errorf("gagal mengupdate data kepala sekolah: %v", err)
	}
	return nil
}

func (r *GuruRepository) GetFotoKepala(id string) (string, error) {
	var kepalaSekolah models.KepalaSekolah
	err := r.GormDB.Select("foto").First(&kepalaSekolah, id).Error
	if err != nil {
		return "", fmt.Errorf("gagal mendapatkan data kepala sekolah: %v", err)
	}
	return kepalaSekolah.Foto, nil
}

// func (r *GuruRepository) CreateKepala(kepalaSekolah models.KepalaSekolah) (models.KepalaSekolah, error) {
// 	result := r.GormDB.Create(&kepalaSekolah)
// 	return kepalaSekolah, result.Error
// }

func (r *GuruRepository) GetKepalaByID(id string) (models.KepalaSekolah, error) {
	var kepalaSekolah models.KepalaSekolah
	err := r.GormDB.First(&kepalaSekolah, id).Error
	if err != nil {
		return kepalaSekolah, fmt.Errorf("gagal menemukan data kepala sekolah: %v", err)
	}
	return kepalaSekolah, nil
}
