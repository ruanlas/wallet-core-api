package service

import "time"

type CreateRequest struct {
	PayIn       time.Time `json:"pay_in"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	IsPassive   bool      `json:"is_passive"`
	Recurrence  uint      `json:"recurrence"`
	CategoryId  uint      `json:"category_id"`
}

type UpdateRequest struct {
	PayIn       time.Time `json:"pay_in"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	IsPassive   bool      `json:"is_passive"`
	CategoryId  uint      `json:"category_id"`
}

type GainProjectionResponse struct {
	Id          string           `json:"id"`
	PayIn       time.Time        `json:"pay_in"`
	Description string           `json:"description"`
	Value       float64          `json:"value"`
	IsPassive   bool             `json:"is_passive"`
	Recurrence  uint             `json:"recurrence,omitempty"`
	Category    CategoryResponse `json:"category"`
}

type CategoryResponse struct {
	Id       uint   `json:"id"`
	Category string `json:"category"`
}

type GainProjectionPaginateResponse struct {
	CurrentPage  uint                     `json:"current_page"`
	TotalPages   uint                     `json:"total_pages"`
	TotalRecords uint                     `json:"total_records"`
	PageLimit    uint                     `json:"page_limit"`
	Records      []GainProjectionResponse `json:"records"`
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
