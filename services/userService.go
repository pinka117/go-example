package services

import (
	"errors"
	"log"

	"example/models"
	"example/request"
	"example/utils"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser(userRequest *request.UserSignupRequest) (*models.User, error) {
	userSaved := searchByMail(userRequest.Mail)
	if userSaved.Mail != "" {
		return nil, errors.New("Duplicate User")
	}

	//Faccio l'hash della password dell'utente
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	//Creo un nuovo documento per l'utente
	user := models.NewUser(userRequest.Name, userRequest.Surname, string(hashedPassword), userRequest.Mail)

	// Salvo a db l'utente
	if err := mgm.Coll(user).Create(user); err != nil {
		log.Print(err)
		return nil, err
	}

	//Faccio una query per cercare l'utente appena salvato
	userSaved2 := searchByMail(userRequest.Mail)

	return userSaved2, nil
}

func searchByMail(mail string) *models.User {
	userSaved := &models.User{}
	coll := mgm.Coll(userSaved)
	coll.First(bson.M{"mail": mail}, userSaved)
	return userSaved
}

func CheckUserPassword(userRequest *request.UserLoginRequest) error {
	userSaved := &models.User{}
	coll := mgm.Coll(userSaved)

	if err := coll.First(bson.M{"mail": userRequest.Mail}, userSaved); err != nil && userSaved.Password != "" {
		return err
	}
	if utils.CheckPasswordHash(userRequest.Password, userSaved.Password) {
		return nil
	}
	return errors.New("Wrong User or Password")

}
