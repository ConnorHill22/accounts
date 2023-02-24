package main

import (
	"fmt"
	"log"
	"strconv"

	v1_routes "connorhill/accounts/api/v1/router"
	db_conn "connorhill/accounts/cmd/common"

	"connorhill/accounts/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	/*
		Database Connection
	*/
	_, err := db_conn.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	/*
		Start Server
	*/
	p := config.Config("SERVER_PORT")
	port, err := strconv.ParseInt(p, 0, 0)
	if err != nil {
		log.Fatalln("Issue fetching server port with error=%v", err)
		return
	}

	app := fiber.New()
	v1_routes.Accounts(app.Group("/v1"))
	log.Fatalln(app.Listen(fmt.Sprintf(":%d", *port)))
}
