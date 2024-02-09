package handlers

import (
	"log/slog"
	"net/http"

	e "errors"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h HandlerDuelist) FindDuelist(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("failed to parse id", err)
		response.Encode(w, errors.ErrorInvalidId, http.StatusBadRequest)
		return
	}

	duelist, err := h.svcDuelist.FindDuelist(uid.String())
	if e.Is(err, errors.ErrorDuelistNotFound) {
		slog.Warn("failed to get duelist", err)
		response.Encode(w, err, http.StatusNotFound)
		return
	}

	if err != nil {
		slog.Error("failed to get duelist", err)
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, duelist, http.StatusOK)
}
