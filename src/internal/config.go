package config

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Constants struct {
	PORT                                        string `default:"8000"`
	DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME string
}

type Config struct {
	Constants
	DB *gorm.DB
}

func New() (*Config, error) {
	config := Config{}

	config.Constants = Constants{
		PORT:    "8000",
		DB_HOST: "localhost",
		DB_PORT: "5432",
		DB_USER: "admin",
		DB_PASS: "admin",
		DB_NAME: "conduit",
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.New("could not access database")
	}

	config.DB = db

	return &config, nil
}
