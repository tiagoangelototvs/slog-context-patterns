package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/tiagoangelototvs/slog-context-patterns/internal/log"
	"github.com/tiagoangelototvs/slog-context-patterns/internal/requestctx"
)

func main() {
	slog.SetDefault(slog.New(
		&log.ContextHandler{
			Handler: tint.NewHandler(os.Stderr, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		},
	))

	ctx := requestctx.With(context.Background(), requestctx.RequestInfo{
		RequestID: "req-123",
		UserID:    "user-456",
		TenantID:  "tenant-789",
		TraceID:   "trace-abc",
		SpanID:    "span-def",
	})

	process(ctx)
}

func process(ctx context.Context) {
	slog.InfoContext(ctx, "starting process",
		slog.String("operation", "process"),
	)

	step(ctx)
}

func step(ctx context.Context) {
	slog.InfoContext(ctx, "executing step")
}
