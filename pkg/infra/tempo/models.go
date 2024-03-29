package tempo

type SearchTagsResult struct {
	Traces []Trace `json:"traces"`
}

type Trace struct {
	TraceID           string `json:"traceID"`
	RootServiceName   string `json:"rootServiceName"`
	RootTraceName     string `json:"rootTraceName"`
	StartTimeUnixNano string `json:"startTimeUnixNano"`
	DurationMs        int    `json:"durationMs"`
}

type TraceDetails struct {
	Batches []Batch `json:"batches"`
}

type Batch struct {
	Resource   Resource    `json:"resource"`
	ScopeSpans []ScopeSpan `json:"scopeSpans"`
}

type Resource struct {
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Key   string            `json:"key"`
	Value map[string]string `json:"value"`
}

type ScopeSpan struct {
	Scope Scope  `json:"scope"`
	Spans []Span `json:"spans"`
}

type Scope struct {
	Name string `json:"name"`
}

type Span struct {
	TraceId           string      `json:"traceId"`
	SpanId            string      `json:"spanId"`
	ParentSpanId      string      `json:"parentSpanId"`
	Name              string      `json:"name"`
	Kind              string      `json:"kind"`
	StartTimeUnixNano string      `json:"startTimeUnixNano"`
	EndTimeUnixNano   string      `json:"endTimeUnixNano"`
	Attributes        []Attribute `json:"attributes"`
}

func (b *Batch) FindSpansByName(name string) []Span {
	var spans []Span
	if b == nil || b.ScopeSpans == nil {
		return spans
	}

	for _, ss := range b.ScopeSpans {
		for _, s := range ss.Spans {
			if s.Name == name {
				spans = append(spans, s)
			}
		}
	}

	return spans
}

func (b *Batch) ChildrenOf(spanId string) []Span {
	var spans []Span
	if b == nil || b.ScopeSpans == nil {
		return spans
	}

	for _, ss := range b.ScopeSpans {
		for _, s := range ss.Spans {
			if s.ParentSpanId == spanId {
				spans = append(spans, s)
			}
		}
	}

	return spans
}
