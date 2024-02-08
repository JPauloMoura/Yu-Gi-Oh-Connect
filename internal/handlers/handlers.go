package handlers

import (
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/cep"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/duelist"
)

type HandlerDuelist struct {
	svcDuelist duelist.DuelistService
	svcCep     cep.CepService
}

func NewHandlerDuelist(svcDuelist duelist.DuelistService, svcCep cep.CepService) HandlerDuelist {
	return HandlerDuelist{
		svcDuelist: svcDuelist,
		svcCep:     svcCep,
	}
}
