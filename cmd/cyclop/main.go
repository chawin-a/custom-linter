package main

import (
	"github.com/chawin-a/custom-linter/pkg/analyzer"
	"golang.org/x/tools/go/analysis"

	// "golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/singlechecker"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.NewAnalyzer(),
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin

func main() {
	// multichecker.Main(AnalyzerPlugin.GetAnalyzers()...)
	singlechecker.Main(analyzer.NewAnalyzer())
}
