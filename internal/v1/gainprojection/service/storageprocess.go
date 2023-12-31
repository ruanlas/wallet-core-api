package service

import (
	"context"
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/repository"
	uuid "github.com/satori/go.uuid"
)

type StorageProcess interface {
	Create(createCtx CreateContext) (*GainProjectionResponse, error)
	Update(updateCtx UpdateContext) (*GainProjectionResponse, error)
	Delete(searchCtx SearchContext) error
	CreateGain(createGainCtx CreateGainContext) (*GainStat, error)
}

type storageProcess struct {
	repository   repository.Repository
	generateUUID func() uuid.UUID
}

func NewStorageProcess(repository repository.Repository, generateUUID func() uuid.UUID) StorageProcess {
	return &storageProcess{repository: repository, generateUUID: generateUUID}
}

func (sp *storageProcess) Create(createCtx CreateContext) (*GainProjectionResponse, error) {
	request := createCtx.Request
	user := idpauth.GetUser(createCtx.UserToken)
	createdAt := time.Now()
	gainProjection := repository.NewGainProjectionBuilder().
		AddId(sp.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddIsAlreadyDone(false).
		AddCategory(repository.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value).
		AddUserId(user.Id).
		Build()

	gainProjectionSaved, err := sp.repository.Save(createCtx.Ctx, *gainProjection)
	if err != nil {
		return nil, err
	}

	if request.Recurrence > 1 {
		err = sp.createRecurrence(createCtx.Ctx, request, createdAt, user.Id)
		if err != nil {
			return nil, err
		}
	} else {
		request.Recurrence = 1
	}
	gainProjectionSaved, err = sp.repository.GetById(createCtx.Ctx, gainProjectionSaved.Id, user.Id)
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

func (sp *storageProcess) createRecurrence(ctx context.Context, request CreateRequest, createdAt time.Time, userId string) error {
	for i := 1; i < int(request.Recurrence+1); i++ {
		gainProjection := repository.NewGainProjectionBuilder().
			AddId(sp.generateUUID().String()).
			AddCreatedAt(createdAt).
			AddPayIn(request.PayIn.AddDate(0, i, 0)).
			AddIsPassive(request.IsPassive).
			AddIsAlreadyDone(false).
			AddCategory(repository.GainCategory{Id: request.CategoryId}).
			AddDescription(request.Description).
			AddValue(request.Value).
			AddUserId(userId).
			Build()

		_, err := sp.repository.Save(ctx, *gainProjection)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sp *storageProcess) Update(updateCtx UpdateContext) (*GainProjectionResponse, error) {
	request := updateCtx.Request
	user := idpauth.GetUser(updateCtx.UserToken)
	gainProjectionBuilder := repository.NewGainProjectionBuilder().
		AddId(updateCtx.Id).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddCategory(repository.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value)
	gainProjectionExists, err := sp.repository.GetById(updateCtx.Ctx, updateCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gainProjectionExists == nil {
		return nil, nil
	}
	gainProjectionBuilder.AddIsAlreadyDone(gainProjectionExists.IsAlreadyDone)
	gainProjectionBuilder.AddUserId(user.Id)
	gainProjectionUpdated, err := sp.repository.Edit(updateCtx.Ctx, *gainProjectionBuilder.Build())
	if err != nil {
		return nil, err
	}

	gainProjectionUpdated, err = sp.repository.GetById(updateCtx.Ctx, gainProjectionUpdated.Id, user.Id)
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

func (sp *storageProcess) Delete(searchCtx SearchContext) error {
	user := idpauth.GetUser(searchCtx.UserToken)
	return sp.repository.Remove(searchCtx.Ctx, searchCtx.Id, user.Id)
}

func (sp *storageProcess) CreateGain(createGainCtx CreateGainContext) (*GainStat, error) {
	request := createGainCtx.Request
	user := idpauth.GetUser(createGainCtx.UserToken)
	gainProjection, err := sp.repository.GetById(createGainCtx.Ctx, createGainCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gainProjection == nil {
		return &GainStat{ProjectionIsFound: false, ProjectionIsAlreadyDone: false}, nil
	}
	if gainProjection.IsAlreadyDone == true {
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
	gain, err := sp.repository.SaveGain(createGainCtx.Ctx, *gainBuilder.Build())
	if err != nil {
		return nil, err
	}
	gainProjection.IsAlreadyDone = true
	_, err = sp.repository.Edit(createGainCtx.Ctx, *gainProjection)
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
