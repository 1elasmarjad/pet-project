package main

import (
	"github.com/gofiber/fiber/v2"
	"src/api"
	"src/db"
)

func main() {
	var err error // variable to handle errors

	db.ConnectDB() // connect database

	if err != nil {
		panic("Failed to migrate Pet")
	}

	app := fiber.New()      // create application
	api.SetupPetRoutes(app) // set-up endpoints

	err = app.Listen(":3000")
	if err != nil {
		panic("Failed to listen to port")
	}
}
