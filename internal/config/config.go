package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port              string `envconfig:"PORT" default: ":4000"`
	DatabaseName      string `envconfig:"DATABASE_NAME" default:"calendar.db"`
	SessionCookieName string `envconfig:"SESSION_COOKIE_NAME" default:"sessions"`
}

func loadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}