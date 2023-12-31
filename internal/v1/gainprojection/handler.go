package gainprojection

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
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
	CreateGain(c *gin.Context)
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
// @Param gain_projection body service.CreateRequest true "Modelo de criação da receita prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 201 {object} service.GainProjectionResponse
// @Router /v1/gain-projection [post]
func (h *handler) Create(c *gin.Context) {
	var request service.CreateRequest
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
	createCtx := service.CreateContext{
		Ctx:       ctx,
		UserToken: userToken,
		Request:   request,
	}
	gainCreated, err := h.storageProcess.Create(createCtx)
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
// @Success 200 {object} service.GainProjectionResponse
// @Router /v1/gain-projection/{id} [get]
func (h *handler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")

	span := tx.StartSpan("GainProjection::ReadingProcess::GetById", "Get a gain-projection by id", nil)

	searchCtx := service.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	gainProjection, err := h.readingProcess.GetById(searchCtx)
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
// @Param gain_projection body service.UpdateRequest true "Modelo de edição da receita prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} service.GainProjectionResponse
// @Router /v1/gain-projection/{id} [put]
func (h *handler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request service.UpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::StorageProcess::Update", "Create new gain-projection", nil)
	updateCtx := service.UpdateContext{
		Ctx:       ctx,
		Request:   request,
		Id:        id,
		UserToken: userToken,
	}
	gainUpdated, err := h.storageProcess.Update(updateCtx)
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
// @Success 200 {object} ResponseDefault{status=int,message=string}
// @Router /v1/gain-projection/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	id := c.Param("id")
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	span := tx.StartSpan("GainProjection::StorageProcess::Delete", "Delete a gain-projection", nil)
	searchCtx := service.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	err := h.storageProcess.Delete(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Gain projection removed"})
}

// @Summary Obter uma listagem de Receitas Previstas
// @Description Este endpoint permite obter uma listagem de receitas previstas
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param page_size query string false "O número de registros retornados pela busca"
// @Param page query string false "A página que será buscada"
// @Param month query string true "O mês que será filtrado a busca"
// @Param year query string true "O ano que será filtrado a busca"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} service.GainProjectionPaginateResponse
// @Router /v1/gain-projection [get]
func (h *handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	searchParams, err := validateAndGetSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::ReadingProcess::GetAllPaginated", "Get a gain-projection paginated", nil)
	searchCtx := service.SearchContext{
		UserToken: userToken,
		Params:    *searchParams,
		Ctx:       ctx,
	}
	resultPaginated, err := h.readingProcess.GetAllPaginated(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusOK, resultPaginated)
}

// Create godoc
// @Summary Realizar uma Receita Prevista
// @Description Este endpoint permite realizar uma receita que foi prevista
// @Tags Gain-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da receita prevista"
// @Param gain body service.CreateGainRequest true "Modelo de criação da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} service.GainResponse
// @Router /v1/gain-projection/{id}/create-gain [post]
func (h *handler) CreateGain(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request service.CreateGainRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("GainProjection::StorageProcess::CreateGain", "Create a gain from gain-projection", nil)
	createGainCtx := service.CreateGainContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	stat, err := h.storageProcess.CreateGain(createGainCtx)
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
