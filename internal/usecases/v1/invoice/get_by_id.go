package invoice

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoice"
)

type GetByIdUseCase interface {
	Execute(searchCtx SearchContext) (*InvoiceResponse, error)
}

func NewGetByIdUseCase(repository invoice.Repository) GetByIdUseCase {
	return &getByIdUseCase{repository: repository}
}

type getByIdUseCase struct {
	repository invoice.Repository
}

func (uc *getByIdUseCase) Execute(searchCtx SearchContext) (*InvoiceResponse, error) {
	user := idpauth.GetUser(searchCtx.UserToken)
	invoice, err := uc.repository.GetById(searchCtx.Ctx, searchCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if invoice == nil {
		return nil, nil
	}

	return NewInvoiceResponseBuilder().
		AddId(invoice.Id).
		AddPayAt(invoice.PayAt).
		AddBuyAt(invoice.BuyAt).
		AddDescription(invoice.Description).
		AddValue(invoice.Value).
		AddPaymentType(PaymentTypeResponse{Id: invoice.PaymentType.Id, Type: invoice.PaymentType.Type}).
		AddCategory(CategoryResponse{Id: invoice.Category.Id, Category: invoice.Category.Category}).
		AddInvoiceProjectionId(invoice.InvoiceProjectionId).
		Build(), nil
}
