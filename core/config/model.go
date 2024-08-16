package config

type Environment string

const (
	Production  Environment = "production"
	Staging     Environment = "staging"
	Development Environment = "development"
)

// Discord struct defines the configuration for the Discord bot
// Token is the bot's authentication token
type Discord struct {
	Token string `env_recommended:"true"`
}

// Log struct defines the configuration for logging
// File is the path to the log file
// Color enables or disables colored log output
type Log struct {
	File  string
	Color bool
}

// Database struct defines the configuration for SQL DB provisioning
// Host address to connect
// Port number of the host
// User to authenticate
// Password to authenticate
// Database to connect
type Database struct {
	Host           string
	Port           int
	User           string `env_recommended:"true"`
	Password       string `env_recommended:"true"`
	Database       string `env_recommended:"true"`
	MaxConnections int
	MaxIdle        int
	MaxLifetime    int
}

type HTTP struct {
	Host string
	Port string
}

// Config struct defines the final configuration struct to be unmarshalled.
type Config struct {
	Environment Environment `mapstructure:"environment"`
	Discord     Discord     `mapstructure:"discord"`
	Log         Log         `mapstructure:"log"`
	Database    Database    `mapstructure:"database"`
	HTTP        HTTP        `mapstructure:"http"`
}
