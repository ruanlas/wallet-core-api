package gservice

import (
	"context"
	"time"
)

type CreateContext struct {
	Ctx       context.Context
	Request   CreateRequest
	UserToken string
}

type UpdateContext struct {
	Ctx       context.Context
	Request   UpdateRequest
	UserToken string
	Id        string
}

type SearchContext struct {
	Ctx       context.Context
	Params    SearchParams
	UserToken string
	Id        string
}

type CreateRequest struct {
	PayIn       time.Time `json:"pay_in"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	IsPassive   bool      `json:"is_passive"`
	CategoryId  uint      `json:"category_id"`
}

type UpdateRequest struct {
	PayIn       time.Time `json:"pay_in"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	IsPassive   bool      `json:"is_passive"`
	CategoryId  uint      `json:"category_id"`
}

type CreateGainRequest struct {
	Value float64   `json:"value"`
	PayIn time.Time `json:"pay_in"`
}

type CategoryResponse struct {
	Id       uint   `json:"id"`
	Category string `json:"category"`
}

type GainResponse struct {
	Id               string           `json:"id"`
	GainProjectionId string           `json:"gain_projection_id,omitempty"`
	PayIn            time.Time        `json:"pay_in"`
	Description      string           `json:"description"`
	Value            float64          `json:"value"`
	IsPassive        bool             `json:"is_passive"`
	Category         CategoryResponse `json:"category"`
}

type GainStat struct {
	ProjectionIsFound       bool
	ProjectionIsAlreadyDone bool
	Gain                    *GainResponse
}

type GainPaginateResponse struct {
	CurrentPage  uint           `json:"current_page"`
	TotalPages   uint           `json:"total_pages"`
	TotalRecords uint           `json:"total_records"`
	PageLimit    uint           `json:"page_limit"`
	Records      []GainResponse `json:"records"`
}

type Paginate struct {
	page     *uint
	pagesize *uint
}

type SearchParams struct {
	month    *uint
	year     *uint
	paginate *Paginate
}
