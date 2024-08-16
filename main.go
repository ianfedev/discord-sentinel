package main

import (
	"discord-sentinel/core/config"
	"discord-sentinel/core/database"
	"discord-sentinel/core/http"
	"discord-sentinel/core/logging"
	"discord-sentinel/feature"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	logger.Info("Loaded configuration successfully", zap.String("environment", string(cfg.Environment)))

	db, err := database.SetupDatabaseConnection(&cfg.Database)
	if err != nil {
		logger.Error("Error establishing connection with the database", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("Established connection with the database successfully")

	app, err := http.SetupHTTPServer(cfg)
	if err != nil {
		logger.Error("Error setting up HTTP server", zap.Error(err))
		os.Exit(1)
	}
	feature.SetupFeatures(app, logger, db)

	// Use a wait group to wait for the server to shut down
	var wg sync.WaitGroup
	wg.Add(1)

	// Run the HTTP server in a goroutine
	go func() {
		defer wg.Done()
		logger.Info("HTTP server is now listening requests", zap.String("host", cfg.HTTP.Host), zap.String("port", cfg.HTTP.Port))
		if err := app.Listen(cfg.HTTP.Host + ":" + cfg.HTTP.Port); err != nil {
			logger.Error("Error listening on address", zap.Error(err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		logger.Info("Trying to shut down gracefully", zap.String("signal", sig.String()))

		if err := app.Shutdown(); err != nil {
			logger.Error("Error shutting down HTTP server", zap.Error(err))
		}

		if err := file.Close(); err != nil {
			logger.Error("Error closing log file", zap.Error(err))
		}

		wg.Wait()
		os.Exit(0)
	}()

	// Wait indefinitely until shutdown signal is received
	select {}
}
