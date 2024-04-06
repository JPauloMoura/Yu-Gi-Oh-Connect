package inmemory

import (
	"log/slog"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	repo "github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/repository"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
)

func init() {
	inmemoryDuelistsDb = make(map[string]entities.Duelist)
}

var inmemoryDuelistsDb map[string]entities.Duelist

func NewDuelistRepository() repo.DuelistRepository {
	return repository{}
}

type repository struct {
}

func (r repository) CreateDuelist(duelist entities.Duelist) (*entities.Duelist, error) {
	inmemoryDuelistsDb[duelist.Id] = duelist

	return &duelist, nil
}

func (r repository) DeleteDuelist(id string) error {
	_, exist := inmemoryDuelistsDb[id]
	if !exist {
		slog.Error("failed to check if the duelist", slog.Any("error", errors.ErrorDuelistNotFound), slog.String("id", id))
		return errors.ErrorDuelistNotFound
	}

	delete(inmemoryDuelistsDb, id)

	return nil
}

func (r repository) FindDuelist(id string) (*entities.Duelist, error) {
	d, exist := inmemoryDuelistsDb[id]
	if !exist {
		return nil, errors.ErrorDuelistNotFound
	}

	return &d, nil
}

func (r repository) ListDuelist(pagination *repo.Pagination) ([]entities.Duelist, error) {
	if len(inmemoryDuelistsDb) <= pagination.Limit {
		list := make([]entities.Duelist, 0)
		for _, d := range inmemoryDuelistsDb {
			list = append(list, d)
		}
		return list, nil
	}

	list := make([]entities.Duelist, 0)
	var count int
	for _, d := range inmemoryDuelistsDb {
		count++
		list = append(list, d)

		if count == pagination.Limit {
			break
		}
	}

	return list, nil
}
func (r repository) UpdateDuelist(duelist entities.Duelist) error {
	d, exist := inmemoryDuelistsDb[duelist.Id]
	if !exist {
		return errors.ErrorDuelistNotFound
	}

	if duelist.Name != "" {
		d.Name = duelist.Name
	}

	if duelist.Presentation != "" {
		d.Presentation = duelist.Presentation
	}

	if duelist.Address.Cep != "" {
		d.Address.Cep = duelist.Address.Cep
		d.Address.City = duelist.Address.City
		d.Address.District = duelist.Address.District
		d.Address.State = duelist.Address.State
		d.Address.Street = duelist.Address.Street
	}

	if !duelist.BirthDate.IsZero() {
		d.BirthDate = duelist.BirthDate
	}

	if duelist.Contact.Email != "" {
		d.Contact.Email = duelist.Contact.Email
	}
	if duelist.Contact.Phone != "" {
		d.Contact.Phone = duelist.Contact.Phone
	}

	inmemoryDuelistsDb[duelist.Id] = d

	return nil
}
