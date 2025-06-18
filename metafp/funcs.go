package metafp

import (
	"go/types"

	"golang.org/x/tools/go/packages"
)

func GetFunctionList(pkgs []*packages.Package, excludeFile map[string]bool) map[string]bool {
	funcList := map[string]bool{}
	for _, p := range pkgs {
		s := p.Types.Scope()
		for _, n := range s.Names() {
			o := s.Lookup(n)
			if _, ok := o.Type().(*types.Signature); ok {
				file := p.Fset.Position(o.Pos()).Filename
				if !excludeFile[file] {
					funcList[o.Name()] = true
				}
			}
		}
	}
	return funcList
}
