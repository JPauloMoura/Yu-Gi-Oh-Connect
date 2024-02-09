package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/handlers/dto"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/response"
)

func (h HandlerDuelist) CreateDuelist(w http.ResponseWriter, r *http.Request) {
	var requestBody dto.CreateDuelistDTO

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

	var err error

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

	duelist := createEntityDuelistByDuelistDTO(requestBody)
	duelistCreated, err := h.svcDuelist.CreateDuelist(duelist)
	if err != nil {
		slog.Error("failed to create duelist", slog.String("error", err.Error()), slog.Any("duelist", duelist))
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, duelistCreated, http.StatusCreated)
}

func createEntityDuelistByDuelistDTO(dto dto.CreateDuelistDTO) entities.Duelist {
	duelist := entities.NewDuelist()

	duelist.Name = dto.Name
	duelist.Presentation = dto.Presentation
	duelist.BirthDate = dto.BirthDate
	duelist.Contact.Email = dto.Email
	duelist.Contact.Phone = dto.Phone
	duelist.Address.Cep = dto.Address.Zipcode
	duelist.Address.State = dto.Address.StateShortname
	duelist.Address.City = dto.Address.City
	duelist.Address.District = dto.Address.District
	duelist.Address.Street = dto.Address.Street

	return duelist
}
