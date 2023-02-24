package models

type HashAlgorithm struct {
	ID            uint `gorm:"primaryKey"`
	AlgorithmName string
}
