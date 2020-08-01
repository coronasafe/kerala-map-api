package handler

import (
	"github.com/coronasafe/kerala-map-api/database"
	"github.com/coronasafe/kerala-map-api/model"

	"github.com/gofiber/fiber"
)

// GetAllFeatures query all features
func GetAllFeatures(c *fiber.Ctx) {
	db := database.DB
	var features []model.Feature
	db.Find(&features)
	c.JSON(fiber.Map{"status": "success", "message": "All features", "data": features})
}

// GetFeature query feature
func GetFeature(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var feature model.Feature
	db.Find(&feature, id)
	if feature.Data == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No feature found with ID", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Feature found", "data": feature})
}

// CreateFeature new feature
func CreateFeature(c *fiber.Ctx) {
	db := database.DB
	feature := new(model.Feature)
	if err := c.BodyParser(feature); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create feature", "data": err})
		return
	}
	db.Create(&feature)
	c.JSON(fiber.Map{"status": "success", "message": "Created feature", "data": feature})
}

// UpdateFeature updates feature
func UpdateFeature(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	type UpdateFeatureInput struct {
		Data        string `json:"data"`
		Description string `json:"description"`
	}
	var upi UpdateFeatureInput
	if err := c.BodyParser(&upi); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		return
	}

	var feature model.Feature
	db.First(&feature, id)
	if feature.Data == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No feature found with ID", "data": nil})
		return
	}
	feature.Data = upi.Data
	feature.Description = upi.Description
	db.Save(&feature)
	c.JSON(fiber.Map{"status": "success", "message": "Feature successfully updated", "data": nil})
}

// DeleteFeature delete feature
func DeleteFeature(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	var feature model.Feature
	db.First(&feature, id)
	if feature.Data == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No feature found with ID", "data": nil})
		return
	}
	db.Delete(&feature)
	c.JSON(fiber.Map{"status": "success", "message": "Feature successfully deleted", "data": nil})
}
