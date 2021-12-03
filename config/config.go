package config

import (
	"os"
	"path/filepath"

	errors "github.com/FernandoGal25/academy-go-q42021/error"
	"github.com/spf13/viper"
)

type Config struct {
	Rest struct {
		Api string
	}
	CSV struct {
		Path string
	}
	Server struct {
		Port string
	}
}

// Extracts the config variables.
func ReadConfig() (*Config, error) {
	var c Config

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	path, err := os.Getwd()

	if err != nil {
		return nil, errors.ErrSystemConfig{Message: "Could not find directory path", Err: err}
	}

	viper.AddConfigPath(filepath.Join(path, "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.ErrSystemConfig{Message: "Could not load configuration", Err: err}
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, errors.ErrSystemConfig{Message: "Failed to parse configuration", Err: err}
	}

	return &c, nil
}
