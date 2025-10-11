package gain

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/handlers"
	"github.com/ruanlas/wallet-core-api/internal/tracing"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/gain"
	"go.elastic.co/apm"
)

type updateHandler struct {
	updateUseCase gain.UpdateUseCase
}

func NewUpdateHandler(updateUseCase gain.UpdateUseCase) handlers.Handler {
	return &updateHandler{updateUseCase: updateUseCase}
}

// Create godoc
// @Summary Editar uma Receita
// @Description Este endpoint permite editar uma receita
// @Tags Gain
// @Accept json
// @Produce json
// @Param gain body gservice.UpdateRequest true "Modelo de edição da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} gservice.GainResponse
// @Router /v1/gain/{id} [put]
func (h *updateHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request gain.UpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Gain::StorageProcess::Update", "Create new gain", nil)
	updateCtx := gain.UpdateContext{
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
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Gain not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, gainUpdated)
}
