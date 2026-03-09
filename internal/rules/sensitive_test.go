package rules_test

import (
	"testing"

	"github.com/Davidianol/loglint/internal/rules"
)

func TestCheckSensitive(t *testing.T) {
	runCases(t, rules.CheckSensitive, []testCase{
		// OK
		{"user authenticated successfully", false},
		{"api request completed", false},
		{"session started", false},
		{"user logged in", false},
		{"", false},

		// want
		{"user password: secret", true},
		{"api_key=abc123", true},
		{"token: xyz", true},
		{"client_secret expired", true},
		{"private_key not found", true},
		{"credential missing", true},
		{"access_key expired", true},
		{"auth failed", true},
	})
}
