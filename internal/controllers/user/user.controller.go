package user

import (
	"log"
	"strconv"

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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(), 
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": `User created successfully`, "user": createdUser})
}

func GetUserHandler(c fiber.Ctx, db *gorm.DB) error {
	userIdStr := c.Params("id")
	log.Printf("Getting user with id:%s", userIdStr)

	// parse user id to int
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid user id",
		})
	}
	// get user by id
	user, err := service.GetUserById(db, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(), 
			"message": "Failed to get user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
}