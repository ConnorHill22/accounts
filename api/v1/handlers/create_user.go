package accounts_handler

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Got Here"})
}
