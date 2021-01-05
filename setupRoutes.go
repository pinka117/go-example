package main

import (
	_ "example/docs"
	"example/routes"
	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.Handler) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	routes.InitUserRoutes()
	app.Post("/signup", routes.PostSignup)

	app.Post("/login", routes.PostLogin)

}
