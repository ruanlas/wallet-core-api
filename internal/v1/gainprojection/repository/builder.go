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

type GainBuilder struct {
	id               string
	createdAt        time.Time
	payIn            time.Time
	description      string
	value            float64
	isPassive        bool
	userId           string
	category         GainCategory
	gainProjectionId string
}

func NewGainBuilder() *GainBuilder {
	return &GainBuilder{}
}
func (builder *GainBuilder) AddId(id string) *GainBuilder {
	builder.id = id
	return builder
}
func (builder *GainBuilder) AddCreatedAt(createdAt time.Time) *GainBuilder {
	builder.createdAt = createdAt
	return builder
}
func (builder *GainBuilder) AddPayIn(payIn time.Time) *GainBuilder {
	builder.payIn = payIn
	return builder
}
func (builder *GainBuilder) AddDescription(description string) *GainBuilder {
	builder.description = description
	return builder
}
func (builder *GainBuilder) AddValue(value float64) *GainBuilder {
	builder.value = value
	return builder
}
func (builder *GainBuilder) AddIsPassive(isPassive bool) *GainBuilder {
	builder.isPassive = isPassive
	return builder
}
func (builder *GainBuilder) AddGainProjectionId(gainProjectionId string) *GainBuilder {
	builder.gainProjectionId = gainProjectionId
	return builder
}
func (builder *GainBuilder) AddUserId(userId string) *GainBuilder {
	builder.userId = userId
	return builder
}
func (builder *GainBuilder) AddCategory(category GainCategory) *GainBuilder {
	builder.category = category
	return builder
}
func (builder *GainBuilder) Build() *Gain {
	gain := Gain{}

	gain.Id = builder.id
	gain.CreatedAt = builder.createdAt
	gain.PayIn = builder.payIn
	gain.Description = builder.description
	gain.Value = builder.value
	gain.IsPassive = builder.isPassive
	gain.GainProjectionId = builder.gainProjectionId
	gain.UserId = builder.userId
	gain.Category = builder.category

	return &gain
}
