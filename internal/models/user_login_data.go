package models

import (
	"time"
)

type UserLoginData struct {
	ID                      uint   `gorm:"primaryKey"`
	OpaqueId                string `gorm:"unique;not null;index"`
	LoginName               string `gorm:"unique;not null;index"`
	PasswordHash            string `gorm:"not null"`
	HashCostID              int
	HashCost                HashCost
	EmailAddress            string `gorm:"not null"`
	ConfirmationToken       string
	TokenGeneratedTime      time.Time
	EmailValidationStatusID int
	EmailValidationStatus   EmailValidationStatus
	PasswordRecoveryToken   string
	RecoveryTokenTime       time.Time
}
