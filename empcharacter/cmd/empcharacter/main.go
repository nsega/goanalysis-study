package main

import (
	"github.com/nsega/goanalysis_study/empcharacter"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(empcharacter.Analyzer) }
