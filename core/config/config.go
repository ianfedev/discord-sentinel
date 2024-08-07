package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/hcl"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"go.uber.org/zap"
	"log"
	"os"
	"reflect"
)

// checkEnvRecommended inspects the parsed configuration for any fields with the env_recommended tag and logs a warning.
func checkEnvRecommended(val reflect.Value, typ reflect.Type, logger *zap.Logger) {
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		// Check for the env_recommended tag
		tag := field.Tag.Get("env_recommended")

		if tag == "true" {

			if fieldVal.Kind() == reflect.String && fieldVal.String() == "" {
				continue
			}

			logger.Warn("Field is recommended to be set via environment variable in production",
				zap.String("key_type", typ.Name()),
				zap.String("key_id", field.Name),
			)

		}

		// Recursively check if the field is a struct
		if fieldVal.Kind() == reflect.Struct {
			checkEnvRecommended(fieldVal, fieldVal.Type(), logger)
		}
	}
}

// ParseConfig gets the configuration from the .hcl file
func ParseConfig(logger *zap.Logger) (*Config, error) {

	k := koanf.New(".")

	if err := k.Load(file.Provider("config.hcl"), hcl.Parser(true)); err != nil {
		return nil, err
	}

	// Create a temporary Config struct to inspect for env_recommended tags
	var tempCfg Config
	if err := k.Unmarshal("", &tempCfg); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	// Check for env_recommended tags
	val := reflect.ValueOf(&tempCfg).Elem()
	typ := reflect.TypeOf(tempCfg)
	checkEnvRecommended(val, typ, logger)

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
