package domain

import "gorm.io/gorm"

type MemberDepartment struct {
	gorm.Model
	MemberID     uint
	Member       Member
	DepartmentID uint
	Department   Department
}
