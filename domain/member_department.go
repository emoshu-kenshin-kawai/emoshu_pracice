package domain

import "gorm.io/gorm"

type MemberDepartment struct {
	gorm.Model
	Member     Member
	Department Department
}
