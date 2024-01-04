package repository

import "time"

type QueryParamsBuilder struct {
	userId string
	limit  uint
	offset uint
	month  uint
	year   uint
}

func NewQueryParamsBuilder() *QueryParamsBuilder {
	return &QueryParamsBuilder{}
}
func (builder *QueryParamsBuilder) AddMonth(month uint) *QueryParamsBuilder {
	builder.month = month
	return builder
}
func (builder *QueryParamsBuilder) AddUserId(userId string) *QueryParamsBuilder {
	builder.userId = userId
	return builder
}
func (builder *QueryParamsBuilder) AddYear(year uint) *QueryParamsBuilder {
	builder.year = year
	return builder
}
func (builder *QueryParamsBuilder) AddLimit(limit uint) *QueryParamsBuilder {
	builder.limit = limit
	return builder
}
func (builder *QueryParamsBuilder) AddOffset(offset uint) *QueryParamsBuilder {
	builder.offset = offset
	return builder
}
func (builder *QueryParamsBuilder) Build() QueryParams {
	return QueryParams{
		userId: builder.userId,
		month:  builder.month,
		year:   builder.year,
		limit:  builder.limit,
		offset: builder.offset,
	}
}

type InvoiceBuilder struct {
	id                  string
	createdAt           time.Time
	buyAt               time.Time
	payAt               time.Time
	description         string
	value               float64
	userId              string
	category            InvoiceCategory
	paymentType         PaymentType
	invoiceProjectionId string
}

func NewInvoiceBuilder() *InvoiceBuilder {
	return &InvoiceBuilder{}
}
func (builder *InvoiceBuilder) AddId(id string) *InvoiceBuilder {
	builder.id = id
	return builder
}
func (builder *InvoiceBuilder) AddCreatedAt(createdAt time.Time) *InvoiceBuilder {
	builder.createdAt = createdAt
	return builder
}
func (builder *InvoiceBuilder) AddPayAt(payAt time.Time) *InvoiceBuilder {
	builder.payAt = payAt
	return builder
}
func (builder *InvoiceBuilder) AddBuyAt(buyAt time.Time) *InvoiceBuilder {
	builder.buyAt = buyAt
	return builder
}
func (builder *InvoiceBuilder) AddDescription(description string) *InvoiceBuilder {
	builder.description = description
	return builder
}
func (builder *InvoiceBuilder) AddValue(value float64) *InvoiceBuilder {
	builder.value = value
	return builder
}
func (builder *InvoiceBuilder) AddPaymentType(paymentType PaymentType) *InvoiceBuilder {
	builder.paymentType = paymentType
	return builder
}
func (builder *InvoiceBuilder) AddInvoiceProjectionId(invoiceProjectionId string) *InvoiceBuilder {
	builder.invoiceProjectionId = invoiceProjectionId
	return builder
}
func (builder *InvoiceBuilder) AddUserId(userId string) *InvoiceBuilder {
	builder.userId = userId
	return builder
}
func (builder *InvoiceBuilder) AddCategory(category InvoiceCategory) *InvoiceBuilder {
	builder.category = category
	return builder
}
func (builder *InvoiceBuilder) Build() *Invoice {
	invoice := Invoice{}

	invoice.Id = builder.id
	invoice.CreatedAt = builder.createdAt
	invoice.PayAt = builder.payAt
	invoice.BuyAt = builder.buyAt
	invoice.Description = builder.description
	invoice.Value = builder.value
	invoice.PaymentType = builder.paymentType
	invoice.InvoiceProjectionId = builder.invoiceProjectionId
	invoice.UserId = builder.userId
	invoice.Category = builder.category

	return &invoice
}
