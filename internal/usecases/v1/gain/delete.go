package gain

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gain"
)

type DeleteUseCase interface {
	Execute(searchCtx SearchContext) error
}

type deleteUseCase struct {
	repository gain.Repository
}

func NewDeleteUseCase(repository gain.Repository) DeleteUseCase {
	return &deleteUseCase{
		repository: repository,
	}
}

func (uc *deleteUseCase) Execute(searchCtx SearchContext) error {
	user := idpauth.GetUser(searchCtx.UserToken)
	return uc.repository.Remove(searchCtx.Ctx, searchCtx.Id, user.Id)
}
