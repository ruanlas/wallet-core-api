package invoiceprojection

import (
	"context"
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoiceprojection"
	uuid "github.com/satori/go.uuid"
)

type CreateUseCase interface {
	Execute(createCtx CreateContext) (*InvoiceProjectionResponse, error)
}

func NewCreateUseCase(repository invoiceprojection.Repository, generateUUID func() uuid.UUID) CreateUseCase {
	return &createUseCase{
		repository:   repository,
		generateUUID: generateUUID,
	}
}

type createUseCase struct {
	repository   invoiceprojection.Repository
	generateUUID func() uuid.UUID
}

func (uc *createUseCase) Execute(createCtx CreateContext) (*InvoiceProjectionResponse, error) {
	request := createCtx.Request
	user := idpauth.GetUser(createCtx.UserToken)
	createdAt := time.Now()
	invoiceProjectionBuilder := invoiceprojection.NewInvoiceProjectionBuilder().
		AddId(uc.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayIn(request.PayIn).
		AddBuyAt(request.BuyAt).
		AddPaymentType(invoiceprojection.PaymentType{Id: request.PaymentTypeId}).
		AddIsAlreadyDone(false).
		AddCategory(invoiceprojection.InvoiceCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value).
		AddUserId(user.Id)
	if request.BuyAt.IsZero() {
		invoiceProjectionBuilder.AddBuyAt(createdAt)
	}
	invoiceProjection := invoiceProjectionBuilder.Build()
	invoiceProjectionSaved, err := uc.repository.Save(createCtx.Ctx, *invoiceProjection)
	if err != nil {
		return nil, err
	}

	if request.Recurrence > 1 {
		err = uc.createRecurrence(createCtx.Ctx, request, createdAt, user.Id)
		if err != nil {
			return nil, err
		}
	} else {
		request.Recurrence = 1
	}
	invoiceProjectionSaved, err = uc.repository.GetById(createCtx.Ctx, invoiceProjectionSaved.Id, user.Id)
	if err != nil {
		return nil, err
	}

	return NewInvoiceProjectionResponseBuilder().
		AddId(invoiceProjection.Id).
		AddPayIn(invoiceProjection.PayIn).
		AddBuyAt(invoiceProjection.BuyAt).
		AddDescription(invoiceProjection.Description).
		AddValue(invoiceProjection.Value).
		AddPaymentType(PaymentTypeResponse{Id: invoiceProjectionSaved.PaymentType.Id, Type: invoiceProjectionSaved.PaymentType.Type}).
		AddCategory(CategoryResponse{Id: invoiceProjectionSaved.Category.Id, Category: invoiceProjectionSaved.Category.Category}).
		AddRecurrence(request.Recurrence).
		Build(), nil
}

func (uc *createUseCase) createRecurrence(ctx context.Context, request CreateRequest, createdAt time.Time, userId string) error {
	for i := 1; i < int(request.Recurrence+1); i++ {
		invoiceProjectionBuilder := invoiceprojection.NewInvoiceProjectionBuilder().
			AddId(uc.generateUUID().String()).
			AddCreatedAt(createdAt).
			AddPayIn(request.PayIn.AddDate(0, i, 0)).
			AddBuyAt(request.BuyAt).
			AddPaymentType(invoiceprojection.PaymentType{Id: request.PaymentTypeId}).
			AddIsAlreadyDone(false).
			AddCategory(invoiceprojection.InvoiceCategory{Id: request.CategoryId}).
			AddDescription(request.Description).
			AddValue(request.Value).
			AddUserId(userId)
		if request.BuyAt.IsZero() {
			invoiceProjectionBuilder.AddBuyAt(createdAt)
		}
		invoiceProjection := invoiceProjectionBuilder.Build()
		_, err := uc.repository.Save(ctx, *invoiceProjection)
		if err != nil {
			return err
		}
	}
	return nil
}
