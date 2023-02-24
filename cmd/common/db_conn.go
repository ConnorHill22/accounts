package db_conn

import (
	"fmt"
	"log"
	"strconv"

	"connorhill/accounts/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	var err error

	p := config.Config("DATABASE_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return nil, err
	}

	conn_string := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Config("DATABASE_HOST"),
		port,
		config.Config("DATABASE_USER"),
		config.Config("DATABASE_PASS"),
		config.Config("DATABASE_NAME"),
		config.Config("DATABASE_SSL_MODE"),
	)
	DB, err = gorm.Open(postgres.Open(conn_string))
	if err != nil {
		return nil, err
	}

	log.Println("Connection Opened to Database")
	return DB, nil
}
