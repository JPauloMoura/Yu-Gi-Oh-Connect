package entities

import (
	"time"

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
	State    string `json:"state"`
	City     string `json:"city"`
	Street   string `json:"street"`
	District string `json:"district"`
	Cep      string `json:"cep"`
}

type Contact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}
