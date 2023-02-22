package database

import (
	"emoshu_practice/domain"
)

type MemberRepository struct {
	DBHandler
}

func (repo *MemberRepository) FindAll() (members []domain.Member, err error) {
	if err = repo.Joins("EmploymentStatus").Joins("Status").Find(&members).Error; err != nil {
		return members, err
	}
	return members, nil
}

func (repo *MemberRepository) FindById(id string) (member domain.Member, err error) {
	if err = repo.Joins("EmploymentStatus").Joins("Status").First(&member, id).Error; err != nil {
		return member, err
	}
	return member, nil
}
