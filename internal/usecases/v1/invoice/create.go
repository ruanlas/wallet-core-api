package invoice

import (
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoice"
	uuid "github.com/satori/go.uuid"
)

type CreateUseCase interface {
	Execute(createCtx CreateContext) (*InvoiceResponse, error)
}

func NewCreateUseCase(repository invoice.Repository, generateUUID func() uuid.UUID) CreateUseCase {
	return &createUseCase{
		repository:   repository,
		generateUUID: generateUUID,
	}
}

type createUseCase struct {
	repository   invoice.Repository
	generateUUID func() uuid.UUID
}

func (uc *createUseCase) Execute(createCtx CreateContext) (*InvoiceResponse, error) {
	request := createCtx.Request
	user := idpauth.GetUser(createCtx.UserToken)
	createdAt := time.Now()
	invoiceBuilder := invoice.NewInvoiceBuilder().
		AddId(uc.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayAt(request.PayAt).
		AddBuyAt(request.BuyAt).
		AddPaymentType(invoice.PaymentType{Id: request.PaymentTypeId}).
		AddCategory(invoice.InvoiceCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value).
		AddUserId(user.Id)
	if request.PayAt.IsZero() {
		request.PayAt = createdAt
		invoiceBuilder.AddPayAt(request.PayAt)
	}
	if request.BuyAt.IsZero() {
		invoiceBuilder.AddBuyAt(request.PayAt)
	}
	invoice := invoiceBuilder.Build()
	invoiceSaved, err := uc.repository.Save(createCtx.Ctx, *invoice)
	if err != nil {
		return nil, err
	}

	invoiceSaved, err = uc.repository.GetById(createCtx.Ctx, invoiceSaved.Id, user.Id)
	if err != nil {
		return nil, err
	}

	return NewInvoiceResponseBuilder().
		AddId(invoice.Id).
		AddPayAt(invoice.PayAt).
		AddBuyAt(invoice.BuyAt).
		AddDescription(invoice.Description).
		AddValue(invoice.Value).
		AddPaymentType(PaymentTypeResponse{Id: invoiceSaved.PaymentType.Id, Type: invoiceSaved.PaymentType.Type}).
		AddCategory(CategoryResponse{Id: invoiceSaved.Category.Id, Category: invoiceSaved.Category.Category}).
		Build(), nil
}
