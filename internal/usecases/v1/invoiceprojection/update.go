package invoiceprojection

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoiceprojection"
)

type UpdateUseCase interface {
	Execute(updateCtx UpdateContext) (*InvoiceProjectionResponse, error)
}

func NewUpdateUseCase(repository invoiceprojection.Repository) UpdateUseCase {
	return &updateUseCase{
		repository: repository,
	}
}

type updateUseCase struct {
	repository invoiceprojection.Repository
}

func (uc *updateUseCase) Execute(updateCtx UpdateContext) (*InvoiceProjectionResponse, error) {
	request := updateCtx.Request
	user := idpauth.GetUser(updateCtx.UserToken)
	invoiceProjectionBuilder := invoiceprojection.NewInvoiceProjectionBuilder().
		AddId(updateCtx.Id).
		AddPayIn(request.PayIn).
		AddBuyAt(request.BuyAt).
		AddPaymentType(invoiceprojection.PaymentType{Id: request.PaymentTypeId}).
		AddCategory(invoiceprojection.InvoiceCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value)
	invoiceProjectionExists, err := uc.repository.GetById(updateCtx.Ctx, updateCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if invoiceProjectionExists == nil {
		return nil, nil
	}
	invoiceProjectionBuilder.AddIsAlreadyDone(invoiceProjectionExists.IsAlreadyDone)
	invoiceProjectionBuilder.AddUserId(user.Id)
	invoiceProjectionUpdated, err := uc.repository.Edit(updateCtx.Ctx, *invoiceProjectionBuilder.Build())
	if err != nil {
		return nil, err
	}

	invoiceProjectionUpdated, err = uc.repository.GetById(updateCtx.Ctx, invoiceProjectionUpdated.Id, user.Id)
	if err != nil {
		return nil, err
	}

	return NewInvoiceProjectionResponseBuilder().
		AddId(invoiceProjectionUpdated.Id).
		AddPayIn(invoiceProjectionUpdated.PayIn).
		AddBuyAt(invoiceProjectionExists.BuyAt).
		AddDescription(invoiceProjectionUpdated.Description).
		AddValue(invoiceProjectionUpdated.Value).
		AddPaymentType(PaymentTypeResponse{Id: invoiceProjectionUpdated.PaymentType.Id, Type: invoiceProjectionUpdated.PaymentType.Type}).
		AddCategory(CategoryResponse{Id: invoiceProjectionUpdated.Category.Id, Category: invoiceProjectionUpdated.Category.Category}).
		Build(), nil
}
