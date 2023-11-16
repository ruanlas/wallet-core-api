package gainprojection

func NewGainProjectionBuilder() *GainProjectionBuilder {
	return &GainProjectionBuilder{}
}

// TODO: Terminar o builder. Inserir os outros campos do GainProjection
type GainProjectionBuilder struct {
	id string
}

func (builder *GainProjectionBuilder) AddId(id string) *GainProjectionBuilder {
	builder.id = id
	return builder
}

func (builder *GainProjectionBuilder) Build() *GainProjection {
	gainProjection := GainProjection{}
	gainProjection.Id = builder.id

	return &gainProjection
}
