package autoimports

import (
	"flag"

	"golang.org/x/tools/go/analysis"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:             "autoimports",
		Doc:              "import custom packages into your package",
		Flags:            flagSet,
		Run:              run,
		RunDespiteErrors: true,
	}
}

const packageText = `_ "go.uber.org/automaxprocs"`

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.Name == "main" {
			isImport := false
			for _, imp := range file.Imports {
				if imp.Name != nil && imp.Path != nil {
					if imp.Name.Name == "_" && imp.Path.Value == `"go.uber.org/automaxprocs"` {
						isImport = true
					}
				}
			}
			if !isImport {
				// fmt.Println(file.Pos(), file.Name.Pos(), file.Name.End())
				pass.Report(analysis.Diagnostic{
					Pos:     file.Pos(),
					Message: "no import",
					SuggestedFixes: []analysis.SuggestedFix{{
						Message: "xxxxxx",
						TextEdits: []analysis.TextEdit{{
							Pos:     file.Pos(),
							NewText: []byte(packageText),
						}},
					}},
				})
			}
		}
	}

	return nil, nil
}
