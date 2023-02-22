package usecase

import "emoshu_practice/domain"

type MemberRepository interface {
	FindById(id string) (domain.Member, error)
	FindAll() ([]domain.Member, error)
	New(member domain.Member) (domain.Member, error)
	Update(member domain.Member) error
}
