package main

import "github.com/kelseyhightower/envconfig"

// Config stores the application configurations.
type Config struct {
	Port string `envconfig:"PORT" default:"8080"`
}

// ReadConfig populates configurations from environment variables.
func ReadConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
