package analyzer

import (
	"flag"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "cyclop",
		Doc:   "calculates cyclomatic complexity",
		Run:   run,
		Flags: flagSet,
	}
}

func complexity(fn *ast.FuncDecl) int {
	v := complexityVisitor{Complexity: 1}
	ast.Walk(&v, fn)
	return v.Complexity
}

type complexityVisitor struct {
	Complexity int
}

func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.FuncDecl, *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
		v.Complexity++
	case *ast.BinaryExpr:
		if n.Op == token.LAND || n.Op == token.LOR {
			v.Complexity++
		}
	}
	return v
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			f, ok := node.(*ast.FuncDecl)
			if !ok {
				return true
			}
			comp := complexity(f)
			pass.Reportf(node.Pos(), "calculated cyclomatic complexity for function %s is %d", f.Name.Name, comp)
			return true
		})
	}
	return nil, nil
}
