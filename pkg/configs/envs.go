package configs

import (
	"log"
	"os"
)

// BuildConfig imports the necessary environment variables and makes them available in a config structure
func BuildConfig() *Config {
	cfg := &Config{
		dbUser:     os.Getenv("DB_USER"),
		dbName:     os.Getenv("DB_NAME"),
		dbPassword: os.Getenv("DB_PASSWORD"),
		dbHost:     os.Getenv("DB_HOST"),
		dbSSlMode:  os.Getenv("DB_SSLMODE"),
		serverPort: os.Getenv("SERVER_PORT"),
		logType:    os.Getenv("LOG_TYPE"),
	}

	cfg.validate()

	return cfg
}

// Config contains the application variables
type Config struct {
	dbUser     string
	dbName     string
	dbPassword string
	dbHost     string
	dbSSlMode  string
	serverPort string
	logType    string
}

func (c *Config) validate() {
	if c.dbUser == "" {
		log.Fatal("DB_USER is required")
	}
	if c.dbName == "" {
		log.Fatal("DB_NAME is required")
	}
	if c.dbPassword == "" {
		log.Fatal("DB_PASSWORD is required")
	}
	if c.dbHost == "" {
		log.Fatal("DB_HOST is required")
	}
	if c.dbSSlMode == "" {
		log.Fatal("DB_SSLMODE is required")
	}
	if c.serverPort == "" {
		log.Fatal("SERVER_PORT is required")
	}
}

func (c *Config) DbUser() string {
	return c.dbUser
}

func (c *Config) DbName() string {
	return c.dbName
}

func (c *Config) DbPassword() string {
	return c.dbPassword
}

func (c *Config) DbHost() string {
	return c.dbHost
}

func (c *Config) DbSSlMode() string {
	return c.dbSSlMode
}

func (c *Config) ServerPort() string {
	return c.serverPort
}

func (c *Config) LogType() string {
	return c.logType
}
