package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/nsega/goanalysis_study/channel_close"
)

func main() { unitchecker.Main(channel_close.Analyzer) }
