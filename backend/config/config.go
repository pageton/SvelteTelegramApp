package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	BotToken string
	Port     string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	config := &Config{}

	var missingVars []string
	if config.BotToken = os.Getenv("BOT_TOKEN"); config.BotToken == "" {
		missingVars = append(missingVars, "BOT_TOKEN")
	}

	if config.Port = os.Getenv("PORT"); config.Port == "" {
		missingVars = append(missingVars, "PORT")
	}

	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return config, nil
}

func (c *Config) Validate() error {
	if c.BotToken == "" {
		return errors.New("bot token is required")
	}
	if c.Port == "" {
		return errors.New("port is required")
	}
	return nil
}
