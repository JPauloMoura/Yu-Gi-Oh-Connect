package repository

import (
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
)

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
