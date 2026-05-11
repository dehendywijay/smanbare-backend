package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}