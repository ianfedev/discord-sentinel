package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// LoggerConfig contains the logger setup configuration.
type LoggerConfig struct {
	LogFile   *os.File
	UseColors bool
	LogLevel  zapcore.Level
}

// NewLogger configures a new logger based on LoggerConfig data.
func NewLogger(config LoggerConfig) (*zap.Logger, error) {

	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if config.UseColors {
		consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	jsonEncoderConfig := zap.NewProductionEncoderConfig()
	jsonEncoder := zapcore.NewJSONEncoder(jsonEncoderConfig)

	// Create the core based on the presence of a log file
	var cores []zapcore.Core
	cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), config.LogLevel))
	if config.LogFile != nil {
		fileCore := zapcore.NewCore(jsonEncoder, zapcore.AddSync(config.LogFile), config.LogLevel)
		cores = append(cores, fileCore)
	}

	// Create the logger
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller())

	return logger, nil
}
