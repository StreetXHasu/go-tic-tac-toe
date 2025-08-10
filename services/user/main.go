package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"service": "user-service",
		})
	})

	// User endpoints
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "User Service",
			"status": "running",
		})
	})

	app.Get("/profile", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get profile endpoint",
			"status": "not implemented yet",
		})
	})

	app.Put("/profile", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Update profile endpoint",
			"status": "not implemented yet",
		})
	})

	log.Fatal(app.Listen(":8082"))
} 