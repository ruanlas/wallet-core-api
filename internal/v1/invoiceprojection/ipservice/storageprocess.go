package ipservice

import (
	"context"
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/v1/invoiceprojection/repository"
	uuid "github.com/satori/go.uuid"
)

type StorageProcess interface {
	Create(createCtx CreateContext) (*InvoiceProjectionResponse, error)
	Update(updateCtx UpdateContext) (*InvoiceProjectionResponse, error)
	Delete(searchCtx SearchContext) error
	CreateInvoice(createInvoiceCtx CreateInvoiceContext) (*InvoiceStat, error)
}

type storageProcess struct {
	repository   repository.Repository
	generateUUID func() uuid.UUID
}

func NewStorageProcess(repository repository.Repository, generateUUID func() uuid.UUID) StorageProcess {
	return &storageProcess{repository: repository, generateUUID: generateUUID}
}

func (sp *storageProcess) Create(createCtx CreateContext) (*InvoiceProjectionResponse, error) {
	request := createCtx.Request
	user := idpauth.GetUser(createCtx.UserToken)
	createdAt := time.Now()
	invoiceProjectionBuilder := repository.NewInvoiceProjectionBuilder().
		AddId(sp.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayIn(request.PayIn).
		AddBuyAt(request.BuyAt).
		AddPaymentType(repository.PaymentType{Id: request.PaymentTypeId}).
		AddIsAlreadyDone(false).
		AddCategory(repository.InvoiceCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value).
		AddUserId(user.Id)
	if request.BuyAt.IsZero() {
		invoiceProjectionBuilder.AddBuyAt(createdAt)
	}
	invoiceProjection := invoiceProjectionBuilder.Build()
	invoiceProjectionSaved, err := sp.repository.Save(createCtx.Ctx, *invoiceProjection)
	if err != nil {
		return nil, err
	}

	if request.Recurrence > 1 {
		err = sp.createRecurrence(createCtx.Ctx, request, createdAt, user.Id)
		if err != nil {
			return nil, err
		}
	} else {
		request.Recurrence = 1
	}
	invoiceProjectionSaved, err = sp.repository.GetById(createCtx.Ctx, invoiceProjectionSaved.Id, user.Id)
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

func (sp *storageProcess) createRecurrence(ctx context.Context, request CreateRequest, createdAt time.Time, userId string) error {
	for i := 1; i < int(request.Recurrence+1); i++ {
		invoiceProjectionBuilder := repository.NewInvoiceProjectionBuilder().
			AddId(sp.generateUUID().String()).
			AddCreatedAt(createdAt).
			AddPayIn(request.PayIn.AddDate(0, i, 0)).
			AddBuyAt(request.BuyAt).
			AddPaymentType(repository.PaymentType{Id: request.PaymentTypeId}).
			AddIsAlreadyDone(false).
			AddCategory(repository.InvoiceCategory{Id: request.CategoryId}).
			AddDescription(request.Description).
			AddValue(request.Value).
			AddUserId(userId)
		if request.BuyAt.IsZero() {
			invoiceProjectionBuilder.AddBuyAt(createdAt)
		}
		invoiceProjection := invoiceProjectionBuilder.Build()
		_, err := sp.repository.Save(ctx, *invoiceProjection)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sp *storageProcess) Update(updateCtx UpdateContext) (*InvoiceProjectionResponse, error) {
	request := updateCtx.Request
	user := idpauth.GetUser(updateCtx.UserToken)
	invoiceProjectionBuilder := repository.NewInvoiceProjectionBuilder().
		AddId(updateCtx.Id).
		AddPayIn(request.PayIn).
		AddBuyAt(request.BuyAt).
		AddPaymentType(repository.PaymentType{Id: request.PaymentTypeId}).
		AddCategory(repository.InvoiceCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value)
	invoiceProjectionExists, err := sp.repository.GetById(updateCtx.Ctx, updateCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if invoiceProjectionExists == nil {
		return nil, nil
	}
	invoiceProjectionBuilder.AddIsAlreadyDone(invoiceProjectionExists.IsAlreadyDone)
	invoiceProjectionBuilder.AddUserId(user.Id)
	invoiceProjectionUpdated, err := sp.repository.Edit(updateCtx.Ctx, *invoiceProjectionBuilder.Build())
	if err != nil {
		return nil, err
	}

	invoiceProjectionUpdated, err = sp.repository.GetById(updateCtx.Ctx, invoiceProjectionUpdated.Id, user.Id)
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

func (sp *storageProcess) Delete(searchCtx SearchContext) error {
	user := idpauth.GetUser(searchCtx.UserToken)
	return sp.repository.Remove(searchCtx.Ctx, searchCtx.Id, user.Id)
}

func (sp *storageProcess) CreateInvoice(createInvoiceCtx CreateInvoiceContext) (*InvoiceStat, error) {
	request := createInvoiceCtx.Request
	user := idpauth.GetUser(createInvoiceCtx.UserToken)
	invoiceProjection, err := sp.repository.GetById(createInvoiceCtx.Ctx, createInvoiceCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if invoiceProjection == nil {
		return &InvoiceStat{ProjectionIsFound: false, ProjectionIsAlreadyDone: false}, nil
	}
	if invoiceProjection.IsAlreadyDone == true {
		return &InvoiceStat{ProjectionIsFound: true, ProjectionIsAlreadyDone: true}, nil
	}
	invoiceBuilder := repository.NewInvoiceBuilder().
		AddId(sp.generateUUID().String()).
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
	invoice, err := sp.repository.SaveInvoice(createInvoiceCtx.Ctx, *invoiceBuilder.Build())
	if err != nil {
		return nil, err
	}
	invoiceProjection.IsAlreadyDone = true
	_, err = sp.repository.Edit(createInvoiceCtx.Ctx, *invoiceProjection)
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
