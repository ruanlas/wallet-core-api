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

type updateHandler struct {
	updateUseCase gainprojection.UpdateUseCase
}

func NewUpdateHandler(updateUseCase gainprojection.UpdateUseCase) handlers.Handler {
	return &updateHandler{updateUseCase: updateUseCase}
}

// Create godoc
// @Summary Editar uma Receita Prevista
// @Description Este endpoint permite editar uma receita prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param gain_projection body gpservice.UpdateRequest true "Modelo de edição da receita prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} gpservice.GainProjectionResponse
// @Router /v1/gain-projection/{id} [put]
func (h *updateHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request gainprojection.UpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::StorageProcess::Update", "Create new gain-projection", nil)
	updateCtx := gainprojection.UpdateContext{
		Ctx:       ctx,
		Request:   request,
		Id:        id,
		UserToken: userToken,
	}
	gainUpdated, err := h.updateUseCase.Execute(updateCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if gainUpdated == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Gain projection not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, gainUpdated)
}
