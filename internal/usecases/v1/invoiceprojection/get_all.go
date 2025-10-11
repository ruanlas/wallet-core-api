package invoiceprojection

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/invoiceprojection"
)

type GetAllUseCase interface {
	Execute(searchCtx SearchContext) (*InvoiceProjectionPaginateResponse, error)
}

func NewGetAllUseCase(repository invoiceprojection.Repository) GetAllUseCase {
	return &getAllUseCase{
		repository: repository,
	}
}

type getAllUseCase struct {
	repository invoiceprojection.Repository
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

func (uc *getAllUseCase) Execute(searchCtx SearchContext) (*InvoiceProjectionPaginateResponse, error) {
	search := searchCtx.Params
	user := idpauth.GetUser(searchCtx.UserToken)
	offset := uc.getOffset(*search.paginate.page, *search.paginate.pagesize)
	queryParam := invoiceprojection.NewQueryParamsBuilder().
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
	invoiceProjectionList, err := uc.repository.GetAll(searchCtx.Ctx, queryParam)
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
