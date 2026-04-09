package requestctx

import "context"

type ctxKey int

const requestInfoKey ctxKey = iota

type RequestInfo struct {
	RequestID string
	UserID    string
	TenantID  string
	TraceID   string
	SpanID    string
}

func With(ctx context.Context, info RequestInfo) context.Context {
	return context.WithValue(ctx, requestInfoKey, info)
}

func From(ctx context.Context) (RequestInfo, bool) {
	info, ok := ctx.Value(requestInfoKey).(RequestInfo)
	return info, ok
}
