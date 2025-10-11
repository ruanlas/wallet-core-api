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

type getByIdHandler struct {
	getByIdUseCase gain.GetByIdUseCase
}

func NewGetByIdHandler(getByIdUseCase gain.GetByIdUseCase) handlers.Handler {
	return &getByIdHandler{
		getByIdUseCase: getByIdUseCase,
	}
}

// @Summary Obter uma Receita
// @Description Este endpoint permite obter uma receita
// @Tags Gain
// @Accept json
// @Produce json
// @Param id path string true "Id da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} gservice.GainResponse
// @Router /v1/gain/{id} [get]
func (h *getByIdHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")

	span := tx.StartSpan("Gain::ReadingProcess::GetById", "Get a gain by id", nil)

	searchCtx := gain.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	Gain, err := h.getByIdUseCase.Execute(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if Gain == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Object not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, Gain)
}
