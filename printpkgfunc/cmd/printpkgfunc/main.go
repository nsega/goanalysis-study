package main

import (
	"github.com/nsega/goanalysis_study/printpkgfunc"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(printpkgfunc.Analyzer) }
