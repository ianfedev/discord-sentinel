package http

import (
	"discord-sentinel/core/config"
	"github.com/gofiber/fiber/v2"
)

// SetupHTTPServer establishes the HTTP server to be used
func SetupHTTPServer(cfg *config.Config) (*fiber.App, error) {

	startMsg := cfg.Environment == config.Production

	app := fiber.New(fiber.Config{
		DisableStartupMessage: startMsg,
	})

	return app, nil

}
