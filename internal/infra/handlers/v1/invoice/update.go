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

type updateHandler struct {
	updateUseCase invoice.UpdateUseCase
}

func NewUpdateHandler(updateUseCase invoice.UpdateUseCase) handlers.Handler {
	return &updateHandler{updateUseCase: updateUseCase}
}

// Create godoc
// @Summary Editar uma Despesa
// @Description Este endpoint permite editar uma despesa
// @Tags Invoice
// @Accept json
// @Produce json
// @Param invoice body iservice.UpdateRequest true "Modelo de edição da despesa"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} iservice.InvoiceResponse
// @Router /v1/invoice/{id} [put]
func (h *updateHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request invoice.UpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Invoice::StorageProcess::Update", "Create new invoice", nil)
	updateCtx := invoice.UpdateContext{
		Ctx:       ctx,
		Request:   request,
		Id:        id,
		UserToken: userToken,
	}
	invoiceUpdated, err := h.updateUseCase.Execute(updateCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if invoiceUpdated == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Invoice projection not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, invoiceUpdated)
}
