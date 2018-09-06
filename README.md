[![Go Report Card](https://goreportcard.com/badge/github.com/go-toolsmith/astinfo)](https://goreportcard.com/report/github.com/go-toolsmith/astinfo)
[![GoDoc](https://godoc.org/github.com/go-toolsmith/astinfo?status.svg)](https://godoc.org/github.com/go-toolsmith/astinfo)


# astinfo

Package astinfo records useful AST information like node parents and such.

## Installation:

```bash
go get github.com/go-toolsmith/astinfo
```

## Example

```go
package main

import (
	"fmt"
	"go/ast"

	"github.com/go-toolsmith/astinfo"
)

func main() {
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
```
