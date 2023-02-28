package controller

import (
	"emoshu_practice/domain"
	"emoshu_practice/usecase/mock_usecase"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMemberCreate(t *testing.T) {

}

func TestMemberShow(t *testing.T) {
	// dbHandler := InitDB()
	// memberController := NewMemberController(dbHandler)
	// con := MemberController{}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := domain.Member{}

	mockRepo := mock_usecase.NewMockMemberRepository(mockCtrl)
	mockRepo.EXPECT().New(m).Return(nil, nil)
}
