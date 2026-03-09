package config_test

import (
	"testing"

	"github.com/Davidianol/loglint/internal/config"
)

func TestParseNil(t *testing.T) {
	cfg, err := config.Parse(nil)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.DisableLowercase || cfg.DisableSensitive {
		t.Error("default config should have all rules enabled")
	}
}

func TestParseSettings(t *testing.T) {
	settings := map[string]any{
		"disable_lowercase":        true,
		"extra_sensitive_keywords": []any{"db_pass", "master_key"},
		"extra_forbidden_patterns": []any{"TODO"},
	}

	cfg, err := config.Parse(settings)
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.DisableLowercase {
		t.Error("disable_lowercase should be true")
	}
	if len(cfg.ExtraSensitiveKeywords) != 2 {
		t.Errorf("expected 2 extra keywords, got %d", len(cfg.ExtraSensitiveKeywords))
	}
	if len(cfg.ExtraForbiddenPatterns) != 1 {
		t.Errorf("expected 1 extra pattern, got %d", len(cfg.ExtraForbiddenPatterns))
	}
}
