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

type createHandler struct {
	createUseCase gainprojection.CreateUseCase
}

func NewCreateHandler(createUseCase gainprojection.CreateUseCase) handlers.Handler {
	return &createHandler{createUseCase: createUseCase}
}

// Create godoc
// @Summary Criar uma Receita Prevista
// @Description Este endpoint permite criar uma receita prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param gain_projection body gpservice.CreateRequest true "Modelo de criação da receita prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 201 {object} gpservice.GainProjectionResponse
// @Router /v1/gain-projection [post]
func (h *createHandler) Handle(c *gin.Context) {
	var request gainprojection.CreateRequest
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	// apm.TransactionFromContext(ctx)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::StorageProcess::Create", "Create new gain-projection", nil)
	createCtx := gainprojection.CreateContext{
		Ctx:       ctx,
		UserToken: userToken,
		Request:   request,
	}
	gainCreated, err := h.createUseCase.Execute(createCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusCreated, gainCreated)
}
