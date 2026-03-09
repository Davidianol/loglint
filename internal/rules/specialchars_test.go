package rules_test

import (
	"testing"

	"github.com/Davidianol/loglint/internal/rules"
)

func TestCheckSpecialChars(t *testing.T) {
	runCases(t, rules.CheckSpecialChars, []testCase{
		// OK
		{"server started", false},
		{"connection failed", false},
		{"something went wrong", false},
		{"", false},

		// want
		{"server started!", true},
		{"connection failed!!!", true},
		{"warning: disk full", true},
		{"request failed???", true},
		{"retrying...", true},
		{"status: ok", true},
	})
}
