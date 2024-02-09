package duelist

import (
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/repository"
)

func (p duelistService) ListDuelist(pagination *repository.Pagination) ([]entities.Duelist, error) {
	list, err := p.database.ListDuelist(pagination)
	if err != nil {
		slog.Error("failed to list duelist", slog.Any("error", err))
		return nil, err
	}

	return list, nil
}
