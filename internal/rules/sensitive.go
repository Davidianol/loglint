package rules

import (
	"go/token"
	"regexp"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var sensitiveKeywords = []string{
	"password", "passwd", "pwd",
	"api_key", "apikey", "api-key",
	"token", "secret", "auth",
	"credential", "private_key", "privatekey",
	"access_key", "client_secret",
}

var sensitivePatterns []*regexp.Regexp

func init() {
	for _, kw := range sensitiveKeywords {
		// \b — граница слова: "auth" не совпадет, например, с "authenticated"
		sensitivePatterns = append(sensitivePatterns,
			regexp.MustCompile(`(?i)\b`+regexp.QuoteMeta(kw)+`\b`),
		)
	}
}

func CheckSensitive(pass *analysis.Pass, pos token.Pos, msg string) {
	lower := strings.ToLower(msg)
	for i, pattern := range sensitivePatterns {
		if pattern.MatchString(lower) {
			pass.Reportf(pos,
				"log message may contain sensitive data (keyword %q found): %q",
				sensitiveKeywords[i], msg,
			)
			return
		}
	}
}
