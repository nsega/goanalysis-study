package main

import (
	"codegen_exp"
	"github.com/gostaticanalysis/codegen/singlegenerator"
)

func main() {
	singlegenerator.Main(codegen_exp.Generator)
}
