package routes

import "github.com/gofiber/fiber/v2"

func PostSignup(c *fiber.Ctx) error {
	return c.SendString("hello")
}
