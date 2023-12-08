package show

//go:generate go run github.com/csgura/fp/cmd/gombok

import (
	"fmt"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/seq"
)

// indent two space and omit empty
var Pretty = fp.ShowOption{
	Indent:           "  ",
	OmitEmpty:        true,
	SpaceAfterComma:  true,
	SpaceAfterColon:  true,
	SpaceBeforeBrace: true,
	SpaceWithinBrace: true,
	TrailingComma:    false,
}

var Space = fp.ShowOption{
	OmitEmpty:        true,
	SpaceAfterComma:  true,
	SpaceAfterColon:  true,
	SpaceBeforeBrace: false,
	SpaceWithinBrace: true,
}

var Json = fp.ShowOption{
	OmitEmpty:             true,
	OmitTypeName:          true,
	SquareBracketForArray: true,
	NullForNil:            true,
	QouteNames:            true,
}

var JsonSpace = fp.ShowOption{
	OmitEmpty:             true,
	OmitTypeName:          true,
	SquareBracketForArray: true,
	NullForNil:            true,
	SpaceAfterComma:       true,
	SpaceAfterColon:       true,
	SpaceBeforeBrace:      false,
	SpaceWithinBrace:      true,
	QouteNames:            true,
}

var PrettyJson = fp.ShowOption{
	Indent:                "  ",
	OmitEmpty:             true,
	OmitTypeName:          true,
	SquareBracketForArray: true,
	NullForNil:            true,
	SpaceAfterComma:       true,
	SpaceAfterColon:       true,
	SpaceBeforeBrace:      true,
	SpaceWithinBrace:      true,
	QouteNames:            true,
}

type Derives[T any] interface {
}

func ContraMap[T, U any](instance fp.Show[T], fn func(U) T) fp.Show[U] {
	return NewAppend(func(buf []string, u U, opt fp.ShowOption) []string {
		return instance.Append(buf, fn(u), opt)
	})
}

func New[T any](f func(T) string) fp.Show[T] {
	return fp.ShowFunc[T](f)
}

func NewIndent[T any](f func(T, fp.ShowOption) string) fp.Show[T] {
	return fp.ShowIndentFunc[T](f)
}

func NewAppend[T any](f func(buf []string, t T, opt fp.ShowOption) []string) fp.Show[T] {
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
		return append(buf, nullForNil(opt))
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

func nullForNil(opt fp.ShowOption) string {
	if opt.NullForNil {
		return "null"
	}
	return "nil"
}

// func spaceBeforeBrace(opt fp.ShowOption) string {
// 	if opt.SpaceBeforeBrace {
// 		return " "
// 	}
// 	return ""
// }

func spaceBetweenTypeAndBrace(opt fp.ShowOption) string {
	if opt.SpaceBeforeBrace && !opt.OmitTypeName {
		return " "
	}
	return ""
}

func spaceAfterComma(opt fp.ShowOption) string {
	if opt.SpaceAfterComma {
		return ", "
	}
	return ","
}

func spaceAfterColon(opt fp.ShowOption) string {
	if opt.SpaceAfterColon {
		return ": "
	}
	return ":"
}

func spaceWithinBrace(opt fp.ShowOption) string {
	if opt.SpaceWithinBrace {
		return " "
	}
	return ""
}

func spaceBeforeHCons(opt fp.ShowOption) string {
	if opt.SpaceAfterComma {
		return " "
	}
	return ""
}

func spaceAfterHCons(opt fp.ShowOption) string {
	if opt.SpaceAfterComma {
		return " "
	}
	return ""
}

func omitTypeName(name string, opt fp.ShowOption) string {
	if opt.OmitTypeName {
		return ""
	}
	return name
}

func quoteNames(name string, opt fp.ShowOption) string {
	if opt.QouteNames {
		return `"` + name + `"`
	}
	return name
}

func arrayOpen(opt fp.ShowOption) string {
	if opt.SquareBracketForArray {
		return "["
	}
	return "{"
}

func arrayClose(opt fp.ShowOption) string {
	if opt.SquareBracketForArray {
		return "]"
	}
	return "}"
}

func trailingComma(opt fp.ShowOption) string {
	if opt.TrailingComma {
		return ","
	}
	return ""
}

func Seq[T any](tshow fp.Show[T]) fp.Show[fp.Seq[T]] {
	return NewAppend(func(buf []string, s fp.Seq[T], opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()
		childStr := iterator.Map(iterator.FromSeq(s), fp.Flip(as.Curried3(tshow.Append)(nil))(childOpt))
		return appendSeq(buf, "Seq", childStr, opt)
	})
}

func Set[V any](showv fp.Show[V]) fp.Show[fp.Set[V]] {
	return NewAppend(func(buf []string, v fp.Set[V], opt fp.ShowOption) []string {
		opt = opt.IncreaseIndent()

		showset := iterator.Map(v.Iterator(), func(v V) []string {
			return showv.Append(nil, v, opt)
		})

		return appendSeq(buf, "Set", showset, opt)

	})
}

func isZero(s []string) bool {
	return len(s) == 0
}

func Map[K, V any](showk fp.Show[K], showv fp.Show[V]) fp.Show[fp.Map[K, V]] {
	return NewAppend(func(buf []string, v fp.Map[K, V], opt fp.ShowOption) []string {

		childOpt := opt.IncreaseIndent()

		keyshow := seq.Sort(iterator.Map(v.Iterator(), as.Func2(product.MapKey[K, V, string]).ApplyLast(showk.Show)).ToSeq(), ord.GivenField(fp.Tuple2[string, V].Head))

		showmap := iterator.Map(iterator.FromSeq(keyshow), func(t fp.Tuple2[string, V]) []string {
			valuestr := showv.Append(nil, t.I2, childOpt)
			if isEmptyString(valuestr) {
				return nil
			}
			return append([]string{t.I1, spaceAfterColon(opt)}, valuestr...)
		}).FilterNot(isZero)

		return appendMap(buf, "Map", showmap, opt)

	})
}

func GoMap[K comparable, V any](showk fp.Show[K], showv fp.Show[V]) fp.Show[map[K]V] {
	return ContraMap(Map(showk, showv), func(u map[K]V) fp.Map[K, V] {
		return mutable.MapOf(u)
	})
}

func Slice[T any](tshow fp.Show[T]) fp.Show[[]T] {
	return ContraMap(Seq(tshow), func(u []T) fp.Seq[T] {
		return u
	})
}

func Option[T any](tshow fp.Show[T]) fp.Show[fp.Option[T]] {
	return NewAppend(func(buf []string, s fp.Option[T], opt fp.ShowOption) []string {
		if s.IsDefined() {
			if opt.OmitTypeName {
				return tshow.Append(buf, s.Get(), opt)
			}
			return append(tshow.Append(append(buf, "Some("), s.Get(), opt.IncreaseIndent()), ")")
		}
		if opt.OmitEmpty {
			return nil
		}
		if opt.OmitTypeName {
			return append(buf, nullForNil(opt))
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
		return append(append(buf, quoteNames(s.Name(), opt), spaceAfterColon(opt)), valuestr...)

	})
}

func structFieldSeparator(opt fp.ShowOption) string {
	if opt.Indent != "" {
		return ",\n" + opt.CurrentIndent()
	}
	return spaceAfterComma(opt)
}

func Labelled2[N1, N2 fp.Named](ins1 fp.Show[N1], ins2 fp.Show[N2]) fp.Show[fp.Labelled2[N1, N2]] {
	return NewAppend(func(buf []string, t fp.Labelled2[N1, N2], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(AsAppender(ins1, t.I1)(nil, opt), AsAppender(ins2, t.I2)(nil, opt)).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

// @fp.Generate
var _ = genfp.GenerateFromUntil{
	File: "show_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/iterator", Name: "iterator"},
	},
	From:  3,
	Until: genfp.MaxProduct,
	Template: `
func Labelled{{.N}}[{{TypeArgs 1 .N}} fp.Named]({{DeclTypeClassArgs 1 .N "fp.Show"}}) fp.Show[fp.Labelled{{.N}}[{{TypeArgs 1 .N}}]] {
	return NewAppend(func(buf []string, t fp.Labelled{{.N}}[{{TypeArgs 1 .N}}], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			{{- range $idx := Range 1 .N}}
			ins{{$idx}}.Append(nil, t.I{{$idx}}, opt),
			{{- end}}
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}
	`,
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
			}
			return append(append(append(buf, hstr...), spaceAfterComma(opt)), tstr...)
		}
		return append(buf, hstr...)
	})
}

func TupleHCons[H any, T hlist.HList](hshow fp.Show[H], tshow fp.Show[T]) fp.Show[hlist.Cons[H, T]] {
	return NewAppend(func(buf []string, list hlist.Cons[H, T], opt fp.ShowOption) []string {

		hstr := hshow.Append(buf, list.Head(), opt)
		tstr := tshow.Append(nil, list.Tail(), opt)

		if !list.Tail().IsNil() {
			return append(append(hstr, spaceAfterComma(opt)), tstr...)
		}
		return hstr
	})
}

func HCons[H any, T hlist.HList](hshow fp.Show[H], tshow fp.Show[T]) fp.Show[hlist.Cons[H, T]] {
	return NewAppend(func(buf []string, list hlist.Cons[H, T], opt fp.ShowOption) []string {

		if opt.SquareBracketForArray {
			if buf == nil {
				buf = append(buf, "[", spaceWithinBrace(opt))
			}

			if !list.Tail().IsNil() {
				hstr := hshow.Append(buf, list.Head(), opt)
				hstr = append(hstr, spaceAfterComma(opt))
				return tshow.Append(hstr, list.Tail(), opt)
			}

			hstr := hshow.Append(buf, list.Head(), opt)
			return append(hstr, spaceWithinBrace(opt), "]")
		} else {
			hstr := hshow.Append(buf, list.Head(), opt)
			tstr := tshow.Append(nil, list.Tail(), opt)

			return append(append(hstr, spaceBeforeHCons(opt), "::", spaceAfterHCons(opt)), tstr...)
		}
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
			if opt.OmitTypeName {
				return append(buf, valueStr...)
			}
			return append(append(append(buf, omitTypeName(gen.Type, opt), spaceBetweenTypeAndBrace(opt), "(", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), ")")
		} else if gen.Kind == fp.GenericKindTuple {
			if opt.SquareBracketForArray {
				return append(append(append(buf, omitTypeName(gen.Type, opt), spaceBetweenTypeAndBrace(opt), "[", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), "]")

			} else {
				return append(append(append(buf, omitTypeName(gen.Type, opt), spaceBetweenTypeAndBrace(opt), "(", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), ")")

			}
		}

		if opt.Indent != "" {
			return append(append(append(buf, omitTypeName(gen.Type, opt), spaceBetweenTypeAndBrace(opt), "{\n", childOpt.CurrentIndent()), valueStr...), trailingComma(opt), "\n", opt.CurrentIndent(), "}")

		} else {
			return append(append(append(buf, omitTypeName(gen.Type, opt), spaceBetweenTypeAndBrace(opt), "{", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), "}")

		}
	})
}
