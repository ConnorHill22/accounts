package models

import (
	"time"
)

type UserLoginData struct {
	ID                      uint   `gorm:"primaryKey"`
	OpaqueId                string `gorm:"unique;not null"`
	LoginName               string `gorm:"unique;not null"`
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
