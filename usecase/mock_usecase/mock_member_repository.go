// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/member_repository.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	domain "emoshu_practice/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMemberRepository is a mock of MemberRepository interface.
type MockMemberRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMemberRepositoryMockRecorder
}

// MockMemberRepositoryMockRecorder is the mock recorder for MockMemberRepository.
type MockMemberRepositoryMockRecorder struct {
	mock *MockMemberRepository
}

// NewMockMemberRepository creates a new mock instance.
func NewMockMemberRepository(ctrl *gomock.Controller) *MockMemberRepository {
	mock := &MockMemberRepository{ctrl: ctrl}
	mock.recorder = &MockMemberRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemberRepository) EXPECT() *MockMemberRepositoryMockRecorder {
	return m.recorder
}

// DeleteById mocks base method.
func (m *MockMemberRepository) DeleteById(arg0 domain.Member) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockMemberRepositoryMockRecorder) DeleteById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockMemberRepository)(nil).DeleteById), arg0)
}

// FindAll mocks base method.
func (m *MockMemberRepository) FindAll() ([]domain.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]domain.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockMemberRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockMemberRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockMemberRepository) FindById(id string) (domain.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(domain.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockMemberRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockMemberRepository)(nil).FindById), id)
}

// New mocks base method.
func (m *MockMemberRepository) New(member domain.Member) (domain.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", member)
	ret0, _ := ret[0].(domain.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockMemberRepositoryMockRecorder) New(member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockMemberRepository)(nil).New), member)
}

// Update mocks base method.
func (m *MockMemberRepository) Update(member domain.Member) (domain.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", member)
	ret0, _ := ret[0].(domain.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMemberRepositoryMockRecorder) Update(member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMemberRepository)(nil).Update), member)
}
