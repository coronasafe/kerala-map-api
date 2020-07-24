package main

import (
	"github.com/coronasafe/kerala-map-api/config"
	"github.com/coronasafe/kerala-map-api/database"
	"github.com/coronasafe/kerala-map-api/router"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	app.Listen(config.Config.Port)

	defer database.DB.Close()
}
