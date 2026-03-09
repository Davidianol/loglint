package config

import "encoding/json"

type Config struct {
	// Отключение отдельных правил
	DisableLowercase    bool `json:"disable_lowercase"`
	DisableEnglish      bool `json:"disable_english"`
	DisableSpecialChars bool `json:"disable_special_chars"`
	DisableSensitive    bool `json:"disable_sensitive"`

	// Кастомные паттерны
	ExtraForbiddenChars    string   `json:"extra_forbidden_chars"`
	ExtraForbiddenPatterns []string `json:"extra_forbidden_patterns"`
	ExtraSensitiveKeywords []string `json:"extra_sensitive_keywords"`
}

func Default() *Config {
	return &Config{}
}

// Parse декодирует settings из golangci-lint
func Parse(settings any) (*Config, error) {
	cfg := Default()
	if settings == nil {
		return cfg, nil
	}
	b, err := json.Marshal(settings)
	if err != nil {
		return cfg, err
	}
	if err := json.Unmarshal(b, cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
