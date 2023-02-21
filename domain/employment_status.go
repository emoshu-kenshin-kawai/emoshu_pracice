package domain

import "time"

type EmploymentStatus struct {
	ID               uint      `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	EmploymentStatus string    `gorm:"not null" json:"employment_status"`
}
