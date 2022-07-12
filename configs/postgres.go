package configs

import (
	"fmt"
	"os"
)

// Config holds the configuration used for instantiating a connection to database
type Config struct {
	Host     string `env:"HOST"`
	Port     string `env:"DBPORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Database string `env:"NAME"`
}

// Dialect returns "postgres"
func (c Config) Dialect() string {
	return os.Getenv("DIALECT")
}

// GetConnectionURI returns a Postgres URI for connection
func (c Config) GetConnectionURI() string {
	if c.Password == "" {
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Database,
		)
	} else {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Password, c.Database)
	}
}

// GetPostgresConfig returns a struct config from the .env file
func GetPostgresConfig() Config {
	return Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Database: os.Getenv("NAME"),
	}
}
