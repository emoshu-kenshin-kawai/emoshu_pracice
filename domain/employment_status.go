package domain

import "gorm.io/gorm"

type EmploymentStatus struct {
	gorm.Model
	EmploymentStatus string `gorm:"not null"`
}
