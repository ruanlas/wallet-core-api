package ipservice

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/v1/invoiceprojection/repository"
)

type ReadingProcess interface {
	GetById(searchCtx SearchContext) (*InvoiceProjectionResponse, error)
	GetAllPaginated(searchCtx SearchContext) (*InvoiceProjectionPaginateResponse, error)
}

type readingProcess struct {
	repository repository.Repository
}

func NewReadingProcess(repository repository.Repository) ReadingProcess {
	return &readingProcess{repository: repository}
}

func (rp *readingProcess) GetById(searchCtx SearchContext) (*InvoiceProjectionResponse, error) {
	user := idpauth.GetUser(searchCtx.UserToken)
	invoiceProjection, err := rp.repository.GetById(searchCtx.Ctx, searchCtx.Id, user.Id)
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

func (rp *readingProcess) GetAllPaginated(searchCtx SearchContext) (*InvoiceProjectionPaginateResponse, error) {
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
	invoiceProjectionList, err := rp.repository.GetAll(searchCtx.Ctx, queryParam)
	if err != nil {
		return nil, err
	}

	var invoiceProjectionResponseList []InvoiceProjectionResponse
	for _, invoiceProjection := range *invoiceProjectionList {
		category := CategoryResponse{
			Id:       invoiceProjection.Category.Id,
			Category: invoiceProjection.Category.Category,
		}
		paymentType := PaymentTypeResponse{
			Id:   invoiceProjection.PaymentType.Id,
			Type: invoiceProjection.PaymentType.Type,
		}
		invoiceProjectionResponse := NewInvoiceProjectionResponseBuilder().
			AddId(invoiceProjection.Id).
			AddCategory(category).
			AddDescription(invoiceProjection.Description).
			AddPaymentType(paymentType).
			AddPayIn(invoiceProjection.PayIn).
			AddBuyAt(invoiceProjection.BuyAt).
			AddValue(invoiceProjection.Value).
			Build()
		invoiceProjectionResponseList = append(invoiceProjectionResponseList, *invoiceProjectionResponse)
	}

	return &InvoiceProjectionPaginateResponse{
		CurrentPage:  *search.paginate.page,
		PageLimit:    *search.paginate.pagesize,
		TotalRecords: *totalRecords,
		TotalPages:   totalPages,
		Records:      invoiceProjectionResponseList,
	}, nil
}
