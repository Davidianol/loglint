package config_patterns

import "log/slog"

func examples() {
	slog.Info("fix TODO later") // want `log message contains forbidden pattern "TODO"`
	slog.Info("server started") // ok
}
