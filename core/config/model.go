package config

// Discord struct defines the configuration for the Discord bot
// Token is the bot's authentication token
type Discord struct {
	Token string `koanf:"token" env_recommended:"true"`
}

// Log struct defines the configuration for logging
// File is the path to the log file
// Color enables or disables colored log output
type Log struct {
	File  string `koanf:"file"`
	Color bool   `koanf:"color"`
}

// Config struct defines the final configuration struct to be unmarshalled.
type Config struct {
	Discord
	Log
}
