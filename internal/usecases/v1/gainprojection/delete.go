package gainprojection

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gainprojection"
)

type DeleteUseCase interface {
	Execute(searchCtx SearchContext) error
}

func NewDeleteUseCase(repository gainprojection.Repository) DeleteUseCase {
	return &deleteUseCase{
		repository: repository,
	}
}

type deleteUseCase struct {
	repository gainprojection.Repository
}

func (uc *deleteUseCase) Execute(searchCtx SearchContext) error {
	user := idpauth.GetUser(searchCtx.UserToken)
	return uc.repository.Remove(searchCtx.Ctx, searchCtx.Id, user.Id)
}
