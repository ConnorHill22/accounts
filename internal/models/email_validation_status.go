package models

type EmailValidationStatus struct {
	ID                uint `gorm:"primaryKey"`
	StatusDescription string
}
