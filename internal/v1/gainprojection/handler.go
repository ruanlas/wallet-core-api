package gainprojection

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(c *gin.Context)
}

type handler struct {
	storageProcess StorageProcess
}

func NewHandler(storageProcess StorageProcess) Handler {
	return &handler{storageProcess: storageProcess}
}

func (h *handler) Create(c *gin.Context) {
	var request CreateRequest
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
