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
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
}

type ResponseDefault interface {
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

// Create godoc
// @Summary Editar uma Receita Prevista
// @Description Este endpoint permite editar uma receita prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param gain_projection body service.UpdateRequest true "Modelo de edição da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Param   X-Userinfo	header	string	true	"Informações do usuário em base64"
// @Success 200 {object} service.GainProjectionResponse
// @Router /v1/gain-projection/{id} [put]
func (h *handler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	id := c.Param("id")
	var request service.UpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::StorageProcess::Update", "Create new gain-projection", nil)
	gainUpdated, err := h.storageProcess.Update(c.Request.Context(), id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if gainUpdated == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Gain projection not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, gainUpdated)
}

// @Summary Remove uma Receita Prevista
// @Description Este endpoint permite remover uma receita prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da receita prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Param   X-Userinfo	header	string	true	"Informações do usuário em base64"
// @Success 200 {object} ResponseDefault{status=int,message=string}
// @Router /v1/gain-projection/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	id := c.Param("id")
	span := tx.StartSpan("GainProjection::StorageProcess::Delete", "Delete a gain-projection", nil)
	err := h.storageProcess.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Gain projection removed"})
}

// @Summary Obter uma Receita Prevista
// @Description Este endpoint permite obter uma receita prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param page_size query string false "O número de registros retornados pela busca"
// @Param page query string false "A página que será buscada"
// @Param month query string true "O mês que será filtrado a busca"
// @Param year query string true "O ano que será filtrado a busca"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Param   X-Userinfo	header	string	true	"Informações do usuário em base64"
// @Success 200 {object} service.GainProjectionPaginateResponse
// @Router /v1/gain-projection [get]
func (h *handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	searchParams, err := validateAndGetSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::StorageProcess::GetAllPaginated", "Get a gain-projection paginated", nil)
	resultPaginated, err := h.readingProcess.GetAllPaginated(ctx, *searchParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusOK, resultPaginated)
}
