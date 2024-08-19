package punishment

import (
	"discord-sentinel/core/database"
	"discord-sentinel/core/http"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	app.Get("/punishment/:id", rh.GetByID)
	app.Put("/punishment/:id", rh.Update)

	return rh
}

// Create handles POST requests to create a new punishment.
func (h *RouteHandler) Create(c *fiber.Ctx) error {

	var punishment Punishment

	// Parse the JSON body into the Punishment struct
	if err := c.BodyParser(&punishment); err != nil {
		return http.PerformError(c, h.logger, "Error while parsing request body", 400, &err)
	}

	punishment.CreatedAt = time.Now()
	punishment.UpdatedAt = time.Now()

	// Every punishment created via API should not be automatic.
	punishment.Automatic = false

	if err := (*h.service).Create(c.Context(), &punishment); err != nil {
		return http.PerformError(c, h.logger, "Error while creating punishment", 500, &err)
	}

	return c.Status(fiber.StatusCreated).JSON(punishment)

}

// GetByID obtains from database a specific id.
func (h *RouteHandler) GetByID(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return http.PerformError(c, h.logger, "No valid ID was provided", 400, &err)
	}

	punishment, err := (*h.service).GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("punishment not found")
		}
		return http.PerformError(c, h.logger, "Error while retrieving punishment", 500, &err)
	}

	return c.Status(fiber.StatusOK).JSON(punishment)
}

func (h *RouteHandler) Update(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return http.PerformError(c, h.logger, "No valid ID was provided", 400, &err)
	}

	var punishment Punishment

	punRec, err := (*h.service).GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("punishment not found")
		}
		return http.PerformError(c, h.logger, "Error while retrieving punishment", 500, &err)
	}

	if err := c.BodyParser(&punishment); err != nil {
		return http.PerformError(c, h.logger, "Error while parsing request body", 400, &err)
	}

	// TODO: Set this at service level
	punishment.Id = punRec.Id
	punishment.CreatedAt = punRec.CreatedAt
	punishment.UpdatedAt = time.Now()

	if err := (*h.service).Update(c.Context(), &punishment); err != nil {
		return http.PerformError(c, h.logger, "Error while updating punishment", 500, &err)
	}

	return c.Status(fiber.StatusOK).JSON(punishment)

}
