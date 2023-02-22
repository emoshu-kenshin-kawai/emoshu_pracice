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

func (interactor *MemberInteractor) CreateMember(member domain.Member) error {
	err := interactor.New(member)
	return err
}

func (interactor *MemberInteractor) UpdateMember(member domain.Member) error {
	err := interactor.Update(member)
	return err
}
