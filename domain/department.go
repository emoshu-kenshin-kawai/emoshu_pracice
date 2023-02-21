package domain

import "time"

type Department struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Department string    `gorm:"not null" json:"department"`
}
