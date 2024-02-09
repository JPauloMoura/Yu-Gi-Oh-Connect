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
		log.Panic("failed to loading .env file")
	}

	cfg := configs.BuildConfig()
	loggers.ConfigLogger(cfg)

	handlers := buildHandlers(cfg)

	slog.Info("server is running in port " + cfg.ServerPort())
	http.ListenAndServe(":"+cfg.ServerPort(), handlers)
}

func buildHandlers(cfg *configs.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	db := database.ConnectDb(cfg)
	duelistRepository := repository.NewDuelistRepository(db)
	duelistService := duelist.NewDuelistService(duelistRepository)
	cepServive := cep.NewCepServive(http.DefaultClient)

	// routers
	handler := handlers.NewHandlerDuelist(duelistService, cepServive)
	router.Post("/duelist", handler.CreateDuelist)
	router.Put("/duelist/{id}", handler.UpadateDuelist)
	router.Get("/duelist", handler.ListDuelist)
	router.Get("/duelist/{id}", handler.FindDuelist)
	router.Delete("/duelist/{id}", handler.DeleteDuelist)

	return router
}
