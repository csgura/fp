package generator

import (
	"bytes"
	"fmt"
	"go/types"

	"github.com/csgura/fp/genfp"
)

func WriteTraverseFunctions(w Writer, md GenerateMonadFunctionsDirective, definedFunc map[string]bool) {

	tp := md.TargetType.TypeArgs()
	tpargs := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return fmt.Sprintf("A %s", w.TypeName(md.Package, tp.Constraint()))
			} else {
				return fmt.Sprintf("%s %s", tp.Obj().Name(), w.TypeName(md.Package, tp.Constraint()))
			}
		}
		return ""

	}), func(v string) bool { return v != "" }), ",")

	tpargs1 := seqMakeString(seqFilter(iterate(tp.Len(), tp.At, func(i int, t types.Type) string {
		if tp, ok := t.(*types.TypeParam); ok {
			if tp.Obj().Name() == md.TypeParm.Obj().Name() {
				return fmt.Sprintf("A1 %s", w.TypeName(md.Package, tp.Constraint()))
			} else {
				return fmt.Sprintf("%s %s", tp.Obj().Name(), w.TypeName(md.Package, tp.Constraint()))
			}
		}
		return ""

	}), func(v string) bool { return v != "" }), ",")

	rettype := NameParamReplaced(w, md.Package, md.TargetType, md.TypeParm)

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

	if _, ok := definedFunc["FoldSliceM"]; ok {
		w.Render(`
			func TraverseSeq[{{.tpargs}}, R any](sa fp.Seq[A], fa func(A) {{monad "R"}}) {{monad "fp.Seq[R]"}} {
				return FoldSliceM(sa, fp.Seq[R]{}, func(acc fp.Seq[R], a A) {{monad "fp.Seq[R]"}} {
						return Map(fa(a), acc.Add)
				})
			}
			
			func TraverseSlice[{{.tpargs}}, R any](sa fp.Slice[A], fa func(A) {{monad "R"}}) {{monad "fp.Slice[R]"}} {
				return FoldSliceM(sa, fp.Slice[R]{}, func(acc fp.Slice[R], a A) {{monad "fp.Slice[R]"}} {
					return Map(fa(a), func(v R) fp.Slice[R] {
						return append(acc,v)
					})
				})
			}

			func Sequence[{{.tpargs}}](tsa []{{monad "A"}}) {{monad "fp.Slice[A]"}} {
				ret := FoldSliceM(tsa, fp.Slice[A]{}, func(t1 fp.Slice[A], t2 {{monad "A"}}) {{monad "fp.Slice[A]"}} {
					return Map(t2, func(v A) fp.Slice[A] {
						return append(t1, v)
					})
				})
	
				return ret
			}

		`, funcs, param)
	} else {
		w.Render(`
			func TraverseSeq[{{.tpargs}}, R any](sa fp.Seq[A], fa func(A) {{monad "R"}}) {{monad "fp.Seq[R]"}} {
				return FoldM(fp.IteratorOfSeq(sa), fp.Seq[R]{}, func(acc fp.Seq[R], a A) {{monad "fp.Seq[R]"}} {
						return Map(fa(a), acc.Add)
				})
			}
			
			func TraverseSlice[{{.tpargs}}, R any](sa fp.Slice[A], fa func(A) {{monad "R"}}) {{monad "fp.Slice[R]"}} {
				return FoldM(fp.IteratorOfSeq(sa), fp.Slice[R]{}, func(acc fp.Slice[R], a A) {{monad "fp.Slice[R]"}} {
					return Map(fa(a), func(v R) fp.Slice[R] {
						return append(acc,v)
					})
				})
			}

			func Sequence[{{.tpargs}}](tsa []{{monad "A"}}) {{monad "fp.Slice[A]"}} {
				ret := FoldM(iterator.FromSlice(tsa), fp.Slice[A]{}, func(t1 fp.Slice[A], t2 {{monad "A"}}) {{monad "fp.Slice[A]"}} {
					return Map(t2, func(v A) fp.Slice[A] {
						return append(t1, v)
					})
				})
	
				return ret
			}
		`, funcs, param)
	}

	w.Render(`

	func Traverse[{{.tpargs}}, R any](ia fp.Iterator[A], fn func(A) {{monad "R"}}) {{monad "fp.Iterator[R]"}} {
		return Map(FoldM(ia, fp.Seq[R]{}, func(acc fp.Seq[R], a A) {{monad "fp.Seq[R]"}} {
			return Map(fn(a), acc.Add)
		}), iterator.FromSeq)
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
	
	func TraverseSliceFunc[{{.tpargs}}, R any](far func(A) {{monad "R"}}) func(fp.Slice[A]) {{monad "fp.Slice[R]"}} {
		return func(seqA fp.Slice[A]) {{monad "fp.Slice[R]"}} {
			return TraverseSlice(seqA, far)
		}
	}

	func FlatMapTraverseSeq[{{.tpargs}}, B any](ta {{monad "fp.Seq[A]"}}, f func(v A) {{monad "B"}}) {{monad "fp.Seq[B]"}} {
		return FlatMap(ta, TraverseSeqFunc(f))
	}
	
	func FlatMapTraverseSlice[{{.tpargs}}, B any](ta {{monad "fp.Slice[A]"}}, f func(v A) {{monad "B"}}) {{monad "fp.Slice[B]"}} {
		return FlatMap(ta, TraverseSliceFunc(f))
	}


	func SequenceIterator[{{.tpargs}}](ita fp.Iterator[{{monad "A"}}]) {{monad "fp.Iterator[A]"}} {
		ret := FoldM(ita, fp.Seq[A]{}, func(t1 fp.Seq[A], t2 {{monad "A"}}) {{monad "fp.Seq[A]"}} {
			return Map(t2, t1.Add)
		})
		return Map(ret, iterator.FromSeq)

	}
	
	`, funcs, param)

}
