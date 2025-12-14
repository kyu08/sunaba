package finderr

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const (
	doc = "finderr is ..."
)

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "finderr",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

// run 3文字以下のパッケージ変数を探す
func run(pass *analysis.Pass) (any, error) {
	for _, p := range pass.Pkg.Imports() {
		if p.Path() == "fmt" {
			obj := p.Scope().Lookup("Stringer")
			fmt.Printf("obj: %v\n", obj)
		}
	}

	return nil, nil
}
