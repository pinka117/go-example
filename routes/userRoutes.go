package routes

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func PostSignup(c *fiber.Ctx) error {
	validate = validator.New()
	userRequest := new(UserSignupRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := validate.Struct(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	log.Print(userRequest)
	return c.SendString("hello")
}
