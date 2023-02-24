package routes

import (
	accounts_handler "connorhill/accounts/api/v1/handlers"

	"github.com/gofiber/fiber/v2"
)

// Endpoints
//
// Accounts
// POST     .../user
// GET      .../user/<opaque_id>
// PUT      .../user/<opaque_id>
// DELETE   .../user/<opaque_id>
//
// AuthZ
// POST     .../login
//
// AuthN
// GET      .../user/<opaque_id>/permissions
//

func Accounts(router fiber.Router) {
	router.Post("/user", accounts_handler.CreateUser)
}
