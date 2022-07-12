package configs

import (
	"fmt"
	"os"
)

// Config holds the configuration used for instantiating a connection to database
// backticks are struct tags for further info and usage
type Config struct {
	Dialect  string `env:"DIALECT"`
	Host     string `env:"HOST"`
	Port     string `env:"DBPORT"`
	User     string `env:"USER"`
	DBname   string `env:"NAME"`
	Password string `env:"PASSWORD"`
}

// Dialect returns "postgres"
func (c Config) DbDialect() string {
	return c.Dialect
}

// GetConnectionURI returns a Postgres URI for connection

func (c Config) GetConnectionURI() string {
	if c.Password == "" {
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.DBname)
	} else {
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			c.Host, c.Port, c.User, c.DBname, c.Password)
	}
}

// GetPostgresConfig returns a struct config from the .env file
func GetPostgresConfig() Config {
	return Config{
		Dialect:  os.Getenv("DIALECT"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("USER"),
		DBname:   os.Getenv("NAME"),
		Password: os.Getenv("PASSWORD"),
	}

}
