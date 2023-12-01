package gainprojection

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/tracing"
	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/service"
	"go.elastic.co/apm"
)

type Handler interface {
	Create(c *gin.Context)
	GetById(c *gin.Context)
}

type handler struct {
	storageProcess service.StorageProcess
	readingProcess service.ReadingProcess
}

func NewHandler(storageProcess service.StorageProcess, readingProcess service.ReadingProcess) Handler {
	return &handler{storageProcess: storageProcess, readingProcess: readingProcess}
}

// Create godoc
// @Summary Criar uma Receita Prevista
// @Description Este endpoint permite criar uma receita prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param gain_projection body service.CreateRequest true "Modelo de criação da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Param   X-Userinfo	header	string	true	"Informações do usuário em base64"
// @Success 201 {object} service.GainProjectionResponse
// @Router /v1/gain-projection [post]
func (h *handler) Create(c *gin.Context) {
	var request service.CreateRequest
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	// apm.TransactionFromContext(ctx)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::StorageProcess::Create", "Create new gain-projection", nil)
	gainCreated, err := h.storageProcess.Create(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusCreated, gainCreated)
}

// @Summary Obter uma Receita Prevista
// @Description Este endpoint permite obter uma receita prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da receita prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Param   X-Userinfo	header	string	true	"Informações do usuário em base64"
// @Success 200 {object} service.GainProjectionResponse
// @Router /v1/gain-projection/{id} [get]
func (h *handler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	id := c.Param("id")

	span := tx.StartSpan("GainProjection::ReadingProcess::GetById", "Get a gain-projection by id", nil)
	gainProjection, err := h.readingProcess.GetById(c.Request.Context(), id)
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
