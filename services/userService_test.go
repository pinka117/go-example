package services

import (
	"example/mocks"
	"example/models"
	"example/utils"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestSave(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockIUserRepository(mockCtrl)

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
