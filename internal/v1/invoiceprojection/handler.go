package invoiceprojection

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/tracing"
	"github.com/ruanlas/wallet-core-api/internal/v1/invoiceprojection/ipservice"
	"go.elastic.co/apm"
)

type Handler interface {
	Create(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
	CreateInvoice(c *gin.Context)
}

type ResponseDefault interface {
}

type handler struct {
	storageProcess ipservice.StorageProcess
	readingProcess ipservice.ReadingProcess
}

func NewHandler(storageProcess ipservice.StorageProcess, readingProcess ipservice.ReadingProcess) Handler {
	return &handler{storageProcess: storageProcess, readingProcess: readingProcess}
}

// Create godoc
// @Summary Criar uma Despesa Prevista
// @Description Este endpoint permite criar uma despesa prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param invoice_projection body ipservice.CreateRequest true "Modelo de criação da despesa prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 201 {object} ipservice.InvoiceProjectionResponse
// @Router /v1/invoice-projection [post]
func (h *handler) Create(c *gin.Context) {
	var request ipservice.CreateRequest
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	// apm.TransactionFromContext(ctx)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("InvoiceProjection::StorageProcess::Create", "Create new invoice-projection", nil)
	createCtx := ipservice.CreateContext{
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

// @Summary Obter uma Despesa Prevista
// @Description Este endpoint permite obter uma despesa prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da despesa prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ipservice.InvoiceProjectionResponse
// @Router /v1/invoice-projection/{id} [get]
func (h *handler) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")

	span := tx.StartSpan("InvoiceProjection::ReadingProcess::GetById", "Get a invoice-projection by id", nil)

	searchCtx := ipservice.SearchContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	invoiceProjection, err := h.readingProcess.GetById(searchCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if invoiceProjection == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Object not found"})
		return
	}
	span.End()
	c.JSON(http.StatusOK, invoiceProjection)
}

// Create godoc
// @Summary Editar uma Despesa Prevista
// @Description Este endpoint permite editar uma despesa prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param invoice_projection body ipservice.UpdateRequest true "Modelo de edição da despesa prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ipservice.InvoiceProjectionResponse
// @Router /v1/invoice-projection/{id} [put]
func (h *handler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request ipservice.UpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("InvoiceProjection::StorageProcess::Update", "Create new invoice-projection", nil)
	updateCtx := ipservice.UpdateContext{
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

// @Summary Remove uma Despesa Prevista
// @Description Este endpoint permite remover uma despesa prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da despesa prevista"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ResponseDefault{status=int,message=string}
// @Router /v1/invoice-projection/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	id := c.Param("id")
	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	span := tx.StartSpan("InvoiceProjection::StorageProcess::Delete", "Delete a invoice-projection", nil)
	searchCtx := ipservice.SearchContext{
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

// @Summary Obter uma listagem de Despesas Previstas
// @Description Este endpoint permite obter uma listagem de despesas previstas
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param page_size query string false "O número de registros retornados pela busca"
// @Param page query string false "A página que será buscada"
// @Param month query string true "O mês que será filtrado a busca"
// @Param year query string true "O ano que será filtrado a busca"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ipservice.InvoiceProjectionPaginateResponse
// @Router /v1/invoice-projection [get]
func (h *handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	searchParams, err := validateAndGetSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("InvoiceProjection::ReadingProcess::GetAllPaginated", "Get a invoice-projection paginated", nil)
	searchCtx := ipservice.SearchContext{
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
// @Summary Realizar uma Despesa Prevista
// @Description Este endpoint permite realizar uma despesa que foi prevista
// @Tags Invoice-Projection
// @Accept json
// @Produce json
// @Param id path string true "Id da despesa prevista"
// @Param invoice body ipservice.CreateInvoiceRequest true "Modelo de criação da despesa"
// @Param   X-Access-Token	header	string	true	"Token de autenticação do usuário"
// @Success 200 {object} ipservice.InvoiceResponse
// @Router /v1/invoice-projection/{id}/create-invoice [post]
func (h *handler) CreateInvoice(c *gin.Context) {
	ctx := c.Request.Context()
	tx := apm.TransactionFromContext(ctx)

	userToken := c.GetHeader(idpauth.AUTH_HEADER)
	id := c.Param("id")
	var request ipservice.CreateInvoiceRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	span := tx.StartSpan("InvoiceProjection::StorageProcess::CreateInvoice", "Create a invoice from invoice-projection", nil)
	createInvoiceCtx := ipservice.CreateInvoiceContext{
		Ctx:       ctx,
		Id:        id,
		UserToken: userToken,
	}
	stat, err := h.storageProcess.CreateInvoice(createInvoiceCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		tracing.SendSpanErr(span, err)
		return
	}
	if !stat.ProjectionIsFound {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Invoice-projection not found"})
		return
	}
	if stat.ProjectionIsAlreadyDone {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "A invoice is already created"})
		return
	}
	span.End()
	c.JSON(http.StatusCreated, stat.Invoice)
}
