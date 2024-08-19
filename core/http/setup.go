package http

import (
	"discord-sentinel/core/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

// SetupHTTPServer establishes the HTTP server to be used
func SetupHTTPServer(cfg *config.Config) (*fiber.App, error) {

	startMsg := cfg.Environment == config.Production

	app := fiber.New(fiber.Config{
		DisableStartupMessage: startMsg,
	})
	app.Use(requestid.New(requestid.Config{
		Header:     "X-Sentinel-Ray",
		Generator:  utils.UUID,
		ContextKey: "ray_id",
	}))

	return app, nil

}
