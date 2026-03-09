package rules_test

import (
	"go/token"
	"testing"

	"golang.org/x/tools/go/analysis"
)

type ruleFunc func(pass *analysis.Pass, pos token.Pos, msg string)

type testCase struct {
	msg     string
	wantErr bool
}

func runCases(t *testing.T, fn ruleFunc, cases []testCase) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.msg, func(t *testing.T) {
			var reports []analysis.Diagnostic
			pass := &analysis.Pass{
				Report: func(d analysis.Diagnostic) {
					reports = append(reports, d)
				},
			}
			fn(pass, token.NoPos, tc.msg)
			got := len(reports) > 0
			if got != tc.wantErr {
				t.Errorf("msg=%q: wantErr=%v got=%v (report: %v)",
					tc.msg, tc.wantErr, got, reports)
			}
		})
	}
}
