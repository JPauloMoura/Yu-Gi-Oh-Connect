package handlers

import (
	e "errors"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h HandlerDuelist) DeleteDuelist(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("failed to parse id", slog.Any("error", err))
		response.Encode(w, errors.ErrorInvalidId, http.StatusBadRequest)
		return
	}

	err = h.svcDuelist.DeleteDuelist(uid.String())
	if e.Is(err, errors.ErrorDuelistNotFound) {
		slog.Warn("failed to delete duelist", slog.Any("error", err))
		response.Encode(w, err, http.StatusNotFound)
		return
	}

	if err != nil {
		slog.Error("failed to get duelist", slog.Any("error", err))
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, "deleted", http.StatusOK)
}
