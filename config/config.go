package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Viper     *viper.Viper
	APIKey    string
	APISecret string
}

func NewConfig() *Config {
	c := &Config{
		Viper: viper.New(),
	}

	c.Viper.AddConfigPath("config")
	c.Viper.SetConfigName("config")
	c.Viper.SetConfigType("yaml")

	err := c.Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return c
}
