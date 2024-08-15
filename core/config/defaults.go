package config

import (
	"github.com/spf13/viper"
)

// SetDefaultValues provides the default configuration values if not present.
func SetDefaultValues(v *viper.Viper) {

	v.SetDefault("discord.token", "")

	v.SetDefault("log.file", "")
	v.SetDefault("log.color", true)

	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.user", "")
	v.SetDefault("database.password", "")
	v.SetDefault("database.database", "psql")
	v.SetDefault("database.maxconnections", 10)
	v.SetDefault("database.maxidle", 5)
	v.SetDefault("database.maxlifetime", 30)

	v.SetDefault("http.port", "3000")
	v.SetDefault("http.address", "0.0.0.0")

}
