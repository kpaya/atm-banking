package router

import (
	"github.com/gofiber/fiber/v2"
)

func Initialize() {
	// Create a new Fiber instance
	app := fiber.New()

	// Initialize the routes
	initializeRoutes(app)

	app.Listen(":8000")
}
