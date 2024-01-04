package iservice

import (
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/v1/invoice/repository"
	uuid "github.com/satori/go.uuid"
)

type StorageProcess interface {
	Create(createCtx CreateContext) (*InvoiceResponse, error)
	Update(updateCtx UpdateContext) (*InvoiceResponse, error)
	Delete(searchCtx SearchContext) error
}

type storageProcess struct {
	repository   repository.Repository
	generateUUID func() uuid.UUID
}

func NewStorageProcess(repository repository.Repository, generateUUID func() uuid.UUID) StorageProcess {
	return &storageProcess{repository: repository, generateUUID: generateUUID}
}

func (sp *storageProcess) Create(createCtx CreateContext) (*InvoiceResponse, error) {
	request := createCtx.Request
	user := idpauth.GetUser(createCtx.UserToken)
	createdAt := time.Now()
	invoiceBuilder := repository.NewInvoiceBuilder().
		AddId(sp.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayAt(request.PayAt).
		AddBuyAt(request.BuyAt).
		AddPaymentType(repository.PaymentType{Id: request.PaymentTypeId}).
		AddCategory(repository.InvoiceCategory{Id: request.CategoryId}).
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
	invoiceSaved, err := sp.repository.Save(createCtx.Ctx, *invoice)
	if err != nil {
		return nil, err
	}

	invoiceSaved, err = sp.repository.GetById(createCtx.Ctx, invoiceSaved.Id, user.Id)
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

func (sp *storageProcess) Update(updateCtx UpdateContext) (*InvoiceResponse, error) {
	request := updateCtx.Request
	user := idpauth.GetUser(updateCtx.UserToken)
	invoiceBuilder := repository.NewInvoiceBuilder().
		AddId(updateCtx.Id).
		AddPayAt(request.PayAt).
		AddBuyAt(request.BuyAt).
		AddPaymentType(repository.PaymentType{Id: request.PaymentTypeId}).
		AddCategory(repository.InvoiceCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value)
	invoiceExists, err := sp.repository.GetById(updateCtx.Ctx, updateCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if invoiceExists == nil {
		return nil, nil
	}
	invoiceBuilder.AddUserId(user.Id)
	invoiceUpdated, err := sp.repository.Edit(updateCtx.Ctx, *invoiceBuilder.Build())
	if err != nil {
		return nil, err
	}

	invoiceUpdated, err = sp.repository.GetById(updateCtx.Ctx, invoiceUpdated.Id, user.Id)
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

func (sp *storageProcess) Delete(searchCtx SearchContext) error {
	user := idpauth.GetUser(searchCtx.UserToken)
	return sp.repository.Remove(searchCtx.Ctx, searchCtx.Id, user.Id)
}
