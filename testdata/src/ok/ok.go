package ok

import (
	"context"
	"log/slog"
)

func examples() {
	ctx := context.Background()

	slog.Debug("debug message")
	slog.Info("starting server")
	slog.Warn("something went wrong")
	slog.Error("failed to connect")
	slog.DebugContext(ctx, "debug context")
	slog.InfoContext(ctx, "info context")
	slog.WarnContext(ctx, "warn context")
	slog.ErrorContext(ctx, "error context")
	slog.Log(ctx, slog.LevelInfo, "log message")
	slog.LogAttrs(ctx, slog.LevelInfo, "logattrs message")
}
