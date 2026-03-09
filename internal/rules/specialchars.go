package rules

import (
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var disallowedChars = "!?@#$%^&*:;"

var disallowedPatterns = []string{"..."}

func CheckSpecialChars(pass *analysis.Pass, pos token.Pos, msg string) {
	for _, r := range msg {
		if r > unicode.MaxASCII {
			continue
		}
		if strings.ContainsRune(disallowedChars, r) {
			pass.Reportf(pos, "log message contains forbidden char %q: %q", r, msg)
			return
		}
	}
	for _, pattern := range disallowedPatterns {
		if strings.Contains(msg, pattern) {
			pass.Reportf(pos, "log message contains forbidden pattern %q: %q", pattern, msg)
			return
		}
	}
}
