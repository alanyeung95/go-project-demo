package config

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

// AppConfig container
type AppConfig struct {
	API     API
	MongoDB MongoDB
	Sentry  Sentry
	v       *viper.Viper
	GRPC    GRPC
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
	v.BindEnv("mongodb.itemCollection", "MONGODB_ITEM_COLLECTION")
	v.BindEnv("mongodb.userCollection", "MONGODB_USER_COLLECTION")

	v.BindEnv("sentry.dsn", "SENTRY_DSN")

	v.BindEnv("GRPC.address", "GRPC_ADDRESS")

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
