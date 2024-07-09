package main

import (
	"log"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/database/mysql"
	// "github.com/0xsenzel/go-fiber-boilerplate/internal/middlewares"
	"github.com/0xsenzel/go-fiber-boilerplate/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	db := mysql.ConnectDB()
	defer mysql.CloseDB(db)
	mysql.Migrate(db)

	// app.Use(middlewares.JWTAuth())
	routes.SetupRoutes(app, db)
    
	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
