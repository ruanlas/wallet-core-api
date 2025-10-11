package gainprojection

import (
	"context"
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gainprojection"
	uuid "github.com/satori/go.uuid"
)

type CreateUseCase interface {
	Execute(createCtx CreateContext) (*GainProjectionResponse, error)
}

func NewCreateUseCase(repository gainprojection.Repository, generateUUID func() uuid.UUID) CreateUseCase {
	return &createUseCase{
		repository:   repository,
		generateUUID: generateUUID,
	}
}

type createUseCase struct {
	repository   gainprojection.Repository
	generateUUID func() uuid.UUID
}

func (uc *createUseCase) Execute(createCtx CreateContext) (*GainProjectionResponse, error) {
	request := createCtx.Request
	user := idpauth.GetUser(createCtx.UserToken)
	createdAt := time.Now()
	gainProjection := gainprojection.NewGainProjectionBuilder().
		AddId(uc.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddIsAlreadyDone(false).
		AddCategory(gainprojection.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value).
		AddUserId(user.Id).
		Build()

	gainProjectionSaved, err := uc.repository.Save(createCtx.Ctx, *gainProjection)
	if err != nil {
		return nil, err
	}

	if request.Recurrence > 1 {
		err = uc.createRecurrence(createCtx.Ctx, request, createdAt, user.Id)
		if err != nil {
			return nil, err
		}
	} else {
		request.Recurrence = 1
	}
	gainProjectionSaved, err = uc.repository.GetById(createCtx.Ctx, gainProjectionSaved.Id, user.Id)
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

func (uc *createUseCase) createRecurrence(ctx context.Context, request CreateRequest, createdAt time.Time, userId string) error {
	for i := 1; i < int(request.Recurrence+1); i++ {
		gainProjection := gainprojection.NewGainProjectionBuilder().
			AddId(uc.generateUUID().String()).
			AddCreatedAt(createdAt).
			AddPayIn(request.PayIn.AddDate(0, i, 0)).
			AddIsPassive(request.IsPassive).
			AddIsAlreadyDone(false).
			AddCategory(gainprojection.GainCategory{Id: request.CategoryId}).
			AddDescription(request.Description).
			AddValue(request.Value).
			AddUserId(userId).
			Build()

		_, err := uc.repository.Save(ctx, *gainProjection)
		if err != nil {
			return err
		}
	}
	return nil
}
