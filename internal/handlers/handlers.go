package handlers

import "github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/services/duelist"

type HandlerDuelist struct {
	svcDuelist duelist.DuelistService
}

func NewHandlerDuelist(svc duelist.DuelistService) HandlerDuelist {
	return HandlerDuelist{
		svcDuelist: svc,
	}
}
