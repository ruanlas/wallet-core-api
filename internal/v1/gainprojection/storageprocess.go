package gainprojection

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type StorageProcess interface {
	Create(request CreateRequest) (*GainProjectionResponse, error)
}

type storageProcess struct {
	repository   Repository
	generateUUID func() uuid.UUID
}

func NewStorageProcess(repository Repository) StorageProcess {
	return &storageProcess{repository: repository}
}

func (sp *storageProcess) Create(request CreateRequest) (*GainProjectionResponse, error) {
	createdAt := time.Now()
	gainProjection := NewGainProjectionBuilder().
		AddId(sp.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddIsDone(false).
		AddCategory(GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value).
		AddUserId("User1").
		Build()

	gainProjectionSaved, err := sp.repository.Save(*gainProjection)
	if err != nil {
		return nil, err
	}

	if request.Recurrence > 1 {
		// Isolar em um método recurrenceProcess
		for i := 1; i < int(request.Recurrence+1); i++ {
			gainProjection := NewGainProjectionBuilder().
				AddId(sp.generateUUID().String()).
				AddCreatedAt(createdAt).
				AddPayIn(request.PayIn.AddDate(0, i, 0)).
				AddIsPassive(request.IsPassive).
				AddIsDone(false).
				AddCategory(GainCategory{Id: request.CategoryId}).
				AddDescription(request.Description).
				AddValue(request.Value).
				AddUserId("User1").
				Build()

			_, err = sp.repository.Save(*gainProjection)
			if err != nil {
				return nil, err
			}
		}
	} else {
		request.Recurrence = 1
	}
	gainProjectionSaved, err = sp.repository.GetById(gainProjectionSaved.Id)
	if err != nil {
		return nil, err
	}

	return NewGainProjectionResponseBuilder().
		AddId(gainProjection.Id).
		AddPayIn(gainProjection.PayIn).
		AddDescription(gainProjection.Description).
		AddValue(gainProjection.Value).
		AddIsPassive(gainProjection.IsPassive).
		AddCategory(CategoryResponse{Id: gainProjection.Category.Id, Category: gainProjection.Category.Category}).
		AddRecurrence(request.Recurrence).
		Build(), nil
}
