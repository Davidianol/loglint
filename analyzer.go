package loglint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/Davidianol/loglint/internal/config"
	"github.com/Davidianol/loglint/internal/rules"
)

// Analyzer - дефолтный анализатор без конфига
var Analyzer = NewAnalyzer(config.Default())

func NewAnalyzer(cfg *config.Config) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "loglint",
		Doc:      "checks log messages for style and security rules",
		Run:      makeRun(cfg),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func makeRun(cfg *config.Config) func(*analysis.Pass) (any, error) {
	return func(pass *analysis.Pass) (any, error) {
		insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

		insp.Preorder(nodeFilter, func(n ast.Node) {
			call := n.(*ast.CallExpr)
			msg, pos, ok := extractLogMessage(pass, call)
			if !ok {
				return
			}
			if !cfg.DisableLowercase {
				rules.CheckLowercase(pass, pos, msg)
			}
			if !cfg.DisableEnglish {
				rules.CheckEnglish(pass, pos, msg)
			}
			if !cfg.DisableSpecialChars {
				rules.CheckSpecialCharsConfig(pass, pos, msg,
					cfg.ExtraForbiddenChars,
					cfg.ExtraForbiddenPatterns,
				)
			}
			if !cfg.DisableSensitive {
				rules.CheckSensitiveConfig(pass, pos, msg,
					cfg.ExtraSensitiveKeywords,
				)
			}
		})
		return nil, nil
	}
}
