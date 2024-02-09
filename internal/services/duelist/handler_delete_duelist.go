package duelist

import "log/slog"

func (p duelistService) DeleteDuelist(id string) error {
	if err := p.database.DeleteDuelist(id); err != nil {
		slog.Error("failed to delete duelist", slog.Any("error", err))
		return err
	}

	return nil
}
