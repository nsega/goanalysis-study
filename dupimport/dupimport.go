package dupimport

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
)

const doc = "dupimport is to find duplicated imports in the same file"

// Analyzer is to find duplicated imports in the same file.
var Analyzer = &analysis.Analyzer{
	Name: "dupimport",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, f := range pass.Files {
		paths := make(map[string]bool)
		for _, i := range f.Imports {
			path, err := strconv.Unquote(i.Path.Value)
			if err != nil {
				return nil, err
			}
			if paths[path] {
				pass.Reportf(i.Pos(), "%s is duplicated improt", path)
			} else {
				paths[path] = true
			}
		}
	}
	return nil, nil
}
