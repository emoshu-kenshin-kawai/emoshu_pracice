package domain

import "gorm.io/gorm"

type MemberRole struct {
	gorm.Model
	MemberID uint
	Member   Member
	RoleID   uint
	Role     Role
}
