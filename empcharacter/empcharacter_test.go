package empcharacter_test

import (
	"testing"

	"github.com/nsega/goanalysis_study/empcharacter"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, empcharacter.Analyzer, "a")
}
