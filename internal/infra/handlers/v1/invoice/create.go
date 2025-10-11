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

type createHandler struct {
	createUseCase invoice.CreateUseCase
}

func NewCreateHandler(createUseCase invoice.CreateUseCase) handlers.Handler {
	return &createHandler{createUseCase: createUseCase}
}

// Create godoc
// @Summary Criar uma Despesa
// @Description Este endpoint permite criar uma despesa
// @Tags Invoice
// @Accept json
// @Produce json
// @Param invoice body iservice.CreateRequest true "Modelo de criação da despesa"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 201 {object} iservice.InvoiceResponse
// @Router /v1/invoice [post]
func (h *createHandler) Handle(c *gin.Context) {
	var request invoice.CreateRequest
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	// apm.TransactionFromContext(ctx)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Invoice::StorageProcess::Create", "Create new invoice", nil)
	createCtx := invoice.CreateContext{
		Ctx:       ctx,
		UserToken: userToken,
		Request:   request,
	}
	invoiceCreated, err := h.createUseCase.Execute(createCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusCreated, invoiceCreated)
}
