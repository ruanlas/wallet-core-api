package gain

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gain"
)

type GetByIdUseCase interface {
	Execute(searchCtx SearchContext) (*GainResponse, error)
}

func NewGetByIdUseCase(repository gain.Repository) GetByIdUseCase {
	return &getByIdUseCase{
		repository: repository,
	}
}

type getByIdUseCase struct {
	repository gain.Repository
}

func (uc *getByIdUseCase) Execute(searchCtx SearchContext) (*GainResponse, error) {
	user := idpauth.GetUser(searchCtx.UserToken)
	gain, err := uc.repository.GetById(searchCtx.Ctx, searchCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gain == nil {
		return nil, nil
	}

	return NewGainResponseBuilder().
		AddId(gain.Id).
		AddPayIn(gain.PayIn).
		AddDescription(gain.Description).
		AddValue(gain.Value).
		AddIsPassive(gain.IsPassive).
		AddGainProjectionId(gain.GainProjectionId).
		AddCategory(CategoryResponse{Id: gain.Category.Id, Category: gain.Category.Category}).
		Build(), nil
}
