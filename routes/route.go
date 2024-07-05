package routes

import (
	"github.com/0xsenzel/go-fiber-boilerplate/internal/controllers/user"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString(user.HelloWorldHandler())
	})
}