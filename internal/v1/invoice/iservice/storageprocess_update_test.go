package iservice

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ruanlas/wallet-core-api/internal/v1/invoice/repository"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSuccess(t *testing.T) {

	now := time.Now()
	invoiceProjectMock := repository.NewInvoiceBuilder().
		AddId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		AddPayAt(now).
		AddBuyAt(now).
		AddCategory(repository.InvoiceCategory{Id: 2, Category: "Alimentação"}).
		AddDescription("Description teste").
		AddPaymentType(repository.PaymentType{Id: 2, Type: "Transferência"}).
		AddValue(750.50).
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string, userId string) (*repository.Invoice, error) {
		return invoiceProjectMock, nil
	})
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string, userId string) (*repository.Invoice, error) {
		return invoiceProjectMock, nil
	})
	_mockRepository.AddEditCall(func(ctx context.Context, invoice repository.Invoice) (*repository.Invoice, error) {
		return invoiceProjectMock, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	now = time.Now()
	request := UpdateRequest{
		PayAt:         now,
		BuyAt:         now,
		Description:   "Description teste",
		Value:         750.50,
		CategoryId:    2,
		PaymentTypeId: 2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRSnVoUjlFSFBIWTZFT195VjV4M1BTZWUzakRLNUs4M0lQMjJwYjFxZXFvIn0.eyJleHAiOjE3MDM4ODk3NTIsImlhdCI6MTcwMzg4OTQ1MiwianRpIjoiNTE4ZDM2MDctZjQ2NC00MDI5LTkwN2ItYjRjNzI1OWY0ZjU0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgxL3JlYWxtcy93YWxsZXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiNTgzMmE1MDItYmVkZS00OTJkLThkYzEtYjEzYjMyYzMwZjI5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2FsbGV0LWFwaSIsInNlc3Npb25fc3RhdGUiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLXdhbGxldCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RldXNlciJ9.AcRSnpgzjsuJL2n_QaRF1idkwDzwNpWNX3wiEOFXkqTG35lr4PYVYPxnhryvRvVVOvN_CUY-AaVmF_YSgR4s6JM3Oca5JFFf7T6fX5lXgj0SbQCUbbyh7Em3BemiNKr_T3wucAyO824MjGXP0smciCnnlWvq-apJDTB_R4EisDJubY_E_zpCmTfYMm0NcJ8aKB2ku8mACKgE2ZJ7WsHkKNmjaFeyU9KjGMmNKtFthYISKqRQW-6u2xPjCkpFt4_HoJ01PgjFrrJacWDlUHxVoSILcaH_Vg-WHrKppIkzgdOg5phB2zVtcakRhPhqzV4EX_jXJp2SgK4umf6ivTC3lg"
	updateCtx := UpdateContext{
		Ctx:       ctx,
		Id:        "cd1cc27b-28a1-47dc-ac76-70e8185e159d",
		Request:   request,
		UserToken: token,
	}
	response, err := _storageProcess.Update(updateCtx)
	assert.NoError(t, err)
	assert.Equal(t, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", response.Id)
}

func TestUpdateFirstGetByIdNotFound(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string, userId string) (*repository.Invoice, error) {
		return nil, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	now := time.Now()
	request := UpdateRequest{
		PayAt:         now,
		BuyAt:         now,
		Description:   "Description teste",
		Value:         750.50,
		CategoryId:    2,
		PaymentTypeId: 2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRSnVoUjlFSFBIWTZFT195VjV4M1BTZWUzakRLNUs4M0lQMjJwYjFxZXFvIn0.eyJleHAiOjE3MDM4ODk3NTIsImlhdCI6MTcwMzg4OTQ1MiwianRpIjoiNTE4ZDM2MDctZjQ2NC00MDI5LTkwN2ItYjRjNzI1OWY0ZjU0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgxL3JlYWxtcy93YWxsZXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiNTgzMmE1MDItYmVkZS00OTJkLThkYzEtYjEzYjMyYzMwZjI5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2FsbGV0LWFwaSIsInNlc3Npb25fc3RhdGUiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLXdhbGxldCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RldXNlciJ9.AcRSnpgzjsuJL2n_QaRF1idkwDzwNpWNX3wiEOFXkqTG35lr4PYVYPxnhryvRvVVOvN_CUY-AaVmF_YSgR4s6JM3Oca5JFFf7T6fX5lXgj0SbQCUbbyh7Em3BemiNKr_T3wucAyO824MjGXP0smciCnnlWvq-apJDTB_R4EisDJubY_E_zpCmTfYMm0NcJ8aKB2ku8mACKgE2ZJ7WsHkKNmjaFeyU9KjGMmNKtFthYISKqRQW-6u2xPjCkpFt4_HoJ01PgjFrrJacWDlUHxVoSILcaH_Vg-WHrKppIkzgdOg5phB2zVtcakRhPhqzV4EX_jXJp2SgK4umf6ivTC3lg"
	updateCtx := UpdateContext{
		Ctx:       ctx,
		Id:        "cd1cc27b-28a1-47dc-ac76-70e8185e159d",
		Request:   request,
		UserToken: token,
	}
	response, err := _storageProcess.Update(updateCtx)
	assert.NoError(t, err)
	assert.Empty(t, response)
}

func TestUpdateFirstGetByIdFail(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string, userId string) (*repository.Invoice, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	now := time.Now()
	request := UpdateRequest{
		PayAt:         now,
		BuyAt:         now,
		Description:   "Description teste",
		Value:         750.50,
		CategoryId:    2,
		PaymentTypeId: 2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRSnVoUjlFSFBIWTZFT195VjV4M1BTZWUzakRLNUs4M0lQMjJwYjFxZXFvIn0.eyJleHAiOjE3MDM4ODk3NTIsImlhdCI6MTcwMzg4OTQ1MiwianRpIjoiNTE4ZDM2MDctZjQ2NC00MDI5LTkwN2ItYjRjNzI1OWY0ZjU0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgxL3JlYWxtcy93YWxsZXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiNTgzMmE1MDItYmVkZS00OTJkLThkYzEtYjEzYjMyYzMwZjI5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2FsbGV0LWFwaSIsInNlc3Npb25fc3RhdGUiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLXdhbGxldCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RldXNlciJ9.AcRSnpgzjsuJL2n_QaRF1idkwDzwNpWNX3wiEOFXkqTG35lr4PYVYPxnhryvRvVVOvN_CUY-AaVmF_YSgR4s6JM3Oca5JFFf7T6fX5lXgj0SbQCUbbyh7Em3BemiNKr_T3wucAyO824MjGXP0smciCnnlWvq-apJDTB_R4EisDJubY_E_zpCmTfYMm0NcJ8aKB2ku8mACKgE2ZJ7WsHkKNmjaFeyU9KjGMmNKtFthYISKqRQW-6u2xPjCkpFt4_HoJ01PgjFrrJacWDlUHxVoSILcaH_Vg-WHrKppIkzgdOg5phB2zVtcakRhPhqzV4EX_jXJp2SgK4umf6ivTC3lg"
	updateCtx := UpdateContext{
		Ctx:       ctx,
		Id:        "cd1cc27b-28a1-47dc-ac76-70e8185e159d",
		Request:   request,
		UserToken: token,
	}
	_, err := _storageProcess.Update(updateCtx)
	assert.Error(t, err)
}

func TestUpdateEditFail(t *testing.T) {

	now := time.Now()
	invoiceProjectMock := repository.NewInvoiceBuilder().
		AddId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		AddPayAt(now).
		AddBuyAt(now).
		AddCategory(repository.InvoiceCategory{Id: 2, Category: "Alimentação"}).
		AddDescription("Description teste").
		AddPaymentType(repository.PaymentType{Id: 2, Type: "Transferência"}).
		AddValue(750.50).
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string, userId string) (*repository.Invoice, error) {
		return invoiceProjectMock, nil
	})
	_mockRepository.AddEditCall(func(ctx context.Context, invoice repository.Invoice) (*repository.Invoice, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	now = time.Now()
	request := UpdateRequest{
		PayAt:         now,
		BuyAt:         now,
		Description:   "Description teste",
		Value:         750.50,
		CategoryId:    2,
		PaymentTypeId: 2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRSnVoUjlFSFBIWTZFT195VjV4M1BTZWUzakRLNUs4M0lQMjJwYjFxZXFvIn0.eyJleHAiOjE3MDM4ODk3NTIsImlhdCI6MTcwMzg4OTQ1MiwianRpIjoiNTE4ZDM2MDctZjQ2NC00MDI5LTkwN2ItYjRjNzI1OWY0ZjU0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgxL3JlYWxtcy93YWxsZXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiNTgzMmE1MDItYmVkZS00OTJkLThkYzEtYjEzYjMyYzMwZjI5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2FsbGV0LWFwaSIsInNlc3Npb25fc3RhdGUiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLXdhbGxldCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RldXNlciJ9.AcRSnpgzjsuJL2n_QaRF1idkwDzwNpWNX3wiEOFXkqTG35lr4PYVYPxnhryvRvVVOvN_CUY-AaVmF_YSgR4s6JM3Oca5JFFf7T6fX5lXgj0SbQCUbbyh7Em3BemiNKr_T3wucAyO824MjGXP0smciCnnlWvq-apJDTB_R4EisDJubY_E_zpCmTfYMm0NcJ8aKB2ku8mACKgE2ZJ7WsHkKNmjaFeyU9KjGMmNKtFthYISKqRQW-6u2xPjCkpFt4_HoJ01PgjFrrJacWDlUHxVoSILcaH_Vg-WHrKppIkzgdOg5phB2zVtcakRhPhqzV4EX_jXJp2SgK4umf6ivTC3lg"
	updateCtx := UpdateContext{
		Ctx:       ctx,
		Id:        "cd1cc27b-28a1-47dc-ac76-70e8185e159d",
		Request:   request,
		UserToken: token,
	}
	_, err := _storageProcess.Update(updateCtx)
	assert.Error(t, err)
}

func TestUpdateSecondGetByIdFail(t *testing.T) {

	now := time.Now()
	invoiceProjectMock := repository.NewInvoiceBuilder().
		AddId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		AddPayAt(now).
		AddBuyAt(now).
		AddCategory(repository.InvoiceCategory{Id: 2, Category: "Alimentação"}).
		AddDescription("Description teste").
		AddPaymentType(repository.PaymentType{Id: 2, Type: "Transferência"}).
		AddValue(750.50).
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string, userId string) (*repository.Invoice, error) {
		return invoiceProjectMock, nil
	})
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string, userId string) (*repository.Invoice, error) {
		return nil, errors.New("An error has been ocurred")
	})
	_mockRepository.AddEditCall(func(ctx context.Context, invoice repository.Invoice) (*repository.Invoice, error) {
		return invoiceProjectMock, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	now = time.Now()
	request := UpdateRequest{
		PayAt:         now,
		BuyAt:         now,
		Description:   "Description teste",
		Value:         750.50,
		CategoryId:    2,
		PaymentTypeId: 2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRSnVoUjlFSFBIWTZFT195VjV4M1BTZWUzakRLNUs4M0lQMjJwYjFxZXFvIn0.eyJleHAiOjE3MDM4ODk3NTIsImlhdCI6MTcwMzg4OTQ1MiwianRpIjoiNTE4ZDM2MDctZjQ2NC00MDI5LTkwN2ItYjRjNzI1OWY0ZjU0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgxL3JlYWxtcy93YWxsZXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiNTgzMmE1MDItYmVkZS00OTJkLThkYzEtYjEzYjMyYzMwZjI5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2FsbGV0LWFwaSIsInNlc3Npb25fc3RhdGUiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLXdhbGxldCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RldXNlciJ9.AcRSnpgzjsuJL2n_QaRF1idkwDzwNpWNX3wiEOFXkqTG35lr4PYVYPxnhryvRvVVOvN_CUY-AaVmF_YSgR4s6JM3Oca5JFFf7T6fX5lXgj0SbQCUbbyh7Em3BemiNKr_T3wucAyO824MjGXP0smciCnnlWvq-apJDTB_R4EisDJubY_E_zpCmTfYMm0NcJ8aKB2ku8mACKgE2ZJ7WsHkKNmjaFeyU9KjGMmNKtFthYISKqRQW-6u2xPjCkpFt4_HoJ01PgjFrrJacWDlUHxVoSILcaH_Vg-WHrKppIkzgdOg5phB2zVtcakRhPhqzV4EX_jXJp2SgK4umf6ivTC3lg"
	updateCtx := UpdateContext{
		Ctx:       ctx,
		Id:        "cd1cc27b-28a1-47dc-ac76-70e8185e159d",
		Request:   request,
		UserToken: token,
	}
	_, err := _storageProcess.Update(updateCtx)
	assert.Error(t, err)
}
