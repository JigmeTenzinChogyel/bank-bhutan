package util

import (
	"time"

	"github.com/spf13/viper"
)

// these values are read by viper from env file
type Config struct {
	Enviroment          string        `mapstructure:"ENVIROMENT"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress   string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	MigrationURL        string        `mapstructure:"MIGRATION_URL"`
}

// load configuration from env
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
