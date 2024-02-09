package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/infrastructure/database"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/handlers"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/repository"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/cep"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/duelist"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/configs"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/loggers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic("Failed to load the .env file")
	}

	cfg := configs.BuildConfig()
	loggers.ConfigLogger(cfg)

	router := setupHandlers(cfg)

	startServer(cfg, router)
}

// setupHandlers configures the handlers for the application
func setupHandlers(cfg *configs.Config) http.Handler {
	router := chi.NewRouter()

	// middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// dependencies
	db := database.ConnectDb(cfg)
	duelistRepository := repository.NewDuelistRepository(db)
	duelistService := duelist.NewDuelistService(duelistRepository)
	cepService := cep.NewCepServive(http.DefaultClient)
	duelistHandler := handlers.NewHandlerDuelist(duelistService, cepService)

	// routes
	router.Post("/duelist", duelistHandler.CreateDuelist)
	router.Put("/duelist/{id}", duelistHandler.UpdateDuelist)
	router.Get("/duelist", duelistHandler.ListDuelist)
	router.Get("/duelist/{id}", duelistHandler.FindDuelist)
	router.Delete("/duelist/{id}", duelistHandler.DeleteDuelist)

	return router
}

// startServer starts the HTTP server
func startServer(cfg *configs.Config, router http.Handler) {
	serverAddr := ":" + cfg.ServerPort()
	slog.Info("The server is running on port " + cfg.ServerPort())

	err := http.ListenAndServe(serverAddr, router)
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
