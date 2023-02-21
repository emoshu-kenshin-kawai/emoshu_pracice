package domain

import "time"

type MemberDepartment struct {
	ID           uint       `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime:milli" json:"updated_at"`
	MemberID     uint       `json:"member_id"`
	Member       Member     `json:"member"`
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department"`
}
