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
	sensitivePatterns = compilePatterns(sensitiveKeywords)
}

func compilePatterns(keywords []string) []*regexp.Regexp {
	patterns := make([]*regexp.Regexp, 0, len(keywords))
	for _, kw := range keywords {
		patterns = append(patterns,
			regexp.MustCompile(`(?i)\b`+regexp.QuoteMeta(kw)+`\b`),
		)
	}
	return patterns
}

// CheckSensitive - оригинальная функция, используется в unit-тестах
func CheckSensitive(pass *analysis.Pass, pos token.Pos, msg string) {
	CheckSensitiveConfig(pass, pos, msg, nil)
}

// CheckSensitiveConfig - с поддержкой кастомных ключевых слов из конфига
func CheckSensitiveConfig(pass *analysis.Pass, pos token.Pos, msg string, extraKeywords []string) {
	lower := strings.ToLower(msg)

	allKeywords := sensitiveKeywords
	allPatterns := sensitivePatterns

	if len(extraKeywords) > 0 {
		allKeywords = append(allKeywords, extraKeywords...)
		allPatterns = append(allPatterns, compilePatterns(extraKeywords)...)
	}

	for i, pattern := range allPatterns {
		if pattern.MatchString(lower) {
			pass.Reportf(pos,
				"log message may contain sensitive data (keyword %q found): %q",
				allKeywords[i], msg,
			)
			return
		}
	}
}
