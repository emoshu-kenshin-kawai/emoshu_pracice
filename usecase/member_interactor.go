package usecase

import "emoshu_practice/domain"

type MemberInteractor struct {
	MemberRepository
}

func (interactor *MemberInteractor) GetMembers() ([]domain.Member, error) {
	members, err := interactor.FindAll()
	return members, err
}

func (interactor *MemberInteractor) GetMemberById(id string) (domain.Member, error) {
	member, err := interactor.FindById(id)
	return member, err
}

func (interactor *MemberInteractor) CreateMember(member domain.Member) (domain.Member, error) {
	member, err := interactor.New(member)
	return member, err
}

func (interactor *MemberInteractor) UpdateMember(member domain.Member) (domain.Member, error) {
	member, err := interactor.Update(member)
	return member, err
}
