package tracing

import "go.elastic.co/apm"

func SendSpanErr(span *apm.Span, err error) {
	apmErr := apm.DefaultTracer.NewError(err)
	apmErr.SetSpan(span)
	apmErr.Send()
	span.End()
}
