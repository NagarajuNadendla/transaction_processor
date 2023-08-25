package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	BufferSize int `json:"value"`
}

var AppConfig *Config

// LoadConfig Load config
func (c *Config) LoadConfig(env string) error {
	we, _ := os.Getwd()

	viper.SetConfigName(env + "_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(we + "/resources/")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error while loading config file", err)
		return err
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

// InitConfig init config
func (c *Config) InitConfig() error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	config := &Config{}
	if err := config.LoadConfig(env); err != nil {
		return fmt.Errorf("error while loading config")
	}
	AppConfig = config

	return nil
}
