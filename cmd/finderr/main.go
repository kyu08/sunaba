package main

import (
	"github.com/kyu08/finderr"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(finderr.Analyzer) }
