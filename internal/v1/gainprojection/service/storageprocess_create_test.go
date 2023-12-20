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

type mockRepository struct {
	saveCallsMock            []func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error)
	getByIdCallsMock         []func(ctx context.Context, id string) (*repository.GainProjection, error)
	editCallsMock            []func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error)
	removeCallsMock          []func(ctx context.Context, id string) error
	getTotalRecordsCallsMock []func(ctx context.Context, params repository.QueryParams) (*uint, error)
	getAllCallsMock          []func(ctx context.Context, params repository.QueryParams) (*[]repository.GainProjection, error)
	saveGainCallsMock        []func(ctx context.Context, gain repository.Gain) (*repository.Gain, error)
}

func (r *mockRepository) AddSaveCall(
	save func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error)) *mockRepository {
	r.saveCallsMock = append(r.saveCallsMock, save)
	return r
}

func (r *mockRepository) AddGetByIdCall(
	getById func(ctx context.Context, id string) (*repository.GainProjection, error)) *mockRepository {
	r.getByIdCallsMock = append(r.getByIdCallsMock, getById)
	return r
}

func (r *mockRepository) AddEditCall(
	edit func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error)) *mockRepository {
	r.editCallsMock = append(r.editCallsMock, edit)
	return r
}

func (r *mockRepository) AddRemoveCall(
	remove func(ctx context.Context, id string) error) *mockRepository {
	r.removeCallsMock = append(r.removeCallsMock, remove)
	return r
}

func (r *mockRepository) AddGetTotalRecordsCalls(
	getTotalRecords func(ctx context.Context, params repository.QueryParams) (*uint, error)) *mockRepository {
	r.getTotalRecordsCallsMock = append(r.getTotalRecordsCallsMock, getTotalRecords)
	return r
}

func (r *mockRepository) AddGetAllCalls(
	getAll func(ctx context.Context, params repository.QueryParams) (*[]repository.GainProjection, error)) *mockRepository {
	r.getAllCallsMock = append(r.getAllCallsMock, getAll)
	return r
}

func (r *mockRepository) AddSaveGainCalls(
	saveGain func(ctx context.Context, gain repository.Gain) (*repository.Gain, error)) *mockRepository {
	r.saveGainCallsMock = append(r.saveGainCallsMock, saveGain)
	return r
}

func (r *mockRepository) Save(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
	if len(r.saveCallsMock) >= 1 {
		save := r.saveCallsMock[0]
		r.saveCallsMock = r.saveCallsMock[1:]
		return save(ctx, gainProjection)
	}
	return nil, nil
}

func (r *mockRepository) Edit(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
	if len(r.editCallsMock) >= 1 {
		edit := r.editCallsMock[0]
		r.editCallsMock = r.editCallsMock[1:]
		return edit(ctx, gainProjection)
	}
	return nil, nil
}

func (r *mockRepository) GetById(ctx context.Context, id string) (*repository.GainProjection, error) {
	if len(r.getByIdCallsMock) >= 1 {
		getById := r.getByIdCallsMock[0]
		r.getByIdCallsMock = r.getByIdCallsMock[1:]
		return getById(ctx, id)
	}
	return nil, nil
}

func (r *mockRepository) Remove(ctx context.Context, id string) error {
	if len(r.removeCallsMock) >= 1 {
		remove := r.removeCallsMock[0]
		r.removeCallsMock = r.removeCallsMock[1:]
		return remove(ctx, id)
	}
	return nil
}

func (r *mockRepository) GetTotalRecords(ctx context.Context, params repository.QueryParams) (*uint, error) {
	if len(r.getTotalRecordsCallsMock) >= 1 {
		getTotalRecords := r.getTotalRecordsCallsMock[0]
		r.getTotalRecordsCallsMock = r.getTotalRecordsCallsMock[1:]
		return getTotalRecords(ctx, params)
	}
	return nil, nil
}

func (r *mockRepository) GetAll(ctx context.Context, params repository.QueryParams) (*[]repository.GainProjection, error) {
	if len(r.getAllCallsMock) >= 1 {
		getAll := r.getAllCallsMock[0]
		r.getAllCallsMock = r.getAllCallsMock[1:]
		return getAll(ctx, params)
	}
	return nil, nil
}

func (r *mockRepository) SaveGain(ctx context.Context, gain repository.Gain) (*repository.Gain, error) {
	if len(r.saveGainCallsMock) >= 1 {
		saveGain := r.saveGainCallsMock[0]
		r.saveGainCallsMock = r.saveGainCallsMock[1:]
		return saveGain(ctx, gain)
	}
	return nil, nil
}

func TestCreateSuccessWithoutRecurrence(t *testing.T) {

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
	_mockRepository.AddSaveCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := CreateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	response, err := _storageProcess.Create(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", response.Id)
	assert.Equal(t, uint(1), response.Recurrence)
}

func TestCreateSuccessWithRecurrence(t *testing.T) {

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
	_mockRepository.AddSaveCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddSaveCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := CreateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		Recurrence:  2,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	response, err := _storageProcess.Create(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, "cd1cc27b-28a1-47dc-ac76-70e8185e159d", response.Id)
	assert.Equal(t, uint(2), response.Recurrence)
}

func TestCreateWithoutRecurrenceSaveFail(t *testing.T) {

	_mockRepository := &mockRepository{}
	_mockRepository.AddSaveCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := CreateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	_, err := _storageProcess.Create(ctx, request)
	assert.Error(t, err)

}

func TestCreateWithoutRecurrenceGetByIdFail(t *testing.T) {

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
	_mockRepository.AddSaveCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddGetByIdCall(func(ctx context.Context, id string) (*repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := CreateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	_, err := _storageProcess.Create(ctx, request)
	assert.Error(t, err)
}

func TestCreateWithRecurrenceSaveFail(t *testing.T) {

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
	_mockRepository.AddSaveCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return gainProjectMock, nil
	})
	_mockRepository.AddSaveCall(func(ctx context.Context, gainProjection repository.GainProjection) (*repository.GainProjection, error) {
		return nil, errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}

	request := CreateRequest{
		PayIn:       time.Now(),
		Description: "Description teste",
		Value:       750.50,
		IsPassive:   false,
		Recurrence:  2,
		CategoryId:  2,
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	_, err := _storageProcess.Create(ctx, request)
	assert.Error(t, err)
}
