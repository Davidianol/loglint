package loglint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/Davidianol/loglint/internal/rules"
)

var Analyzer = &analysis.Analyzer{
	Name:     "loglint",
	Doc:      "checks log messages for style and security rules",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)
		msg, pos, ok := extractLogMessage(pass, call)
		if !ok {
			return
		}

		rules.CheckLowercase(pass, pos, msg)
		rules.CheckEnglish(pass, pos, msg)
		rules.CheckSpecialChars(pass, pos, msg)
		rules.CheckSensitive(pass, pos, msg)
	})

	return nil, nil
}
