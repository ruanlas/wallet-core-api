package service

import (
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/repository"
)

type ReadingProcess interface {
	GetById(searchCtx SearchContext) (*GainProjectionResponse, error)
	GetAllPaginated(searchCtx SearchContext) (*GainProjectionPaginateResponse, error)
}

type readingProcess struct {
	repository repository.Repository
}

func NewReadingProcess(repository repository.Repository) ReadingProcess {
	return &readingProcess{repository: repository}
}

func (rp *readingProcess) GetById(searchCtx SearchContext) (*GainProjectionResponse, error) {
	user := idpauth.GetUser(searchCtx.UserToken)
	gainProjection, err := rp.repository.GetById(searchCtx.Ctx, searchCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gainProjection == nil {
		return nil, nil
	}

	return NewGainProjectionResponseBuilder().
		AddId(gainProjection.Id).
		AddPayIn(gainProjection.PayIn).
		AddDescription(gainProjection.Description).
		AddValue(gainProjection.Value).
		AddIsPassive(gainProjection.IsPassive).
		AddCategory(CategoryResponse{Id: gainProjection.Category.Id, Category: gainProjection.Category.Category}).
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

func (rp *readingProcess) GetAllPaginated(searchCtx SearchContext) (*GainProjectionPaginateResponse, error) {
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
	gainProjectionList, err := rp.repository.GetAll(searchCtx.Ctx, queryParam)
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
