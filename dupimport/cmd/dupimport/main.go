package main

import (
	"github.com/nsega/analysis_study/dupimport"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(dupimport.Analyzer) }
