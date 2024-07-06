package routes

import (
	"github.com/0xsenzel/go-fiber-boilerplate/internal/controllers/user"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/user/create", func(c fiber.Ctx) error {
		return user.CreateUserHandler(c, db)
	})
}