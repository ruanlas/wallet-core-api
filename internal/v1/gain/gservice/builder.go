package gservice

import "time"

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

type GainResponseBuilder struct {
	id               string
	payIn            time.Time
	description      string
	value            float64
	isPassive        bool
	gainProjectionId string
	category         CategoryResponse
}

func NewGainResponseBuilder() *GainResponseBuilder {
	return &GainResponseBuilder{}
}
func (builder *GainResponseBuilder) AddId(id string) *GainResponseBuilder {
	builder.id = id
	return builder
}
func (builder *GainResponseBuilder) AddPayIn(payIn time.Time) *GainResponseBuilder {
	builder.payIn = payIn
	return builder
}
func (builder *GainResponseBuilder) AddDescription(description string) *GainResponseBuilder {
	builder.description = description
	return builder
}
func (builder *GainResponseBuilder) AddValue(value float64) *GainResponseBuilder {
	builder.value = value
	return builder
}
func (builder *GainResponseBuilder) AddIsPassive(isPassive bool) *GainResponseBuilder {
	builder.isPassive = isPassive
	return builder
}
func (builder *GainResponseBuilder) AddGainProjectionId(gainProjectionId string) *GainResponseBuilder {
	builder.gainProjectionId = gainProjectionId
	return builder
}
func (builder *GainResponseBuilder) AddCategory(category CategoryResponse) *GainResponseBuilder {
	builder.category = category
	return builder
}
func (builder *GainResponseBuilder) Build() *GainResponse {
	gainResponse := GainResponse{}

	gainResponse.Id = builder.id
	gainResponse.Description = builder.description
	gainResponse.Value = builder.value
	gainResponse.PayIn = builder.payIn
	gainResponse.IsPassive = builder.isPassive
	gainResponse.GainProjectionId = builder.gainProjectionId
	gainResponse.Category = builder.category

	return &gainResponse
}
