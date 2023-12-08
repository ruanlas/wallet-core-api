package gainprojection

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/service"
	"github.com/stretchr/testify/assert"
)

type storageProcessMock struct {
	err      error
	response *service.GainProjectionResponse
}

func (sp *storageProcessMock) Create(ctx context.Context, request service.CreateRequest) (*service.GainProjectionResponse, error) {
	if sp.err != nil {
		return nil, sp.err
	}
	return sp.response, nil
}

func (sp *storageProcessMock) Update(ctx context.Context, id string, request service.UpdateRequest) (*service.GainProjectionResponse, error) {
	if sp.err != nil {
		return nil, sp.err
	}
	return sp.response, nil
}

func (sp *storageProcessMock) Delete(ctx context.Context, id string) error {
	if sp.err != nil {
		return sp.err
	}
	return nil
}

type readingProcessMock struct {
	err      error
	response *service.GainProjectionResponse
}

func (rp *readingProcessMock) GetById(ctx context.Context, gainProjectionId string) (*service.GainProjectionResponse, error) {
	if rp.err != nil {
		return nil, rp.err
	}
	return rp.response, nil
}

func TestCreateSuccess(t *testing.T) {
	_storageProcessMock := &storageProcessMock{
		response: &service.GainProjectionResponse{},
	}

	handler := NewHandler(_storageProcessMock, nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.POST("/gain-projection", handler.Create)

	body := []byte(`
	{
		"pay_in": "2023-12-30T00:00:00+00:00",
		"description": "Teste",
		"value": 500,
		"is_passive": false,
		"recurrence": 1,
		"category_id": 2
	}`)
	req, _ := http.NewRequest("POST", "/v1/gain-projection", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	bodyExpected := `{"id":"","pay_in":"0001-01-01T00:00:00Z","description":"","value":0,"is_passive":false,"category":{"id":0,"category":""}}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateInvalidBody(t *testing.T) {
	_storageProcessMock := &storageProcessMock{
		response: &service.GainProjectionResponse{},
	}

	handler := NewHandler(_storageProcessMock, nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.POST("/gain-projection", handler.Create)

	body := []byte(`
	{
		"pay_in": "2023-12-30T00:00:00+00:00",
		"description": "Teste",
		"value": 500,
		"is_passive": false,
		"recurrence": 1,
		"category_id": 2
	broken`)
	req, _ := http.NewRequest("POST", "/v1/gain-projection", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	bodyExpected := `{"message":"invalid character 'b' after object key:value pair","status":400}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateError(t *testing.T) {
	_storageProcessMock := &storageProcessMock{
		err: errors.New("An error has been ocurred"),
	}

	handler := NewHandler(_storageProcessMock, nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.POST("/gain-projection", handler.Create)

	body := []byte(`
	{
		"pay_in": "2023-12-30T00:00:00+00:00",
		"description": "Teste",
		"value": 500,
		"is_passive": false,
		"recurrence": 1,
		"category_id": 2
	}`)
	req, _ := http.NewRequest("POST", "/v1/gain-projection", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	bodyExpected := `{"message":"An error has been ocurred","status":500}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetByIdSuccess(t *testing.T) {
	_readingProcessMock := &readingProcessMock{
		response: &service.GainProjectionResponse{Id: "9b15034f-85fe-4476-82b1-a95f438aadd5"},
	}

	handler := NewHandler(nil, _readingProcessMock)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.GET("/gain-projection/:id", handler.GetById)

	req, _ := http.NewRequest("GET", "/v1/gain-projection/9b15034f-85fe-4476-82b1-a95f438aadd5", nil)

	router.ServeHTTP(w, req)
	bodyExpected := `{"id":"9b15034f-85fe-4476-82b1-a95f438aadd5","pay_in":"0001-01-01T00:00:00Z","description":"","value":0,"is_passive":false,"category":{"id":0,"category":""}}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetByIdNotFound(t *testing.T) {
	_readingProcessMock := &readingProcessMock{}

	handler := NewHandler(nil, _readingProcessMock)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.GET("/gain-projection/:id", handler.GetById)

	req, _ := http.NewRequest("GET", "/v1/gain-projection/9b15034f-85fe-4476-82b1-a95f438aadd5", nil)

	router.ServeHTTP(w, req)
	bodyExpected := `{"message":"Object not found","status":404}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetByIdError(t *testing.T) {
	_readingProcessMock := &readingProcessMock{
		err: errors.New("An error has been ocurred"),
	}

	handler := NewHandler(nil, _readingProcessMock)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.GET("/gain-projection/:id", handler.GetById)

	req, _ := http.NewRequest("GET", "/v1/gain-projection/9b15034f-85fe-4476-82b1-a95f438aadd5", nil)

	router.ServeHTTP(w, req)
	bodyExpected := `{"message":"An error has been ocurred","status":500}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateSuccess(t *testing.T) {
	_storageProcess := &storageProcessMock{
		response: &service.GainProjectionResponse{Id: "9b15034f-85fe-4476-82b1-a95f438aadd5"},
	}

	handler := NewHandler(_storageProcess, nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.PUT("/gain-projection/:id", handler.Update)

	body := []byte(`
	{
		"pay_in": "2023-12-30T00:00:00+00:00",
		"description": "Teste",
		"value": 500,
		"is_passive": false,
		"category_id": 2
	}`)
	req, _ := http.NewRequest("PUT", "/v1/gain-projection/9b15034f-85fe-4476-82b1-a95f438aadd5", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	bodyExpected := `{"id":"9b15034f-85fe-4476-82b1-a95f438aadd5","pay_in":"0001-01-01T00:00:00Z","description":"","value":0,"is_passive":false,"category":{"id":0,"category":""}}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateInvalidBody(t *testing.T) {
	_storageProcess := &storageProcessMock{}

	handler := NewHandler(_storageProcess, nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.PUT("/gain-projection/:id", handler.Update)

	body := []byte(`
	{
		"pay_in": "2023-12-30T00:00:00+00:00",
		"description": "Teste",
		"value": 500,
		"is_passive": false,
		"category_id": 2
	broken`)
	req, _ := http.NewRequest("PUT", "/v1/gain-projection/9b15034f-85fe-4476-82b1-a95f438aadd5", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	bodyExpected := `{"message":"invalid character 'b' after object key:value pair","status":400}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateNotFound(t *testing.T) {
	_storageProcess := &storageProcessMock{
		response: nil,
	}

	handler := NewHandler(_storageProcess, nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.PUT("/gain-projection/:id", handler.Update)

	body := []byte(`
	{
		"pay_in": "2023-12-30T00:00:00+00:00",
		"description": "Teste",
		"value": 500,
		"is_passive": false,
		"category_id": 2
	}`)
	req, _ := http.NewRequest("PUT", "/v1/gain-projection/9b15034f-85fe-4476-82b1-a95f438aadd5", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	bodyExpected := `{"message":"Gain projection not found","status":404}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateFail(t *testing.T) {
	_storageProcess := &storageProcessMock{
		err: errors.New("An error has been ocurred"),
	}

	handler := NewHandler(_storageProcess, nil)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.PUT("/gain-projection/:id", handler.Update)

	body := []byte(`
	{
		"pay_in": "2023-12-30T00:00:00+00:00",
		"description": "Teste",
		"value": 500,
		"is_passive": false,
		"category_id": 2
	}`)
	req, _ := http.NewRequest("PUT", "/v1/gain-projection/9b15034f-85fe-4476-82b1-a95f438aadd5", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	bodyExpected := `{"message":"An error has been ocurred","status":500}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
