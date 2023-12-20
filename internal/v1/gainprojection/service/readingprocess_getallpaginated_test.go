package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPaginatedSuccess(t *testing.T) {
	createdAt := time.Now()
	gainProjectMock := repository.NewGainProjectionBuilder().
		AddId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		AddCreatedAt(createdAt).
		AddPayIn(createdAt).
		AddIsPassive(true).
		AddIsAlreadyDone(false).
		AddCategory(repository.GainCategory{Id: 2, Category: "Salário"}).
		AddDescription("Description teste").
		AddValue(750.50).
		AddUserId("User1").
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetTotalRecordsCalls(func(ctx context.Context, params repository.QueryParams) (*uint, error) {
		totalRecords := uint(5)
		return &totalRecords, nil
	})
	_mockRepository.AddGetAllCalls(func(ctx context.Context, params repository.QueryParams) (*[]repository.GainProjection, error) {
		return &[]repository.GainProjection{*gainProjectMock}, nil
	})

	ctx := context.TODO()
	_readingProcess := NewReadingProcess(_mockRepository)

	searchParams := NewSearchParamsBuilder().
		AddMonth(10).
		AddYear(2023).
		AddPage(3).
		AddPageSize(2).
		Build()
	gainProjectPaginated, err := _readingProcess.GetAllPaginated(ctx, *searchParams)
	assert.NoError(t, err)
	assert.Equal(t, uint(3), gainProjectPaginated.TotalPages)
	assert.Equal(t, uint(5), gainProjectPaginated.TotalRecords)
}

func TestGetAllPaginatedGetTotalRecordsFail(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetTotalRecordsCalls(func(ctx context.Context, params repository.QueryParams) (*uint, error) {
		return nil, errors.New("An error has been ocurred")
	})

	ctx := context.TODO()
	_readingProcess := NewReadingProcess(_mockRepository)

	searchParams := NewSearchParamsBuilder().
		AddMonth(10).
		AddYear(2023).
		AddPage(3).
		AddPageSize(2).
		Build()
	_, err := _readingProcess.GetAllPaginated(ctx, *searchParams)
	assert.Error(t, err)
}

func TestGetAllPaginatedGetAllFail(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetTotalRecordsCalls(func(ctx context.Context, params repository.QueryParams) (*uint, error) {
		totalRecords := uint(5)
		return &totalRecords, nil
	})
	_mockRepository.AddGetAllCalls(func(ctx context.Context, params repository.QueryParams) (*[]repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})

	ctx := context.TODO()
	_readingProcess := NewReadingProcess(_mockRepository)

	searchParams := NewSearchParamsBuilder().
		AddMonth(10).
		AddYear(2023).
		AddPage(3).
		AddPageSize(2).
		Build()
	_, err := _readingProcess.GetAllPaginated(ctx, *searchParams)
	assert.Error(t, err)
}
