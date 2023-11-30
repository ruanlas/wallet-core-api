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
	err error
	*service.GainProjectionResponse
}

func (sp *storageProcessMock) Create(ctx context.Context, request service.CreateRequest) (*service.GainProjectionResponse, error) {
	if sp.err != nil {
		return nil, sp.err
	}
	return sp.GainProjectionResponse, nil
}

func TestCreateSuccess(t *testing.T) {
	_storageProcessMock := &storageProcessMock{
		GainProjectionResponse: &service.GainProjectionResponse{},
	}

	handler := NewHandler(_storageProcessMock)
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
		GainProjectionResponse: &service.GainProjectionResponse{},
	}

	handler := NewHandler(_storageProcessMock)
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

	handler := NewHandler(_storageProcessMock)
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
