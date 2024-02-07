package duelist

import (
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
)

func (d duelistService) CreateDuelist(duelist entities.Duelist) (*entities.Duelist, error) {
	duelistCreated, err := d.database.CreateDuelist(duelist)
	if err != nil {
		slog.Error("failed to create duelist", slog.String("error", err.Error()))
		return nil, err
	}

	return duelistCreated, nil
}
