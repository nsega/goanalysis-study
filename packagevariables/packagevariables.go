package packagevariables

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "packagevariables is ..."

// Analyzer is to count the package variables
var Analyzer = &analysis.Analyzer{
	Name: "packagevariables",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	var count int
	for _, f := range pass.Files {
		for _, decl := range f.Decls {

			d, ok := decl.(*ast.GenDecl)
			if !ok || d.Tok != token.VAR {
				continue
			}

			fmt.Println(d)
			for _, spec := range d.Specs {
				spec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				count += len(spec.Names)
				fmt.Println(spec.Names)
			}
		}
	}
	fmt.Println(count)

	return nil, nil
}
