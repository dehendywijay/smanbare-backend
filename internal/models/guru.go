package models

import "gorm.io/gorm"

type Guru struct {
	gorm.Model
	Nama    string `json:"nama" gorm:"not null"`
	Jabatan string `json:"jabatan" gorm:"not null"`
	Foto    string `json:"foto" gorm:"not null"`
}


type KepalaSekolah struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Foto    string `json:"foto" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
}
