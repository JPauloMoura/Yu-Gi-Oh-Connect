package duelist

import (
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/repository"
)

type DuelistService interface {
	CreateDuelist(duelist entities.Duelist) (*entities.Duelist, error)
	ListDuelist(pagination *repository.Pagination) ([]entities.Duelist, error)
	FindDuelist(id string) (*entities.Duelist, error)
	UpdateDuelist(duelist entities.Duelist) error
	DeleteDuelist(id string) error
}

func NewDuelistService(repo repository.DuelistRepository) DuelistService {
	return duelistService{
		database: repo,
	}
}

type duelistService struct {
	database repository.DuelistRepository
}
