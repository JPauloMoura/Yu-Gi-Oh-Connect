package duelist

import (
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
)

func (p duelistService) FindDuelist(id string) (*entities.Duelist, error) {
	duelist, err := p.database.FindDuelist(id)
	if err != nil {
		slog.Error("failed to get duelist", slog.Any("error", err))
		return nil, err
	}

	return duelist, nil
}
