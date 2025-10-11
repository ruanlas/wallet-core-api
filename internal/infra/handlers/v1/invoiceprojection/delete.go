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

type deleteHandler struct {
	deleteUseCase invoiceprojection.DeleteUseCase
}

func NewDeleteHandler(deleteUseCase invoiceprojection.DeleteUseCase) handlers.Handler {
	return &deleteHandler{deleteUseCase: deleteUseCase}
}

// @Summary Remove uma Despesa Prevista
// @Description Este endpoint permite remover uma despesa prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da despesa prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ResponseDefault{status=int,message=string}
// @Router /v1/invoice-projection/{id} [delete]
func (h *deleteHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	id := c.Param("id")
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	span := tx.StartSpan("InvoiceProjection::StorageProcess::Delete", "Delete a invoice-projection", nil)
	searchCtx := invoiceprojection.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	err := h.deleteUseCase.Execute(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Invoice projection removed"})
}
