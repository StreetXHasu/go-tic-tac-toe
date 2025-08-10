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
			"service": "auth-service",
		})
	})

	// Auth endpoints
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Auth Service",
			"status": "running",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Login endpoint",
			"status": "not implemented yet",
		})
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Register endpoint",
			"status": "not implemented yet",
		})
	})

	log.Fatal(app.Listen(":8081"))
} 