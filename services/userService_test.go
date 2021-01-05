package services

import (
	"example/mocks"
	"example/models"
	"example/utils"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCheckPasswordOk(t *testing.T) {
	mock := mockUserRepository(t)
	hashedPassword, _ := utils.HashPassword("password")
	mock.
		EXPECT().
		SearchByMail("mail").
		Return(models.NewUser("name", "surname", hashedPassword, "mail"))

	SUT := UserService{mock}
	if err := SUT.CheckUserPassword("mail", "password"); err != nil {
		t.Error("Check Hash Password Failed")
	}
}

func TestCheckPasswordWrong(t *testing.T) {
	mock := mockUserRepository(t)
	mock.
		EXPECT().
		SearchByMail("mail").
		Return(models.NewUser("name", "surname", "notOkPassword", "mail"))

	SUT := UserService{mock}
	if err := SUT.CheckUserPassword("mail", "password"); err == nil {
		t.Error("Check wrong password failed")
	}
}

func mockUserRepository(t *testing.T) *mocks.MockIUserRepository {
	mockCtrl := gomock.NewController(t)
	mock := mocks.NewMockIUserRepository(mockCtrl)
	return mock
}
