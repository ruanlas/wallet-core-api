package gain

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gain"
)

type UpdateUseCase interface {
	Execute(updateCtx UpdateContext) (*GainResponse, error)
}

type updateUseCase struct {
	repository gain.Repository
}

func NewUpdateUseCase(repository gain.Repository) UpdateUseCase {
	return &updateUseCase{
		repository: repository,
	}
}

func (uc *updateUseCase) Execute(updateCtx UpdateContext) (*GainResponse, error) {
	request := updateCtx.Request
	user := idpauth.GetUser(updateCtx.UserToken)
	gainBuilder := gain.NewGainBuilder().
		AddId(updateCtx.Id).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddCategory(gain.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value)
	gainExists, err := uc.repository.GetById(updateCtx.Ctx, updateCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gainExists == nil {
		return nil, nil
	}
	gainBuilder.AddUserId(user.Id)
	gainUpdated, err := uc.repository.Edit(updateCtx.Ctx, *gainBuilder.Build())
	if err != nil {
		return nil, err
	}

	gainUpdated, err = uc.repository.GetById(updateCtx.Ctx, gainUpdated.Id, user.Id)
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
