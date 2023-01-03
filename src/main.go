package main

import (
	"github.com/gofiber/fiber/v2"
	"src/api"
	"src/db"
)

func main() {
	var err error

	app := fiber.New()

	// --- DATABASE ---

	db.ConnectDB()

	if err != nil {
		panic("Failed to migrate Pet")
	}

	// --- ENDPOINTS ---

	api.SetupPetRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		panic("Failed to listen to port")
	}
}
