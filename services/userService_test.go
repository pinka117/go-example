package services

import (
	"testing"
	"github.com/golang/mock/gomock"

	"example/models"
)

func TestSave(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockRepo := NewMockRepository(ctrl)

	u := models.NewUser("jack", "surname","password","mail@mail.com")
	mockRepo.

}
