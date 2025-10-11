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

type createInvoiceHandler struct {
	createInvoiceUseCase invoiceprojection.CreateInvoiceUseCase
}

func NewCreateInvoiceHandler(createInvoiceUseCase invoiceprojection.CreateInvoiceUseCase) handlers.Handler {
	return &createInvoiceHandler{createInvoiceUseCase: createInvoiceUseCase}
}

// Create godoc
// @Summary Realizar uma Despesa Prevista
// @Description Este endpoint permite realizar uma despesa que foi prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da despesa prevista"
// @Param invoice body ipservice.CreateInvoiceRequest true "Modelo de criação da despesa"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ipservice.InvoiceResponse
// @Router /v1/invoice-projection/{id}/create-invoice [post]
func (h *createInvoiceHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request invoiceprojection.CreateInvoiceRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("InvoiceProjection::StorageProcess::CreateInvoice", "Create a invoice from invoice-projection", nil)
	createInvoiceCtx := invoiceprojection.CreateInvoiceContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	stat, err := h.createInvoiceUseCase.Execute(createInvoiceCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if !stat.ProjectionIsFound {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Invoice-projection not found"})
		return
	}
	if stat.ProjectionIsAlreadyDone {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "A invoice is already created"})
		return
	}
	span.End()
	c.JSON(http.StatusCreated, stat.Invoice)
}
