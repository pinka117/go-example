package routes

import (
	"log"

	"example/models"
	"example/utils"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

func PostSignup(c *fiber.Ctx) error {
	//Prendo il body in JSON e lo metto dentro un oggetto
	userRequest := new(UserSignupRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	//Valido il body
	validate = validator.New()
	if err := validate.Struct(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//Faccio l'hash della password dell'utente
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	//Creo un nuovo documento per l'utente
	user := models.NewUser(userRequest.Name, userRequest.Surname, string(hashedPassword), userRequest.Mail)

	// Salvo a db l'utente
	if err := mgm.Coll(user).Create(user); err != nil {
		log.Print(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	//Faccio una query per cercare l'utente appena salvato
	userSaved := &models.User{}
	coll := mgm.Coll(userSaved)
	coll.First(bson.M{"mail": userRequest.Mail}, userSaved)

	//Faccio una risposta con l'utente salvato
	return c.JSON(userSaved)
}

func PostLogin(c *fiber.Ctx) error {
	userRequest := new(UserLoginRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	//Valido il body
	validate = validator.New()
	if err := validate.Struct(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	userSaved := &models.User{}
	coll := mgm.Coll(userSaved)

	if err := coll.First(bson.M{"mail": userRequest.Mail}, userSaved); err != nil && userSaved.Password != "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if checkPasswordHash(userRequest.Password, userSaved.Password) {
		sess, _ := utils.Store.Get(c)

		sess.Set("name", []byte(userRequest.Mail))
		defer sess.Save()
		return c.SendStatus(fiber.StatusOK)
	}

	return c.SendStatus(fiber.StatusBadRequest)

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
