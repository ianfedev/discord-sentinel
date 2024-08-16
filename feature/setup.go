package feature

import (
	"discord-sentinel/core/database"
	"discord-sentinel/feature/punishment"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SetupFeatures configures every transport and service used by Sentinel
// features with loggers, apps and required databases.
func SetupFeatures(app *fiber.App, logger *zap.Logger, db *gorm.DB) {

	punishRepo := database.NewGormRepository[punishment.Punishment](db)
	punishSvc := database.NewService[punishment.Punishment](punishRepo)
	punishment.NewPunishmentHandler(&punishSvc, logger, app)

}
