package main

import (
	"github.com/nsega/goanalysis_study/unexportedpkgfunc"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unexportedpkgfunc.Analyzer) }
