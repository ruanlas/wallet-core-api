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

type createHandler struct {
	createUseCase gain.CreateUseCase
}

func NewCreateHandler(createUseCase gain.CreateUseCase) handlers.Handler {
	return &createHandler{
		createUseCase: createUseCase,
	}
}

// Create godoc
// @Summary Criar uma Receita
// @Description Este endpoint permite criar uma receita
// @Tags Gain
// @Accept json
// @Produce json
// @Param gain body gservice.CreateRequest true "Modelo de criação da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 201 {object} gservice.GainResponse
// @Router /v1/gain [post]
func (h *createHandler) Handle(c *gin.Context) {
	var request gain.CreateRequest
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	// apm.TransactionFromContext(ctx)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Gain::StorageProcess::Create", "Create new gain", nil)
	createCtx := gain.CreateContext{
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
