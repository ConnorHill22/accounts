package accounts_handler

import (
	uuid "connorhill/accounts/internal/common"
	"connorhill/accounts/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequest {
	user_name 		string		`json:"user_name"`
	email_address	string 		`json:"email_address`
	password 		string 		`json:"password"`
}

func CreateUser(c *fiber.Ctx) error {

	request := new(CreateUserRequest)
	err := c.BodyParser(request)
	if err != nil {
		return err
	}

	user_opaque_id = uuid.Generate(models.UserLoginData.Name())

	user := model.UserLoginData{
		OpaqueId: user_opaque_id,
		LoginName: request.user_name
	}

	`type UserLoginData struct {
		ID                      uint   `gorm:"primaryKey"`
		OpaqueId                string `gorm:"unique;not null;index"`
		LoginName               string `gorm:"unique;not null;index"`
		PasswordHash            string `gorm:"not null"`
		PasswordSalt            string `gorm:"not null"`
		HashAlgorithmID         int
		HashAlgorithm           HashAlgorithm
		EmailAddress            string `gorm:"not null"`
		ConfirmationToken       string
		TokenGeneratedTime      time.Time
		EmailValidationStatusID int
		EmailValidationStatus   EmailValidationStatus
		PasswordRecoveryToken   string
		RecoveryTokenTime       time.Time
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Got Here"})
}
