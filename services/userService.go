package services

import (
	"errors"
	"log"

	"example/models"
	"example/repositories"
	"example/utils"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	SaveUser(mail, name, password, surname string) (*models.User, error)
	CheckUserPassword(mail, password string) error
}

type UserService struct {
	UserRepository *repositories.UserRepository
}

func (p UserService) SaveUser(mail, name, password, surname string) (*models.User, error) {
	userDuplicate := p.UserRepository.SearchByMail(mail)
	if userDuplicate.Mail != "" {
		return nil, errors.New("Duplicate User")
	}

	//Faccio l'hash della password dell'utente
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	//Creo un nuovo documento per l'utente
	if err := p.UserRepository.SaveUser(name, surname, string(hashedPassword), mail); err != nil {
		return nil, err
	}

	//Faccio una query per cercare l'utente appena salvato
	userSaved := p.UserRepository.SearchByMail(mail)

	return userSaved, nil
}

func (p UserService) CheckUserPassword(mail, password string) error {
	userSaved := p.UserRepository.SearchByMail(mail)

	if userSaved.Password == "" {
		return errors.New("Wrong User or Password")
	}

	if utils.CheckPasswordHash(password, userSaved.Password) {
		return nil
	}
	return errors.New("Wrong User or Password")

}
