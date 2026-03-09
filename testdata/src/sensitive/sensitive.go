package sensitive

import (
	"context"
	"log/slog"
)

func examples(password, apiKey, token, secret string) {
	ctx := context.Background()

	// slog - базовые
	slog.Debug("api_key=" + apiKey)         // want `log message may contain sensitive data`
	slog.Info("user password: " + password) // want `log message contains forbidden char` `log message may contain sensitive data`
	slog.Warn("secret: " + secret)          // want `log message contains forbidden char` `log message may contain sensitive data`
	slog.Error("token: " + token)           // want `log message contains forbidden char` `log message may contain sensitive data`

	// slog - Context
	slog.DebugContext(ctx, "api_key="+apiKey)       // want `log message may contain sensitive data`
	slog.InfoContext(ctx, "token: "+token)          // want `log message contains forbidden char` `log message may contain sensitive data`
	slog.WarnContext(ctx, "client_secret expired")  // want `log message may contain sensitive data`
	slog.ErrorContext(ctx, "private_key not found") // want `log message may contain sensitive data`

	// slog - Log / LogAttrs
	slog.Log(ctx, slog.LevelInfo, "password reset")          // want `log message may contain sensitive data`
	slog.LogAttrs(ctx, slog.LevelInfo, "auth token: "+token) // want `log message contains forbidden char` `log message may contain sensitive data`

}
