package gservice

import (
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/v1/gain/repository"
	uuid "github.com/satori/go.uuid"
)

type StorageProcess interface {
	Create(createCtx CreateContext) (*GainResponse, error)
	Update(updateCtx UpdateContext) (*GainResponse, error)
	Delete(searchCtx SearchContext) error
}

type storageProcess struct {
	repository   repository.Repository
	generateUUID func() uuid.UUID
}

func NewStorageProcess(repository repository.Repository, generateUUID func() uuid.UUID) StorageProcess {
	return &storageProcess{repository: repository, generateUUID: generateUUID}
}

func (sp *storageProcess) Create(createCtx CreateContext) (*GainResponse, error) {
	request := createCtx.Request
	user := idpauth.GetUser(createCtx.UserToken)
	createdAt := time.Now()
	gainBuilder := repository.NewGainBuilder().
		AddId(sp.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddCategory(repository.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value).
		AddUserId(user.Id)
	if request.PayIn.IsZero() {
		gainBuilder.AddPayIn(createdAt)
	}

	gain := gainBuilder.Build()
	gainSaved, err := sp.repository.Save(createCtx.Ctx, *gain)
	if err != nil {
		return nil, err
	}

	gainSaved, err = sp.repository.GetById(createCtx.Ctx, gainSaved.Id, user.Id)
	if err != nil {
		return nil, err
	}

	return NewGainResponseBuilder().
		AddId(gain.Id).
		AddPayIn(gain.PayIn).
		AddDescription(gain.Description).
		AddValue(gain.Value).
		AddIsPassive(gain.IsPassive).
		AddCategory(CategoryResponse{Id: gainSaved.Category.Id, Category: gainSaved.Category.Category}).
		Build(), nil
}

func (sp *storageProcess) Update(updateCtx UpdateContext) (*GainResponse, error) {
	request := updateCtx.Request
	user := idpauth.GetUser(updateCtx.UserToken)
	gainBuilder := repository.NewGainBuilder().
		AddId(updateCtx.Id).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddCategory(repository.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value)
	gainExists, err := sp.repository.GetById(updateCtx.Ctx, updateCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gainExists == nil {
		return nil, nil
	}
	gainBuilder.AddUserId(user.Id)
	gainUpdated, err := sp.repository.Edit(updateCtx.Ctx, *gainBuilder.Build())
	if err != nil {
		return nil, err
	}

	gainUpdated, err = sp.repository.GetById(updateCtx.Ctx, gainUpdated.Id, user.Id)
	if err != nil {
		return nil, err
	}

	return NewGainResponseBuilder().
		AddId(gainUpdated.Id).
		AddPayIn(gainUpdated.PayIn).
		AddDescription(gainUpdated.Description).
		AddValue(gainUpdated.Value).
		AddIsPassive(gainUpdated.IsPassive).
		AddCategory(CategoryResponse{Id: gainUpdated.Category.Id, Category: gainUpdated.Category.Category}).
		Build(), nil
}

func (sp *storageProcess) Delete(searchCtx SearchContext) error {
	user := idpauth.GetUser(searchCtx.UserToken)
	return sp.repository.Remove(searchCtx.Ctx, searchCtx.Id, user.Id)
}
