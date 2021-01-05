package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"example/utils"
)

func main() {
	app := fiber.New()

	utils.InitRedisSession()

	utils.InitDb()

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
