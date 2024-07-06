package main

import (
	"log"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/database/mysql"
	"github.com/0xsenzel/go-fiber-boilerplate/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	mysql.ConnectDB()
	defer mysql.CloseDB()

	mysql.Migrate()

	routes.SetupRoutes(app)
    
	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
