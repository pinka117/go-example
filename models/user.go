package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Surname          string `json:"surname" bson:"surname"`
	Password         string `json:"-" bson:"password"`
	Mail             string `json:"mail" bson:"mail"`
}

func NewUser(name string, surname string, password string, mail string) *User {
	return &User{
		Name:     name,
		Surname:  surname,
		Password: password,
		Mail:     mail,
	}
}
