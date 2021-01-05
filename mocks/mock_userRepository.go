// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/userRepository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "example/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIUserRepository is a mock of IUserRepository interface
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// SearchByMail mocks base method
func (m *MockIUserRepository) SearchByMail(mail string) *models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByMail", mail)
	ret0, _ := ret[0].(*models.User)
	return ret0
}

// SearchByMail indicates an expected call of SearchByMail
func (mr *MockIUserRepositoryMockRecorder) SearchByMail(mail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByMail", reflect.TypeOf((*MockIUserRepository)(nil).SearchByMail), mail)
}

// SaveUser mocks base method
func (m *MockIUserRepository) SaveUser(name, surname, hashedPassword, mail string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", name, surname, hashedPassword, mail)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser
func (mr *MockIUserRepositoryMockRecorder) SaveUser(name, surname, hashedPassword, mail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockIUserRepository)(nil).SaveUser), name, surname, hashedPassword, mail)
}
