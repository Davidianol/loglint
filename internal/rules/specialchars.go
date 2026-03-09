package rules

import (
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var disallowedChars = "!?@#$%^&*:;"
var disallowedPatterns = []string{"..."}

// CheckSpecialChars - оригинальная функция, используется в unit-тестах
func CheckSpecialChars(pass *analysis.Pass, pos token.Pos, msg string) {
	CheckSpecialCharsConfig(pass, pos, msg, "", nil)
}

// CheckSpecialCharsConfig - с поддержкой кастомных паттернов из конфига
func CheckSpecialCharsConfig(pass *analysis.Pass, pos token.Pos, msg, extraChars string, extraPatterns []string) {
	allChars := disallowedChars + extraChars

	for _, r := range msg {
		if r > unicode.MaxASCII {
			continue
		}
		if strings.ContainsRune(allChars, r) {
			pass.Reportf(pos, "log message contains forbidden char %q: %q", r, msg)
			return
		}
	}

	allPatterns := append(disallowedPatterns, extraPatterns...)
	for _, pattern := range allPatterns {
		if strings.Contains(msg, pattern) {
			pass.Reportf(pos, "log message contains forbidden pattern %q: %q", pattern, msg)
			return
		}
	}
}
