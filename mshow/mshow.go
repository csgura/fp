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
	"github.com/csgura/fp/slice"
)

type Derives[T any] interface {
}

type Buffer []string

func (r Buffer) AppendColon(opt fp.ShowOption) Buffer {
	return show.AppendColon(r, opt)
}

func (r Buffer) AppendSpaceBetweenTypeAndBrace(opt fp.ShowOption) Buffer {
	return show.AppendSpaceBetweenTypeAndBrace(r, opt)
}

func (r Buffer) AppendTypeName(name string, opt fp.ShowOption) Buffer {
	return show.AppendTypeName(r, name, opt)
}

func (r Buffer) AppendSpaceAfterHCons(opt fp.ShowOption) Buffer {
	return show.AppendSpaceAfterHCons(r, opt)
}

func (r Buffer) AppendSpaceBeforeHCons(opt fp.ShowOption) Buffer {
	return show.AppendSpaceBeforeHCons(r, opt)
}

func (r Buffer) AppendSpaceWithinBrace(opt fp.ShowOption) Buffer {
	return show.AppendSpaceWithinBrace(r, opt)
}

func (r Buffer) AppendComma(opt fp.ShowOption) Buffer {
	return show.AppendComma(r, opt)
}

func (r Buffer) AppendStringLiteral(strliteral string, opt fp.ShowOption) Buffer {
	return show.AppendStringLiteral(r, strliteral, opt)
}

func (r Buffer) AppendStruct(name string, opt fp.ShowOption, fields ...fp.Tuple2[string, func(Buffer, fp.ShowOption) Buffer]) Buffer {
	return show.AppendStruct(r, name, opt, slice.MapValue(fields, func(a func(Buffer, fp.ShowOption) Buffer) show.Appender {
		return func(buf []string, opt fp.ShowOption) []string {
			return a(buf, opt)
		}
	})...)
}

func (r Buffer) AppendFieldName(name string, fieldName string, opt fp.ShowOption) Buffer {
	return show.AppendFieldName(r, fieldName, opt)
}

func (r Buffer) Append(v ...string) Buffer {
	return append(r, v...)
}

type Show[T any] = func(buf Buffer, t T, option fp.ShowOption) Buffer

func ContraMap[T, U any](instance Show[T], fn func(U) T) Show[U] {
	return NewAppend(func(buf Buffer, u U, opt fp.ShowOption) Buffer {
		return instance(buf, fn(u), opt)
	})
}

func New[T any](f func(T) string) Show[T] {
	return NewAppend(func(buf Buffer, t T, option fp.ShowOption) Buffer {
		return append(buf, f(t))
	})
}

func NewIndent[T any](f func(T, fp.ShowOption) string) Show[T] {
	return NewAppend(func(buf Buffer, t T, option fp.ShowOption) Buffer {
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
	return func(buf Buffer, t T, option fp.ShowOption) Buffer {
		return buf.Append(fp.Sprint[T]().Show(t))
	}
}

func Number[T fp.ImplicitNum]() Show[T] {
	return func(buf Buffer, t T, option fp.ShowOption) Buffer {
		return buf.Append(fp.Sprint[T]().Show(t))
	}
}

var Bool = New(func(t bool) string {
	if t {
		return "true"
	}
	return "false"
})

func Ptr[T any](tshow lazy.Eval[Show[T]]) Show[*T] {
	return NewAppend(func(buf Buffer, pt *T, opt fp.ShowOption) Buffer {
		if pt != nil {
			return tshow.Get()(buf, *pt, opt)
		}
		if opt.OmitEmpty {
			return append(buf, "")
		}
		return show.AppendNil(buf, opt)
	})
}

func Given[T fmt.Stringer]() Show[T] {
	return NewAppend(func(buf Buffer, t T, opt fp.ShowOption) Buffer {
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
			return show.AppendNil(buf, opt)
		}
		return append(buf, t.String())
	})
}

var HNil = New(func(hlist.Nil) string {
	return "Nil"
})

func Set[V any](showv Show[V]) Show[fp.Set[V]] {
	return NewAppend(func(buf Buffer, v fp.Set[V], opt fp.ShowOption) Buffer {
		opt = opt.IncreaseIndent()

		showset := iterator.Map(v.Iterator(), func(v V) show.Appender {
			return AsAppender(showv, v)
		}).ToSeq()

		return show.AppendSlice(buf, "Set", showset, opt)

	})
}

func isZero(s []string) bool {
	return len(s) == 0
}

func FullShow[T any](s Show[T]) fp.Show[T] {
	return fp.ShowAppendFunc[T](func(buf []string, t T, opt fp.ShowOption) []string {
		return s(buf, t, opt)
	})
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
	return NewAppend(func(buf Buffer, v fp.Map[K, V], opt fp.ShowOption) Buffer {

		childOpt := opt.IncreaseIndent()

		keyshow := seq.Sort(iterator.Map(v.Iterator(), as.Func2(product.MapKey[K, V, string]).ApplyLast(pshow(showk))).ToSeq(), ord.GivenField(fp.Tuple2[string, V].Head))

		showmap := iterator.FilterMap(iterator.FromSeq(keyshow), func(t fp.Tuple2[string, V]) fp.Option[show.Appender] {
			valuestr := showv(nil, t.I2, childOpt)
			if isEmptyString(valuestr) {
				return option.None[show.Appender]()
			}
			if opt.QuoteNames == false && strings.HasPrefix(t.I1, `"`) && strings.HasSuffix(t.I1, `"`) {
				return option.Some[show.Appender](func(buf []string, opt fp.ShowOption) []string {
					return Buffer(buf).Append(t.I1[1 : len(t.I1)-1]).AppendColon(opt).Append(valuestr...)
				})
			}
			return option.Some[show.Appender](func(buf []string, opt fp.ShowOption) []string {
				return Buffer(buf).Append(t.I1).AppendColon(opt).Append(valuestr...)
			})
		}).ToSeq()

		return show.AppendMap(buf, "Map", showmap, opt)

	})
}

func GoMap[K comparable, V any](showk Show[K], showv Show[V]) Show[map[K]V] {
	return ContraMap(Map(showk, showv), func(u map[K]V) fp.Map[K, V] {
		return mutable.MapOf(u)
	})
}

func Seq[T any](tshow Show[T]) Show[fp.Seq[T]] {
	return NewAppend(func(buf Buffer, s fp.Seq[T], opt fp.ShowOption) Buffer {

		var childStr []show.Appender
		for _, v := range s {
			childStr = append(childStr, AsAppender(tshow, v))
		}
		return show.AppendSlice(buf, "Seq", childStr, opt)
	})
}

func Slice[T any](tshow Show[T]) Show[[]T] {
	return NewAppend(func(buf Buffer, s []T, opt fp.ShowOption) Buffer {
		var childStr []show.Appender
		for _, v := range s {
			childStr = append(childStr, AsAppender(tshow, v))
		}
		return show.AppendSlice(buf, "Seq", childStr, opt)
	})
}

func Option[T any](tshow Show[T]) Show[fp.Option[T]] {
	return NewAppend(func(buf Buffer, s fp.Option[T], opt fp.ShowOption) Buffer {
		if s.IsDefined() {
			if opt.OmitTypeName || opt.OmitObjectBrace {
				return tshow(buf, s.Get(), opt)
			}
			return append(tshow(append(buf, "Some("), s.Get(), opt.IncreaseIndent()), ")")
		}
		if opt.OmitEmpty {
			return nil
		}
		if opt.OmitTypeName || opt.OmitObjectBrace {
			return show.AppendNil(buf, opt)
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
	return NewAppend(func(buf Buffer, s T, opt fp.ShowOption) Buffer {
		valuestr := ashow(nil, s, opt)
		if isEmptyString(valuestr) {
			return nil
		}
		buf = show.AppendFieldName(buf, fp.ConvertNaming(name.Name(), opt.NamingCase), opt)
		buf = show.AppendColon(buf, opt)
		return append(buf, valuestr...)
	})
}

func Struct1[N1 any](names []fp.Named, ins1 Show[N1]) Show[minimal.Tuple1[N1]] {
	return NewAppend(func(buf Buffer, t minimal.Tuple1[N1], opt fp.ShowOption) Buffer {
		fields := seq.Of(
			AsAppender(Named(names[0], ins1), t.I1),
		)
		return show.AppendCommaSperated(buf, fields, opt)
	})
}

func Struct2[N1, N2 any](names []fp.Named, ins1 Show[N1], ins2 Show[N2]) Show[minimal.Tuple2[N1, N2]] {
	return NewAppend(func(buf Buffer, t minimal.Tuple2[N1, N2], opt fp.ShowOption) Buffer {
		fields := seq.Of(
			AsAppender(Named(names[0], ins1), t.I1),
			AsAppender(Named(names[1], ins2), t.I2),
		)
		return show.AppendCommaSperated(buf, fields, opt)
	})
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Generate
var _ = genfp.GenerateFromUntil{
	File: "show_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/slice", Name: "slice"},

		{Package: "github.com/csgura/fp/minimal", Name: "minimal"},
		{Package: "github.com/csgura/fp/show", Name: "show"},
	},
	From:  3,
	Until: genfp.MaxProduct,
	Template: `
func Struct{{.N}}[{{TypeArgs 1 .N}} any](names []fp.Named, {{DeclTypeClassArgs 1 .N "Show"}}) Show[minimal.Tuple{{.N}}[{{TypeArgs 1 .N}}]] {
	return NewAppend(func(buf Buffer, t minimal.Tuple{{.N}}[{{TypeArgs 1 .N}}], opt fp.ShowOption) Buffer {
		fields := slice.Of(
			{{- range $idx := Range 1 .N}}
			AsAppender(Named(names[{{dec $idx}}], ins{{$idx}}), t.I{{$idx}}),
			{{- end}}
		)
		return show.AppendCommaSperated(buf, fields, opt)
	})
}
	`,
}

func StructHCons[H any, T minimal.HList](hshow Show[H], tshow Show[T]) Show[minimal.Cons[H, T]] {
	return NewAppend(func(buf Buffer, list minimal.Cons[H, T], opt fp.ShowOption) Buffer {

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
			return buf.Append(hstr...).AppendComma(opt).Append(tstr...)
		}
		return append(buf, hstr...)
	})
}

func TupleHCons[H any, T hlist.HList](hshow Show[H], tshow Show[T]) Show[hlist.Cons[H, T]] {
	return NewAppend(func(buf Buffer, list hlist.Cons[H, T], opt fp.ShowOption) Buffer {

		hstr := hshow(buf, list.Head(), opt)
		tstr := tshow(nil, hlist.Tail(list), opt)

		if !hlist.IsNil(hlist.Tail(list)) {
			return hstr.AppendComma(opt).Append(tstr...)
		}
		return hstr
	})
}

func HCons[H any, T hlist.HList](hshow Show[H], tshow Show[T]) Show[hlist.Cons[H, T]] {
	return NewAppend(func(buf Buffer, list hlist.Cons[H, T], opt fp.ShowOption) Buffer {

		if opt.SquareBracketForArray {
			if buf == nil {
				buf = buf.Append("[").AppendSpaceWithinBrace(opt)
			}

			if !hlist.IsNil(hlist.Tail(list)) {
				hstr := hshow(buf, list.Head(), opt)
				hstr = hstr.AppendComma(opt)
				return tshow(hstr, hlist.Tail(list), opt)
			}

			hstr := hshow(buf, list.Head(), opt)
			return hstr.AppendSpaceWithinBrace(opt).Append("]")
		} else {
			hstr := hshow(buf, list.Head(), opt)
			tstr := tshow(nil, hlist.Tail(list), opt)

			return hstr.AppendSpaceBeforeHCons(opt).Append("::").AppendSpaceAfterHCons(opt).Append(tstr...)
		}
	})
}

func ContraGeneric[A, Repr any](name string, kind string, reprShow Show[Repr], to func(A) Repr) Show[A] {
	return NewAppend(func(buf Buffer, a A, opt fp.ShowOption) Buffer {
		return show.AppendGeneric(buf, name, kind, AsAppender(reprShow, to(a)), opt)
	})
}

func AsAppender[T any](tshow Show[T], t T) show.Appender {
	return func(buf []string, opt fp.ShowOption) []string {
		return tshow(buf, t, opt)
	}
}
