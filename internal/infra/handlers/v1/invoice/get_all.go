package invoice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/handlers"
	"github.com/ruanlas/wallet-core-api/internal/tracing"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/invoice"
	"go.elastic.co/apm"
)

type getAllHandler struct {
	getAllUseCase invoice.GetAllUseCase
}

func NewGetAllHandler(getAllUseCase invoice.GetAllUseCase) handlers.Handler {
	return &getAllHandler{getAllUseCase: getAllUseCase}
}

// @Summary Obter uma listagem de Despesas
// @Description Este endpoint permite obter uma listagem de despesas
// @Tags Invoice
// @Accept json
// @Produce json
// @Param page_size query string false "O número de registros retornados pela busca"
// @Param page query string false "A página que será buscada"
// @Param month query string true "O mês que será filtrado a busca"
// @Param year query string true "O ano que será filtrado a busca"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} iservice.InvoicePaginateResponse
// @Router /v1/invoice [get]
func (h *getAllHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	searchParams, err := validateAndGetSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Invoice::ReadingProcess::GetAllPaginated", "Get a invoice paginated", nil)
	searchCtx := invoice.SearchContext{
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
