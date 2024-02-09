package repository

import (
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
)

func (r repository) ListDuelist(pagination *Pagination) ([]entities.Duelist, error) {
	items, err := r.db.Query(pagination.Query())
	if err != nil {
		slog.Error("failed to get duelists", slog.Any("error", err))
		return nil, errors.ErrorUnableToListDuelists
	}

	list := []entities.Duelist{}

	for items.Next() {
		var d entities.Duelist
		err := items.Scan(
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
			slog.Error("failed to scan duelists", slog.Any("error", err))
			return nil, errors.ErrorUnableToScanDuelist
		}

		list = append(list, d)
	}

	return list, nil
}
