package config

import "github.com/spf13/viper"

type Config struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBName         string `mapstructure:"DB_NAME"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	Server         string `mapstructure:"SERVER"`
	HTTPPort       string `mapstructure:"HTTP_PORT"`
	GRPCPort       string `mapstructure:"GRPC_PORT"`
	GatewayEnabled string `mapstructure:"GATEWAY_ENABLED"`
	GatewayPort    string `mapstructure:"GATEWAY_PORT"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "SERVER", "HTTP_PORT", "GRPC_PORT", "GATEWAY_PORT",
}

func LoadConfigFile() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
