package domain

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Department string `gorm:"not null"`
}
