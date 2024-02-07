package repository

import (
	"database/sql"
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
)

type DuelistRepository interface {
	CreateDuelist(user entities.Duelist) (*entities.Duelist, error)
}

func NewDuelistRepository(db *sql.DB) DuelistRepository {
	return repository{db: db}
}

type repository struct {
	db *sql.DB
}

func (r repository) CreateDuelist(duelist entities.Duelist) (*entities.Duelist, error) {
	queryInsert, err := r.db.Prepare(`
		INSERT INTO duelists (id, name, birthDate, street, city, state, postalCode, complement, email, phone, presentation) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`)

	if err != nil {
		slog.Error("failed to prepare query", slog.String("error", err.Error()))
		return nil, errors.ErrorQueryToCreateDuelistIsInvalid
	}

	_, err = queryInsert.Exec(
		duelist.Id,
		duelist.Name,
		duelist.BirthDate,
		duelist.Address.Street,
		duelist.Address.City,
		duelist.Address.State,
		duelist.Address.PostalCode,
		duelist.Address.Complement,
		duelist.Contact.Email,
		duelist.Contact.Phone,
		duelist.Presentation,
	)

	if err != nil {
		slog.Error("failed to create duelist", slog.String("error", err.Error()))
		return nil, errors.ErrorUnableToCreateDuelist
	}

	return &duelist, nil
}
