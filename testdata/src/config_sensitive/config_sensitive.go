package config_sensitive

import "log/slog"

func examples(dbPass string) {
	slog.Info("db_pass=" + dbPass) // want `log message may contain sensitive data \(keyword "db_pass" found\)`
}
