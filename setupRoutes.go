package main

import (
	"github.com/gofiber/fiber/v2"
	"example/routes"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/signup", routes.PostSignup)

	app.Post("/login", routes.PostLogin)

}
