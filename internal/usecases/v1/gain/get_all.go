package gain

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gain"
)

type GetAllUseCase interface {
	Execute(searchCtx SearchContext) (*GainPaginateResponse, error)
}

func NewGetAllUseCase(repository gain.Repository) GetAllUseCase {
	return &getAllUseCase{
		repository: repository,
	}
}

type getAllUseCase struct {
	repository gain.Repository
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

func (uc *getAllUseCase) Execute(searchCtx SearchContext) (*GainPaginateResponse, error) {
	search := searchCtx.Params
	user := idpauth.GetUser(searchCtx.UserToken)
	offset := uc.getOffset(*search.paginate.page, *search.paginate.pagesize)
	queryParam := gain.NewQueryParamsBuilder().
		AddMonth(*search.month).
		AddYear(*search.year).
		AddUserId(user.Id).
		// AddUserId("user1").
		AddOffset(offset).
		AddLimit(*search.paginate.pagesize).
		Build()

	totalRecords, err := uc.repository.GetTotalRecords(searchCtx.Ctx, queryParam)
	if err != nil {
		return nil, err
	}
	totalPages := uc.getTotalPages(*totalRecords, *search.paginate.pagesize)
	gainList, err := uc.repository.GetAll(searchCtx.Ctx, queryParam)
	if err != nil {
		return nil, err
	}

	var gainResponseList []GainResponse
	for _, gain := range *gainList {
		category := CategoryResponse{
			Id:       gain.Category.Id,
			Category: gain.Category.Category,
		}
		GainResponse := NewGainResponseBuilder().
			AddId(gain.Id).
			AddCategory(category).
			AddDescription(gain.Description).
			AddIsPassive(gain.IsPassive).
			AddPayIn(gain.PayIn).
			AddValue(gain.Value).
			Build()
		gainResponseList = append(gainResponseList, *GainResponse)
	}

	return &GainPaginateResponse{
		CurrentPage:  *search.paginate.page,
		PageLimit:    *search.paginate.pagesize,
		TotalRecords: *totalRecords,
		TotalPages:   totalPages,
		Records:      gainResponseList,
	}, nil
}
