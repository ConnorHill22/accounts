package models

type HashCost struct {
	ID   uint `gorm:"primaryKey"`
	Cost int
}
