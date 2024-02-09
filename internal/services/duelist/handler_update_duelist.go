package duelist

import (
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
)

func (p duelistService) UpdateDuelist(duelist entities.Duelist) error {
	if err := p.database.UpdateDuelist(duelist); err != nil {
		slog.Error("failed tp updade duelist", slog.Any("error", err))
		return err
	}

	return nil
}

type DuelistFieldsUpdated interface {
	ApplyField() string
}
