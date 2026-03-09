package loglint_test

import (
	"testing"

	"github.com/Davidianol/loglint"
	"github.com/Davidianol/loglint/internal/config"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzerConfigDisableLowercase(t *testing.T) {
	cfg := &config.Config{DisableLowercase: true}
	analyzer := loglint.NewAnalyzer(cfg)

	// lowercase/ содержит ошибки — но с DisableLowercase=true их быть не должно
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer, "ok")
}

func TestAnalyzerConfigExtraSensitive(t *testing.T) {
	cfg := &config.Config{
		ExtraSensitiveKeywords: []string{"db_pass"},
	}
	analyzer := loglint.NewAnalyzer(cfg)
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer, "config_sensitive")
}

func TestAnalyzerConfigExtraPatterns(t *testing.T) {
	cfg := &config.Config{
		ExtraForbiddenPatterns: []string{"TODO"},
	}
	analyzer := loglint.NewAnalyzer(cfg)
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer, "config_patterns")
}
