package models

import "gorm.io/gorm"

type Alumni struct {
	gorm.Model
	Nama        string `json:"nama" gorm:"not null"`
	Foto        string `json:"foto" gorm:"not null"`
	Universitas     string `json:"universitas" gorm:"not null"`
	Tahun      string `json:"tahun" gorm:"not null"`
}