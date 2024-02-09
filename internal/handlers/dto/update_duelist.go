package dto

import (
	"encoding/json"
	"time"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/cep"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UpdateDuelistDTO struct {
	Name         string    `json:"name"`
	Presentation string    `json:"presentation"`
	BirthDate    time.Time `json:"birthDate"`
	Cep          string    `json:"cep"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`

	Address *cep.AddressDTO `json:"-"`
}

func (c UpdateDuelistDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.RuneLength(0, 50)),
		validation.Field(&c.Presentation, validation.RuneLength(0, 200)),
		validation.Field(&c.BirthDate, validation.By(isDateBeforeNow)),
		validation.Field(&c.Cep, validation.RuneLength(8, 8)),
		validation.Field(&c.Email, is.Email, validation.RuneLength(0, 50)),
		validation.Field(&c.Phone, validation.RuneLength(10, 11)),
	)
}

func (c *UpdateDuelistDTO) UnmarshalJSON(data []byte) error {
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

	if wrapper.BirthDate != "" {
		birthDate, err := time.Parse("02/01/2006", wrapper.BirthDate)
		if err != nil {
			return errors.ErrorInvalidDateFormat
		}

		c.BirthDate = birthDate
	}

	c.Name = wrapper.Name
	c.Presentation = wrapper.Presentation
	c.Cep = wrapper.Cep
	c.Email = wrapper.Email
	c.Phone = wrapper.Phone

	return nil
}
