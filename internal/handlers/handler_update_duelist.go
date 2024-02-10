package handlers

import (
	"encoding/json"
	e "errors"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/handlers/dto"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h HandlerDuelist) UpdateDuelist(w http.ResponseWriter, r *http.Request) {
	duelistId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		slog.Error("failed to parse id", err)
		response.Encode(w, errors.ErrorInvalidId, http.StatusBadRequest)
		return
	}

	var requestBody dto.UpdateDuelistDTO

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		slog.Error("failed to decode body", slog.String("error", err.Error()))
		response.Encode(w, errors.Join(errors.ErrorInvalidRequest, err), http.StatusBadRequest)
		return
	}

	if err := requestBody.Validate(); err != nil {
		slog.Error("failed to validate duelist", slog.String("error", err.Error()))
		response.Encode(w, err, http.StatusBadRequest)
		return
	}

	if requestBody.Cep != "" {
		requestBody.Address, err = h.svcCep.GetAddress(requestBody.Cep)
		if err == errors.ErrorInvalidCep {
			slog.Error("failed get address", slog.String("error", err.Error()))
			response.Encode(w, err, http.StatusBadRequest)
			return
		}

		if err != nil {
			slog.Error("failed get address", slog.String("error", err.Error()))
			response.Encode(w, err, http.StatusInternalServerError)
			return
		}
	}

	duelist := createEntityDuelistByUpdateDuelistDTO(requestBody)
	duelist.Id = duelistId.String()

	err = h.svcDuelist.UpdateDuelist(duelist)
	if e.Is(err, errors.ErrorDuelistNotFound) {
		slog.Warn("failed to update duelist", slog.Any("error", err))
		response.Encode(w, err, http.StatusNotFound)
		return
	}

	if err != nil {
		slog.Error("failed to update duelist", slog.String("error", err.Error()), slog.Any("duelist", duelist))
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, "updated", http.StatusOK)
}

func createEntityDuelistByUpdateDuelistDTO(dto dto.UpdateDuelistDTO) entities.Duelist {
	var duelist entities.Duelist

	duelist.Name = dto.Name
	duelist.Presentation = dto.Presentation
	duelist.BirthDate = dto.BirthDate
	duelist.Contact.Email = dto.Email
	duelist.Contact.Phone = dto.Phone

	if dto.Address != nil {
		duelist.Address.Cep = dto.Address.Zipcode
		duelist.Address.State = dto.Address.StateShortname
		duelist.Address.City = dto.Address.City
		duelist.Address.District = dto.Address.District
		duelist.Address.Street = dto.Address.Street
	}

	return duelist
}
