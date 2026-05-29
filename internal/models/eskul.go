package models

import "gorm.io/gorm"

type Eskul struct {
	gorm.Model
	Nama        string `json:"nama" gorm:"not null"`
	Foto        string `json:"foto" gorm:"not null"`
	Pembina     string `json:"pembina" gorm:"not null"`
	Jadwal      string `json:"jadwal" gorm:"not null"`
	Prestasi	 string `json:"prestasi" gorm:"not null"`
	Tujuan		 string `json:"tujuan" gorm:"not null"`
	Slug		string `json:"slug" gorm:"uniqueIndex;not null"`
}