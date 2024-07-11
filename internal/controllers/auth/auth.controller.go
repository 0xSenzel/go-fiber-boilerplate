package auth

import (
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/models"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/service/auth"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func LoginUserHandler(c fiber.Ctx, db *gorm.DB) error {
	var userRequestDto models.UserRequestDto
    if err := c.Bind().Body(&userRequestDto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   true,
            "message": "Invalid request payload",
        })
    }

	token, err := auth.LoginUser(db, userRequestDto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(), 
			"message": "Failed to login user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"access_token": token})
}