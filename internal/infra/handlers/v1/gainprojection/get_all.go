package gainprojection

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/handlers"
	"github.com/ruanlas/wallet-core-api/internal/tracing"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/gainprojection"
	"go.elastic.co/apm"
)

type getAllHandler struct {
	getAllUseCase gainprojection.GetAllUseCase
}

func NewGetAllHandler(getAllUseCase gainprojection.GetAllUseCase) handlers.Handler {
	return &getAllHandler{getAllUseCase: getAllUseCase}
}

// @Summary Obter uma listagem de Receitas Previstas
// @Description Este endpoint permite obter uma listagem de receitas previstas
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param page_size query string false "O número de registros retornados pela busca"
// @Param page query string false "A página que será buscada"
// @Param month query string true "O mês que será filtrado a busca"
// @Param year query string true "O ano que será filtrado a busca"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} gpservice.GainProjectionPaginateResponse
// @Router /v1/gain-projection [get]
func (h *getAllHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	searchParams, err := validateAndGetSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::ReadingProcess::GetAllPaginated", "Get a gain-projection paginated", nil)
	searchCtx := gainprojection.SearchContext{
		UserToken: userToken,
		Params:    *searchParams,
		Ctx:       ctx,
	}
	resultPaginated, err := h.getAllUseCase.Execute(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusOK, resultPaginated)
}
