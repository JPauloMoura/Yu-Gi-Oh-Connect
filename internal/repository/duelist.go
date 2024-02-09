package repository

import (
	"database/sql"
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
)

type DuelistRepository interface {
	CreateDuelist(duelist entities.Duelist) (*entities.Duelist, error)
	FindDuelist(id string) (*entities.Duelist, error)
}

func NewDuelistRepository(db *sql.DB) DuelistRepository {
	return repository{db: db}
}

type repository struct {
	db *sql.DB
}

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
func (r repository) FindDuelist(id string) (*entities.Duelist, error) {
	items, err := r.db.Query(`SELECT * FROM duelists WHERE id=$1`, id)
	if err != nil {
		slog.Error("failed to find duelist by id", err, slog.String("id", id))
		return nil, errors.ErrorUnableToFindDuelists
	}

	var d entities.Duelist

	if !items.Next() {
		return nil, errors.ErrorDuelistNotFound
	}

	err = items.Scan(
		&d.Id,
		&d.Name,
		&d.Presentation,
		&d.BirthDate,
		&d.Address.State,
		&d.Address.City,
		&d.Address.Street,
		&d.Address.District,
		&d.Address.Cep,
		&d.Contact.Email,
		&d.Contact.Phone,
	)

	if err != nil {
		slog.Error("failed to scan when find duelist by id", slog.Any("error", err), slog.String("id", id))
		return nil, errors.ErrorUnableToScanDuelist
	}

	return &d, nil
}

func (r repository) UpdateDuelist(duelist entities.Duelist) error {
	q, fields := generateQueryToUpdateDuelist(duelist)
	fmt.Println(q)
	query, err := r.db.Prepare(q)

	if err != nil {
		slog.Error("failed to prepare query to update duelist", slog.Any("error", err), slog.Any("duelist", duelist))
		return errors.ErrorQueryToUpdateDuelistIsInvalid
	}

	_, err = query.Exec(fields...)
	if err != nil {
		slog.Error("failed to update duelist", slog.Any("error", err), slog.Any("duelist", duelist))
		return errors.ErrorUnableToUpdateDuelist
	}

	return nil
}
