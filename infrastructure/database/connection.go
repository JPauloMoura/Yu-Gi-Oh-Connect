package database

import (
	"database/sql"
	"log"
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/configs"
	_ "github.com/lib/pq"
)

// ConnectDb creates a new database connection
func ConnectDb(cfg *configs.Config) *sql.DB {
	connectionString := cfg.DbConnectionStr()

	connect, err := sql.Open("postgres", connectionString)
	if err != nil || connect == nil {
		slog.Error("failed to open conection", err, slog.String("connectionString", connectionString))
		log.Fatal("down service")
	}

	if err := connect.Ping(); err != nil {
		slog.Error("failed to ping on database", err, slog.String("connectionString", connectionString))
		log.Fatal("down service")
	}
	return connect
}
