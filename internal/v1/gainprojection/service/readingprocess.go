package service

import (
	"context"

	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/repository"
	uuid "github.com/satori/go.uuid"
)

type ReadingProcess interface {
	GetById(ctx context.Context, gainProjectionId string) (*GainProjectionResponse, error)
}

type readingProcess struct {
	repository   repository.Repository
	generateUUID func() uuid.UUID
}

func NewReadingProcess(repository repository.Repository) ReadingProcess {
	return &readingProcess{repository: repository}
}

func (rp *readingProcess) GetById(ctx context.Context, gainProjectionId string) (*GainProjectionResponse, error) {
	gainProjection, err := rp.repository.GetById(ctx, gainProjectionId)
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
		Build(), nil
}
