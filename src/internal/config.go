package config

type Constants struct {
	PORT string `default:"8000"`
}

type User struct {
	Username string `json:"username"`
}

type database struct {
	Users []User
}

type Config struct {
	Constants
	DB *database
}

func New() (*Config, error) {
	config := Config{}

	config.DB = &database{
		Users: make([]User, 0),
	}

	config.Constants = Constants{"8000"}

	return &config, nil
}
