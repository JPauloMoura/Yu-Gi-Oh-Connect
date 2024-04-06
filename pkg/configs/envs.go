package configs

import (
	"log"
	"os"
	"strconv"
)

// BuildConfig imports the necessary environment variables and makes them available in a config structure
func BuildConfig() *Config {
	inmemory, err := strconv.ParseBool(os.Getenv("DB_INMEMORY"))
	if err != nil {
		log.Fatal("DB_INMEMORY is no bool")
	}

	cfg := &Config{
		dbInmemory:      inmemory,
		dbConnectionStr: os.Getenv("DB_CONNECTION_STRING"),
		serverPort:      os.Getenv("SERVER_PORT"),
		logType:         os.Getenv("LOG_TYPE"),
	}

	cfg.validate()

	return cfg
}

// Config contains the application variables
type Config struct {
	dbInmemory bool
	dbUser     string
	dbName     string
	dbPassword string
	dbHost     string
	dbSSlMode  string

	dbConnectionStr string
	serverPort      string
	logType         string
}

func (c *Config) validate() {
	if c.serverPort == "" {
		log.Fatal("SERVER_PORT is required")
	}
	if c.dbConnectionStr == "" {
		log.Fatal("DB_CONNECTION_STRING is required")
	}
}

func (c *Config) ServerPort() string {
	return c.serverPort
}

func (c *Config) LogType() string {
	return c.logType
}

func (c *Config) DbConnectionStr() string {
	return c.dbConnectionStr
}

func (c *Config) DbInmemory() bool {
	return c.dbInmemory
}
