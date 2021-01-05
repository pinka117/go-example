package main

import (
	"example/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/signup", routes.PostSignup)

	app.Post("/login", routes.PostLogin)


}
