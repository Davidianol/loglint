package lowercase

import (
	"context"
	"log/slog"
)

func examples() {
	ctx := context.Background()

	// slog - базовые методы
	slog.Debug("Debug message") // want `log message must start with a lowercase letter`
	slog.Info("Info message")   // want `log message must start with a lowercase letter`
	slog.Warn("Warn message")   // want `log message must start with a lowercase letter`
	slog.Error("Error message") // want `log message must start with a lowercase letter`

	// slog - Context-методы
	slog.DebugContext(ctx, "Debug context message") // want `log message must start with a lowercase letter`
	slog.InfoContext(ctx, "Info context message")   // want `log message must start with a lowercase letter`
	slog.WarnContext(ctx, "Warn context message")   // want `log message must start with a lowercase letter`
	slog.ErrorContext(ctx, "Error context message") // want `log message must start with a lowercase letter`

	// slog - Log / LogAttrs
	slog.Log(ctx, slog.LevelInfo, "Log message")           // want `log message must start with a lowercase letter`
	slog.LogAttrs(ctx, slog.LevelInfo, "LogAttrs message") // want `log message must start with a lowercase letter`
}
