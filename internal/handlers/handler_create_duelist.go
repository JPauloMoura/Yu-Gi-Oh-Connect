package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/response"
)

func (h HandlerDuelist) CreateDuelist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Warn("invalid method",
			slog.String("aceppted", http.MethodPost),
			slog.String("received", r.Method),
		)

		response.Encode(w, errors.ErrorInvalidHttpMethod, http.StatusBadRequest)
		return
	}

	var duelist = entities.NewDuelist()

	if err := json.NewDecoder(r.Body).Decode(&duelist); err != nil {
		slog.Error("failed to decode body", slog.String("error", errors.ErrorInvalidDuelistFieldsJson.Error()))
		response.Encode(w, errors.ErrorInvalidDuelistFieldsJson, http.StatusBadRequest)
		return
	}

	if err := duelist.Validate(); err != nil {
		slog.Error("failed to validate duelist", slog.String("error", err.Error()))
		response.Encode(w, err, http.StatusBadRequest)
		return
	}

	duelistCreated, err := h.svcDuelist.CreateDuelist(duelist)
	if err != nil {
		slog.Error("failed to create duelist", slog.String("error", err.Error()))
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, duelistCreated, http.StatusCreated)
}
