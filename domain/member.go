package domain

import (
	"time"
)

type Member struct {
	ID                 uint             `gorm:"primarykey" json:"id"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
	No                 string           `json:"no"`
	ProfileImg         string           `gorm:"not null" json:"profile_img"`
	FullName           string           `gorm:"not null" json:"full_name"`
	KanaName           string           `json:"kana_name"`
	Motto              string           `json:"motto"`
	Biography          string           `json:"biography"`
	StartDate          time.Time        `json:"start_date"`
	EndDate            time.Time        `json:"end_date"`
	EmploymentStatusID uint             `json:"employment_status_id"`
	EmploymentStatus   EmploymentStatus `json:"employment_status"`
	StatusID           uint             `json:"status_id"`
	Status             Status           `json:"status"`
}
