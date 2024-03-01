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

// @Summary FindDuelist
// @Description Busca as informações de um duelista com base no seu uid.
// @Tags Duelist
// @Accept json
// @Produce json
// @Param uid path string true "6457d5dc-6a4b-409f-972e-f8bb8f9f9f67"
// @Success 200 {object} entities.Duelist
// @Router /duelist/{uid} [GET]
func (h HandlerDuelist) FindDuelist(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("failed to parse id", slog.Any("error", err))
		response.Encode(w, errors.ErrorInvalidId, http.StatusBadRequest)
		return
	}

	duelist, err := h.svcDuelist.FindDuelist(uid.String())
	if e.Is(err, errors.ErrorDuelistNotFound) {
		slog.Warn("failed to get duelist", slog.Any("error", err))
		response.Encode(w, err, http.StatusNotFound)
		return
	}

	if err != nil {
		slog.Error("failed to get duelist", slog.Any("error", err))
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, duelist, http.StatusOK)
}
