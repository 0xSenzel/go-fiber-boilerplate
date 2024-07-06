package user

import (
	"log"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/models"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/service"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func CreateUserHandler(c fiber.Ctx, db *gorm.DB) error {
	var userRequestDto models.UserRequestDto
    if err := c.Bind().Body(&userRequestDto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   true,
            "message": "Invalid request payload",
        })
    }
	log.Printf("Creating user with data:%+v", userRequestDto)

	createdUser, err := service.CreateUser(db, userRequestDto)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "message": "Failed to create user"})
		return err
	}

	c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": `User created successfully`, "user": createdUser})
	return nil
}