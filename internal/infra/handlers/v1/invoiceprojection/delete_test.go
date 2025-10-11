package invoiceprojection

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/invoiceprojection"
	"github.com/stretchr/testify/assert"
)

type deleteUseCaseMock struct {
	execute func(searchCtx invoiceprojection.SearchContext) error
}

func (m *deleteUseCaseMock) Execute(searchCtx invoiceprojection.SearchContext) error {
	return m.execute(searchCtx)
}

func TestDeleteSuccess(t *testing.T) {
	useCaseMock := &deleteUseCaseMock{
		execute: func(searchCtx invoiceprojection.SearchContext) error {
			return nil
		},
	}

	handler := NewDeleteHandler(useCaseMock)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.DELETE("/invoice-projection/:id", handler.Handle)

	req, _ := http.NewRequest("DELETE", "/v1/invoice-projection/9b15034f-85fe-4476-82b1-a95f438aadd5", nil)
	userToken := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRSnVoUjlFSFBIWTZFT195VjV4M1BTZWUzakRLNUs4M0lQMjJwYjFxZXFvIn0.eyJleHAiOjE3MDM4ODk3NTIsImlhdCI6MTcwMzg4OTQ1MiwianRpIjoiNTE4ZDM2MDctZjQ2NC00MDI5LTkwN2ItYjRjNzI1OWY0ZjU0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgxL3JlYWxtcy93YWxsZXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiNTgzMmE1MDItYmVkZS00OTJkLThkYzEtYjEzYjMyYzMwZjI5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2FsbGV0LWFwaSIsInNlc3Npb25fc3RhdGUiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLXdhbGxldCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RldXNlciJ9.AcRSnpgzjsuJL2n_QaRF1idkwDzwNpWNX3wiEOFXkqTG35lr4PYVYPxnhryvRvVVOvN_CUY-AaVmF_YSgR4s6JM3Oca5JFFf7T6fX5lXgj0SbQCUbbyh7Em3BemiNKr_T3wucAyO824MjGXP0smciCnnlWvq-apJDTB_R4EisDJubY_E_zpCmTfYMm0NcJ8aKB2ku8mACKgE2ZJ7WsHkKNmjaFeyU9KjGMmNKtFthYISKqRQW-6u2xPjCkpFt4_HoJ01PgjFrrJacWDlUHxVoSILcaH_Vg-WHrKppIkzgdOg5phB2zVtcakRhPhqzV4EX_jXJp2SgK4umf6ivTC3lg"
	req.Header.Add(idpauth.AUTH_HEADER, userToken)

	router.ServeHTTP(w, req)
	bodyExpected := `{"message":"Invoice projection removed","status":200}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteFail(t *testing.T) {
	useCaseMock := &deleteUseCaseMock{
		execute: func(searchCtx invoiceprojection.SearchContext) error {
			return errors.New("An error has been ocurred")
		},
	}

	handler := NewDeleteHandler(useCaseMock)
	w := httptest.NewRecorder()
	router := gin.Default()
	apiRouter := router.Group("/v1")
	apiRouter.DELETE("/invoice-projection/:id", handler.Handle)

	req, _ := http.NewRequest("DELETE", "/v1/invoice-projection/9b15034f-85fe-4476-82b1-a95f438aadd5", nil)
	userToken := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRSnVoUjlFSFBIWTZFT195VjV4M1BTZWUzakRLNUs4M0lQMjJwYjFxZXFvIn0.eyJleHAiOjE3MDM4ODk3NTIsImlhdCI6MTcwMzg4OTQ1MiwianRpIjoiNTE4ZDM2MDctZjQ2NC00MDI5LTkwN2ItYjRjNzI1OWY0ZjU0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgxL3JlYWxtcy93YWxsZXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiNTgzMmE1MDItYmVkZS00OTJkLThkYzEtYjEzYjMyYzMwZjI5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2FsbGV0LWFwaSIsInNlc3Npb25fc3RhdGUiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLXdhbGxldCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RldXNlciJ9.AcRSnpgzjsuJL2n_QaRF1idkwDzwNpWNX3wiEOFXkqTG35lr4PYVYPxnhryvRvVVOvN_CUY-AaVmF_YSgR4s6JM3Oca5JFFf7T6fX5lXgj0SbQCUbbyh7Em3BemiNKr_T3wucAyO824MjGXP0smciCnnlWvq-apJDTB_R4EisDJubY_E_zpCmTfYMm0NcJ8aKB2ku8mACKgE2ZJ7WsHkKNmjaFeyU9KjGMmNKtFthYISKqRQW-6u2xPjCkpFt4_HoJ01PgjFrrJacWDlUHxVoSILcaH_Vg-WHrKppIkzgdOg5phB2zVtcakRhPhqzV4EX_jXJp2SgK4umf6ivTC3lg"
	req.Header.Add(idpauth.AUTH_HEADER, userToken)

	router.ServeHTTP(w, req)
	bodyExpected := `{"message":"An error has been ocurred","status":500}`
	assert.Equal(t, bodyExpected, w.Body.String())
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
