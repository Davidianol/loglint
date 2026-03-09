package rules_test

import (
	"testing"

	"github.com/Davidianol/loglint/internal/rules"
)

func TestCheckLowercase(t *testing.T) {
	runCases(t, rules.CheckLowercase, []testCase{
		// OK
		{"starting server", false},
		{"failed to connect", false},
		{"123 starts with digit", false},
		{"", false},

		// want
		{"Starting server", true},
		{"Failed to connect", true},
		{"ERROR occurred", true},
	})
}
