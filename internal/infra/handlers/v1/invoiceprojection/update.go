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

type updateHandler struct {
	updateUseCase invoiceprojection.UpdateUseCase
}

func NewUpdateHandler(updateUseCase invoiceprojection.UpdateUseCase) handlers.Handler {
	return &updateHandler{updateUseCase: updateUseCase}
}

// Create godoc
// @Summary Editar uma Despesa Prevista
// @Description Este endpoint permite editar uma despesa prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param invoice_projection body ipservice.UpdateRequest true "Modelo de edição da despesa prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ipservice.InvoiceProjectionResponse
// @Router /v1/invoice-projection/{id} [put]
func (h *updateHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request invoiceprojection.UpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("InvoiceProjection::StorageProcess::Update", "Create new invoice-projection", nil)
	updateCtx := invoiceprojection.UpdateContext{
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
