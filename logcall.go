package loglint

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// supportedLogPackages: пакет -> разрешённые имена методов
var supportedLogPackages = map[string]map[string]bool{
	"log/slog": {
		"Debug": true, "Info": true, "Warn": true, "Error": true,
		"DebugContext": true, "InfoContext": true,
		"WarnContext": true, "ErrorContext": true, "Log": true, "LogAttrs": true,
	},
	"go.uber.org/zap": {
		"Debug": true, "Info": true, "Warn": true,
		"Error": true, "Fatal": true, "Panic": true, "DPanic": true,
	},
}

// contextMethods принимают context.Context первым аргументом,
// поэтому сообщение находится на позиции 1
var contextMethods = map[string]bool{
	"DebugContext": true, "InfoContext": true,
	"WarnContext": true, "ErrorContext": true,
}

func extractLogMessage(pass *analysis.Pass, call *ast.CallExpr) (string, token.Pos, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", 0, false
	}

	methodName := sel.Sel.Name
	if !isLogCall(pass, sel, methodName) {
		return "", 0, false
	}

	msgIdx := 0
	if contextMethods[methodName] {
		msgIdx = 1
	}
	if len(call.Args) <= msgIdx {
		return "", 0, false
	}

	return extractStringArg(call.Args[msgIdx])
}

func isLogCall(pass *analysis.Pass, sel *ast.SelectorExpr, methodName string) bool {
	if pass.TypesInfo == nil {
		return false
	}

	// package-level func: slog.Info, slog.Error и т.д.
	if obj := pass.TypesInfo.ObjectOf(sel.Sel); obj != nil {
		if fn, ok := obj.(*types.Func); ok && fn.Pkg() != nil {
			pkgPath := fn.Pkg().Path()
			if funcs, ok := supportedLogPackages[pkgPath]; ok {
				return funcs[methodName]
			}
		}
	}

	// method call: logger.Info где logger — *zap.Logger или *slog.Logger
	if selection, ok := pass.TypesInfo.Selections[sel]; ok {
		recv := selection.Recv()
		if ptr, ok := recv.(*types.Pointer); ok {
			recv = ptr.Elem()
		}
		if named, ok := recv.(*types.Named); ok && named.Obj().Pkg() != nil {
			pkgPath := named.Obj().Pkg().Path()
			if funcs, ok := supportedLogPackages[pkgPath]; ok {
				return funcs[methodName]
			}
		}
	}

	return false
}

func extractStringArg(expr ast.Expr) (string, token.Pos, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind != token.STRING {
			return "", 0, false
		}
		val, err := unquoteString(e.Value)
		if err != nil {
			return "", 0, false
		}
		return val, e.Pos(), true

	case *ast.BinaryExpr:
		if e.Op != token.ADD {
			return "", 0, false
		}
		// рекурсивно собираем все строковые части
		collected, pos := collectStringParts(e)
		if collected == "" {
			return "", 0, false
		}
		return collected, pos, true
	}
	return "", 0, false
}

// collectStringParts рекурсивно обходит дерево конкатенации
// и собирает все строковые литералы в одну строку
func collectStringParts(expr ast.Expr) (string, token.Pos) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			val, err := unquoteString(e.Value)
			if err == nil {
				return val, e.Pos()
			}
		}
		return "", 0

	case *ast.BinaryExpr:
		if e.Op != token.ADD {
			return "", 0
		}
		leftStr, leftPos := collectStringParts(e.X)
		rightStr, rightPos := collectStringParts(e.Y)

		pos := leftPos
		if pos == 0 {
			pos = rightPos
		}
		return leftStr + rightStr, pos
	}
	return "", 0
}

func unquoteString(s string) (string, error) {
	if strings.HasPrefix(s, "`") && strings.HasSuffix(s, "`") {
		return s[1 : len(s)-1], nil
	}
	return strconv.Unquote(s)
}
