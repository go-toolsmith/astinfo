# astinfo

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Package astinfo records useful AST information like node parents and such.

## Installation:

Go version 1.16+

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

## License

[MIT License](LICENSE).

[build-img]: https://github.com/go-toolsmith/astinfo/workflows/build/badge.svg
[build-url]: https://github.com/go-toolsmith/astinfo/actions
[pkg-img]: https://pkg.go.dev/badge/go-toolsmith/astinfo
[pkg-url]: https://pkg.go.dev/github.com/go-toolsmith/astinfo
[reportcard-img]: https://goreportcard.com/badge/go-toolsmith/astinfo
[reportcard-url]: https://goreportcard.com/report/go-toolsmith/astinfo
[coverage-img]: https://codecov.io/gh/go-toolsmith/astinfo/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/go-toolsmith/astinfo
[version-img]: https://img.shields.io/github/v/release/go-toolsmith/astinfo
[version-url]: https://github.com/go-toolsmith/astinfo/releases
