package codegen_exp_test

import (
	"flag"
	"os"
	"testing"

	"codegen_exp"
	"github.com/gostaticanalysis/codegen/codegentest"
)

var flagUpdate bool

func TestMain(m *testing.M) {
	flag.BoolVar(&flagUpdate, "update", false, "update the golden files")
	flag.Parse()
	os.Exit(m.Run())
}

func TestGenerator(t *testing.T) {
	rs := codegentest.Run(t, codegentest.TestData(), codegen_exp.Generator, "a")
	codegentest.Golden(t, rs, flagUpdate)
}
