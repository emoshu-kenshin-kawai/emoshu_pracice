package domain

import "time"

type Role struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `gorm:"not null" json:"role"`
}
