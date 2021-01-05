package repositories

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"

	"example/models"
)

type UserRepository struct {
}

type IUserRepository interface {
	SearchByMail(mail string) *models.User
	SaveUser(name, surname, hashedPassword, mail string) error
}

func (p UserRepository) SearchByMail(mail string) *models.User {
	userSaved := &models.User{}
	coll := mgm.Coll(userSaved)
	coll.First(bson.M{"mail": mail}, userSaved)
	return userSaved
}

func (p UserRepository) SaveUser(name, surname, hashedPassword, mail string) error {
	// Salvo a db l'utente
	user := models.NewUser(name, surname, hashedPassword, mail)

	if err := mgm.Coll(user).Create(user); err != nil {

		return err
	}
	return nil
}
