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

type getByIdHandler struct {
	getByIdUseCase invoiceprojection.GetByIdUseCase
}

func NewGetByIdHandler(getByIdUseCase invoiceprojection.GetByIdUseCase) handlers.Handler {
	return &getByIdHandler{getByIdUseCase: getByIdUseCase}
}

// @Summary Obter uma Despesa Prevista
// @Description Este endpoint permite obter uma despesa prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da despesa prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ipservice.InvoiceProjectionResponse
// @Router /v1/invoice-projection/{id} [get]
func (h *getByIdHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")

	span := tx.StartSpan("InvoiceProjection::ReadingProcess::GetById", "Get a invoice-projection by id", nil)

	searchCtx := invoiceprojection.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	invoiceProjection, err := h.getByIdUseCase.Execute(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if invoiceProjection == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Object not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, invoiceProjection)
}
