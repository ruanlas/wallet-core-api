package service

import (
	"context"
	"errors"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSuccess(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddRemoveCall(func(ctx context.Context, id string, userId string) error {
		return nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRSnVoUjlFSFBIWTZFT195VjV4M1BTZWUzakRLNUs4M0lQMjJwYjFxZXFvIn0.eyJleHAiOjE3MDM4ODk3NTIsImlhdCI6MTcwMzg4OTQ1MiwianRpIjoiNTE4ZDM2MDctZjQ2NC00MDI5LTkwN2ItYjRjNzI1OWY0ZjU0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgxL3JlYWxtcy93YWxsZXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiNTgzMmE1MDItYmVkZS00OTJkLThkYzEtYjEzYjMyYzMwZjI5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2FsbGV0LWFwaSIsInNlc3Npb25fc3RhdGUiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLXdhbGxldCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RldXNlciJ9.AcRSnpgzjsuJL2n_QaRF1idkwDzwNpWNX3wiEOFXkqTG35lr4PYVYPxnhryvRvVVOvN_CUY-AaVmF_YSgR4s6JM3Oca5JFFf7T6fX5lXgj0SbQCUbbyh7Em3BemiNKr_T3wucAyO824MjGXP0smciCnnlWvq-apJDTB_R4EisDJubY_E_zpCmTfYMm0NcJ8aKB2ku8mACKgE2ZJ7WsHkKNmjaFeyU9KjGMmNKtFthYISKqRQW-6u2xPjCkpFt4_HoJ01PgjFrrJacWDlUHxVoSILcaH_Vg-WHrKppIkzgdOg5phB2zVtcakRhPhqzV4EX_jXJp2SgK4umf6ivTC3lg"
	searchCtx := SearchContext{
		Ctx:       ctx,
		Id:        "cd1cc27b-28a1-47dc-ac76-70e8185e159d",
		UserToken: token,
	}
	err := _storageProcess.Delete(searchCtx)
	assert.NoError(t, err)
}

func TestDeleteFail(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddRemoveCall(func(ctx context.Context, id string, userId string) error {
		return errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRSnVoUjlFSFBIWTZFT195VjV4M1BTZWUzakRLNUs4M0lQMjJwYjFxZXFvIn0.eyJleHAiOjE3MDM4ODk3NTIsImlhdCI6MTcwMzg4OTQ1MiwianRpIjoiNTE4ZDM2MDctZjQ2NC00MDI5LTkwN2ItYjRjNzI1OWY0ZjU0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgxL3JlYWxtcy93YWxsZXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiNTgzMmE1MDItYmVkZS00OTJkLThkYzEtYjEzYjMyYzMwZjI5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2FsbGV0LWFwaSIsInNlc3Npb25fc3RhdGUiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLXdhbGxldCJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJhMmE2MDM5YS0zZTQxLTQ0MDEtOWJjNC01NWIyNDlkNmY3ZDYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RldXNlciJ9.AcRSnpgzjsuJL2n_QaRF1idkwDzwNpWNX3wiEOFXkqTG35lr4PYVYPxnhryvRvVVOvN_CUY-AaVmF_YSgR4s6JM3Oca5JFFf7T6fX5lXgj0SbQCUbbyh7Em3BemiNKr_T3wucAyO824MjGXP0smciCnnlWvq-apJDTB_R4EisDJubY_E_zpCmTfYMm0NcJ8aKB2ku8mACKgE2ZJ7WsHkKNmjaFeyU9KjGMmNKtFthYISKqRQW-6u2xPjCkpFt4_HoJ01PgjFrrJacWDlUHxVoSILcaH_Vg-WHrKppIkzgdOg5phB2zVtcakRhPhqzV4EX_jXJp2SgK4umf6ivTC3lg"
	searchCtx := SearchContext{
		Ctx:       ctx,
		Id:        "cd1cc27b-28a1-47dc-ac76-70e8185e159d",
		UserToken: token,
	}
	err := _storageProcess.Delete(searchCtx)
	assert.Error(t, err)
}
