package main

import (
	"github.com/nsega/goanalysis_study/finderror"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(finderror.Analyzer) }
