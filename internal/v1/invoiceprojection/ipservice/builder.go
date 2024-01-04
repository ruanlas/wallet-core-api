package ipservice

import "time"

type InvoiceProjectionResponseBuilder struct {
	id          string
	payIn       time.Time
	buyAt       time.Time
	description string
	value       float64
	recurrence  uint
	category    CategoryResponse
	paymentType PaymentTypeResponse
}

func NewInvoiceProjectionResponseBuilder() *InvoiceProjectionResponseBuilder {
	return &InvoiceProjectionResponseBuilder{}
}
func (builder *InvoiceProjectionResponseBuilder) AddId(id string) *InvoiceProjectionResponseBuilder {
	builder.id = id
	return builder
}
func (builder *InvoiceProjectionResponseBuilder) AddPayIn(payIn time.Time) *InvoiceProjectionResponseBuilder {
	builder.payIn = payIn
	return builder
}
func (builder *InvoiceProjectionResponseBuilder) AddBuyAt(buyAt time.Time) *InvoiceProjectionResponseBuilder {
	builder.buyAt = buyAt
	return builder
}
func (builder *InvoiceProjectionResponseBuilder) AddDescription(description string) *InvoiceProjectionResponseBuilder {
	builder.description = description
	return builder
}
func (builder *InvoiceProjectionResponseBuilder) AddValue(value float64) *InvoiceProjectionResponseBuilder {
	builder.value = value
	return builder
}
func (builder *InvoiceProjectionResponseBuilder) AddPaymentType(paymentType PaymentTypeResponse) *InvoiceProjectionResponseBuilder {
	builder.paymentType = paymentType
	return builder
}
func (builder *InvoiceProjectionResponseBuilder) AddRecurrence(recurrence uint) *InvoiceProjectionResponseBuilder {
	builder.recurrence = recurrence
	return builder
}
func (builder *InvoiceProjectionResponseBuilder) AddCategory(category CategoryResponse) *InvoiceProjectionResponseBuilder {
	builder.category = category
	return builder
}
func (builder *InvoiceProjectionResponseBuilder) Build() *InvoiceProjectionResponse {
	invoiceProjectionResponse := InvoiceProjectionResponse{}

	invoiceProjectionResponse.Id = builder.id
	invoiceProjectionResponse.Description = builder.description
	invoiceProjectionResponse.Value = builder.value
	invoiceProjectionResponse.PayIn = builder.payIn
	invoiceProjectionResponse.BuyAt = builder.buyAt
	invoiceProjectionResponse.PaymentType = builder.paymentType
	invoiceProjectionResponse.Recurrence = builder.recurrence
	invoiceProjectionResponse.Category = builder.category

	return &invoiceProjectionResponse
}

type SearchParamsBuilder struct {
	month    *uint
	year     *uint
	page     *uint
	pagesize *uint
}

func NewSearchParamsBuilder() *SearchParamsBuilder {
	return &SearchParamsBuilder{}
}

func (builder *SearchParamsBuilder) AddMonth(month uint) *SearchParamsBuilder {
	builder.month = &month
	return builder
}
func (builder *SearchParamsBuilder) AddYear(year uint) *SearchParamsBuilder {
	builder.year = &year
	return builder
}
func (builder *SearchParamsBuilder) AddPage(page uint) *SearchParamsBuilder {
	builder.page = &page
	return builder
}
func (builder *SearchParamsBuilder) AddPageSize(pagesize uint) *SearchParamsBuilder {
	builder.pagesize = &pagesize
	return builder
}
func (builder *SearchParamsBuilder) Build() *SearchParams {
	return &SearchParams{
		month: builder.month,
		year:  builder.year,
		paginate: &Paginate{
			page:     builder.page,
			pagesize: builder.pagesize,
		},
	}
}

type InvoiceResponseBuilder struct {
	id                  string
	payAt               time.Time
	buyAt               time.Time
	description         string
	value               float64
	invoiceProjectionId string
	category            CategoryResponse
	paymentType         PaymentTypeResponse
}

func NewInvoiceResponseBuilder() *InvoiceResponseBuilder {
	return &InvoiceResponseBuilder{}
}
func (builder *InvoiceResponseBuilder) AddId(id string) *InvoiceResponseBuilder {
	builder.id = id
	return builder
}
func (builder *InvoiceResponseBuilder) AddPayAt(payAt time.Time) *InvoiceResponseBuilder {
	builder.payAt = payAt
	return builder
}
func (builder *InvoiceResponseBuilder) AddBuyAt(buyAt time.Time) *InvoiceResponseBuilder {
	builder.buyAt = buyAt
	return builder
}
func (builder *InvoiceResponseBuilder) AddDescription(description string) *InvoiceResponseBuilder {
	builder.description = description
	return builder
}
func (builder *InvoiceResponseBuilder) AddValue(value float64) *InvoiceResponseBuilder {
	builder.value = value
	return builder
}
func (builder *InvoiceResponseBuilder) AddPaymentType(paymentType PaymentTypeResponse) *InvoiceResponseBuilder {
	builder.paymentType = paymentType
	return builder
}
func (builder *InvoiceResponseBuilder) AddInvoiceProjectionId(invoiceProjectionId string) *InvoiceResponseBuilder {
	builder.invoiceProjectionId = invoiceProjectionId
	return builder
}
func (builder *InvoiceResponseBuilder) AddCategory(category CategoryResponse) *InvoiceResponseBuilder {
	builder.category = category
	return builder
}
func (builder *InvoiceResponseBuilder) Build() *InvoiceResponse {
	invoiceResponse := InvoiceResponse{}

	invoiceResponse.Id = builder.id
	invoiceResponse.Description = builder.description
	invoiceResponse.Value = builder.value
	invoiceResponse.PayAt = builder.payAt
	invoiceResponse.PaymentType = builder.paymentType
	invoiceResponse.InvoiceProjectionId = builder.invoiceProjectionId
	invoiceResponse.Category = builder.category

	return &invoiceResponse
}
