package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/repository"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSuccess(t *testing.T) {

	gainProjectMock := repository.NewGainProjectionBuilder().
		AddId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		AddPayIn(time.Now()).
		AddIsPassive(true).
		AddCategory(repository.GainCategory{Id: 2, Category: "Salário"}).
		AddDescription("Description teste").
		AddValue(750.50).
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddEditCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := UpdateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	response, err := _storageProcess.Update(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.NoError(t, err)
	assert.Equal(t, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", response.Id)
}

func TestUpdateFirstGetByIdNotFound(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return nil, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := UpdateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	response, err := _storageProcess.Update(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.NoError(t, err)
	assert.Empty(t, response)
}

func TestUpdateFirstGetByIdFail(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := UpdateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	_, err := _storageProcess.Update(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.Error(t, err)
}

func TestUpdateEditFail(t *testing.T) {

	gainProjectMock := repository.NewGainProjectionBuilder().
		AddId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		AddPayIn(time.Now()).
		AddIsPassive(true).
		AddCategory(repository.GainCategory{Id: 2, Category: "Salário"}).
		AddDescription("Description teste").
		AddValue(750.50).
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddEditCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := UpdateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	_, err := _storageProcess.Update(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.Error(t, err)
}

func TestUpdateSecondGetByIdFail(t *testing.T) {

	gainProjectMock := repository.NewGainProjectionBuilder().
		AddId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		AddPayIn(time.Now()).
		AddIsPassive(true).
		AddCategory(repository.GainCategory{Id: 2, Category: "Salário"}).
		AddDescription("Description teste").
		AddValue(750.50).
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})
	_mockRepository.AddEditCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := UpdateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	_, err := _storageProcess.Update(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.Error(t, err)
}
