package genfp

import (
	"fmt"
	"go/types"
	"strings"
)

func asPtr[T any](v T) *T {
	return &v
}

func NameParamReplaced(w Writer, pk *types.Package, realtp *types.Named, p *types.TypeParam) func(string, ...any) string {
	return func(newname string, fmtargs ...any) string {
		tpname := realtp.Origin().Obj().Name()
		nameWithPkg := tpname
		if realtp.Obj().Pkg() != nil && realtp.Obj().Pkg().Path() != pk.Path() {
			alias := w.GetImportedName(realtp.Obj().Pkg())

			nameWithPkg = fmt.Sprintf("%s.%s", alias, tpname)
		}

		if realtp.TypeArgs() != nil {
			args := []string{}
			for i := 0; i < realtp.TypeArgs().Len(); i++ {
				ta := realtp.TypeArgs().At(i)
				if tp, ok := ta.(*types.TypeParam); ok {
					if tp.Obj().Name() == p.Obj().Name() {
						args = append(args, fmt.Sprintf(newname, fmtargs...))
					} else {
						args = append(args, w.TypeName(pk, ta))
					}

				} else {
					args = append(args, w.TypeName(pk, ta))
				}
			}

			argsstr := strings.Join(args, ",")

			return fmt.Sprintf("%s[%s]", nameWithPkg, argsstr)
		} else {

			return nameWithPkg

		}
	}

}

func WriteMonadFunctions(w Writer, md GenerateMonadFunctionsDirective) {
	fmt.Printf("generateo for %s\n", md.TargetType)

	tp := md.TargetType.TypeArgs()
	tpargs := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return fmt.Sprintf("A %s", w.TypeName(md.Package.Types, tp.Constraint()))
			} else {
				return fmt.Sprintf("%s %s", tp.Obj().Name(), w.TypeName(md.Package.Types, tp.Constraint()))
			}
		}
		return ""

	}), func(v string) bool { return v != "" }), ",")

	rettype := NameParamReplaced(w, md.Package.Types, md.TargetType, md.TypeParm)

	srctype := rettype("A")
	rettp := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return "R"

			} else {
				return tp.Obj().Name()
			}
		}
		return ""
	}), func(v string) bool { return v != "" }), ",")

	fmt.Printf("tpargs = %s\n", tpargs)
	fmt.Fprintf(w, `
		func Map[%s, R any](m %s,  f func(%s) R) %s {
			return FlatMap(m, fp.Compose2(f, Pure[%s]))
		}
	`, tpargs, srctype, md.TypeParm, rettype("R"), rettp)

	fmt.Fprintf(w, `
		// haskell 의 <$
		// map . const 와 같은 함수
		func Replace[%s, R any](s %s, b R) %s {
			return Map(s, fp.Const[%s](b))
		}
	`, tpargs, srctype, rettype("R"), md.TypeParm)

	fmt.Fprintf(w, `
		func Map2[%s, B, R any](first %s, second %s, fab func(%s, B) R) %s {
			return FlatMap(first, func(a %s) %s {
				return Map(second, func(b B) R {
					return fab(a, b)
				})
			})
		}
	`, tpargs, srctype, rettype("B"), md.TypeParm, rettype("R"),
		md.TypeParm, rettype("R"),
	)

	w.AddImport(types.NewPackage("github.com/csgura/fp/product", "product"))
	fmt.Fprintf(w, `
		func Zip[%s, B any](first %s, second %s) %s {
			return Map2(first, second, product.Tuple2)
		}

	`, tpargs, srctype, rettype("B"), rettype(fmt.Sprintf("fp.Tuple2[%s,B]", md.TypeParm)),
	)

	fmt.Fprintf(w, `
		func Ap[%s, B any](tfab %s, ta %s) %s {
			return FlatMap(tfab, func(fab fp.Func1[A, B]) %s {
				return Map(ta, fab)
			})
		}

	`, tpargs, rettype("fp.Func1[%s,B]", md.TypeParm), rettype("%s", md.TypeParm), rettype("B"),
		rettype("B"),
	)

}
