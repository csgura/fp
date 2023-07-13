package show

import (
	"fmt"
	"strings"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
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

var Time = New(func(t time.Time) string {
	return t.Format(time.RFC3339)
})

var String = NewIndent(func(t string, opt fp.ShowOption) string {
	if opt.OmitEmpty && t == "" {
		return ""
	}
	return fmt.Sprintf(`"%s"`, t)
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
	return NewIndent(func(pt *T, opt fp.ShowOption) string {
		if pt != nil {
			return tshow.Get().ShowIndent(*pt, opt)
		}
		if opt.OmitEmpty {
			return ""
		}
		return "nil"
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

func showSeq(typeName string, itr fp.Iterator[string], opt fp.ShowOption) string {
	childOpt := opt.IncreaseIndent()

	showseq := as.Seq(itr.ToSeq())
	if opt.OmitEmpty && showseq.IsEmpty() {
		return ""
	}
	if opt.Indent != "" && showseq.Exists(fp.Test(as.Func2(strings.Contains), "\n")) {
		return fmt.Sprintf("%s {\n%s%s\n%s}", typeName, childOpt.CurrentIndent(), showseq.MakeString(",\n"+childOpt.CurrentIndent()), opt.CurrentIndent())
	} else {
		if showseq.IsEmpty() {
			return typeName + " {}"
		}
		if opt.Indent != "" {
			return typeName + " { " + showseq.MakeString(", ") + " }"

		}
		return typeName + "{" + showseq.MakeString(",") + "}"
	}
}

func Seq[T any](tshow fp.Show[T]) fp.Show[fp.Seq[T]] {
	return NewIndent(func(s fp.Seq[T], opt fp.ShowOption) string {
		childOpt := opt.IncreaseIndent()
		return showSeq("Seq", iterator.Map(iterator.FromSeq(s), as.Func2(tshow.ShowIndent).ApplyLast(childOpt)), opt)
	})
}

func Set[V any](showv fp.Show[V]) fp.Show[fp.Set[V]] {
	return NewIndent(func(v fp.Set[V], opt fp.ShowOption) string {
		opt = opt.IncreaseIndent()

		showset := iterator.Map(v.Iterator(), func(v V) string {
			return showv.ShowIndent(v, opt)
		})

		return showSeq("Set", showset, opt)

	})
}

func Map[K, V any](showk fp.Show[K], showv fp.Show[V]) fp.Show[fp.Map[K, V]] {
	return NewIndent(func(v fp.Map[K, V], opt fp.ShowOption) string {

		childOpt := opt.IncreaseIndent()

		showmap := iterator.Map(v.Iterator(), func(t fp.Tuple2[K, V]) string {
			valuestr := showv.ShowIndent(t.I2, childOpt)
			if valuestr == "" {
				return ""
			}
			return fmt.Sprintf("%s: %s", showk.Show(t.I1), valuestr)
		}).FilterNot(eq.GivenValue(""))

		return showSeq("Map", showmap, opt)

	})
}

func GoMap[K comparable, V any](showk fp.Show[K], showv fp.Show[V]) fp.Show[map[K]V] {
	return NewIndent(func(v map[K]V, opt fp.ShowOption) string {
		childOpt := opt.IncreaseIndent()

		showmap := iterator.Map(mutable.MapOf(v).Iterator(), func(t fp.Tuple2[K, V]) string {
			valuestr := showv.ShowIndent(t.I2, childOpt)
			if valuestr == "" {
				return ""
			}
			return fmt.Sprintf("%s: %s", showk.Show(t.I1), valuestr)
		}).FilterNot(eq.GivenValue(""))
		return showSeq("Map", showmap, opt)
	})
}

func Slice[T any](tshow fp.Show[T]) fp.Show[[]T] {
	return NewIndent(func(s []T, opt fp.ShowOption) string {
		childOpt := opt.IncreaseIndent()
		return showSeq("Seq", iterator.Map(iterator.FromSeq(s), as.Func2(tshow.ShowIndent).ApplyLast(childOpt)), opt)
	})
}

func Option[T any](tshow fp.Show[T]) fp.Show[fp.Option[T]] {
	return NewIndent(func(s fp.Option[T], opt fp.ShowOption) string {
		if s.IsDefined() {
			return fmt.Sprintf("Some(%s)", tshow.ShowIndent(s.Get(), opt.IncreaseIndent()))
		}
		if opt.OmitEmpty {
			return ""
		}
		return "None"
	})
}

func Named[T fp.NamedField[A], A any](ashow fp.Show[A]) fp.Show[T] {
	return NewIndent(func(s T, opt fp.ShowOption) string {
		valuestr := ashow.ShowIndent(s.Value(), opt)
		if valuestr == "" {
			return ""
		}
		if opt.Indent != "" {
			return fmt.Sprintf("%s: %s", s.Name(), valuestr)
		}
		return fmt.Sprintf("%s:%s", s.Name(), valuestr)

	})
}

func HConsLabelled[H fp.Named, T hlist.HList](hshow fp.Show[H], tshow fp.Show[T]) fp.Show[hlist.Cons[H, T]] {
	return NewIndent(func(list hlist.Cons[H, T], opt fp.ShowOption) string {

		hstr := hshow.ShowIndent(list.Head(), opt)
		tstr := tshow.ShowIndent(list.Tail(), opt)

		if hstr == "" {
			if tstr == "Nil" {
				return ""
			}
			return tstr
		}
		if tstr != "Nil" && tstr != "" {
			if opt.Indent != "" {
				return fmt.Sprintf("%s,\n%s%s", hstr, opt.CurrentIndent(), tstr)
			}
			return fmt.Sprintf("%s,%s", hstr, tstr)
		}
		return hstr
	})
}

func HCons[H any, T hlist.HList](hshow fp.Show[H], tshow fp.Show[T]) fp.Show[hlist.Cons[H, T]] {
	return NewIndent(func(list hlist.Cons[H, T], opt fp.ShowOption) string {

		childOpt := opt.IncreaseIndent()

		hstr := hshow.ShowIndent(list.Head(), childOpt)
		tstr := tshow.ShowIndent(list.Tail(), opt)

		if opt.Indent != "" {
			return fmt.Sprintf("%s :: \n%s%s", hstr, opt.CurrentIndent(), tstr)
		}
		return fmt.Sprintf("%s :: %s", hstr, tstr)
	})
}

func ContraMap[A, B any](ashow fp.Show[A], fba func(B) A) fp.Show[B] {
	return NewIndent(func(b B, opt fp.ShowOption) string {
		return ashow.ShowIndent(fba(b), opt)
	})
}

func Generic[A, Repr any](gen fp.Generic[A, Repr], reprShow fp.Show[Repr]) fp.Show[A] {
	return NewIndent(func(a A, opt fp.ShowOption) string {
		childOpt := opt.IncreaseIndent()
		valueStr := reprShow.ShowIndent(gen.To(a), childOpt)
		if opt.OmitEmpty && valueStr == "" {
			return ""
		}

		if gen.Kind == fp.GenericKindNewType {

			return fmt.Sprintf("%s(%s)", gen.Type, valueStr)
		}

		if opt.Indent != "" {
			return fmt.Sprintf("%s {\n%s%s\n%s}", gen.Type, childOpt.CurrentIndent(), valueStr, opt.CurrentIndent())
		} else {
			return fmt.Sprintf("%s{%s}", gen.Type, valueStr)

		}
	})
}
