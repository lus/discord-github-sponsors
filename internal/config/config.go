package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel string `default:"info" split_words:"true"`
	BotToken string `required:"true" split_words:"true"`
}

func Load() (*Config, error) {
	_ = godotenv.Overload()
	cfg := new(Config)
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
