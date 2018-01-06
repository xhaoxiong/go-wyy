package models

import "github.com/iqysf/gorm"

type AdminUser struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
