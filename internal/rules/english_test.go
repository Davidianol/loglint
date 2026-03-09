package rules_test

import (
	"testing"

	"github.com/Davidianol/loglint/internal/rules"
)

func TestCheckEnglish(t *testing.T) {
	runCases(t, rules.CheckEnglish, []testCase{
		// OK
		{"starting server", false},
		{"failed to connect to database", false},
		{"résumé uploaded", false},
		{"", false},

		// want
		{"запуск сервера", true},
		{"ошибка подключения", true},
		{"서버 시작", true},
		{"server started🚀", true},
		{"サーバー起動", true},
		{"警告", true},
	})
}
