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

type createGainHandler struct {
	createGainUseCase gainprojection.CreateGainUseCase
}

func NewCreateGainHandler(createGainUseCase gainprojection.CreateGainUseCase) handlers.Handler {
	return &createGainHandler{createGainUseCase: createGainUseCase}
}

// Create godoc
// @Summary Realizar uma Receita Prevista
// @Description Este endpoint permite realizar uma receita que foi prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da receita prevista"
// @Param gain body gpservice.CreateGainRequest true "Modelo de criação da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} gpservice.GainResponse
// @Router /v1/gain-projection/{id}/create-gain [post]
func (h *createGainHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request gainprojection.CreateGainRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::StorageProcess::CreateGain", "Create a gain from gain-projection", nil)
	createGainCtx := gainprojection.CreateGainContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	stat, err := h.createGainUseCase.Execute(createGainCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if !stat.ProjectionIsFound {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Gain-projection not found"})
		return
	}
	if stat.ProjectionIsAlreadyDone {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "A gain is already created"})
		return
	}
	span.End()
	c.JSON(http.StatusCreated, stat.Gain)
}
