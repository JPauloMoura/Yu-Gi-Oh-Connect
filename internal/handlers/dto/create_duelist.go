package dto

import (
	"encoding/json"
	"time"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/cep"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// CreateDuelistDTO represents the data transfer object for creating a new duelist
type CreateDuelistDTO struct {
	Name         string    `json:"name" example:"JP"`
	Presentation string    `json:"presentation" example:"Jogar com Lair of Darkness"`
	BirthDate    time.Time `json:"birthDate" example:"14/09/1992"`
	Cep          string    `json:"cep" example:"72007040"`
	Email        string    `json:"email" example:"jp@gmail.com"`
	Phone        string    `json:"phone" example:"61999876543"`

	Address *cep.AddressDTO `json:"-"` // not serialized
}

// Validate performs validation on the CreateDuelistDTO fields
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

// UnmarshalJSON customizes the unmarshalling behavior for CreateDuelistDTO
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

	if d.After(time.Now().Add(-time.Hour*24)) || d.Equal(time.Now()) {
		return errors.ErrorDateMustBeLessThanTheCurrentDate
	}
	return nil
}
