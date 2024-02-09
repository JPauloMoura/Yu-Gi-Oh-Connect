package cep

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/errors"
	pkghttp "github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/http"
)

type CepService interface {
	GetAddress(cep string) (*AddressDTO, error)
}

func NewCepServive(client pkghttp.HttpClient) CepService {
	return cepService{
		client: client,
	}
}

type cepService struct {
	client pkghttp.HttpClient
}

func (c cepService) GetAddress(cep string) (*AddressDTO, error) {
	req, err := generateRequest(cep)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		slog.Error("failed to do request", slog.Any("error", err), slog.String("cep", cep))
		return nil, errors.ErrorInvalidRequest
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		slog.Error("failed to get address by cep", slog.Any("error", errors.ErrorInvalidCep), slog.String("cep", cep))
		return nil, errors.ErrorInvalidCep
	}

	if resp.StatusCode != http.StatusOK {
		slog.Error("failed to get address by cep", slog.Any("error", errors.ErrorCepServiceIsUnavailable), slog.String("cep", cep))
		return nil, errors.ErrorCepServiceIsUnavailable
	}

	var address AddressDTO

	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		slog.Error("failed to decode body address", slog.Any("error", err), slog.String("cep", cep))
		return nil, errors.ErrorInvalidResponseBody
	}

	if err := address.Validate(); err != nil {
		slog.Error("failed to validate address", slog.Any("error", err), slog.Any("address", address))
		return nil, errors.ErrorInvalidResponseBody
	}

	return &address, nil
}

func generateRequest(cep string) (*http.Request, error) {
	url := fmt.Sprintf("https://brasilaberto.com/api/v1/zipcode/%s", cep)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("failed to generate request", slog.Any("error", err))
		return nil, err
	}

	return req, nil
}
