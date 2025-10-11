package invoice

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoice"
)

type GetAllUseCase interface {
	Execute(searchCtx SearchContext) (*InvoicePaginateResponse, error)
}

func NewGetAllUseCase(repository invoice.Repository) GetAllUseCase {
	return &getAllUseCase{repository: repository}
}

type getAllUseCase struct {
	repository invoice.Repository
}

func (uc *getAllUseCase) getOffset(actualPage uint, pagesize uint) uint {
	return (actualPage - 1) * pagesize
}

func (uc *getAllUseCase) getTotalPages(totalRecords uint, pagesize uint) uint {
	totalPages := totalRecords / pagesize
	if (totalRecords % pagesize) > 0 {
		totalPages++
	}
	return totalPages
}

func (uc *getAllUseCase) Execute(searchCtx SearchContext) (*InvoicePaginateResponse, error) {
	search := searchCtx.Params
	user := idpauth.GetUser(searchCtx.UserToken)
	offset := uc.getOffset(*search.paginate.page, *search.paginate.pagesize)
	queryParam := invoice.NewQueryParamsBuilder().
		AddMonth(*search.month).
		AddYear(*search.year).
		AddUserId(user.Id).
		AddOffset(offset).
		AddLimit(*search.paginate.pagesize).
		Build()

	totalRecords, err := uc.repository.GetTotalRecords(searchCtx.Ctx, queryParam)
	if err != nil {
		return nil, err
	}
	totalPages := uc.getTotalPages(*totalRecords, *search.paginate.pagesize)
	invoiceList, err := uc.repository.GetAll(searchCtx.Ctx, queryParam)
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
