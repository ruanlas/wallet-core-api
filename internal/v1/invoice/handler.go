package invoice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/tracing"
	"github.com/ruanlas/wallet-core-api/internal/v1/invoice/iservice"
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
	storageProcess iservice.StorageProcess
	readingProcess iservice.ReadingProcess
}

func NewHandler(storageProcess iservice.StorageProcess, readingProcess iservice.ReadingProcess) Handler {
	return &handler{storageProcess: storageProcess, readingProcess: readingProcess}
}

// Create godoc
// @Summary Criar uma Despesa
// @Description Este endpoint permite criar uma despesa
// @Tags Invoice
// @Accept json
// @Produce json
// @Param invoice body iservice.CreateRequest true "Modelo de criação da despesa"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 201 {object} iservice.InvoiceResponse
// @Router /v1/invoice [post]
func (h *handler) Create(c *gin.Context) {
	var request iservice.CreateRequest
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	// apm.TransactionFromContext(ctx)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Invoice::StorageProcess::Create", "Create new invoice", nil)
	createCtx := iservice.CreateContext{
		Ctx:       ctx,
		UserToken: userToken,
		Request:   request,
	}
	invoiceCreated, err := h.storageProcess.Create(createCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	span.End()
	c.JSON(http.StatusCreated, invoiceCreated)
}

// @Summary Obter uma Despesa
// @Description Este endpoint permite obter uma despesa
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Id da despesa"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} iservice.InvoiceResponse
// @Router /v1/invoice/{id} [get]
func (h *handler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")

	span := tx.StartSpan("Invoice::ReadingProcess::GetById", "Get a invoice by id", nil)

	searchCtx := iservice.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	invoice, err := h.readingProcess.GetById(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if invoice == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Object not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, invoice)
}

// Create godoc
// @Summary Editar uma Despesa
// @Description Este endpoint permite editar uma despesa
// @Tags Invoice
// @Accept json
// @Produce json
// @Param invoice body iservice.UpdateRequest true "Modelo de edição da despesa"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} iservice.InvoiceResponse
// @Router /v1/invoice/{id} [put]
func (h *handler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request iservice.UpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Invoice::StorageProcess::Update", "Create new invoice", nil)
	updateCtx := iservice.UpdateContext{
		Ctx:       ctx,
		Request:   request,
		Id:        id,
		UserToken: userToken,
	}
	invoiceUpdated, err := h.storageProcess.Update(updateCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if invoiceUpdated == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Invoice projection not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, invoiceUpdated)
}

// @Summary Remove uma Despesa
// @Description Este endpoint permite remover uma despesa
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Id da despesa"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ResponseDefault{status=int,message=string}
// @Router /v1/invoice/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	id := c.Param("id")
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	span := tx.StartSpan("Invoice::StorageProcess::Delete", "Delete a invoice", nil)
	searchCtx := iservice.SearchContext{
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
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Invoice projection removed"})
}

// @Summary Obter uma listagem de Despesas
// @Description Este endpoint permite obter uma listagem de despesas
// @Tags Invoice
// @Accept json
// @Produce json
// @Param page_size query string false "O número de registros retornados pela busca"
// @Param page query string false "A página que será buscada"
// @Param month query string true "O mês que será filtrado a busca"
// @Param year query string true "O ano que será filtrado a busca"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} iservice.InvoicePaginateResponse
// @Router /v1/invoice [get]
func (h *handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	searchParams, err := validateAndGetSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("Invoice::ReadingProcess::GetAllPaginated", "Get a invoice paginated", nil)
	searchCtx := iservice.SearchContext{
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
