package configs

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
			"host=%s dbport=%s user=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Database,
		)
	} else {
		return fmt.Sprintf("host=%s dbport=%s user=%s password=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.Password, c.User, c.Database)
	}
}

// GetPostgresConfig returns a config type for further use
func GetPostgresConfig() Config {
	return Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Database: os.Getenv("NAME"),
	}
}
