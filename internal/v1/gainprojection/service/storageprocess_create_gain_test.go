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

func TestCreateGainSuccess(t *testing.T) {

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
	gainMock := repository.NewGainBuilder().
		AddId("71e31eb6-dde2-4dcb-b7ef-c7e8a699c628").
		AddCreatedAt(createdAt).
		AddPayIn(createdAt).
		AddIsPassive(true).
		AddCategory(repository.GainCategory{Id: 2, Category: "Salário"}).
		AddDescription("Description teste").
		AddValue(750.50).
		AddUserId("User1").
		AddGainProjectionId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddSaveGainCalls(func(ctx context.Context, gain repository.Gain) (*repository.Gain, error) {
		return gainMock, nil
	})
	_mockRepository.AddEditCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return &gainProjection, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("71e31eb6-dde2-4dcb-b7ef-c7e8a699c628")
	}

	request := CreateGainRequest{
		PayIn: time.Now(),
		Value: 750.50,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	response, err := _storageProcess.CreateGain(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.NoError(t, err)
	assert.Equal(t, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", response.Gain.GainProjectionId)
	assert.True(t, response.ProjectionIsFound)
	assert.False(t, response.ProjectionIsAlreadyDone)
}

func TestCreateGainGetByIdFail(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("71e31eb6-dde2-4dcb-b7ef-c7e8a699c628")
	}

	request := CreateGainRequest{
		PayIn: time.Now(),
		Value: 750.50,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	_, err := _storageProcess.CreateGain(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.Error(t, err)
}

func TestCreateGainGetByIdEmpty(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return nil, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("71e31eb6-dde2-4dcb-b7ef-c7e8a699c628")
	}

	request := CreateGainRequest{
		PayIn: time.Now(),
		Value: 750.50,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	response, err := _storageProcess.CreateGain(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.NoError(t, err)
	assert.False(t, response.ProjectionIsFound)
	assert.False(t, response.ProjectionIsAlreadyDone)
}

func TestCreateGainAlreadyDone(t *testing.T) {

	createdAt := time.Now()
	gainProjectMock := repository.NewGainProjectionBuilder().
		AddId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		AddCreatedAt(createdAt).
		AddPayIn(createdAt).
		AddIsPassive(true).
		AddIsAlreadyDone(true).
		AddCategory(repository.GainCategory{Id: 2, Category: "Salário"}).
		AddDescription("Description teste").
		AddValue(750.50).
		AddUserId("User1").
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("71e31eb6-dde2-4dcb-b7ef-c7e8a699c628")
	}

	request := CreateGainRequest{
		PayIn: time.Now(),
		Value: 750.50,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	response, err := _storageProcess.CreateGain(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.NoError(t, err)
	assert.True(t, response.ProjectionIsFound)
	assert.True(t, response.ProjectionIsAlreadyDone)
}

func TestCreateGainSaveGainFail(t *testing.T) {

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
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddSaveGainCalls(func(ctx context.Context, gain repository.Gain) (*repository.Gain, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("71e31eb6-dde2-4dcb-b7ef-c7e8a699c628")
	}

	request := CreateGainRequest{
		PayIn: time.Now(),
		Value: 750.50,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	_, err := _storageProcess.CreateGain(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.Error(t, err)
}

func TestCreateGainEditFail(t *testing.T) {

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
	gainMock := repository.NewGainBuilder().
		AddId("71e31eb6-dde2-4dcb-b7ef-c7e8a699c628").
		AddCreatedAt(createdAt).
		AddPayIn(createdAt).
		AddIsPassive(true).
		AddCategory(repository.GainCategory{Id: 2, Category: "Salário"}).
		AddDescription("Description teste").
		AddValue(750.50).
		AddUserId("User1").
		AddGainProjectionId("cd1cc27b-28a1-47dc-ac76-70e8185e159d").
		Build()
	_mockRepository := &mockRepository{}
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddSaveGainCalls(func(ctx context.Context, gain repository.Gain) (*repository.Gain, error) {
		return gainMock, nil
	})
	_mockRepository.AddEditCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("71e31eb6-dde2-4dcb-b7ef-c7e8a699c628")
	}

	request := CreateGainRequest{
		PayIn: time.Now(),
		Value: 750.50,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	_, err := _storageProcess.CreateGain(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", request)
	assert.Error(t, err)
}
