package routes

import (
	"github.com/0xsenzel/go-fiber-boilerplate/internal/controllers/auth"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/controllers/user"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/middlewares"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	app.Post("/user/create", func(c fiber.Ctx) error {
		return user.CreateUserHandler(c, db)
	})

	app.Post("auth/login", func(c fiber.Ctx) error {
		return auth.LoginUserHandler(c, db)
	})

	protectedRoute(app, db)
}

func protectedRoute(app *fiber.App, db *gorm.DB) {
	protected := app.Group("/user", func(c fiber.Ctx) error {
		return middlewares.JwtAuth(c)
	})
	protected.Get("/:id", func(c fiber.Ctx) error {
		return user.GetUserHandler(c, db)
	})
}