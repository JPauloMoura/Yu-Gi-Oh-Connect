package repository

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
)

func (r repository) UpdateDuelist(duelist entities.Duelist) error {
	q, fields := generateQueryToUpdateDuelist(duelist)
	query, err := r.db.Prepare(q)

	if err != nil {
		slog.Error("failed to prepare query to update duelist", slog.Any("error", err), slog.Any("duelist", duelist))
		return errors.ErrorQueryToUpdateDuelistIsInvalid
	}

	_, err = query.Exec(fields...)
	if err != nil {
		slog.Error("failed to update duelist", slog.Any("error", err), slog.Any("duelist", duelist))
		return errors.ErrorUnableToUpdateDuelist
	}

	return nil
}

func generateQueryToUpdateDuelist(duelist entities.Duelist) (string, []interface{}) {
	query := "UPDATE duelists SET"
	var values []interface{}
	counter := 1

	if duelist.Name != "" {
		query += " name=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.Name)
		counter++
	}
	if duelist.Presentation != "" {
		query += " presentation=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.Presentation)
		counter++
	}
	if !duelist.BirthDate.IsZero() {
		query += " birthDate=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.BirthDate)
		counter++
	}
	if duelist.Address.State != "" {
		query += " state=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.Address.State)
		counter++
	}
	if duelist.Address.City != "" {
		query += " city=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.Address.City)
		counter++
	}
	if duelist.Address.Street != "" {
		query += " street=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.Address.Street)
		counter++
	}
	if duelist.Address.District != "" {
		query += " district=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.Address.District)
		counter++
	}
	if duelist.Address.Cep != "" {
		query += " cep=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.Address.Cep)
		counter++
	}
	if duelist.Contact.Email != "" {
		query += " email=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.Contact.Email)
		counter++
	}
	if duelist.Contact.Phone != "" {
		query += " phone=$" + strconv.Itoa(counter) + ","
		values = append(values, duelist.Contact.Phone)
		counter++
	}

	query = strings.TrimSuffix(query, ",")

	query += " WHERE id=$" + strconv.Itoa(counter)
	values = append(values, duelist.Id)

	return query, values
}
