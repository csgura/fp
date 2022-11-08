package main

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/internal/max"
)

func main() {
	genfp.Generate("fp", "func_gen.go", func(f genfp.Writer) {
		for i := 3; i < max.Func; i++ {
			fmt.Fprintf(f, "type Func%d[%s, R any] func(%s) R\n", i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, i))

		}
		for i := 3; i < max.Func; i++ {

			// 			fmt.Fprintf(f, `
			// func(r Func%d[%s,R]) Tupled() Func1[Tuple%d[%s],R] {
			// 	return func(t Tuple%d[%s]) R {
			// 		return r(t.Unapply())
			// 	}
			// }
			// `, i, genfp.FuncTypeArgs(1, i), i, genfp.FuncTypeArgs(1, i), i, genfp.FuncTypeArgs(1, i))

			// 			fmt.Fprintf(f, `
			// func(r Func%d[%s,R]) Curried() %s {
			// 	return func(a1 A1) %s {
			// 		return Func%d[%s,R](func(%s) R {
			// 			return r(%s)
			// 		}).Curried()
			// 	}
			// }
			// `, i, genfp.FuncTypeArgs(1, i), curriedType(1, i), curriedType(2, i), i-1, genfp.FuncTypeArgs(2, i), genfp.FuncDeclArgs(2, i), genfp.FuncCallArgs(1, i))

			for j := i - 1; j < i; j++ {
				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyFirst%d(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, genfp.FuncTypeArgs(1, i), j, genfp.FuncDeclArgs(1, j), i-j, genfp.FuncTypeArgs(j+1, i), genfp.FuncDeclArgs(j+1, i), genfp.FuncCallArgs(1, i))
			}

			for j := i - 1; j < i; j++ {
				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyLast%d(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, genfp.FuncTypeArgs(1, i), j, genfp.FuncDeclArgs(i-j+1, i), i-j, genfp.FuncTypeArgs(1, i-j), genfp.FuncDeclArgs(1, i-j), genfp.FuncCallArgs(1, i))
			}

			/* 일부만 아규먼트 적용하는 함수는
			   너무 많은 code 를 생성해내서, 삭제

			   			fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyFirst(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(1, 1), i-1, genfp.FuncTypeArgs(2, i), genfp.FuncDeclArgs(2, i), genfp.FuncCallArgs(1, i))

			   			for j := 2; j < i; j++ {
			   				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyFirst%d(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, genfp.FuncTypeArgs(1, i), j, genfp.FuncDeclArgs(1, j), i-j, genfp.FuncTypeArgs(j+1, i), genfp.FuncDeclArgs(j+1, i), genfp.FuncCallArgs(1, i))
			   			}

			   			fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyLast(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, genfp.FuncTypeArgs(1, i), genfp.FuncDeclArgs(i, i), i-1, genfp.FuncTypeArgs(1, i-1), genfp.FuncDeclArgs(1, i-1), genfp.FuncCallArgs(1, i))

			   			for j := 2; j < i; j++ {
			   				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) ApplyLast%d(%s) Func%d[%s,R] {
			   	return func(%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, genfp.FuncTypeArgs(1, i), j, genfp.FuncDeclArgs(i-j+1, i), i-j, genfp.FuncTypeArgs(1, i-j), genfp.FuncDeclArgs(1, i-j), genfp.FuncCallArgs(1, i))

			   			}

			   			for j := 2; j < i; j++ {
			   				fmt.Fprintf(f, `
			   func(r Func%d[%s,R]) Apply%s(%s) Func%d[%s,%s,R] {
			   	return func(%s,%s) R {
			   		return r(%s)
			   	}
			   }
			   `, i, genfp.FuncTypeArgs(1, i), ordinalName[j], genfp.FuncDeclArgs(j, j), i-1, genfp.FuncTypeArgs(1, j-1), genfp.FuncTypeArgs(j+1, i), genfp.FuncDeclArgs(1, j-1), genfp.FuncDeclArgs(j+1, i), genfp.FuncCallArgs(1, i))

			   			}

			*/

			// 			if i < max.Shift {
			// 				fmt.Fprintf(f, `
			// func(r Func%d[%s,R]) Shift() Func%d[%s,A1,R] {
			// 	return func(%s , a1 A1) R {
			// 		return r(%s)
			// 	}
			// }
			// `, i, genfp.FuncTypeArgs(1, i), i, genfp.FuncTypeArgs(2, i), genfp.FuncDeclArgs(2, i), genfp.FuncCallArgs(1, i))
			// 			}

		}

		for i := 3; i < max.Compose; i++ {
			fmt.Fprintf(f, `
func Compose%d[%s,R any] ( %s ) Func1[A1,R] {
	return Compose2(f1, Compose%d(%s))
}
			`, i, genfp.FuncTypeArgs(1, i), genfp.FuncChain(1, i), i-1, strings.ReplaceAll(genfp.FuncCallArgs(2, i), "a", "f"))
		}

		for i := 2; i < max.Func; i++ {
			fmt.Fprintf(f, `
func Id%d[%s,R any] ( %s, r R) R {
		return r
}
			`, i, genfp.FuncTypeArgs(1, i-1), genfp.FuncDeclArgs(1, i-1))
		}
	})

	genfp.Generate("fp", "tuple_gen.go", func(f genfp.Writer) {
		_ = f.GetImportedName(types.NewPackage("fmt", "fmt"))

		for i := 2; i < max.Product; i++ {
			fmt.Fprintf(f, "type Tuple%d[%s any] struct {\n", i, genfp.TypeArgs("T", 1, i))

			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "    I%d T%d\n", j, j)
			}
			fmt.Fprintf(f, "}\n\n")

			fmt.Fprintf(f, `
func (r Tuple%d[%s]) Head() T1 {
	return r.I1;
}
`, i, genfp.TypeArgs("T", 1, i))

			fmt.Fprintf(f, `
			func (r Tuple%d[%s]) Tail() Tuple%d[%s] {
				return Tuple%d[%s]{%s};
			}
			`, i, genfp.TypeArgs("T", 1, i), i-1, genfp.TypeArgs("T", 2, i), i-1, genfp.TypeArgs("T", 2, i), genfp.FuncCallArgs(2, i, "r.I"))

			// fmt.Fprintf(f, `
			// func (r Tuple%d[%s]) ToHList() %s {
			// 	return hlist.Concat( r.Head(), r.Tail().ToHList())
			// }
			// `, i, genfp.TypeArgs("T", 1, i), consType(1, i))

			fmt.Fprintf(f, `
func (r Tuple%d[%s]) String() string {
	return fmt.Sprintf("(%s)", %s)
}
`, i, genfp.TypeArgs("T", 1, i), genfp.FormatStr(1, i), genfp.FuncCallArgs(1, i, "r.I"))

			fmt.Fprintf(f, `
func (r Tuple%d[%s]) Unapply() (%s) {
	return %s
}
`, i, genfp.TypeArgs("T", 1, i), genfp.TypeArgs("T", 1, i), genfp.FuncCallArgs(1, i, "r.I"))

		}

	})

	genfp.Generate("fp", "labelled_gen.go", func(f genfp.Writer) {
		_ = f.GetImportedName(types.NewPackage("fmt", "fmt"))

		for i := 2; i < max.Product; i++ {
			fmt.Fprintf(f, "type Labelled%d[%s Named] struct {\n", i, genfp.TypeArgs("T", 1, i))

			for j := 1; j <= i; j++ {
				fmt.Fprintf(f, "    I%d T%d\n", j, j)
			}
			fmt.Fprintf(f, "}\n\n")

			fmt.Fprintf(f, `
func (r Labelled%d[%s]) Head() T1 {
	return r.I1;
}
`, i, genfp.TypeArgs("T", 1, i))

			fmt.Fprintf(f, `
			func (r Labelled%d[%s]) Tail() Labelled%d[%s] {
				return Labelled%d[%s]{%s};
			}
			`, i, genfp.TypeArgs("T", 1, i), i-1, genfp.TypeArgs("T", 2, i), i-1, genfp.TypeArgs("T", 2, i), genfp.FuncCallArgs(2, i, "r.I"))

			// fmt.Fprintf(f, `
			// func (r Labelled%d[%s]) ToHList() %s {
			// 	return hlist.Concat( r.Head(), r.Tail().ToHList())
			// }
			// `, i, genfp.TypeArgs("T", 1, i), consType(1, i))

			fmt.Fprintf(f, `
func (r Labelled%d[%s]) String() string {
	return fmt.Sprintf("(%s)", %s)
}
`, i, genfp.TypeArgs("T", 1, i), genfp.FormatStr(1, i), genfp.FuncCallArgs(1, i, "r.I"))

			fmt.Fprintf(f, `
func (r Labelled%d[%s]) Unapply() (%s) {
	return %s
}
`, i, genfp.TypeArgs("T", 1, i), genfp.TypeArgs("T", 1, i), genfp.FuncCallArgs(1, i, "r.I"))

		}

	})

}
