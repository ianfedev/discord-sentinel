package http

import (
	"github.com/gofiber/fiber/v2"
)

// SetupHTTPServer establishes the HTTP server to be used
func SetupHTTPServer() (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	return app, nil

}
