package metafp

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/packages"
)

func isPredefined(pkgs []*packages.Package, pos token.Pos, willGenerated map[string]bool) bool {
	for _, p := range pkgs {
		file := p.Fset.Position(pos)
		if file.Filename != "" {
			if willGenerated[file.Filename] {
				return false
			}
		}
	}
	return true
}

func GetFunctionList(pkgs []*packages.Package, excludeFile map[string]bool) map[string]bool {
	funcList := map[string]bool{}
	for _, p := range pkgs {
		s := p.Types.Scope()
		for _, n := range s.Names() {
			o := s.Lookup(n)
			if _, ok := o.Type().(*types.Signature); ok {
				if isPredefined(pkgs, o.Pos(), excludeFile) {
					funcList[o.Name()] = true
				}
			}
		}
	}
	return funcList
}

func GetMethodList(pkgs []*packages.Package, targetType types.Type, excludeFile map[string]bool) map[string]bool {
	ret := map[string]bool{}

	switch tp := targetType.(type) {
	case *types.Named:
		for m := range tp.Methods() {
			if isPredefined(pkgs, m.Pos(), excludeFile) {
				ret[m.Name()] = true
			}
		}
	}

	return ret
}
