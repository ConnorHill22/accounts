package main

import (
	"log"

	db_conn "connorhill/accounts/cmd/common"
	"connorhill/accounts/internal/models"
)

func main() {
	log.Println("Starting Migrations")
	db, err := db_conn.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := db.AutoMigrate(
		&models.EmailValidationStatus{},
		&models.HashCost{},
		&models.UserLoginData{},
	); err != nil {
		log.Fatalln(err)
	}
	log.Println("Sucessfully made migrations")
}
