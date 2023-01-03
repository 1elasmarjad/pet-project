package api

import (
	"github.com/gofiber/fiber/v2"
)
import "src/db"

func SetupPetRoutes(router fiber.Router) {
	pets := router.Group("/pets")

	pets.Get("/:petID", getPet)
	pets.Get("/", getPets)

	pets.Post("/", createPet)
}

func getPet(c *fiber.Ctx) error {
	c.Type("application/json")
	var pet db.Pet

	// Read the path param
	targetID := c.Params("petID")

	// Find pet
	db.DB.Find(&pet, "id = ?", targetID)

	// Not found?
	if pet.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Pet not found", "data": nil})
	}

	// Found
	return c.JSON(fiber.Map{"status": "success", "message": "Pet found", "data": pet})
}

func getPets(c *fiber.Ctx) error {
	c.Type("application/json")
	var pets []db.Pet

	db.DB.Find(&pets)

	return c.JSON(fiber.Map{"status": "success", "message": "Pets found", "data": pets})
}

func createPet(c *fiber.Ctx) error {
	c.Type("application/json")
	var pet = db.Pet{}

	err := c.BodyParser(pet)

	//if err != nil {
	//	return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	//}

	err = db.DB.Create(&pet).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create pet", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created pet", "data": pet})
}
