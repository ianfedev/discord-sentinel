package logging

import (
	"discord-sentinel/core/config"
	"go.uber.org/zap"
	"os"
)

// SetupInitialLogger creates a temporal logger to be used
// only when Sentinel is starting.
func SetupInitialLogger() *zap.Logger {
	logCfg := &LoggerConfig{
		UseColors: false,
		LogFile:   nil,
	}

	logger, err := NewLogger(*logCfg)
	if err != nil {
		panic(err)
	}

	return logger
}

// SetupEnhancedLogger creates the unmarshalled final configuration logger.
func SetupEnhancedLogger(cfg *config.Config) (*zap.Logger, *os.File, error) {

	var file os.File
	logCfg := &LoggerConfig{
		UseColors: cfg.Log.Color,
	}

	if cfg.Log.File != "" {
		file, err := os.OpenFile(cfg.Log.File, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return nil, nil, err
		}
		logCfg.LogFile = file
	}

	logger, err := NewLogger(*logCfg)
	if err != nil {
		return nil, nil, err
	}

	return logger, &file, nil
}
