package main

import (
	"github.com/gostaticanalysis/constructor"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(constructor.Analyzer) }
