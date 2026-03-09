package english

import (
	"context"
	"log/slog"
)

func examples() {
	ctx := context.Background()

	// slog - базовые
	slog.Debug("отладка")            // want `log message must be in English only`
	slog.Info("запуск сервера")      // want `log message must be in English only`
	slog.Warn("предупреждение")      // want `log message must be in English only`
	slog.Error("ошибка подключения") // want `log message must be in English only`

	// slog - Context
	slog.DebugContext(ctx, "отладка контекст") // want `log message must be in English only`
	slog.InfoContext(ctx, "запуск контекст")   // want `log message must be in English only`
	slog.WarnContext(ctx, "ошибка")            // want `log message must be in English only`
	slog.ErrorContext(ctx, "错误")               // want `log message must be in English only`

	// slog - Log / LogAttrs
	slog.Log(ctx, slog.LevelInfo, "서버 시작")       // want `log message must be in English only`
	slog.LogAttrs(ctx, slog.LevelInfo, "サーバー起動") // want `log message must be in English only`

}
