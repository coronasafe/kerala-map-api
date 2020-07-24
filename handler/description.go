package handler

import (
	"github.com/coronasafe/kerala-map-api/database"
	"github.com/coronasafe/kerala-map-api/model"

	"github.com/gofiber/fiber"
)

// GetAllDescriptions query all descriptions
func GetAllDescriptions(c *fiber.Ctx) {
	db := database.DB
	var descriptions []model.Description
	db.Find(&descriptions)
	c.JSON(fiber.Map{"status": "success", "message": "All descriptions", "data": descriptions})
}

// GetDescription query description
func GetDescription(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var description model.Description
	db.Find(&description, id)
	if description.District == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No description found with ID", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Description found", "data": description})
}

// CreateDescription new description
func CreateDescription(c *fiber.Ctx) {
	db := database.DB
	description := new(model.Description)
	if err := c.BodyParser(description); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create description", "data": err})
		return
	}
	db.Create(&description)
	c.JSON(fiber.Map{"status": "success", "message": "Created description", "data": description})
}

// UpdateDescription updates description
func UpdateDescription(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	type UpdateDescriptionInput struct {
		Data string `json:"data"`
	}
	var upi UpdateDescriptionInput
	if err := c.BodyParser(&upi); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		return
	}

	var description model.Description
	db.First(&description, id)
	if description.District == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No description found with ID", "data": nil})
		return
	}
	description.Data = upi.Data
	db.Save(&description)
	c.JSON(fiber.Map{"status": "success", "message": "Description successfully updated", "data": nil})
}

// DeleteDescription delete description
func DeleteDescription(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	var description model.Description
	db.First(&description, id)
	if description.District == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No description found with ID", "data": nil})
		return
	}
	db.Delete(&description)
	c.JSON(fiber.Map{"status": "success", "message": "Description successfully deleted", "data": nil})
}
