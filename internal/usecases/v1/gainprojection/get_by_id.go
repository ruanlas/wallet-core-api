package gainprojection

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gainprojection"
)

type GetByIdUseCase interface {
	Execute(searchCtx SearchContext) (*GainProjectionResponse, error)
}

type getByIdUseCase struct {
	repository gainprojection.Repository
}

func NewGetByIdUseCase(repository gainprojection.Repository) GetByIdUseCase {
	return &getByIdUseCase{
		repository: repository,
	}
}

func (uc *getByIdUseCase) Execute(searchCtx SearchContext) (*GainProjectionResponse, error) {
	user := idpauth.GetUser(searchCtx.UserToken)
	gainProjection, err := uc.repository.GetById(searchCtx.Ctx, searchCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gainProjection == nil {
		return nil, nil
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
