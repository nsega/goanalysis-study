package main

import (
	"github.com/nsega/goanalysis_study/packagevariables"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(packagevariables.Analyzer) }
