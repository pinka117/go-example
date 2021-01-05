package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"example/utils"
)

// @title Go Example
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	utils.InitRedisSession()

	utils.InitDb()

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
