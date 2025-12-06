package handler

import (
	"github.com/gofiber/fiber/v2"
)

func HealthHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!, Olha como o mundo e' maravilhoso")
}
