package invoiceprojection

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoiceprojection"
)

type GetByIdUseCase interface {
	Execute(searchCtx SearchContext) (*InvoiceProjectionResponse, error)
}

func NewGetByIdUseCase(repository invoiceprojection.Repository) GetByIdUseCase {
	return &getByIdUseCase{
		repository: repository,
	}
}

type getByIdUseCase struct {
	repository invoiceprojection.Repository
}

func (uc *getByIdUseCase) Execute(searchCtx SearchContext) (*InvoiceProjectionResponse, error) {
	user := idpauth.GetUser(searchCtx.UserToken)
	invoiceProjection, err := uc.repository.GetById(searchCtx.Ctx, searchCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if invoiceProjection == nil {
		return nil, nil
	}

	return NewInvoiceProjectionResponseBuilder().
		AddId(invoiceProjection.Id).
		AddPayIn(invoiceProjection.PayIn).
		AddBuyAt(invoiceProjection.BuyAt).
		AddDescription(invoiceProjection.Description).
		AddValue(invoiceProjection.Value).
		AddPaymentType(PaymentTypeResponse{Id: invoiceProjection.PaymentType.Id, Type: invoiceProjection.PaymentType.Type}).
		AddCategory(CategoryResponse{Id: invoiceProjection.Category.Id, Category: invoiceProjection.Category.Category}).
		Build(), nil
}
