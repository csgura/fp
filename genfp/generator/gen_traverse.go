package generator

import (
	"bytes"
	"fmt"
	"go/types"

	"github.com/csgura/fp/genfp"
)

func WriteTraverseFunctions(w Writer, md GenerateMonadFunctionsDirective) {

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

	tpargs1 := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return fmt.Sprintf("A1 %s", w.TypeName(md.Package.Types, tp.Constraint()))
			} else {
				return fmt.Sprintf("%s %s", tp.Obj().Name(), w.TypeName(md.Package.Types, tp.Constraint()))
			}
		}
		return ""

	}), func(v string) bool { return v != "" }), ",")

	rettype := NameParamReplaced(w, md.Package.Types, md.TargetType, md.TypeParm)

	// srctype := rettype("A")
	// rettp := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
	// 	if tp, ok := t.(*types.TypeParam); ok {
	// 		if tp.Obj().Name() == md.TypeParm.Obj().Name() {
	// 			return "R"

	// 		} else {
	// 			return tp.Obj().Name()
	// 		}
	// 	}
	// 	return ""
	// }), func(v string) bool { return v != "" }), ",")
	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp", "fp"))

	funcs := map[string]any{
		"monad": rettype,
		"monadIns": func(start, until int) string {
			f := &bytes.Buffer{}
			for j := start; j <= until; j++ {
				if j != start {
					fmt.Fprintf(f, ", ")
				}
				fmt.Fprintf(f, "ins%d %s", j, rettype("A%d", j))
			}
			return f.String()
		},
		"monadTypes": func(start, until int) string {
			f := &bytes.Buffer{}
			for j := start; j <= until; j++ {
				if j != start {
					fmt.Fprintf(f, ", ")
				}
				fmt.Fprintf(f, "%s", rettype("A%d", j))
			}
			return f.String()
		},
	}
	param := map[string]any{
		"tpargs":  tpargs,
		"tpargs1": tpargs1,

		"tp": md.TypeParm.String(),
	}

	w.AddImport(genfp.NewImportPackage("github.com/csgura/fp/iterator", "iterator"))

	w.Render(`

	func Traverse[{{.tpargs}}, R any](ia fp.Iterator[A], fn func(A) {{monad "R"}}) {{monad "fp.Iterator[R]"}} {
		return Map(FoldM(ia, fp.Seq[R]{}, func(acc fp.Seq[R], a A) {{monad "fp.Seq[R]"}} {
			return Map(fn(a), acc.Add)
		}), iterator.FromSeq)
	}

	func TraverseSeq[{{.tpargs}}, R any](sa fp.Seq[A], fa func(A) {{monad "R"}}) {{monad "fp.Seq[R]"}} {
		return FoldM(fp.IteratorOfSeq(sa), fp.Seq[R]{}, func(acc fp.Seq[R], a A) {{monad "fp.Seq[R]"}} {
			return Map(fa(a), acc.Add)
		})
	}
	
	func TraverseSlice[{{.tpargs}}, R any](sa []A, fa func(A) {{monad "R"}}) {{monad "[]R"}} {
		return Map(TraverseSeq(sa, fa), fp.Seq[R].Widen)
	}

	
	func TraverseFunc[{{.tpargs}}, R any](far func(A) {{monad "R"}}) func(fp.Iterator[A]) {{monad "fp.Iterator[R]"}} {
		return func(iterA fp.Iterator[A]) {{monad "fp.Iterator[R]"}} {
			return Traverse(iterA, far)
		}
	}
	
	func TraverseSeqFunc[{{.tpargs}}, R any](far func(A) {{monad "R"}}) func(fp.Seq[A]) {{monad "fp.Seq[R]"}} {
		return func(seqA fp.Seq[A]) {{monad "fp.Seq[R]"}} {
			return TraverseSeq(seqA, far)
		}
	}
	
	func TraverseSliceFunc[{{.tpargs}}, R any](far func(A) {{monad "R"}}) func([]A) {{monad "[]R"}} {
		return func(seqA []A) {{monad "[]R"}} {
			return TraverseSlice(seqA, far)
		}
	}

	func FlatMapTraverseSeq[{{.tpargs}}, B any](ta {{monad "fp.Seq[A]"}}, f func(v A) {{monad "B"}}) {{monad "fp.Seq[B]"}} {
		return FlatMap(ta, TraverseSeqFunc(f))
	}
	
	func FlatMapTraverseSlice[{{.tpargs}}, B any](ta {{monad "[]A"}}, f func(v A) {{monad "B"}}) {{monad "[]B"}} {
		return FlatMap(ta, TraverseSliceFunc(f))
	}

	func Sequence[{{.tpargs}}](tsa []{{monad "A"}}) {{monad "[]A"}} {
		ret := FoldM(iterator.FromSeq(tsa), fp.Seq[A]{}, func(t1 fp.Seq[A], t2 {{monad "A"}}) {{monad "fp.Seq[A]"}} {
			return Map(t2, t1.Add)
		})
	
		return Map(ret, fp.Seq[A].Widen)
	}


	func SequenceIterator[{{.tpargs}}](ita fp.Iterator[{{monad "A"}}]) {{monad "fp.Iterator[A]"}} {
		ret := FoldM(ita, fp.Seq[A]{}, func(t1 fp.Seq[A], t2 {{monad "A"}}) {{monad "fp.Seq[A]"}} {
			return Map(t2, t1.Add)
		})
		return Map(ret, iterator.FromSeq)

	}
	
	`, funcs, param)

}
