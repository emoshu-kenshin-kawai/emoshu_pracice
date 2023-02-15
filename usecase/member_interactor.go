package usecase

import "emoshu_practice/domain"

type MemberInteractor struct {
	MemberRepository
}

func (interactor *MemberInteractor) GetMemberById(id string) (domain.Member, error) {
	member, err := interactor.FindById(id)
	return member, err
}
