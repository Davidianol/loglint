package loglint_test

import (
	"testing"

	"github.com/Davidianol/loglint"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, loglint.Analyzer,
		"lowercase",
		"english",
		"specialchars",
		"sensitive",
	)
}
