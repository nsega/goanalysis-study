package dupimport_test

import (
	"testing"

	"github.com/nsega/analysis_study/dupimport"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, dupimport.Analyzer, "a")
}
