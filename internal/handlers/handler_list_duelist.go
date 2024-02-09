package handlers

import (
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/repository"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/response"
)

func (h HandlerDuelist) ListDuelist(w http.ResponseWriter, r *http.Request) {
	pagination, err := repository.NewPagination(r)
	if err != nil {
		slog.Error("failed to create pagination", slog.Any("error", err))
		response.Encode(w, err, http.StatusBadRequest)
		return
	}

	duelists, err := h.svcDuelist.ListDuelist(pagination)
	if err != nil {
		slog.Error("failed to list duelists", slog.Any("error", err))
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, duelists, http.StatusOK)
}
