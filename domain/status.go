package domain

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Status string `gorm:"not null"`
}
