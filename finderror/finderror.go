package finderror

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const doc = "finderror is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "finderror",
	Doc:  doc,
	Run:  run,
}

var errTyp = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

func run(pass *analysis.Pass) (interface{}, error) {
	for _, name := range pass.Pkg.Scope().Names() {
		obj := pass.Pkg.Scope().Lookup(name).(*types.TypeName)
		if obj == nil {
			continue
		}

		t := obj.Type()
		if types.Implements(t, errTyp) ||
			types.Implements(types.NewPointer(t), errTyp) {
			fmt.Println(t)
		}
	}
	return nil, nil
}
