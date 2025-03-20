package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const linterName = "utctime"
const linterDoc = "Checks that time.Now() is followed by .UTC()"

var Analyzer = &analysis.Analyzer{
	Name:     linterName,
	Doc:      linterDoc,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// Look for selector expressions (like .UTC())
			sel, ok := n.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// Check if the selector is UTC
			if sel.Sel.Name != "UTC" {
				return true
			}

			// Check if it's being called (has parentheses)
			_, ok = sel.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Get the time.Now() part
			nowCall, ok := sel.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Check if it's a selector expression (time.Now)
			nowSel, ok := nowCall.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// Check if it's from the time package
			var ident *ast.Ident
			ident, ok = nowSel.X.(*ast.Ident)
			if !ok || ident.Name != "time" {
				return true
			}

			// Check if it's Now()
			if nowSel.Sel.Name != "Now" {
				return true
			}

			// If we get here, we found a time.Now().UTC() call, which is fine
			return true
		})

		// Make a second pass to find bare time.Now() calls
		ast.Inspect(file, func(n ast.Node) bool {
			// Look for function calls
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Check if it's a selector expression (time.Now)
			selectorExpr, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// Check if it's from the time package
			var ident *ast.Ident
			ident, ok = selectorExpr.X.(*ast.Ident)
			if !ok || ident.Name != "time" {
				return true
			}

			// Check if it's Now()
			if selectorExpr.Sel.Name != "Now" {
				return true
			}

			// Check if this Now() call is the X part of a SelectorExpr with .UTC()
			if parent, ok := findParentNode(n, file); ok {
				if parentSel, ok := parent.(*ast.SelectorExpr); ok && parentSel.Sel.Name == "UTC" {
					return true
				}
			}

			// If we get here, we found a time.Now() without .UTC()
			pass.Reportf(call.Pos(), "time.Now() must be followed by .UTC()")
			return true
		})
	}

	//nolint:nilnil // ignore for analysis testing.
	return nil, nil
}

// findParentNode checks if the current node is part of a larger expression
func findParentNode(node ast.Node, file *ast.File) (ast.Node, bool) {
	var parent ast.Node
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.SelectorExpr:
			if x.X == node {
				parent = n
				return false
			}
		}
		return true
	})
	return parent, parent != nil
}
