package rules

import (
	"go/token"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func CheckLowercase(pass *analysis.Pass, pos token.Pos, msg string) {
	if len(msg) == 0 {
		return
	}
	if unicode.IsUpper([]rune(msg)[0]) {
		pass.Reportf(pos, "log message must start with a lowercase letter: %q", msg)
	}
}
