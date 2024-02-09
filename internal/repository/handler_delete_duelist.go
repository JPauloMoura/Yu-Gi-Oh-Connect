package repository

import (
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
)

func (r repository) DeleteDuelist(id string) error {
	_, err := r.FindDuelist(id)
	if err != nil {
		slog.Error("failed to check if the duelist", slog.Any("error", err), slog.String("id", id))
		return errors.ErrorDuelistNotFound
	}

	query, err := r.db.Prepare(`DELETE FROM duelists WHERE id=$1`)
	if err != nil {
		slog.Error("failed to prepare query to delete duelist", slog.Any("error", err))
		return errors.ErrorQueryToDeleteDuelistIsInvalid
	}

	if _, err = query.Exec(id); err != nil {
		slog.Error("failed to delete duelist", slog.Any("error", err))
		return errors.ErrorUnableToDeleteDuelist
	}

	return nil
}
