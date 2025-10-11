package gainprojection

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gainprojection"
)

type GetAllUseCase interface {
	Execute(searchCtx SearchContext) (*GainProjectionPaginateResponse, error)
}

type getAllUseCase struct {
	repository gainprojection.Repository
}

func NewGetAllUseCase(repository gainprojection.Repository) GetAllUseCase {
	return &getAllUseCase{
		repository: repository,
	}
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

func (uc *getAllUseCase) Execute(searchCtx SearchContext) (*GainProjectionPaginateResponse, error) {
	search := searchCtx.Params
	user := idpauth.GetUser(searchCtx.UserToken)
	offset := uc.getOffset(*search.paginate.page, *search.paginate.pagesize)
	queryParam := gainprojection.NewQueryParamsBuilder().
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
	gainProjectionList, err := uc.repository.GetAll(searchCtx.Ctx, queryParam)
	if err != nil {
		return nil, err
	}

	var gainProjectionResponseList []GainProjectionResponse
	for _, gainProjection := range *gainProjectionList {
		category := CategoryResponse{
			Id:       gainProjection.Category.Id,
			Category: gainProjection.Category.Category,
		}
		gainProjectionResponse := NewGainProjectionResponseBuilder().
			AddId(gainProjection.Id).
			AddCategory(category).
			AddDescription(gainProjection.Description).
			AddIsPassive(gainProjection.IsPassive).
			AddPayIn(gainProjection.PayIn).
			AddValue(gainProjection.Value).
			Build()
		gainProjectionResponseList = append(gainProjectionResponseList, *gainProjectionResponse)
	}

	return &GainProjectionPaginateResponse{
		CurrentPage:  *search.paginate.page,
		PageLimit:    *search.paginate.pagesize,
		TotalRecords: *totalRecords,
		TotalPages:   totalPages,
		Records:      gainProjectionResponseList,
	}, nil
}
