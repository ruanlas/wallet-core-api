package invoiceprojection

import (
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoiceprojection"
	uuid "github.com/satori/go.uuid"
)

type CreateInvoiceUseCase interface {
	Execute(createInvoiceCtx CreateInvoiceContext) (*InvoiceStat, error)
}

func NewCreateInvoiceUseCase(repository invoiceprojection.Repository, generateUUID func() uuid.UUID) CreateInvoiceUseCase {
	return &createInvoiceUseCase{
		repository:   repository,
		generateUUID: generateUUID,
	}
}

type createInvoiceUseCase struct {
	repository   invoiceprojection.Repository
	generateUUID func() uuid.UUID
}

func (uc *createInvoiceUseCase) Execute(createInvoiceCtx CreateInvoiceContext) (*InvoiceStat, error) {
	request := createInvoiceCtx.Request
	user := idpauth.GetUser(createInvoiceCtx.UserToken)
	invoiceProjection, err := uc.repository.GetById(createInvoiceCtx.Ctx, createInvoiceCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if invoiceProjection == nil {
		return &InvoiceStat{ProjectionIsFound: false, ProjectionIsAlreadyDone: false}, nil
	}
	if invoiceProjection.IsAlreadyDone == true {
		return &InvoiceStat{ProjectionIsFound: true, ProjectionIsAlreadyDone: true}, nil
	}
	invoiceBuilder := invoiceprojection.NewInvoiceBuilder().
		AddId(uc.generateUUID().String()).
		AddCategory(invoiceProjection.Category).
		AddCreatedAt(time.Now()).
		AddDescription(invoiceProjection.Description).
		AddInvoiceProjectionId(invoiceProjection.Id).
		AddPaymentType(invoiceProjection.PaymentType).
		AddUserId(invoiceProjection.UserId).
		AddValue(invoiceProjection.Value).
		AddPayAt(invoiceProjection.PayIn).
		AddBuyAt(invoiceProjection.BuyAt)
	if request.Value != 0 {
		invoiceBuilder.AddValue(request.Value)
	}
	if !request.PayIn.IsZero() {
		invoiceBuilder.AddPayAt(request.PayIn)
	}
	if !request.BuyAt.IsZero() {
		invoiceBuilder.AddBuyAt(request.BuyAt)
	}
	invoice, err := uc.repository.SaveInvoice(createInvoiceCtx.Ctx, *invoiceBuilder.Build())
	if err != nil {
		return nil, err
	}
	invoiceProjection.IsAlreadyDone = true
	_, err = uc.repository.Edit(createInvoiceCtx.Ctx, *invoiceProjection)
	if err != nil {
		return nil, err
	}

	invoiceResponse := NewInvoiceResponseBuilder().
		AddId(invoice.Id).
		AddInvoiceProjectionId(invoice.InvoiceProjectionId).
		AddPayAt(invoice.PayAt).
		AddBuyAt(invoice.BuyAt).
		AddDescription(invoice.Description).
		AddValue(invoice.Value).
		AddPaymentType(PaymentTypeResponse{Id: invoice.PaymentType.Id, Type: invoice.PaymentType.Type}).
		AddCategory(CategoryResponse{Id: invoice.Category.Id, Category: invoice.Category.Category}).
		Build()
	return &InvoiceStat{ProjectionIsFound: true, ProjectionIsAlreadyDone: false, Invoice: invoiceResponse}, nil
}
