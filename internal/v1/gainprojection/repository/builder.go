package repository

import "time"

type GainProjectionBuilder struct {
	id          string
	createdAt   time.Time
	payIn       time.Time
	description string
	value       float64
	isPassive   bool
	isDone      bool
	userId      string
	category    GainCategory
}

func NewGainProjectionBuilder() *GainProjectionBuilder {
	return &GainProjectionBuilder{}
}
func (builder *GainProjectionBuilder) AddId(id string) *GainProjectionBuilder {
	builder.id = id
	return builder
}
func (builder *GainProjectionBuilder) AddCreatedAt(createdAt time.Time) *GainProjectionBuilder {
	builder.createdAt = createdAt
	return builder
}
func (builder *GainProjectionBuilder) AddPayIn(payIn time.Time) *GainProjectionBuilder {
	builder.payIn = payIn
	return builder
}
func (builder *GainProjectionBuilder) AddDescription(description string) *GainProjectionBuilder {
	builder.description = description
	return builder
}
func (builder *GainProjectionBuilder) AddValue(value float64) *GainProjectionBuilder {
	builder.value = value
	return builder
}
func (builder *GainProjectionBuilder) AddIsPassive(isPassive bool) *GainProjectionBuilder {
	builder.isPassive = isPassive
	return builder
}
func (builder *GainProjectionBuilder) AddIsDone(isDone bool) *GainProjectionBuilder {
	builder.isDone = isDone
	return builder
}
func (builder *GainProjectionBuilder) AddUserId(userId string) *GainProjectionBuilder {
	builder.userId = userId
	return builder
}
func (builder *GainProjectionBuilder) AddCategory(category GainCategory) *GainProjectionBuilder {
	builder.category = category
	return builder
}
func (builder *GainProjectionBuilder) Build() *GainProjection {
	gainProjection := GainProjection{}

	gainProjection.Id = builder.id
	gainProjection.CreatedAt = builder.createdAt
	gainProjection.PayIn = builder.payIn
	gainProjection.Description = builder.description
	gainProjection.Value = builder.value
	gainProjection.IsPassive = builder.isPassive
	gainProjection.IsDone = builder.isDone
	gainProjection.UserId = builder.userId
	gainProjection.Category = builder.category

	return &gainProjection
}

type QueryParamsBuilder struct {
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
		month:  builder.month,
		year:   builder.year,
		limit:  builder.limit,
		offset: builder.offset,
	}
}
