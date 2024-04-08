package main

import (
	"log"
	"log/slog"
	"net/http"

	_ "github.com/JPauloMoura/Yu-Gi-Oh-Connect/docs"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/infrastructure/database"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/handlers"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/repository"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/repository/inmemory"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/cep"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/duelist"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/configs"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/loggers"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Yu-Gi-Oh-Connect
// @version 1.0
// @description A API Yu-Gi-Oh! Connect permite criar, atualizar, listar, recuperar e excluir informações de duelistas. Duelistas são como são chamados os jogadores de Yu-Gi-oh TCG. E para ajudar a conectar esses jogadores novas batalhas esse projeto foi criado!
// @host localhost:3001
// @BasePath /
// @schemes http
func main() {

	cfg := configs.BuildConfig()
	loggers.ConfigLogger(cfg)

	router := setupHandlers(cfg)

	startServer(cfg, router)
}

// setupHandlers configures the handlers for the application
func setupHandlers(cfg *configs.Config) http.Handler {
	router := chi.NewRouter()

	// middlewares
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)

	// dependencies
	var duelistRepository repository.DuelistRepository
	if cfg.DbInmemory() {
		duelistRepository = inmemory.NewDuelistRepository()
	} else {
		db := database.ConnectDb(cfg)
		duelistRepository = repository.NewDuelistRepository(db)
	}

	duelistService := duelist.NewDuelistService(duelistRepository)
	cepService := cep.NewCepServive(http.DefaultClient)
	duelistHandler := handlers.NewHandlerDuelist(duelistService, cepService)

	return defineRouters(duelistHandler)
}

func defineRouters(duelistHandler handlers.HandlerDuelist) *chi.Mux {
	router := chi.NewRouter()

	// middlewares
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)
	router.Use(middleware.Cors())

	router.Get("/docs/*", httpSwagger.WrapHandler)

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
