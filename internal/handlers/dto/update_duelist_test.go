package dto

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdateDuelistDTO_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		jsonData   string
		wantFields UpdateDuelistDTO
		wantErr    bool
	}{

		{
			name:       "invalid json",
			jsonData:   `"name": "João", "presentation": "Apresentação do João"`,
			wantFields: UpdateDuelistDTO{},
			wantErr:    true,
		},
		{
			name:       "invalid birthDate",
			jsonData:   `{"name": "João", "presentation": "Apresentação do João", "birthDate": "31/02/2000"}`,
			wantFields: UpdateDuelistDTO{},
			wantErr:    true,
		},
		{
			name: "json valid",
			jsonData: `{
				"name": "João",
				"presentation": "Presenter",
				"birthDate": "01/01/2000",
				"cep": "12345678",
				"email": "joao@example.com",
				"phone": "1234567890"
			}`,
			wantFields: UpdateDuelistDTO{
				Name:         "João",
				Presentation: "Presenter",
				BirthDate: func() time.Time {
					date, _ := time.Parse("02/01/2006", "01/01/2000")
					return date
				}(),
				Cep:   "12345678",
				Email: "joao@example.com",
				Phone: "1234567890",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dto := UpdateDuelistDTO{}
			err := json.Unmarshal([]byte(tt.jsonData), &dto)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantFields, dto)
		})
	}
}

func TestUpdateDuelistDTO_Validate(t *testing.T) {
	tests := []struct {
		name    string
		dto     UpdateDuelistDTO
		wantErr bool
	}{
		{
			name:    "no should return an error when all fields is empty",
			dto:     UpdateDuelistDTO{},
			wantErr: false,
		},
		{
			name: "should return an error when the name has more than 50 characters",
			dto: UpdateDuelistDTO{
				Name:         strings.Repeat("a", 51),
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        "test@example.com",
				Phone:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the presentation has more than 200 characters",
			dto: UpdateDuelistDTO{
				Name:         "João",
				Presentation: strings.Repeat("a", 201),
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        "test@example.com",
				Phone:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the ZIP code has more than 8 characters",
			dto: UpdateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          strings.Repeat("a", 9),
				Email:        "test@example.com",
				Phone:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the email is not valid",
			dto: UpdateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        "testexample.com",
				Phone:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the email has more than 50 characters",
			dto: UpdateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        strings.Repeat("a", 51),
				Phone:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return error when phone has less than 10 characters",
			dto: UpdateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        "test@example.com",
				Phone:        strings.Repeat("a", 9),
			},
			wantErr: true,
		},
		{
			name: "should return an error when the phone has more than 11 characters",
			dto: UpdateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        "test@example.com",
				Phone:        strings.Repeat("a", 12),
			},
			wantErr: true,
		},
		{
			name: "No errors should be returned",
			dto: UpdateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        "test@example.com",
				Phone:        "1234567890",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dto.Validate()
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
