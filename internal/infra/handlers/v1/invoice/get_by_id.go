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

type getByIdHandler struct {
	getByIdUseCase invoice.GetByIdUseCase
}

func NewGetByIdHandler(getByIdUseCase invoice.GetByIdUseCase) handlers.Handler {
	return &getByIdHandler{getByIdUseCase: getByIdUseCase}
}

// @Summary Obter uma Despesa
// @Description Este endpoint permite obter uma despesa
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Id da despesa"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} iservice.InvoiceResponse
// @Router /v1/invoice/{id} [get]
func (h *getByIdHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")

	span := tx.StartSpan("Invoice::ReadingProcess::GetById", "Get a invoice by id", nil)

	searchCtx := invoice.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	invoice, err := h.getByIdUseCase.Execute(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if invoice == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Object not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, invoice)
}
