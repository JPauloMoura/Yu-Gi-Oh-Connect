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

// @Summary DeleteDuelist
// @Description Realiza a deleção de um duelista com base no seu uid.
// @Tags Duelist
// @Param uid path string true "6457d5dc-6a4b-409f-972e-f8bb8f9f9f67"
// @Produce json
// @Success 200 {object} response.Response
// @Router /duelist/{uid} [DELETE]
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
