package cep

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type AddressDTO struct {
	Address `json:"result"`
}

type Address struct {
	Street         string `json:"street"`
	District       string `json:"district"`
	City           string `json:"city"`
	StateShortname string `json:"stateShortname"`
	Zipcode        string `json:"zipcode"`
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Street, validation.RuneLength(0, 50)),
		validation.Field(&a.District, validation.RuneLength(0, 50)),
		validation.Field(&a.City, validation.Required, validation.RuneLength(0, 50)),
		validation.Field(&a.StateShortname, validation.Required, validation.RuneLength(0, 2)),
		validation.Field(&a.Zipcode, validation.Required, validation.RuneLength(0, 50)),
	)
}
