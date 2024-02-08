package handlers

import (
	"encoding/json"
	"time"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/cep"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CreateDuelistDTO struct {
	Name         string    `json:"name"`
	Presentation string    `json:"presentation"`
	BirthDate    time.Time `json:"birthDate"`
	Cep          string    `json:"cep"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`

	address *cep.AddressDTO
}

func (c CreateDuelistDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.RuneLength(0, 50)),
		validation.Field(&c.Presentation, validation.RuneLength(0, 200)),
		validation.Field(&c.BirthDate, validation.Required, validation.By(isDateBeforeNow)),
		validation.Field(&c.Cep, validation.Required, validation.RuneLength(8, 8)),
		validation.Field(&c.Email, validation.Required, is.Email, validation.RuneLength(0, 50)),
		validation.Field(&c.Phone, validation.Required, validation.RuneLength(10, 11)),
	)
}

func (c *CreateDuelistDTO) UnmarshalJSON(data []byte) error {
	var wrapper struct {
		Name         string `json:"name"`
		Presentation string `json:"presentation"`
		BirthDate    string `json:"birthDate"`
		Cep          string `json:"cep"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
	}

	if err := json.Unmarshal(data, &wrapper); err != nil {
		return err
	}

	birthDate, err := time.Parse("02/01/2006", wrapper.BirthDate)
	if err != nil {
		return errors.ErrorInvalidDateFormat
	}

	c.Name = wrapper.Name
	c.Presentation = wrapper.Presentation
	c.BirthDate = birthDate
	c.Cep = wrapper.Cep
	c.Email = wrapper.Email
	c.Phone = wrapper.Phone

	return nil
}

func isDateBeforeNow(date any) error {
	d, implemt := date.(time.Time)
	if !implemt {
		return errors.ErrorInvalidDateFormat
	}

	if !d.Before(time.Now()) {
		return errors.ErrorDateMustBeLessThanTheCurrentDate
	}

	return nil
}
