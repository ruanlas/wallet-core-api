package invoiceprojection

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/handlers"
	"github.com/ruanlas/wallet-core-api/internal/tracing"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/invoiceprojection"
	"go.elastic.co/apm"
)

type createHandler struct {
	createUseCase invoiceprojection.CreateUseCase
}

func NewCreateHandler(createUseCase invoiceprojection.CreateUseCase) handlers.Handler {
	return &createHandler{createUseCase: createUseCase}
}

// Create godoc
// @Summary Criar uma Despesa Prevista
// @Description Este endpoint permite criar uma despesa prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param invoice_projection body ipservice.CreateRequest true "Modelo de criação da despesa prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 201 {object} ipservice.InvoiceProjectionResponse
// @Router /v1/invoice-projection [post]
func (h *createHandler) Handle(c *gin.Context) {
	var request invoiceprojection.CreateRequest
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	// apm.TransactionFromContext(ctx)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("InvoiceProjection::StorageProcess::Create", "Create new invoice-projection", nil)
	createCtx := invoiceprojection.CreateContext{
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
