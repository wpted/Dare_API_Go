package configs

// MongoConfig object
type MongoConfig struct {
	User       string `env:"USER"`
	Password   string `env:"PASSWORD"`
	Port       string `env:"PORT"`
	Database   string `env:"DATABASE"`
	Collection string `env:"COLLECTION"`
}

// Admin object
type Admin struct {
	UserName  string `env:"USERNAME"`
	Password  string `env:"PASSWORD"`
	SecretKey string `env:"KEY"`
}
