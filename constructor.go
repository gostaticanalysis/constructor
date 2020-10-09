package constructor

import (
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

const doc = `constructor reports whether name of a constructor like function does not begin "New"`

var Analyzer = &analysis.Analyzer{
	Name: "constructor",
	Doc:  doc,
	Run:  run,
}

var errType = types.Universe.Lookup("error").Type()

func run(pass *analysis.Pass) (interface{}, error) {
	scope := pass.Pkg.Scope()

	var tns []*types.TypeName
	for _, n := range scope.Names() {
		tn, _ := scope.Lookup(n).(*types.TypeName)
		if tn != nil && tn.Exported() && isStruct(tn.Type()) {
			tns = append(tns, tn)
		}
	}

	for _, n := range scope.Names() {
		fun, _ := scope.Lookup(n).(*types.Func)
		if fun == nil {
			continue
		}

		if strings.HasPrefix(fun.Name(), "New") || !fun.Exported() {
			continue
		}

		sig, _ := fun.Type().(*types.Signature)
		if sig == nil || sig.Recv() != nil {
			continue
		}

		rets := sig.Results()
		if rets.Len() != 1 && rets.Len() != 2 {
			continue
		}

		if rets.Len() == 2 && !types.Identical(rets.At(1).Type(), errType) {
			continue
		}

		typ := rets.At(0).Type()
		if ptr, _ := typ.(*types.Pointer); ptr != nil {
			typ = ptr.Elem()
		}

		for _, tn := range tns {
			if types.Identical(typ, tn.Type()) {
				pass.Reportf(fun.Pos(), `name of a constructor like function must begin "New"`)
			}
		}
	}

	return nil, nil
}

func isStruct(typ types.Type) bool {
	_, ok := analysisutil.Under(typ).(*types.Struct)
	return ok
}
