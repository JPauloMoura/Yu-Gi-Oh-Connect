package dto

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_isDateBeforeNow(t *testing.T) {
	type args struct {
		date any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return an error when the date entered is greater than the current date",
			args: args{
				date: time.Now().Add(time.Hour * 24),
			},
			wantErr: true,
		},
		{
			name: "should return an error when the date entered is equal to the current date",
			args: args{
				date: time.Now(),
			},
			wantErr: true,
		},
		{
			name: "must accept dates prior to the current date",
			args: args{
				date: time.Now().Add(-time.Hour * 24),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := isDateBeforeNow(tt.args.date)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCreateDuelistDTO_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		jsonData   string
		wantFields CreateDuelistDTO
		wantErr    bool
	}{

		{
			name:       "invalid json",
			jsonData:   `"name": "João", "presentation": "Apresentação do João"`,
			wantFields: CreateDuelistDTO{},
			wantErr:    true,
		},
		{
			name:       "invalid birthDate",
			jsonData:   `{"name": "João", "presentation": "Apresentação do João", "birthDate": "31/02/2000"}`,
			wantFields: CreateDuelistDTO{},
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
			wantFields: CreateDuelistDTO{
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
			dto := CreateDuelistDTO{}
			err := json.Unmarshal([]byte(tt.jsonData), &dto)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantFields, dto)
		})
	}
}

func TestCreateDuelistDTO_Validate(t *testing.T) {
	tests := []struct {
		name    string
		dto     CreateDuelistDTO
		wantErr bool
	}{
		{
			name: "should return error when name is empty",
			dto: CreateDuelistDTO{
				Name:         "",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        "test@example.com",
				Phone:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the name has more than 50 characters",
			dto: CreateDuelistDTO{
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
			dto: CreateDuelistDTO{
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
			name: "should return error when date of birth is not defined",
			dto: CreateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				Cep:          "12345678",
				Email:        "test@example.com",
				Phone:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the CEP is empty",
			dto: CreateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "",
				Email:        "test@example.com",
				Phone:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the ZIP code has more than 8 characters",
			dto: CreateDuelistDTO{
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
			name: "should return an error when the email is empty",
			dto: CreateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        "",
				Phone:        "1234567890",
			},
			wantErr: true,
		},
		{
			name: "should return an error when the email is not valid",
			dto: CreateDuelistDTO{
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
			dto: CreateDuelistDTO{
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
			name: "should return error when phone is empty",
			dto: CreateDuelistDTO{
				Name:         "João",
				Presentation: "Apresentação",
				BirthDate:    time.Now().Add(-time.Hour * 48),
				Cep:          "12345678",
				Email:        "test@example.com",
				Phone:        "",
			},
			wantErr: true,
		},
		{
			name: "should return error when phone has less than 10 characters",
			dto: CreateDuelistDTO{
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
			dto: CreateDuelistDTO{
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
			dto: CreateDuelistDTO{
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
