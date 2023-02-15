package domain

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	No                 string
	ProfileImg         string `gorm:"not null"`
	FullName           string `gorm:"not null"`
	KanaName           string
	Motto              string
	Biography          string
	StartDate          time.Time
	EndDate            time.Time
	EmploymentStatusID uint
	EmploymentStatus   EmploymentStatus
	StatusID           uint
	Status             Status
}
