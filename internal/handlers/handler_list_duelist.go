package handlers

import (
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/repository"
	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/pkg/response"
)

// @Summary ListDuelist
// @Description Realiza a listagem de duelista. Podemos obter a lista de duelistas utilizando paginação e ordenação dos resultados.
// @Tags Duelist
// @Accept  json
// @Produce  json
// @Param sort query string false "A ordem de classificação dos Duelistas (asc ou desc) "
// @Param field query string false "O campo pelo qual os Duelistas devem ser classificados (name, birthDate)"
// @Param limit query integer false "O número máximo de Duelistas a serem retornados. O padrão é 10."
// @Param page query integer false "O número da página de resultados"
// @Success 200 {object} []entities.Duelist
// @Router /duelist [GET]
func (h HandlerDuelist) ListDuelist(w http.ResponseWriter, r *http.Request) {
	pagination, err := repository.NewPagination(r)
	if err != nil {
		slog.Error("failed to create pagination", slog.Any("error", err))
		response.Encode(w, err, http.StatusBadRequest)
		return
	}

	duelists, err := h.svcDuelist.ListDuelist(pagination)
	if err != nil {
		slog.Error("failed to list duelists", slog.Any("error", err))
		response.Encode(w, err, http.StatusInternalServerError)
		return
	}

	response.Encode(w, duelists, http.StatusOK)
}
