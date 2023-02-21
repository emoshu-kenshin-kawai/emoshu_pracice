package domain

import "time"

type MemberRole struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	MemberID  uint      `json:"member_id"`
	Member    Member    `json:"member"`
	RoleID    uint      `json:"role_id"`
	Role      Role      `json:"role"`
}
