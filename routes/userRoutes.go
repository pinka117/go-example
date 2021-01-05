package routes

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"

	"example/request"
	"example/services"
	"example/utils"
)

var validate *validator.Validate

func PostSignup(c *fiber.Ctx) error {
	//Prendo il body in JSON e lo metto dentro un oggetto
	userRequest := new(request.UserSignupRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	//Valido il body
	validate = validator.New()
	if err := validate.Struct(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//Salvo l'utente
	userSaved, err := services.SaveUser(userRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//Faccio una risposta con l'utente salvato
	return c.JSON(userSaved)
}

func PostLogin(c *fiber.Ctx) error {
	userRequest := new(request.UserLoginRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	//Valido il body
	validate = validator.New()
	if err := validate.Struct(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := services.CheckUserPassword(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	sess, err := utils.Store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	sess.Set("name", []byte(userRequest.Mail))
	defer sess.Save()

	return c.SendStatus(fiber.StatusOK)

}
