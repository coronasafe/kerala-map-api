package router

import (
	"github.com/coronasafe/kerala-map-api/handler"
	"github.com/coronasafe/kerala-map-api/middleware"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Post("/", middleware.KeyauthProtected(), handler.CreateUser)
	user.Delete("/:id", middleware.KeyauthProtected(), handler.DeleteUser)

	// Descriptions
	description := api.Group("/description")
	description.Get("/", handler.GetAllDescriptions)
	description.Get("/:id", handler.GetDescription)
	description.Post("/", middleware.Protected(), handler.CreateDescription)
	description.Patch("/:id", middleware.Protected(), handler.UpdateDescription)
	description.Delete("/:id", middleware.Protected(), handler.DeleteDescription)
}
