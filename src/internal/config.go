package config

import "example/realworld-api/src/schemas"

type Constants struct {
	PORT string `default:"8000"`
}

type database struct {
	Users []schemas.User
}

type Config struct {
	Constants
	DB *database
}

func New() (*Config, error) {
	config := Config{}

	config.DB = &database{
		Users: make([]schemas.User, 0),
	}

	config.Constants = Constants{"8000"}

	return &config, nil
}
