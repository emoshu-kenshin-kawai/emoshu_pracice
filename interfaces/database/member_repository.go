package database

import (
	"emoshu_practice/domain"
)

type MemberRepository struct {
	DBHandler
}

func (repo *MemberRepository) FindById(id string) (member domain.Member, err error) {
	if err = repo.Joins("EmploymentStatus").Joins("Status").First(&member, id).Error; err != nil {
		return member, err
	}
	return member, err
}
