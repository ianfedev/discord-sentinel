package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/hcl"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"log"
	"os"
)

// ParseConfig gets the configuration from the .hcl file
func ParseConfig() (*Config, error) {

	k := koanf.New(".")

	if err := k.Load(file.Provider("config.hcl"), hcl.Parser(true)); err != nil {
		return nil, err
	}

	// Load environment variables
	if err := k.Load(env.Provider("", ".", func(s string) string {
		return s
	}), nil); err != nil {
		return nil, err
	}

	// Load .env file variables
	if err := k.Load(env.Provider(".env", ".", func(s string) string {
		return s
	}), nil); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	// Unmarshal the configuration into the Config struct.
	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	return &cfg, nil

}
