package repository

import (
	"database/sql"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
)

type DuelistRepository interface {
	CreateDuelist(duelist entities.Duelist) (*entities.Duelist, error)
	ListDuelist(pagination *Pagination) ([]entities.Duelist, error)
	FindDuelist(id string) (*entities.Duelist, error)
	UpdateDuelist(duelist entities.Duelist) error
	DeleteDuelist(id string) error
}

func NewDuelistRepository(db *sql.DB) DuelistRepository {
	return repository{db: db}
}

type repository struct {
	db *sql.DB
}
