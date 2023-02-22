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

func (repo *MemberRepository) New(m domain.Member) (domain.Member, error) {
	if err := repo.Create(&m).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (repo *MemberRepository) Update(m domain.Member) (err error) {
	if err = repo.Model(&domain.Member{}).Where("id = ?", m.ID).Updates(map[string]interface{}{
		"No":                 m.No,
		"ProfileImg":         m.ProfileImg,
		"FullName":           m.FullName,
		"KanaName":           m.KanaName,
		"Motto":              m.Motto,
		"Biography":          m.Biography,
		"StartDate":          m.StartDate,
		"EndDate":            m.EndDate,
		"EmploymentStatusID": m.EmploymentStatusID,
		"StatusID":           m.StatusID,
	}).Error; err != nil {
		return err
	}
	return nil
}
