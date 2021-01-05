package services

import (
	"testing"
	"github.com/golang/mock/gomock"

	"example/models"
)

func TestSave(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	
	m := NewMockUserRepository(ctrl)

}
