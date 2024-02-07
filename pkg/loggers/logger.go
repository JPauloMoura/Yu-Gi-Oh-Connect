package loggers

import (
	"log/slog"
	"os"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/configs"
)

func ConfigLogger(cfg *configs.Config) {
	var opts = &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	var logger *slog.Logger

	switch cfg.LogType() {
	case "text":
		logger = slog.New(slog.NewTextHandler(os.Stdout, opts))
	case "json", "":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	}

	slog.SetDefault(logger)
}
