package gain

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/tracing"
	"github.com/ruanlas/wallet-core-api/internal/v1/gain/gservice"
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
	storageProcess gservice.StorageProcess
	readingProcess gservice.ReadingProcess
}

func NewHandler(storageProcess gservice.StorageProcess, readingProcess gservice.ReadingProcess) Handler {
	return &handler{storageProcess: storageProcess, readingProcess: readingProcess}
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
func (h *handler) Create(c *gin.Context) {
	var request gservice.CreateRequest
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
	createCtx := gservice.CreateContext{
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

// @Summary Obter uma Receita
// @Description Este endpoint permite obter uma receita
// @Tags Gain
// @Accept json
// @Produce json
// @Param id path string true "Id da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} gservice.GainResponse
// @Router /v1/gain/{id} [get]
func (h *handler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")

	span := tx.StartSpan("Gain::ReadingProcess::GetById", "Get a gain by id", nil)

	searchCtx := gservice.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	Gain, err := h.readingProcess.GetById(searchCtx)
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

// Create godoc
// @Summary Editar uma Receita
// @Description Este endpoint permite editar uma receita
// @Tags Gain
// @Accept json
// @Produce json
// @Param gain body gservice.UpdateRequest true "Modelo de edição da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} gservice.GainResponse
// @Router /v1/gain/{id} [put]
func (h *handler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request gservice.UpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Gain::StorageProcess::Update", "Create new gain", nil)
	updateCtx := gservice.UpdateContext{
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
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Gain not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, gainUpdated)
}

// @Summary Remove uma Receita
// @Description Este endpoint permite remover uma receita
// @Tags Gain
// @Accept json
// @Produce json
// @Param id path string true "Id da receita"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ResponseDefault{status=int,message=string}
// @Router /v1/gain/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	id := c.Param("id")
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	span := tx.StartSpan("Gain::StorageProcess::Delete", "Delete a gain", nil)
	searchCtx := gservice.SearchContext{
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
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Gain removed"})
}

// @Summary Obter uma listagem de Receitas
// @Description Este endpoint permite obter uma listagem de receitas
// @Tags Gain
// @Accept json
// @Produce json
// @Param page_size query string false "O número de registros retornados pela busca"
// @Param page query string false "A página que será buscada"
// @Param month query string true "O mês que será filtrado a busca"
// @Param year query string true "O ano que será filtrado a busca"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} gservice.GainPaginateResponse
// @Router /v1/gain [get]
func (h *handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	searchParams, err := validateAndGetSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Gain::ReadingProcess::GetAllPaginated", "Get a gain paginated", nil)
	searchCtx := gservice.SearchContext{
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
