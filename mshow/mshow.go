package mshow

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/minimal"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/show"
)

type Derives[T any] interface {
}

type Show[T any] = func(buf []string, t T, option fp.ShowOption) []string

func ContraMap[T, U any](instance Show[T], fn func(U) T) Show[U] {
	return NewAppend(func(buf []string, u U, opt fp.ShowOption) []string {
		return instance(buf, fn(u), opt)
	})
}

func New[T any](f func(T) string) Show[T] {
	return NewAppend(func(buf []string, t T, option fp.ShowOption) []string {
		return append(buf, f(t))
	})
}

func NewIndent[T any](f func(T, fp.ShowOption) string) Show[T] {
	return NewAppend(func(buf []string, t T, option fp.ShowOption) []string {
		return append(buf, f(t, option))
	})
}

func NewAppend[T any](f Show[T]) Show[T] {
	return f
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

func Int[T fp.ImplicitInt]() Show[T] {
	return fp.Sprint[T]().Append
}

func Number[T fp.ImplicitNum]() Show[T] {
	return fp.Sprint[T]().Append
}

var Bool = New(func(t bool) string {
	if t {
		return "true"
	}
	return "false"
})

func Ptr[T any](tshow lazy.Eval[Show[T]]) Show[*T] {
	return NewAppend(func(buf []string, pt *T, opt fp.ShowOption) []string {
		if pt != nil {
			return tshow.Get()(buf, *pt, opt)
		}
		if opt.OmitEmpty {
			return append(buf, "")
		}
		return append(buf, nullForNil(opt))
	})
}

func Given[T fmt.Stringer]() Show[T] {
	return NewAppend(func(buf []string, t T, opt fp.ShowOption) []string {
		if opt.UserOptions.Get("format") == option.Some("json") {
			b, err := json.Marshal(t)
			if err == nil {
				if len(b) > 0 {
					return append(buf, string(b))
				}
			}
			if opt.OmitEmpty {
				return append(buf, "")
			}
			return append(buf, nullForNil(opt))
		}
		return append(buf, t.String())
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
	if opt.QuoteNames {
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

func Seq[T any](tshow Show[T]) Show[fp.Seq[T]] {
	return NewAppend(func(buf []string, s fp.Seq[T], opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()

		var childStr [][]string
		for _, v := range s {
			childStr = append(childStr, tshow(nil, v, childOpt))
		}
		return appendSeq(buf, "Seq", iterator.FromSlice(childStr), opt)
	})
}

func Set[V any](showv Show[V]) Show[fp.Set[V]] {
	return NewAppend(func(buf []string, v fp.Set[V], opt fp.ShowOption) []string {
		opt = opt.IncreaseIndent()

		showset := iterator.Map(v.Iterator(), func(v V) []string {
			return showv(nil, v, opt)
		})

		return appendSeq(buf, "Set", showset, opt)

	})
}

func isZero(s []string) bool {
	return len(s) == 0
}

func FullShow[T any](s Show[T]) fp.Show[T] {
	return fp.ShowAppendFunc[T](s)
}

func MakeString[T any](a Show[T], t T, option fp.ShowOption) string {
	return strings.Join(a(nil, t, option), "")
}

func Stringer[T any](a Show[T], t T, option fp.ShowOption) fmt.Stringer {
	return fp.StringerFunc(as.Supplier3(MakeString[T], a, t, option))
}

func pshow[T any](ins Show[T]) func(t T) string {
	return func(t T) string {
		return MakeString(ins, t, fp.ShowOption{})
	}
}

func Map[K, V any](showk Show[K], showv Show[V]) Show[fp.Map[K, V]] {
	return NewAppend(func(buf []string, v fp.Map[K, V], opt fp.ShowOption) []string {

		childOpt := opt.IncreaseIndent()

		keyshow := seq.Sort(iterator.Map(v.Iterator(), as.Func2(product.MapKey[K, V, string]).ApplyLast(pshow(showk))).ToSeq(), ord.GivenField(fp.Tuple2[string, V].Head))

		showmap := iterator.Map(iterator.FromSeq(keyshow), func(t fp.Tuple2[string, V]) []string {
			valuestr := showv(nil, t.I2, childOpt)
			if isEmptyString(valuestr) {
				return nil
			}
			return append([]string{t.I1, spaceAfterColon(opt)}, valuestr...)
		}).FilterNot(isZero)

		return appendMap(buf, "Map", showmap, opt)

	})
}

func GoMap[K comparable, V any](showk Show[K], showv Show[V]) Show[map[K]V] {
	return ContraMap(Map(showk, showv), func(u map[K]V) fp.Map[K, V] {
		return mutable.MapOf(u)
	})
}

func Slice[T any](tshow Show[T]) Show[[]T] {
	return NewAppend(func(buf []string, s []T, opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()

		var childStr [][]string
		for _, v := range s {
			childStr = append(childStr, tshow(nil, v, childOpt))
		}
		return appendSeq(buf, "Seq", iterator.FromSlice(childStr), opt)
	})
}

func Option[T any](tshow Show[T]) Show[fp.Option[T]] {
	return NewAppend(func(buf []string, s fp.Option[T], opt fp.ShowOption) []string {
		if s.IsDefined() {
			if opt.OmitTypeName {
				return tshow(buf, s.Get(), opt)
			}
			return append(tshow(append(buf, "Some("), s.Get(), opt.IncreaseIndent()), ")")
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

func Named[T any](name fp.Named, ashow Show[T]) Show[T] {
	return NewAppend(func(buf []string, s T, opt fp.ShowOption) []string {
		valuestr := ashow(nil, s, opt)
		if isEmptyString(valuestr) {
			return nil
		}
		return append(append(buf, quoteNames(fp.ConvertNaming(name.Name(), opt.NamingCase), opt), spaceAfterColon(opt)), valuestr...)
	})
}

func structFieldSeparator(opt fp.ShowOption) string {
	if opt.Indent != "" {
		return ",\n" + opt.CurrentIndent()
	}
	return spaceAfterComma(opt)
}

func Struct1[N1 any](names []fp.Named, ins1 Show[N1]) Show[minimal.Tuple1[N1]] {
	return NewAppend(func(buf []string, t minimal.Tuple1[N1], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(Named(names[0], ins1)(nil, t.I1, opt)).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

func Struct2[N1, N2 any](names []fp.Named, ins1 Show[N1], ins2 Show[N2]) Show[minimal.Tuple2[N1, N2]] {
	return NewAppend(func(buf []string, t minimal.Tuple2[N1, N2], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(Named(names[0], ins1)(nil, t.I1, opt), Named(names[1], ins2)(nil, t.I2, opt)).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Generate
var _ = genfp.GenerateFromUntil{
	File: "show_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/minimal", Name: "minimal"},

		{Package: "github.com/csgura/fp/iterator", Name: "iterator"},
	},
	From:  3,
	Until: genfp.MaxProduct,
	Template: `
func Struct{{.N}}[{{TypeArgs 1 .N}} any](names []fp.Named, {{DeclTypeClassArgs 1 .N "Show"}}) Show[minimal.Tuple{{.N}}[{{TypeArgs 1 .N}}]] {
	return NewAppend(func(buf []string, t minimal.Tuple{{.N}}[{{TypeArgs 1 .N}}], opt fp.ShowOption) []string {
		return append(buf, makeString(iterator.Of(
			{{- range $idx := Range 1 .N}}
			Named(names[{{dec $idx}}],ins{{$idx}})(nil, t.I{{$idx}}, opt),
			{{- end}}
		).FilterNot(isEmptyString).ToSeq(), structFieldSeparator(opt))...)
	})
}
	`,
}

func StructHCons[H any, T minimal.HList](hshow Show[H], tshow Show[T]) Show[minimal.Cons[H, T]] {
	return NewAppend(func(buf []string, list minimal.Cons[H, T], opt fp.ShowOption) []string {

		hstr := hshow(nil, list.Head, opt)
		tstr := tshow(nil, list.Tail, opt)

		if isEmptyString(hstr) {
			if minimal.IsNil(list.Tail) {
				return nil
			}
			return tstr
		}
		if !minimal.IsNil(list.Tail) && !isEmptyString(tstr) {
			if opt.Indent != "" {
				return append(append(append(buf, hstr...), ",\n", opt.CurrentIndent()), tstr...)
			}
			return append(append(append(buf, hstr...), spaceAfterComma(opt)), tstr...)
		}
		return append(buf, hstr...)
	})
}

func TupleHCons[H any, T hlist.HList](hshow Show[H], tshow Show[T]) Show[hlist.Cons[H, T]] {
	return NewAppend(func(buf []string, list hlist.Cons[H, T], opt fp.ShowOption) []string {

		hstr := hshow(buf, list.Head(), opt)
		tstr := tshow(nil, hlist.Tail(list), opt)

		if !hlist.IsNil(hlist.Tail(list)) {
			return append(append(hstr, spaceAfterComma(opt)), tstr...)
		}
		return hstr
	})
}

func HCons[H any, T hlist.HList](hshow Show[H], tshow Show[T]) Show[hlist.Cons[H, T]] {
	return NewAppend(func(buf []string, list hlist.Cons[H, T], opt fp.ShowOption) []string {

		if opt.SquareBracketForArray {
			if buf == nil {
				buf = append(buf, "[", spaceWithinBrace(opt))
			}

			if !hlist.IsNil(hlist.Tail(list)) {
				hstr := hshow(buf, list.Head(), opt)
				hstr = append(hstr, spaceAfterComma(opt))
				return tshow(hstr, hlist.Tail(list), opt)
			}

			hstr := hshow(buf, list.Head(), opt)
			return append(hstr, spaceWithinBrace(opt), "]")
		} else {
			hstr := hshow(buf, list.Head(), opt)
			tstr := tshow(nil, hlist.Tail(list), opt)

			return append(append(hstr, spaceBeforeHCons(opt), "::", spaceAfterHCons(opt)), tstr...)
		}
	})
}

func ContraGeneric[A, Repr any](name string, kind string, reprShow Show[Repr], to func(A) Repr) Show[A] {
	return NewAppend(func(buf []string, a A, opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()
		valueStr := reprShow.Append(nil, to(a), childOpt)
		if opt.OmitEmpty && isEmptyString(valueStr) {
			return nil
		}

		if kind == fp.GenericKindNewType {
			if opt.OmitTypeName {
				return append(buf, valueStr...)
			}
			return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "(", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), ")")
		} else if kind == fp.GenericKindTuple {
			if opt.SquareBracketForArray {
				return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "[", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), "]")

			} else {
				return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "(", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), ")")

			}
		}

		if opt.Indent != "" {
			return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "{\n", childOpt.CurrentIndent()), valueStr...), trailingComma(opt), "\n", opt.CurrentIndent(), "}")

		} else {
			return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "{", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), "}")

		}
	})
}

func ContraGeneric[A, Repr any](name string, kind string, reprShow Show[Repr], to func(A) Repr) Show[A] {
	return NewAppend(func(buf []string, a A, opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()
		valueStr := reprShow(nil, to(a), childOpt)
		if opt.OmitEmpty && isEmptyString(valueStr) {
			return nil
		}

		if kind == fp.GenericKindNewType {
			if opt.OmitTypeName {
				return append(buf, valueStr...)
			}
			return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "(", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), ")")
		} else if kind == fp.GenericKindTuple {
			if opt.SquareBracketForArray {
				return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "[", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), "]")

			} else {
				return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "(", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), ")")

			}
		}

		if opt.Indent != "" {
			return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "{\n", childOpt.CurrentIndent()), valueStr...), trailingComma(opt), "\n", opt.CurrentIndent(), "}")

		} else {
			return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "{", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), "}")

		}
	})
}

func appendSeq(buf []string, typeName string, itr fp.Iterator[[]string], opt fp.ShowOption) []string {
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
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), arrayOpen(opt), "\n", childOpt.CurrentIndent()),
				makeString(showseq, ",\n"+childOpt.CurrentIndent())...,
			),
			trailingComma(opt), "\n", opt.CurrentIndent(), arrayClose(opt),
		)
		//		return fmt.Sprintf("%s {\n%s%s\n%s}", typeName, childOpt.CurrentIndent(), showseq.MakeString(",\n"+childOpt.CurrentIndent()), opt.CurrentIndent())
	} else {

		if showseq.IsEmpty() {
			return append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), arrayOpen(opt), arrayClose(opt))
		}

		return append(
			append(
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), arrayOpen(opt), spaceWithinBrace(opt)),
				makeString(showseq, spaceAfterComma(opt))...,
			),
			spaceWithinBrace(opt), arrayClose(opt),
		)
	}
}

func appendMap(buf []string, typeName string, itr fp.Iterator[[]string], opt fp.ShowOption) []string {
	childOpt := opt.IncreaseIndent()

	showseq := as.Seq(itr.ToSeq())
	if opt.OmitEmpty && showseq.IsEmpty() {
		return nil
	}

	if opt.Indent != "" {
		return append(
			append(
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), "{\n", childOpt.CurrentIndent()),
				makeString(showseq, ",\n"+childOpt.CurrentIndent())...,
			),
			trailingComma(opt), "\n", opt.CurrentIndent(), "}",
		)
		//		return fmt.Sprintf("%s {\n%s%s\n%s}", typeName, childOpt.CurrentIndent(), showseq.MakeString(",\n"+childOpt.CurrentIndent()), opt.CurrentIndent())
	} else {

		if showseq.IsEmpty() {
			return append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), " {}")
		}

		return append(
			append(
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), "{", spaceWithinBrace(opt)),
				makeString(showseq, spaceAfterComma(opt))...,
			),
			spaceWithinBrace(opt), "}",
		)
	}
}

func AsAppender[T any](tshow Show[T], t T) show.Appender {
	return func(buf []string, opt fp.ShowOption) []string {
		return tshow(buf, t, opt)
	}
}
