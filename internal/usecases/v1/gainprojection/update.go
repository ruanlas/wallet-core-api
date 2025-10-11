package gainprojection

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gainprojection"
)

type UpdateUseCase interface {
	Execute(updateCtx UpdateContext) (*GainProjectionResponse, error)
}

func NewUpdateUseCase(repository gainprojection.Repository) UpdateUseCase {
	return &updateUseCase{
		repository: repository,
	}
}

type updateUseCase struct {
	repository gainprojection.Repository
}

func (uc *updateUseCase) Execute(updateCtx UpdateContext) (*GainProjectionResponse, error) {
	request := updateCtx.Request
	user := idpauth.GetUser(updateCtx.UserToken)
	gainProjectionBuilder := gainprojection.NewGainProjectionBuilder().
		AddId(updateCtx.Id).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddCategory(gainprojection.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value)
	gainProjectionExists, err := uc.repository.GetById(updateCtx.Ctx, updateCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gainProjectionExists == nil {
		return nil, nil
	}
	gainProjectionBuilder.AddIsAlreadyDone(gainProjectionExists.IsAlreadyDone)
	gainProjectionBuilder.AddUserId(user.Id)
	gainProjectionUpdated, err := uc.repository.Edit(updateCtx.Ctx, *gainProjectionBuilder.Build())
	if err != nil {
		return nil, err
	}

	gainProjectionUpdated, err = uc.repository.GetById(updateCtx.Ctx, gainProjectionUpdated.Id, user.Id)
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
