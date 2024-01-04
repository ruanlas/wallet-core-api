package iservice

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/v1/invoice/repository"
)

type ReadingProcess interface {
	GetById(searchCtx SearchContext) (*InvoiceResponse, error)
	GetAllPaginated(searchCtx SearchContext) (*InvoicePaginateResponse, error)
}

type readingProcess struct {
	repository repository.Repository
}

func NewReadingProcess(repository repository.Repository) ReadingProcess {
	return &readingProcess{repository: repository}
}

func (rp *readingProcess) GetById(searchCtx SearchContext) (*InvoiceResponse, error) {
	user := idpauth.GetUser(searchCtx.UserToken)
	invoice, err := rp.repository.GetById(searchCtx.Ctx, searchCtx.Id, user.Id)
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

func (rp *readingProcess) getOffset(actualPage uint, pagesize uint) uint {
	return (actualPage - 1) * pagesize
}

func (rp *readingProcess) getTotalPages(totalRecords uint, pagesize uint) uint {
	totalPages := totalRecords / pagesize
	if (totalRecords % pagesize) > 0 {
		totalPages++
	}
	return totalPages
}

func (rp *readingProcess) GetAllPaginated(searchCtx SearchContext) (*InvoicePaginateResponse, error) {
	search := searchCtx.Params
	user := idpauth.GetUser(searchCtx.UserToken)
	offset := rp.getOffset(*search.paginate.page, *search.paginate.pagesize)
	queryParam := repository.NewQueryParamsBuilder().
		AddMonth(*search.month).
		AddYear(*search.year).
		AddUserId(user.Id).
		AddOffset(offset).
		AddLimit(*search.paginate.pagesize).
		Build()

	totalRecords, err := rp.repository.GetTotalRecords(searchCtx.Ctx, queryParam)
	if err != nil {
		return nil, err
	}
	totalPages := rp.getTotalPages(*totalRecords, *search.paginate.pagesize)
	invoiceList, err := rp.repository.GetAll(searchCtx.Ctx, queryParam)
	if err != nil {
		return nil, err
	}

	var invoiceResponseList []InvoiceResponse
	for _, invoice := range *invoiceList {
		category := CategoryResponse{
			Id:       invoice.Category.Id,
			Category: invoice.Category.Category,
		}
		paymentType := PaymentTypeResponse{
			Id:   invoice.PaymentType.Id,
			Type: invoice.PaymentType.Type,
		}
		invoiceResponse := NewInvoiceResponseBuilder().
			AddId(invoice.Id).
			AddCategory(category).
			AddDescription(invoice.Description).
			AddPaymentType(paymentType).
			AddPayAt(invoice.PayAt).
			AddBuyAt(invoice.BuyAt).
			AddValue(invoice.Value).
			AddInvoiceProjectionId(invoice.InvoiceProjectionId).
			Build()
		invoiceResponseList = append(invoiceResponseList, *invoiceResponse)
	}

	return &InvoicePaginateResponse{
		CurrentPage:  *search.paginate.page,
		PageLimit:    *search.paginate.pagesize,
		TotalRecords: *totalRecords,
		TotalPages:   totalPages,
		Records:      invoiceResponseList,
	}, nil
}
