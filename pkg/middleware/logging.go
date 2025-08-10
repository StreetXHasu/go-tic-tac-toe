package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/your-org/go-tic-tac-toe/pkg/utils"
)

// LoggingMiddleware creates middleware for logging using logger
func LoggingMiddleware(logger *utils.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		latency := time.Since(start)
		
		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()
		remoteAddr := c.IP()
		
		// Log request
		logger.LogRequest(method, path, remoteAddr, status, latency)
		
		if err != nil {
			logger.LogError("HTTP Request", err)
			c.App().Config().ErrorHandler(c, err)
		}
		
		return err
	}
}

// Simple logger for backward compatibility
func Logging() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		latency := time.Since(start)
		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()
		if err != nil {
			c.App().Config().ErrorHandler(c, err)
		}
		println(method, path, status, latency.String())
		return err
	}
} 