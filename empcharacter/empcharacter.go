package empcharacter

import (
	"go/ast"
	"go/token"
	"strconv"

	"fmt"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "empcharacter is ..."

// Analyzer is to return an empty character
var Analyzer = &analysis.Analyzer{
	Name: "empcharacter",
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

			for _, stmt := range decl.Body.List {
				stmt, ok := stmt.(*ast.ReturnStmt)
				if !ok && len(stmt.Results) != 1 {
					continue
				}

				lit, ok := stmt.Results[0].(*ast.BasicLit)
				if !ok {
					continue
				}

				if lit.Kind != token.STRING {
					continue
				}

				unquoted, err := strconv.Unquote(lit.Value)
				if err != nil {
					return nil, err
				}
				if unquoted == "" {
					fmt.Printf("found: %s\n", decl.Name)
				}
			}
		}
	}
	return nil, nil
}
