package repository

import (
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
)

func (r repository) CreateDuelist(duelist entities.Duelist) (*entities.Duelist, error) {
	queryInsert, err := r.db.Prepare(`
		INSERT INTO duelists (id, name, presentation, birthDate, state, city, street, district, cep, email, phone) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`)

	if err != nil {
		slog.Error("failed to prepare query", slog.String("error", err.Error()))
		return nil, errors.ErrorQueryToCreateDuelistIsInvalid
	}

	_, err = queryInsert.Exec(
		duelist.Id,
		duelist.Name,
		duelist.Presentation,
		duelist.BirthDate,
		duelist.Address.State,
		duelist.Address.City,
		duelist.Address.Street,
		duelist.Address.District,
		duelist.Address.Cep,
		duelist.Contact.Email,
		duelist.Contact.Phone,
	)

	if err != nil {
		slog.Error("failed to create duelist", slog.String("error", err.Error()))
		return nil, errors.ErrorUnableToCreateDuelist
	}

	return &duelist, nil
}
