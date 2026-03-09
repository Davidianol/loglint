package rules

import (
	"go/token"
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

func CheckSensitive(pass *analysis.Pass, pos token.Pos, msg string) {
	lower := strings.ToLower(msg)
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(lower, keyword) {
			pass.Reportf(pos,
				"log message may contain sensitive data (keyword %q found): %q",
				keyword, msg,
			)
			return
		}
	}
}
