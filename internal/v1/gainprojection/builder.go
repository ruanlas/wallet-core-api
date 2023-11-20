package gainprojection

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

type GainProjectionResponseBuilder struct {
	id          string
	payIn       time.Time
	description string
	value       float64
	isPassive   bool
	recurrence  uint
	category    CategoryResponse
}

func NewGainProjectionResponseBuilder() *GainProjectionResponseBuilder {
	return &GainProjectionResponseBuilder{}
}
func (builder *GainProjectionResponseBuilder) AddId(id string) *GainProjectionResponseBuilder {
	builder.id = id
	return builder
}
func (builder *GainProjectionResponseBuilder) AddPayIn(payIn time.Time) *GainProjectionResponseBuilder {
	builder.payIn = payIn
	return builder
}
func (builder *GainProjectionResponseBuilder) AddDescription(description string) *GainProjectionResponseBuilder {
	builder.description = description
	return builder
}
func (builder *GainProjectionResponseBuilder) AddValue(value float64) *GainProjectionResponseBuilder {
	builder.value = value
	return builder
}
func (builder *GainProjectionResponseBuilder) AddIsPassive(isPassive bool) *GainProjectionResponseBuilder {
	builder.isPassive = isPassive
	return builder
}
func (builder *GainProjectionResponseBuilder) AddRecurrence(recurrence uint) *GainProjectionResponseBuilder {
	builder.recurrence = recurrence
	return builder
}
func (builder *GainProjectionResponseBuilder) AddCategory(category CategoryResponse) *GainProjectionResponseBuilder {
	builder.category = category
	return builder
}
func (builder *GainProjectionResponseBuilder) Build() *GainProjectionResponse {
	gainProjectionResponse := GainProjectionResponse{}

	gainProjectionResponse.Id = builder.id
	gainProjectionResponse.Description = builder.description
	gainProjectionResponse.Value = builder.value
	gainProjectionResponse.PayIn = builder.payIn
	gainProjectionResponse.IsPassive = builder.isPassive
	gainProjectionResponse.Recurrence = builder.recurrence
	gainProjectionResponse.Category = builder.category

	return &gainProjectionResponse
}
