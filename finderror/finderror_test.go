package finderror_test

import (
	"testing"

	"github.com/nsega/goanalysis_study/finderror"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, finderror.Analyzer, "a")
}
