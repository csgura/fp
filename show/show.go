package show

import (
	"fmt"
	"strings"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/mutable"
)

type Derives[T any] interface {
}

func New[T any](f func(T) string) fp.Show[T] {
	return fp.ShowFunc[T](f)
}

func NewIndent[T any](f func(T, fp.ShowOption) string) fp.Show[T] {
	return fp.ShowIndentFunc[T](f)
}

func NewAppend[T any](f func(buf []string, t T, option fp.ShowOption) []string) fp.Show[T] {
	return fp.ShowAppendFunc[T](f)
}

var Time = New(func(t time.Time) string {
	return t.Format(time.RFC3339)
})

var String = NewIndent(func(t string, opt fp.ShowOption) string {
	if opt.OmitEmpty && t == "" {
		return ""
	}
	return `"` + t + `"`
})

func Int[T fp.ImplicitInt]() fp.Show[T] {
	return fp.Sprint[T]()
}

func Number[T fp.ImplicitNum]() fp.Show[T] {
	return fp.Sprint[T]()
}

var Bool = New(func(t bool) string {
	if t {
		return "true"
	}
	return "false"
})

func Ptr[T any](tshow lazy.Eval[fp.Show[T]]) fp.Show[*T] {
	return NewAppend(func(buf []string, pt *T, opt fp.ShowOption) []string {
		if pt != nil {
			return tshow.Get().Append(buf, *pt, opt)
		}
		if opt.OmitEmpty {
			return append(buf, "")
		}
		return append(buf, "nil")
	})
}

func Given[T fmt.Stringer]() fp.Show[T] {
	return New(func(t T) string {
		return t.String()
	})
}

var HNil = New(func(hlist.Nil) string {
	return "Nil"
})

func appendSeq(buf []string, typeName string, itr fp.Iterator[string], opt fp.ShowOption) []string {
	childOpt := opt.IncreaseIndent()

	showseq := as.Seq(itr.ToSeq())
	if opt.OmitEmpty && showseq.IsEmpty() {
		return nil
	}
	if opt.Indent != "" && showseq.Exists(fp.Test(as.Func2(strings.Contains), "\n")) {
		return append(buf, typeName, "", "{\n", childOpt.CurrentIndent(), showseq.MakeString(",\n"+childOpt.CurrentIndent()), "\n", opt.CurrentIndent(), "}")
		//		return fmt.Sprintf("%s {\n%s%s\n%s}", typeName, childOpt.CurrentIndent(), showseq.MakeString(",\n"+childOpt.CurrentIndent()), opt.CurrentIndent())
	} else {
		if showseq.IsEmpty() {
			return append(buf, typeName, " {}")
		}
		if opt.Indent != "" {
			return append(buf, typeName, " { ", showseq.MakeString(", "), " }")

		}
		return append(buf, typeName, "{", showseq.MakeString(","), "}")
	}
}

func makeString(s fp.Seq[[]string], sep string) []string {
	ret := make([]string, 0, len(s)*2)

	for i, v := range s {
		if i != 0 {
			ret = append(ret, sep)
		}
		ret = append(ret, v...)
	}
	return ret
}
func appendSeq2(buf []string, typeName string, itr fp.Iterator[[]string], opt fp.ShowOption) []string {
	childOpt := opt.IncreaseIndent()

	showseq := as.Seq(itr.ToSeq())
	if opt.OmitEmpty && showseq.IsEmpty() {
		return nil
	}
	if opt.Indent != "" && showseq.Exists(func(v []string) bool {
		return as.Seq(v).Exists(fp.Test(as.Func2(strings.Contains), "\n"))
	}) {
		return append(
			append(
				append(buf, typeName, "", "{\n", childOpt.CurrentIndent()),
				makeString(showseq, ",\n"+childOpt.CurrentIndent())...,
			),
			"\n", opt.CurrentIndent(), "}",
		)
		//		return fmt.Sprintf("%s {\n%s%s\n%s}", typeName, childOpt.CurrentIndent(), showseq.MakeString(",\n"+childOpt.CurrentIndent()), opt.CurrentIndent())
	} else {
		if showseq.IsEmpty() {
			return append(buf, typeName, " {}")
		}
		if opt.Indent != "" {
			return append(
				append(
					append(buf, typeName, " { "),
					makeString(showseq, ", ")...,
				),
				" }",
			)

		}
		return append(
			append(
				append(buf, typeName, "{"),
				makeString(showseq, ",")...,
			),
			"}",
		)
	}
}

func Seq[T any](tshow fp.Show[T]) fp.Show[fp.Seq[T]] {
	return NewAppend(func(buf []string, s fp.Seq[T], opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()
		childStr := iterator.Map(iterator.FromSeq(s), fp.Flip(as.Curried3(tshow.Append)(nil))(childOpt))
		return appendSeq2(buf, "Seq", childStr, opt)
	})
}

func Set[V any](showv fp.Show[V]) fp.Show[fp.Set[V]] {
	return NewAppend(func(buf []string, v fp.Set[V], opt fp.ShowOption) []string {
		opt = opt.IncreaseIndent()

		showset := iterator.Map(v.Iterator(), func(v V) []string {
			return showv.Append(nil, v, opt)
		})

		return appendSeq2(buf, "Set", showset, opt)

	})
}

func isZero(s []string) bool {
	return len(s) == 0
}

func Map[K, V any](showk fp.Show[K], showv fp.Show[V]) fp.Show[fp.Map[K, V]] {
	return NewAppend(func(buf []string, v fp.Map[K, V], opt fp.ShowOption) []string {

		childOpt := opt.IncreaseIndent()

		showmap := iterator.Map(v.Iterator(), func(t fp.Tuple2[K, V]) []string {
			valuestr := showv.Append(nil, t.I2, childOpt)
			if isEmptyString(valuestr) {
				return nil
			}
			return append([]string{showk.Show(t.I1), ": "}, valuestr...)
		}).FilterNot(isZero)

		return appendSeq2(buf, "Map", showmap, opt)

	})
}

func GoMap[K comparable, V any](showk fp.Show[K], showv fp.Show[V]) fp.Show[map[K]V] {
	return NewAppend(func(buf []string, v map[K]V, opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()

		showmap := iterator.Map(mutable.MapOf(v).Iterator(), func(t fp.Tuple2[K, V]) []string {
			valuestr := showv.Append(nil, t.I2, childOpt)
			if isEmptyString(valuestr) {
				return nil
			}
			return append([]string{showk.Show(t.I1), ": "}, valuestr...)
		}).FilterNot(isZero)
		return appendSeq2(buf, "Map", showmap, opt)
	})
}

func Slice[T any](tshow fp.Show[T]) fp.Show[[]T] {
	return NewAppend(func(buf []string, s []T, opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()
		childStr := iterator.Map(iterator.FromSeq(s), fp.Flip(as.Curried3(tshow.Append)(nil))(childOpt))
		return appendSeq2(buf, "Seq", childStr, opt)
	})
}

func Option[T any](tshow fp.Show[T]) fp.Show[fp.Option[T]] {
	return NewAppend(func(buf []string, s fp.Option[T], opt fp.ShowOption) []string {
		if s.IsDefined() {
			return append(tshow.Append(append(buf, "Some("), s.Get(), opt.IncreaseIndent()), ")")
		}
		if opt.OmitEmpty {
			return nil
		}
		return append(buf, "None")
	})
}

func isEmptyString(s []string) bool {
	if s == nil {
		return true
	}
	if len(s) == 1 && s[0] == "" {
		return true
	}

	return false
}

func Named[T fp.NamedField[A], A any](ashow fp.Show[A]) fp.Show[T] {
	return NewAppend(func(buf []string, s T, opt fp.ShowOption) []string {
		valuestr := ashow.Append(nil, s.Value(), opt)
		if isEmptyString(valuestr) {
			return nil
		}
		if opt.Indent != "" {
			return append(append(buf, s.Name(), ": "), valuestr...)
		}
		return append(append(buf, s.Name(), ":"), valuestr...)

	})
}

func HConsLabelled[H fp.Named, T hlist.HList](hshow fp.Show[H], tshow fp.Show[T]) fp.Show[hlist.Cons[H, T]] {
	return NewAppend(func(buf []string, list hlist.Cons[H, T], opt fp.ShowOption) []string {

		hstr := hshow.Append(nil, list.Head(), opt)
		tstr := tshow.Append(nil, list.Tail(), opt)

		if isEmptyString(hstr) {
			if list.Tail().IsNil() {
				return nil
			}
			return tstr
		}
		if !list.Tail().IsNil() && !isEmptyString(tstr) {
			if opt.Indent != "" {
				return append(append(append(buf, hstr...), ",\n", opt.CurrentIndent()), tstr...)
				//return fmt.Sprintf("%s,\n%s%s", hstr, opt.CurrentIndent(), tstr)
			}
			return append(append(append(buf, hstr...), ","), tstr...)
		}
		return append(buf, hstr...)
	})
}

func TupleHCons[H any, T hlist.HList](hshow fp.Show[H], tshow fp.Show[T]) fp.Show[hlist.Cons[H, T]] {
	return NewAppend(func(buf []string, list hlist.Cons[H, T], opt fp.ShowOption) []string {

		hstr := hshow.Append(buf, list.Head(), opt)
		tstr := tshow.Append(nil, list.Tail(), opt)

		if !list.Tail().IsNil() {
			if opt.Indent != "" {
				return append(append(hstr, ", "), tstr...)

			}
			return append(append(hstr, ","), tstr...)
		}
		return hstr
	})
}

func HCons[H any, T hlist.HList](hshow fp.Show[H], tshow fp.Show[T]) fp.Show[hlist.Cons[H, T]] {
	return NewAppend(func(buf []string, list hlist.Cons[H, T], opt fp.ShowOption) []string {

		hstr := hshow.Append(buf, list.Head(), opt)
		tstr := tshow.Append(nil, list.Tail(), opt)

		if opt.Indent != "" {
			return append(append(hstr, " :: "), tstr...)

		}
		return append(append(hstr, "::"), tstr...)

	})
}

func Generic[A, Repr any](gen fp.Generic[A, Repr], reprShow fp.Show[Repr]) fp.Show[A] {
	return NewAppend(func(buf []string, a A, opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()
		valueStr := reprShow.Append(nil, gen.To(a), childOpt)
		if opt.OmitEmpty && isEmptyString(valueStr) {
			return nil
		}

		if gen.Kind == fp.GenericKindNewType {
			return append(append(append(buf, gen.Type, "("), valueStr...), ")")
			//return append(buf, fmt.Sprintf("%s(%s)", gen.Type, valueStr))
		} else if gen.Kind == fp.GenericKindTuple {
			return append(append(append(buf, gen.Type, "("), valueStr...), ")")
			//			return append(buf, fmt.Sprintf("%s(%s)", gen.Type, valueStr))
		}

		if opt.Indent != "" {
			return append(append(append(buf, gen.Type, " {\n", childOpt.CurrentIndent()), valueStr...), "\n", opt.CurrentIndent(), "}")

			//return append(buf, fmt.Sprintf("%s {\n%s%s\n%s}", gen.Type, childOpt.CurrentIndent(), valueStr, opt.CurrentIndent()))
		} else {
			return append(append(append(buf, gen.Type, "{"), valueStr...), "}")
			//return append(buf, fmt.Sprintf("%s{%s}", gen.Type, valueStr))

		}
	})
}
