package utils

import (
	"example/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/signup", routes.PostSignup)

}
