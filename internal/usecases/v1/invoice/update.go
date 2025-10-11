package invoice

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoice"
)

type UpdateUseCase interface {
	Execute(updateCtx UpdateContext) (*InvoiceResponse, error)
}

func NewUpdateUseCase(repository invoice.Repository) UpdateUseCase {
	return &updateUseCase{
		repository: repository,
	}
}

type updateUseCase struct {
	repository invoice.Repository
}

func (uc *updateUseCase) Execute(updateCtx UpdateContext) (*InvoiceResponse, error) {
	request := updateCtx.Request
	user := idpauth.GetUser(updateCtx.UserToken)
	invoiceBuilder := invoice.NewInvoiceBuilder().
		AddId(updateCtx.Id).
		AddPayAt(request.PayAt).
		AddBuyAt(request.BuyAt).
		AddPaymentType(invoice.PaymentType{Id: request.PaymentTypeId}).
		AddCategory(invoice.InvoiceCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value)
	invoiceExists, err := uc.repository.GetById(updateCtx.Ctx, updateCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if invoiceExists == nil {
		return nil, nil
	}
	invoiceBuilder.AddUserId(user.Id)
	invoiceUpdated, err := uc.repository.Edit(updateCtx.Ctx, *invoiceBuilder.Build())
	if err != nil {
		return nil, err
	}

	invoiceUpdated, err = uc.repository.GetById(updateCtx.Ctx, invoiceUpdated.Id, user.Id)
	if err != nil {
		return nil, err
	}

	return NewInvoiceResponseBuilder().
		AddId(invoiceUpdated.Id).
		AddPayAt(invoiceUpdated.PayAt).
		AddBuyAt(invoiceExists.BuyAt).
		AddDescription(invoiceUpdated.Description).
		AddValue(invoiceUpdated.Value).
		AddPaymentType(PaymentTypeResponse{Id: invoiceUpdated.PaymentType.Id, Type: invoiceUpdated.PaymentType.Type}).
		AddCategory(CategoryResponse{Id: invoiceUpdated.Category.Id, Category: invoiceUpdated.Category.Category}).
		Build(), nil
}
