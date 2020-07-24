package middleware

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/keyauth"
)

// Keyauth Protected protect routes
func KeyauthProtected() func(*fiber.Ctx) {
	return keyauth.New()
}
