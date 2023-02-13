package domain

import "gorm.io/gorm"

type MemberRole struct {
	gorm.Model
	Member Member
	Role   Role
}
