package printpkgfunc

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "printpkgfunc is ..."

// Analyzer is to print package function.
var Analyzer = &analysis.Analyzer{
	Name: "printpkgfunc",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			decl, ok := decl.(*ast.FuncDecl)
			if !ok || decl.Recv != nil {
				continue
			}
			fmt.Println(decl.Name)
		}
	}

	return nil, nil
}
