package main

import (
	"discord-sentinel/core/config"
	"discord-sentinel/core/database"
	"discord-sentinel/core/logging"
	"go.uber.org/zap"
	"os"
)

func main() {

	// Setup basic logger while loading config
	logger := logging.SetupInitialLogger()
	logger.Info("Starting up Sentinel")

	// Parse config and set up enhanced logger
	cfg, err := config.ParseConfig(logger)
	if err != nil {
		logger.Error("Error reading configuration", zap.Error(err))
		os.Exit(1)
	}

	logger, file, err := logging.SetupEnhancedLogger(cfg)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	logger.Info("Loaded configuration successfully")

	_, err = database.SetupDatabaseConnection(&cfg.Database)
	if err != nil {
		logger.Error("Error establishing connection with the database", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("Established connection with the database successfully")

}
