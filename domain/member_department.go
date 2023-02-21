package domain

import "time"

type MemberDepartment struct {
	ID           uint       `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	MemberID     uint       `json:"member_id"`
	Member       Member     `json:"member"`
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department"`
}
