package punishment

import (
	"discord-sentinel/core/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

// RouteHandler defines the transport logic for the Punishment resource.
// It encapsulates the service logic for handling Punishment-related operations.
type RouteHandler struct {
	service *database.Service[Punishment]
	logger  *zap.Logger
}

// NewPunishmentHandler returns a new handler for Punishment related operations.
func NewPunishmentHandler(service *database.Service[Punishment], logger *zap.Logger, app *fiber.App) *RouteHandler {

	rh := &RouteHandler{
		service: service,
		logger:  logger,
	}

	app.Post("/punishment", rh.Create)

	return rh
}

// Create handles POST requests to create a new punishment.
func (h *RouteHandler) Create(c *fiber.Ctx) error {

	var punishment Punishment

	// Parse the JSON body into the Punishment struct
	if err := c.BodyParser(&punishment); err != nil {
		(*h.logger).Error("Error while parsing request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid_request_payload",
			"message": err.Error(),
		})
	}

	punishment.CreatedAt = time.Now()
	punishment.UpdatedAt = time.Now()

	if err := (*h.service).Create(c.Context(), &punishment); err != nil {
		(*h.logger).Error("Error while creating punishment", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "internal_server_error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(punishment)

}
