package domain

import "time"

type Department struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Department string    `gorm:"not null" json:"department"`
}
