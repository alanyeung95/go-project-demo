package config

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

// AppConfig container
type AppConfig struct {
	API     API
	MongoDB MongoDB
	v       *viper.Viper
}

// Validate the config struct
func (config AppConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.MongoDB),
	)
}

// LoadFromEnv loads env vars into config
func (config *AppConfig) LoadFromEnv() error {
	v := config.v

	v.BindEnv("api.port", "API_PORT")

	v.BindEnv("mongodb.addresses", "MONGODB_ADDRESSES")
	v.BindEnv("mongodb.username", "MONGODB_USERNAME")
	v.BindEnv("mongodb.password", "MONGODB_PASSWORD")
	v.BindEnv("mongodb.database", "MONGODB_DATABASE")

	err := v.Unmarshal(&config)

	if err != nil {
		return err
	}

	return nil
}

// NewAppConfig returns config
func NewAppConfig() AppConfig {
	return AppConfig{
		v: viper.New(),
	}
}
