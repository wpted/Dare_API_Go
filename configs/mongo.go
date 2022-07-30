package configs

import (
	"fmt"
	"os"
)

// GetMongoURI returns the Mongo connection URI
func (c *MongoConfig) GetMongoURI() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@localhost:%s",
		c.User,
		c.Password,
		c.Port)
}

// NewMongoConfig returns a new MongoConfig object
func NewMongoConfig() MongoConfig {
	return MongoConfig{
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_DATABASE"),
		os.Getenv("MONGO_COLLECTION"),
	}
}
