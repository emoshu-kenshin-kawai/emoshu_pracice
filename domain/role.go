package domain

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Role string `gorm:"not null"`
}
