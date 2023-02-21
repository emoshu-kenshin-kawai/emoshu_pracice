package domain

import "time"

type EmploymentStatus struct {
	ID               uint      `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	EmploymentStatus string    `gorm:"not null" json:"employment_status"`
}
