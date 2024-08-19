package http

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// PerformError logs error with sentinel ray_id and returns
// error performed by Fiber.
func PerformError(ctx *fiber.Ctx, logger *zap.Logger, message string, status int, err *error) error {
	id := ctx.Locals("ray_id").(string)
	logger.Error(message, zap.String("ray_id", id), zap.Error(*err))
	return ctx.Status(status).JSON(fiber.Map{
		"message": (*err).Error(),
		"ray_id":  id,
	})
}
