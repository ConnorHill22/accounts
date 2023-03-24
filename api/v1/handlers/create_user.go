package accounts_handler

import (
	"connorhill/accounts/internal/models"

	"github.com/gofiber/fiber/v2"
)

var User struct {
	LoginName    string
	Password     string
	EmailAddress string
}

func CreateUser(c *fiber.Ctx) error {
	request := new(CreateUserRequest)
	err := c.BodyParser(request)
	if err != nil {
		return err
	}

	user_opaque_id = uuid.Generate(models.UserLoginData.Name())

	user := model.UserLoginData{
		OpaqueId:  user_opaque_id,
		LoginName: request.user_name,
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Got Here"})
}
