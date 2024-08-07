package main

import (
	"discord-sentinel/core/config"
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

	enhancedLogger, file, err := logging.SetupEnhancedLogger(cfg)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	enhancedLogger.Info("Loaded configuration successfully")

}
