package config

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

	v := viper.New()

	// Set the file name of the configurations file
	v.SetConfigName("config")
	v.SetConfigType("hcl")
	v.AddConfigPath(".")

	SetDefaultValues(v)

	// Read in environment variables that match
	v.AutomaticEnv()

	// If a config file is found, read it in
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			logger.Warn("Config file not found, using default values")
		}
	}

	// Add config viper decoder to prevent errors when mapping HCL redefinition
	configOption := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		SliceOfMapsToMapHook(),
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	))

	// Unmarshal the configuration into a Config struct
	var cfg Config
	if err := v.Unmarshal(&cfg, configOption); err != nil {
		return nil, err
	}

	// Check for recommended environment variables
	checkEnvRecommended(reflect.ValueOf(cfg), reflect.TypeOf(cfg), logger)
	return &cfg, nil

}
