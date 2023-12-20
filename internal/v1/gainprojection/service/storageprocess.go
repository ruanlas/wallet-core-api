package service

import (
	"context"
	"time"

	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/repository"
	uuid "github.com/satori/go.uuid"
)

type StorageProcess interface {
	Create(ctx context.Context, request CreateRequest) (*GainProjectionResponse, error)
	Update(ctx context.Context, id string, request UpdateRequest) (*GainProjectionResponse, error)
	Delete(ctx context.Context, id string) error
	CreateGain(ctx context.Context, id string, request CreateGainRequest) (*GainStat, error)
}

type storageProcess struct {
	repository   repository.Repository
	generateUUID func() uuid.UUID
}

func NewStorageProcess(repository repository.Repository, generateUUID func() uuid.UUID) StorageProcess {
	return &storageProcess{repository: repository, generateUUID: generateUUID}
}

func (sp *storageProcess) Create(ctx context.Context, request CreateRequest) (*GainProjectionResponse, error) {
	createdAt := time.Now()
	gainProjection := repository.NewGainProjectionBuilder().
		AddId(sp.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddIsDone(false).
		AddCategory(repository.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value).
		AddUserId("User1").
		Build()

	gainProjectionSaved, err := sp.repository.Save(ctx, *gainProjection)
	if err != nil {
		return nil, err
	}

	if request.Recurrence > 1 {
		err = sp.createRecurrence(ctx, request, createdAt)
		if err != nil {
			return nil, err
		}
	} else {
		request.Recurrence = 1
	}
	gainProjectionSaved, err = sp.repository.GetById(ctx, gainProjectionSaved.Id)
	if err != nil {
		return nil, err
	}

	return NewGainProjectionResponseBuilder().
		AddId(gainProjection.Id).
		AddPayIn(gainProjection.PayIn).
		AddDescription(gainProjection.Description).
		AddValue(gainProjection.Value).
		AddIsPassive(gainProjection.IsPassive).
		AddCategory(CategoryResponse{Id: gainProjectionSaved.Category.Id, Category: gainProjectionSaved.Category.Category}).
		AddRecurrence(request.Recurrence).
		Build(), nil
}

func (sp *storageProcess) createRecurrence(ctx context.Context, request CreateRequest, createdAt time.Time) error {
	for i := 1; i < int(request.Recurrence+1); i++ {
		gainProjection := repository.NewGainProjectionBuilder().
			AddId(sp.generateUUID().String()).
			AddCreatedAt(createdAt).
			AddPayIn(request.PayIn.AddDate(0, i, 0)).
			AddIsPassive(request.IsPassive).
			AddIsDone(false).
			AddCategory(repository.GainCategory{Id: request.CategoryId}).
			AddDescription(request.Description).
			AddValue(request.Value).
			AddUserId("User1").
			Build()

		_, err := sp.repository.Save(ctx, *gainProjection)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sp *storageProcess) Update(ctx context.Context, id string, request UpdateRequest) (*GainProjectionResponse, error) {
	gainProjectionBuilder := repository.NewGainProjectionBuilder().
		AddId(id).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddCategory(repository.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value)
	gainProjectionExists, err := sp.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if gainProjectionExists == nil {
		return nil, nil
	}
	gainProjectionBuilder.AddIsDone(gainProjectionExists.IsDone)
	gainProjectionUpdated, err := sp.repository.Edit(ctx, *gainProjectionBuilder.Build())
	if err != nil {
		return nil, err
	}

	gainProjectionUpdated, err = sp.repository.GetById(ctx, gainProjectionUpdated.Id)
	if err != nil {
		return nil, err
	}

	return NewGainProjectionResponseBuilder().
		AddId(gainProjectionUpdated.Id).
		AddPayIn(gainProjectionUpdated.PayIn).
		AddDescription(gainProjectionUpdated.Description).
		AddValue(gainProjectionUpdated.Value).
		AddIsPassive(gainProjectionUpdated.IsPassive).
		AddCategory(CategoryResponse{Id: gainProjectionUpdated.Category.Id, Category: gainProjectionUpdated.Category.Category}).
		Build(), nil
}

func (sp *storageProcess) Delete(ctx context.Context, id string) error {
	return sp.repository.Remove(ctx, id)
}

func (sp *storageProcess) CreateGain(ctx context.Context, id string, request CreateGainRequest) (*GainStat, error) {
	gainProjection, err := sp.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if gainProjection == nil {
		return &GainStat{ProjectionIsFound: false, ProjectionIsAlreadyDone: false}, nil
	}
	if gainProjection.IsDone == true {
		return &GainStat{ProjectionIsFound: true, ProjectionIsAlreadyDone: true}, nil
	}
	gainBuilder := repository.NewGainBuilder().
		AddId(sp.generateUUID().String()).
		AddCategory(gainProjection.Category).
		AddCreatedAt(time.Now()).
		AddDescription(gainProjection.Description).
		AddGainProjectionId(gainProjection.Id).
		AddIsPassive(gainProjection.IsPassive).
		AddUserId(gainProjection.UserId).
		AddValue(gainProjection.Value).
		AddPayIn(gainProjection.PayIn)
	if request.Value != 0 {
		gainBuilder.AddValue(request.Value)
	}
	if !request.PayIn.IsZero() {
		gainBuilder.AddPayIn(request.PayIn)
	}
	gain, err := sp.repository.SaveGain(ctx, *gainBuilder.Build())
	if err != nil {
		return nil, err
	}
	gainProjection.IsDone = true
	_, err = sp.repository.Edit(ctx, *gainProjection)
	if err != nil {
		return nil, err
	}

	gainResponse := NewGainResponseBuilder().
		AddId(gain.Id).
		AddGainProjectionId(gain.GainProjectionId).
		AddPayIn(gain.PayIn).
		AddDescription(gain.Description).
		AddValue(gain.Value).
		AddIsPassive(gain.IsPassive).
		AddCategory(CategoryResponse{Id: gain.Category.Id, Category: gain.Category.Category}).
		Build()
	return &GainStat{ProjectionIsFound: true, ProjectionIsAlreadyDone: false, Gain: gainResponse}, nil
}
