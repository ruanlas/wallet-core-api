package service

import "time"

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
