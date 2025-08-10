package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"service": "chat-service",
		})
	})

	// Chat endpoints
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Chat Service",
			"status": "running",
		})
	})

	// WebSocket endpoint
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// Simple WebSocket handler
		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			
			// Echo the message back
			err = c.WriteMessage(messageType, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	app.Get("/messages", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get messages endpoint",
			"status": "not implemented yet",
		})
	})

	app.Post("/messages", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Send message endpoint",
			"status": "not implemented yet",
		})
	})

	log.Fatal(app.Listen(":8084"))
} 