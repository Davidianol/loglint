package rules

import (
	"go/token"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func CheckEnglish(pass *analysis.Pass, pos token.Pos, msg string) {
	for _, r := range msg {
		if r > unicode.MaxASCII && !unicode.Is(unicode.Latin, r) {
			pass.Reportf(pos, "log message must be in English only: %q", msg)
			return
		}
	}
}
