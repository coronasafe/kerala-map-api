package handler

import (
	"time"

	"github.com/coronasafe/kerala-map-api/config"
	"github.com/coronasafe/kerala-map-api/database"
	"github.com/coronasafe/kerala-map-api/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByUsername(u string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: u}).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Login get user and password
func Login(c *fiber.Ctx) {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var input LoginInput
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
		return
	}
	username := input.Username
	pass := input.Password

	user, err := getUserByUsername(username)
	if err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
		return
	}

	if user == nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
		return
	}

	ud = UserData{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}

	if !CheckPasswordHash(pass, ud.Password) {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config.JWT))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
