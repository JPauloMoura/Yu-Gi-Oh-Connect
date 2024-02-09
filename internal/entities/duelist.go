package entities

import (
	"time"

	"github.com/google/uuid"
)

// NewDuelist creates a new Duelist instance with a generated UUID as the ID
func NewDuelist() Duelist {
	return Duelist{
		Id: uuid.New().String(),
	}
}

// Duelist represents a duelist entity
type Duelist struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Presentation string    `json:"presentation"`
	BirthDate    time.Time `json:"birthDate"`
	Address      Address   `json:"address"`
	Contact      Contact   `json:"contact"`
}

// Address represents the address information of a duelist
type Address struct {
	State    string `json:"state"`
	City     string `json:"city"`
	Street   string `json:"street"`
	District string `json:"district"`
	Cep      string `json:"cep"`
}

// Contact represents the contact information of a duelist.
type Contact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}
