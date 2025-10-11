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

type getByIdHandler struct {
	getByIdUseCase gainprojection.GetByIdUseCase
}

func NewGetByIdHandler(getByIdUseCase gainprojection.GetByIdUseCase) handlers.Handler {
	return &getByIdHandler{getByIdUseCase: getByIdUseCase}
}

// @Summary Obter uma Receita Prevista
// @Description Este endpoint permite obter uma receita prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da receita prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} gpservice.GainProjectionResponse
// @Router /v1/gain-projection/{id} [get]
func (h *getByIdHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")

	span := tx.StartSpan("GainProjection::ReadingProcess::GetById", "Get a gain-projection by id", nil)

	searchCtx := gainprojection.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	gainProjection, err := h.getByIdUseCase.Execute(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if gainProjection == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Object not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, gainProjection)
}
