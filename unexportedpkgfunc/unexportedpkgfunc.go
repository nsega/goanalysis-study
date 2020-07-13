package unexportedpkgfunc

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "unexportedpkgfunc is to count unexported package functions"

// Analyzer is to count unexported package functions
var Analyzer = &analysis.Analyzer{
	Name: "unexportedpkgfunc",
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

			if !decl.Name.IsExported() {
				fmt.Printf("found: %s\n", decl.Name)
			}

		}

	}
	return nil, nil
}
