package log

import (
	"context"
	"log/slog"

	"github.com/tiagoangelototvs/slog-context-patterns/internal/requestctx"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	info, ok := requestctx.From(ctx)
	if !ok {
		return h.Handler.Handle(ctx, r)
	}

	if info.RequestID != "" {
		r.AddAttrs(slog.String("request_id", info.RequestID))
	}

	if info.UserID != "" {
		r.AddAttrs(slog.String("user_id", info.UserID))
	}

	if info.TenantID != "" {
		r.AddAttrs(slog.String("tenant_id", info.TenantID))
	}

	if info.TraceID != "" {
		r.AddAttrs(slog.String("trace_id", info.TraceID))
	}

	if info.SpanID != "" {
		r.AddAttrs(slog.String("span_id", info.SpanID))
	}

	return h.Handler.Handle(ctx, r)
}
