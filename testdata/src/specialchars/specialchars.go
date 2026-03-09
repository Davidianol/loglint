package specialchars

import (
	"context"
	"log/slog"
)

func examples() {
	ctx := context.Background()

	// slog - базовые
	slog.Debug("debug!")               // want `log message contains forbidden char`
	slog.Info("server started!")       // want `log message contains forbidden char`
	slog.Warn("warning: disk full")    // want `log message contains forbidden char`
	slog.Error("connection failed!!!") // want `log message contains forbidden char`

	// slog - Context
	slog.DebugContext(ctx, "debug context!")           // want `log message contains forbidden char`
	slog.InfoContext(ctx, "info context?")             // want `log message contains forbidden char`
	slog.WarnContext(ctx, "warn context...")           // want `log message contains forbidden pattern`
	slog.ErrorContext(ctx, "error context: something") // want `log message contains forbidden char`

	// slog - Log / LogAttrs
	slog.Log(ctx, slog.LevelInfo, "log message!")           // want `log message contains forbidden char`
	slog.LogAttrs(ctx, slog.LevelInfo, "logattrs message🚀") // want `log message must be in English only`
}
