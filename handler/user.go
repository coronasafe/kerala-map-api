package handler

import (
	"strconv"

	"github.com/coronasafe/kerala-map-api/config"
	"github.com/coronasafe/kerala-map-api/database"
	"github.com/coronasafe/kerala-map-api/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	if uid != n {
		return false
	}

	return true
}

func validUser(id string, p string) bool {
	db := database.DB
	var user model.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) {
	token := c.Locals("token").(string)
	if token != config.Config.KEY {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid API key", "data": nil})
		return
	}
	type NewUser struct {
		Username string `json:"username"`
	}

	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		return
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
		return
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
		return
	}

	newUser := NewUser{
		Username: user.Username,
	}

	c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// DeleteUser delete user
func DeleteUser(c *fiber.Ctx) {
	token := c.Locals("token").(string)
	if token != config.Config.KEY {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid API key", "data": nil})
		return
	}
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		return
	}
	id := c.Params("id")

	if !validUser(id, pi.Password) {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})
		return
	}

	db := database.DB
	var user model.User

	db.First(&user, id)

	db.Delete(&user)
	c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
