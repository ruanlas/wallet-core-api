package ipservice

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

type CreateInvoiceContext struct {
	Ctx       context.Context
	Request   CreateInvoiceRequest
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
	PayIn         time.Time `json:"pay_in"`
	BuyAt         time.Time `json:"buy_at"`
	Description   string    `json:"description"`
	Value         float64   `json:"value"`
	Recurrence    uint      `json:"recurrence"`
	CategoryId    uint      `json:"category_id"`
	PaymentTypeId uint      `json:"payment_type_id"`
}

type UpdateRequest struct {
	PayIn         time.Time `json:"pay_in"`
	BuyAt         time.Time `json:"buy_at"`
	Description   string    `json:"description"`
	Value         float64   `json:"value"`
	CategoryId    uint      `json:"category_id"`
	PaymentTypeId uint      `json:"payment_type_id"`
}

type CreateInvoiceRequest struct {
	Value float64   `json:"value"`
	PayIn time.Time `json:"pay_in"`
	BuyAt time.Time `json:"buy_at"`
}

type InvoiceProjectionResponse struct {
	Id          string              `json:"id"`
	PayIn       time.Time           `json:"pay_in"`
	BuyAt       time.Time           `json:"buy_at"`
	Description string              `json:"description"`
	Value       float64             `json:"value"`
	Recurrence  uint                `json:"recurrence,omitempty"`
	Category    CategoryResponse    `json:"category"`
	PaymentType PaymentTypeResponse `json:"payment_type"`
}

type CategoryResponse struct {
	Id       uint   `json:"id"`
	Category string `json:"category"`
}
type PaymentTypeResponse struct {
	Id   uint   `json:"id"`
	Type string `json:"type"`
}

type InvoiceResponse struct {
	Id                  string              `json:"id"`
	InvoiceProjectionId string              `json:"invoice_projection_id"`
	PayAt               time.Time           `json:"pay_at"`
	BuyAt               time.Time           `json:"buy_at"`
	Description         string              `json:"description"`
	Value               float64             `json:"value"`
	Category            CategoryResponse    `json:"category"`
	PaymentType         PaymentTypeResponse `json:"payment_type"`
}

type InvoiceStat struct {
	ProjectionIsFound       bool
	ProjectionIsAlreadyDone bool
	Invoice                 *InvoiceResponse
}

type InvoiceProjectionPaginateResponse struct {
	CurrentPage  uint                        `json:"current_page"`
	TotalPages   uint                        `json:"total_pages"`
	TotalRecords uint                        `json:"total_records"`
	PageLimit    uint                        `json:"page_limit"`
	Records      []InvoiceProjectionResponse `json:"records"`
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
