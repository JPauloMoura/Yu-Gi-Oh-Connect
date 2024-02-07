package configs

import (
	"log"
	"os"
)

type Config struct {
	DbUser     func() string
	DbName     func() string
	DbPassword func() string
	DbHost     func() string
	DbSSlMode  func() string
	ServerPort func() string
	LogType    func() string
}

func (c Config) validate() {
	if c.DbUser() == "" {
		log.Fatal("DB_USER is required")
	}
	if c.DbName() == "" {
		log.Fatal("DB_NAME is required")
	}
	if c.DbPassword() == "" {
		log.Fatal("DB_PASSWORD is required")
	}
	if c.DbHost() == "" {
		log.Fatal("DB_HOST is required")
	}
	if c.DbSSlMode() == "" {
		log.Fatal("DB_SSLMODE is required")
	}
	if c.ServerPort() == "" {
		log.Fatal("SERVER_PORT is required")
	}
}

func BuildConfig() *Config {
	var (
		DbUser     = os.Getenv("DB_USER")
		DbName     = os.Getenv("DB_NAME")
		DbPassword = os.Getenv("DB_PASSWORD")
		DbHost     = os.Getenv("DB_HOST")
		DbSSlMode  = os.Getenv("DB_SSLMODE")
		ServerPort = os.Getenv("SERVER_PORT")
		LogType    = os.Getenv("LOG_TYPE")
	)

	cfg := Config{
		DbUser:     func() string { return DbUser },
		DbName:     func() string { return DbName },
		DbPassword: func() string { return DbPassword },
		DbHost:     func() string { return DbHost },
		DbSSlMode:  func() string { return DbSSlMode },
		ServerPort: func() string { return ServerPort },
		LogType:    func() string { return LogType },
	}

	cfg.validate()

	return &cfg
}
