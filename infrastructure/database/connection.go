package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/configs"
	_ "github.com/lib/pq"
)

func ConnectDb(cfg *configs.Config) *sql.DB {
	connectionString := getConnectStr(cfg)

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

func getConnectStr(cfg *configs.Config) string {
	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", cfg.DbUser(), cfg.DbName(), cfg.DbPassword(), cfg.DbHost(), cfg.DbSSlMode())
}
