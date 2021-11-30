package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
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

/*
	Extracts the config variables.
*/
func ReadConfig() (*config, error) {
	var C config
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	path, _ := os.Getwd()
	viper.AddConfigPath(filepath.Join(path, "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return nil, err
	}

	return Config, nil
}
