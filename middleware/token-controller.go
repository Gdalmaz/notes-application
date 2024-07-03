package middleware

import (
	"log"
	"notes-application/database"
	"notes-application/models"

	"github.com/gofiber/fiber/v2"
)

func TokenControl(c *fiber.Ctx) (models.User, error) {
	db := database.DB.Db
	authorizationHeader := c.Get("Authorization")

	if authorizationHeader == "" || len(authorizationHeader) < 7 || authorizationHeader[:7] != "Bearer " {
		return models.User{}, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}
	token := authorizationHeader[7:]
	var user models.User
	session := new(models.Session)
	err := db.Where("token =?", token).First(&session).Error
	if err != nil {
		return models.User{}, c.Status(500).JSON(fiber.Map{"status": "error", "message": "you don't have session", "data": err})
	}

	userID := session.UserID
	log.Println("userid ========", userID)
	err = db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return models.User{}, c.Status(500).JSON(fiber.Map{"status": "error", "message": "user not found it is illegall !!", "data": err})
	}

	return user, nil

}
