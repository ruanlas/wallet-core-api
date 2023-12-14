package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetByIdSuccess(t *testing.T) {
	createdAt := time.Now()
	gainProjectMock := repository.NewGainProjectionBuilder().
		AddId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		AddCreatedAt(createdAt).
		AddPayIn(createdAt).
		AddIsPassive(true).
		AddIsDone(false).
		AddCategory(repository.GainCategory{Id: 2, Category: "Salário"}).
		AddDescription("Description teste").
		AddValue(750.50).
		AddUserId("User1").
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})

	ctx := context.TODO()
	_readingProcess := NewReadingProcess(_mockRepository)

	gainProject, err := _readingProcess.GetById(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	assert.NoError(t, err)
	assert.Equal(t, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", gainProject.Id)
}

func TestGetByIdNotFound(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return nil, nil
	})

	ctx := context.TODO()
	_readingProcess := NewReadingProcess(_mockRepository)

	gainProject, err := _readingProcess.GetById(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	assert.NoError(t, err)
	assert.Empty(t, gainProject)
}

func TestGetByIdError(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})

	ctx := context.TODO()
	_readingProcess := NewReadingProcess(_mockRepository)

	_, err := _readingProcess.GetById(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	assert.Error(t, err)
}
