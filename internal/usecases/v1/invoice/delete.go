package invoice

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoice"
)

type DeleteUseCase interface {
	Execute(searchCtx SearchContext) error
}

func NewDeleteUseCase(repository invoice.Repository) DeleteUseCase {
	return &deleteUseCase{repository: repository}
}

type deleteUseCase struct {
	repository invoice.Repository
}

func (uc *deleteUseCase) Execute(searchCtx SearchContext) error {
	user := idpauth.GetUser(searchCtx.UserToken)
	return uc.repository.Remove(searchCtx.Ctx, searchCtx.Id, user.Id)
}
