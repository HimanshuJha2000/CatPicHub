package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	// config : this will hold all the application configuration
	config AppConfig
)

type AppConfig struct {
	Application Application `toml:"application"`
	Database    Database    `toml:"database"`
}

// LoadConfig will load the configuration values available in the config directory
func LoadConfig(env string) {
	log.Println("Loading environment for " + env)
	config.Application.Environment = env

	filePath := fmt.Sprintf("./config/env.%s.toml", env)
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		fmt.Println("Error loading TOML file:", err)
		return
	}
}

// GetConfig : will give the struct as value so that the actual config doesn't get tampered
func GetConfig() AppConfig {
	return config
}
