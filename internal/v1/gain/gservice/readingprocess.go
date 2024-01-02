package gservice

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/v1/gain/repository"
)

type ReadingProcess interface {
	GetById(searchCtx SearchContext) (*GainResponse, error)
	GetAllPaginated(searchCtx SearchContext) (*GainPaginateResponse, error)
}

type readingProcess struct {
	repository repository.Repository
}

func NewReadingProcess(repository repository.Repository) ReadingProcess {
	return &readingProcess{repository: repository}
}

func (rp *readingProcess) GetById(searchCtx SearchContext) (*GainResponse, error) {
	user := idpauth.GetUser(searchCtx.UserToken)
	gain, err := rp.repository.GetById(searchCtx.Ctx, searchCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gain == nil {
		return nil, nil
	}

	return NewGainResponseBuilder().
		AddId(gain.Id).
		AddPayIn(gain.PayIn).
		AddDescription(gain.Description).
		AddValue(gain.Value).
		AddIsPassive(gain.IsPassive).
		AddGainProjectionId(gain.GainProjectionId).
		AddCategory(CategoryResponse{Id: gain.Category.Id, Category: gain.Category.Category}).
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

func (rp *readingProcess) GetAllPaginated(searchCtx SearchContext) (*GainPaginateResponse, error) {
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
	gainList, err := rp.repository.GetAll(searchCtx.Ctx, queryParam)
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
