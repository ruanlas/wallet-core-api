package gainprojection

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/service"
)

type Handler interface {
	Create(c *gin.Context)
}

type handler struct {
	storageProcess service.StorageProcess
}

func NewHandler(storageProcess service.StorageProcess) Handler {
	return &handler{storageProcess: storageProcess}
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
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	gainCreated, err := h.storageProcess.Create(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gainCreated)
}
