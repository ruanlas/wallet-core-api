package gainprojection

import "github.com/gin-gonic/gin"

type Handler interface {
	Create(c *gin.Context)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Create(c *gin.Context) {

}
