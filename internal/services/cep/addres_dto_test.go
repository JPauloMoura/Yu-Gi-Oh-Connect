package cep

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddress_Validate(t *testing.T) {
	tests := []struct {
		name    string
		addr    Address
		wantErr bool
	}{
		{
			name:    "should return error when all fields are empty",
			addr:    Address{},
			wantErr: true,
		},
		{
			name: "should return an error when the street field has more than 50 characters",
			addr: Address{
				Street:         strings.Repeat("a", 51),
				District:       "district",
				City:           "city",
				StateShortname: "PI",
				Zipcode:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the district field has more than 50 characters",
			addr: Address{
				Street:         strings.Repeat("a", 50),
				District:       strings.Repeat("a", 51),
				City:           strings.Repeat("a", 50),
				StateShortname: strings.Repeat("a", 2),
				Zipcode:        strings.Repeat("a", 50),
			},
			wantErr: true,
		},
		{
			name: "should return an error when the city field is empty",
			addr: Address{
				Street:         strings.Repeat("a", 50),
				District:       strings.Repeat("a", 51),
				City:           "",
				StateShortname: strings.Repeat("a", 2),
				Zipcode:        strings.Repeat("a", 50),
			},
			wantErr: true,
		},
		{
			name: "should return an error when the city field has more than 50 characters",
			addr: Address{
				Street:         strings.Repeat("a", 50),
				District:       strings.Repeat("a", 50),
				City:           strings.Repeat("a", 51),
				StateShortname: strings.Repeat("a", 2),
				Zipcode:        strings.Repeat("a", 50),
			},
			wantErr: true,
		},
		{
			name: "should return an error when the state shot name field is empty",
			addr: Address{
				Street:         strings.Repeat("a", 50),
				District:       strings.Repeat("a", 50),
				City:           strings.Repeat("a", 50),
				StateShortname: "",
				Zipcode:        strings.Repeat("a", 50),
			},
			wantErr: true,
		},
		{
			name: "should return an error when the state shot name field has more than 2 characters",
			addr: Address{
				Street:         strings.Repeat("a", 50),
				District:       strings.Repeat("a", 50),
				City:           strings.Repeat("a", 50),
				StateShortname: strings.Repeat("a", 3),
				Zipcode:        strings.Repeat("a", 50),
			},
			wantErr: true,
		},
		{
			name: "should return an error when the zipcode field is empty",
			addr: Address{
				Street:         strings.Repeat("a", 50),
				District:       strings.Repeat("a", 50),
				City:           strings.Repeat("a", 50),
				StateShortname: strings.Repeat("a", 2),
				Zipcode:        "",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the zipcode field has more than 50 characters",
			addr: Address{
				Street:         strings.Repeat("a", 50),
				District:       strings.Repeat("a", 50),
				City:           strings.Repeat("a", 50),
				StateShortname: strings.Repeat("a", 2),
				Zipcode:        strings.Repeat("a", 51),
			},
			wantErr: true,
		},
		{
			name: "no errors should be returned",
			addr: Address{
				Street:         strings.Repeat("a", 50),
				District:       strings.Repeat("a", 50),
				City:           strings.Repeat("a", 50),
				StateShortname: strings.Repeat("a", 2),
				Zipcode:        strings.Repeat("a", 50),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr := Address{
				Street:         tt.addr.Street,
				District:       tt.addr.District,
				City:           tt.addr.City,
				StateShortname: tt.addr.StateShortname,
				Zipcode:        tt.addr.Zipcode,
			}

			assert.Equal(t, tt.wantErr, addr.Validate() != nil)

		})
	}
}
