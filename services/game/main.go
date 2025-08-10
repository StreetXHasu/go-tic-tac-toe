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
			"service": "game-service",
		})
	})

	// Game endpoints
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Game Service",
			"status": "running",
		})
	})

	app.Post("/games", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Create game endpoint",
			"status": "not implemented yet",
		})
	})

	app.Get("/games/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get game endpoint",
			"status": "not implemented yet",
		})
	})

	app.Post("/games/:id/move", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Make move endpoint",
			"status": "not implemented yet",
		})
	})

	log.Fatal(app.Listen(":8083"))
} 