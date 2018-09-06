package astinfo_test

import (
	"fmt"
	"go/ast"

	"github.com/go-toolsmith/astinfo"
)

func Example() {
	innermost := &ast.Ident{}
	root := &ast.ExprStmt{
		X: &ast.BinaryExpr{
			X: &ast.BasicLit{},
			Y: &ast.UnaryExpr{X: innermost},
		},
	}

	info := astinfo.Info{Parents: make(map[ast.Node]ast.Node)}
	info.Origin = root
	info.Resolve()

	for p := info.Parents[innermost]; p != nil; p = info.Parents[p] {
		fmt.Printf("%T\n", p)
	}

	// Output:
	// *ast.UnaryExpr
	// *ast.BinaryExpr
	// *ast.ExprStmt
}
