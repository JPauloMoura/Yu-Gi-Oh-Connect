package entities

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

func NewDuelist() Duelist {
	return Duelist{
		Id: uuid.New().String(),
	}
}

type Duelist struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Presentation string    `json:"presentation"`
	BirthDate    time.Time `json:"birthDate"`
	Address      Address   `json:"address"`
	Contact      Contact   `json:"contact"`
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Complement string `json:"complement"`
}

type Contact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (d Duelist) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required, validation.RuneLength(0, 50)),
		validation.Field(&d.BirthDate, validation.Required),
		validation.Field(&d.Address),
		validation.Field(&d.Contact),
	)
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Street, validation.Required, validation.RuneLength(0, 50)),
		validation.Field(&a.City, validation.Required, validation.RuneLength(0, 50)),
		validation.Field(&a.State, validation.Required, validation.RuneLength(0, 50)),
		validation.Field(&a.PostalCode, validation.Required, validation.RuneLength(0, 50)),
	)
}

func (c Contact) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required, is.Email, validation.RuneLength(0, 50)),
		validation.Field(&c.Phone, validation.Required, validation.RuneLength(0, 50)),
	)
}
