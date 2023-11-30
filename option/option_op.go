//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package option

import (
	"reflect"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/product"
)

func Some[T any](v T) fp.Option[T] {
	return fp.None[T]().Recover(func() T {
		return v
	})
}

func None[T any]() fp.Option[T] {
	return fp.None[T]()
}

// 아규먼트를 무시하고 항상 None 을 리턴
func ConstNone[A, B any](a A) fp.Option[B] {
	return fp.None[B]()
}

func isNil(v reflect.Value) bool {
	k := v.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.Interface, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func Of[T any](v T) fp.Option[T] {
	var i any = v
	if i == nil {
		return None[T]()
	}

	rv := reflect.ValueOf(i)
	if isNil(rv) {
		return None[T]()
	}
	return Some(v)
}

func Ptr[T any](v *T) fp.Option[T] {
	if v == nil {
		return None[T]()
	}
	return Some(*v)
}

func String(v string) fp.Option[string] {
	return NonZero(v)
}

func NonZero[T comparable](t T) fp.Option[T] {
	if t == fp.Zero[T]() {
		return None[T]()
	}
	return Some(t)
}

func FromTry[T any](t fp.Try[T]) fp.Option[T] {
	if t.IsSuccess() {
		return Some(t.Get())
	}
	return None[T]()
}

func Ap[T, U any](t fp.Option[fp.Func1[T, U]], a fp.Option[T]) fp.Option[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Option[U] {
		return Map(a, f)
	})
}

func ApFunc[T, U any](t fp.Option[fp.Func1[T, U]], fopt func() fp.Option[T]) fp.Option[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Option[U] {
		return Map(fopt(), f)
	})
}

func Map[T, U any](opt fp.Option[T], f func(v T) U) fp.Option[U] {
	return FlatMap(opt, func(v T) fp.Option[U] {
		return Some(f(v))
	})
}

func MapSeqMap[A, B any](ta fp.Option[fp.Seq[A]], f func(v A) B) fp.Option[fp.Seq[B]] {
	return Map(ta, func(a fp.Seq[A]) fp.Seq[B] {
		return iterator.Map(iterator.FromSeq(a), f).ToSeq()
	})
}

func MapSliceMap[A, B any](ta fp.Option[[]A], f func(v A) B) fp.Option[[]B] {
	return Map(ta, func(a []A) []B {
		return iterator.Map(iterator.FromSeq(a), f).ToSeq()
	})
}

func Map2[A, B, U any](a fp.Option[A], b fp.Option[B], f func(A, B) U) fp.Option[U] {
	return FlatMap(a, func(v1 A) fp.Option[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		})
	})
}

// fp.With 의 option 버젼
// fp.With 가 Flip 과 사실상 같은 것처럼
// FlapMap 의 Flip 버젼과 동일
// var b fp.Option[B]
// a := option.Some(A{})
// a.FlatMap( option.With(A.WithB, b))
// 형태로 코딩 가능
func With[A, B any](withf func(A, B) A, v fp.Option[B]) func(A) fp.Option[A] {
	return Flap(Map(v, fp.Flip2(withf)))
}

func Lift[T, U any](f func(v T) U) func(fp.Option[T]) fp.Option[U] {
	return func(opt fp.Option[T]) fp.Option[U] {
		return Map(opt, f)
	}
}

func LiftA2[A1, A2, R any](f func(A1, A2) R) func(fp.Option[A1], fp.Option[A2]) fp.Option[R] {
	return func(a1 fp.Option[A1], a2 fp.Option[A2]) fp.Option[R] {
		return Map2(a1, a2, f)
	}
}

func LiftM[A, R any](fa func(v A) fp.Option[R]) func(fp.Option[A]) fp.Option[R] {
	return func(ta fp.Option[A]) fp.Option[R] {
		return Flatten(Map(ta, fa))
	}
}

// (a -> b -> m r) -> m a -> m b -> m r
// 하스켈에서는  liftM2 와 liftA2 는 같은 함수이고
// 위와 같은 함수는 존재하지 않음.
// hoogle 에서 검색해 보면 , liftJoin2 , bindM2 등의 이름으로 정의된 것이 있음.
// 하지만 ,  fp 패키지에서도   LiftA2 와 LiftM2 를 동일하게 하는 것은 낭비이고
// M 은 Monad 라는 뜻인데, Monad는 Flatten, FlatMap 의 의미가 있으니까
// LiftM2 를 다음과 같이 정의함.
func LiftM2[A, B, R any](fab func(A, B) fp.Option[R]) func(fp.Option[A], fp.Option[B]) fp.Option[R] {
	return func(a fp.Option[A], b fp.Option[B]) fp.Option[R] {
		return Flatten(Map2(a, b, fab))
	}
}

func Compose[A, B, C any](f1 func(A) fp.Option[B], f2 func(B) fp.Option[C]) func(A) fp.Option[C] {
	return func(a A) fp.Option[C] {
		return FlatMap(f1(a), f2)
	}
}

func Compose2[A, B, C any](f1 func(A) fp.Option[B], f2 func(B) fp.Option[C]) func(A) fp.Option[C] {
	return func(a A) fp.Option[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposePure[A, B any](fab func(A) B) func(A) fp.Option[B] {
	return fp.Compose(fab, Some)
}

func FlatMap[T, U any](opt fp.Option[T], fn func(v T) fp.Option[U]) fp.Option[U] {
	if opt.IsDefined() {
		return fn(opt.Get())
	}
	return None[U]()
}

func FlatMapTraverseSeq[A, B any](ta fp.Option[fp.Seq[A]], f func(v A) fp.Option[B]) fp.Option[fp.Seq[B]] {
	return FlatMap(ta, TraverseSeqFunc(f))
}

func FlatMapTraverseSlice[A, B any](ta fp.Option[[]A], f func(v A) fp.Option[B]) fp.Option[[]B] {
	return FlatMap(ta, TraverseSliceFunc(f))
}

func Flatten[T any](opt fp.Option[fp.Option[T]]) fp.Option[T] {
	return FlatMap(opt, func(v fp.Option[T]) fp.Option[T] {
		return v
	})
}

func FlatPtr[T any](opt fp.Option[*T]) fp.Option[T] {
	return FlatMap(opt, func(v *T) fp.Option[T] {
		return Ptr(v)
	})
}

// 하스켈 : m( a -> r ) -> a -> m r
// 스칼라 : M[ A => R ] => A => M[R]
// 하스켈이나 스칼라의 기본 패키지에는 이런 기능을 하는 함수가 없는데,
// hoogle 에서 검색해 보면
// https://hoogle.haskell.org/?hoogle=m%20(%20a%20-%3E%20b)%20-%3E%20a%20-%3E%20m%20b
// ?? 혹은 flap 이라는 이름으로 정의된 함수가 있음
func Flap[A, R any](tfa fp.Option[fp.Func1[A, R]]) func(A) fp.Option[R] {
	return func(a A) fp.Option[R] {
		return Ap(tfa, Some(a))
	}
}

// 하스켈 : m( a -> b -> r ) -> a -> b -> m r
func Flap2[A, B, R any](tfab fp.Option[fp.Func1[A, fp.Func1[B, R]]]) fp.Func1[A, fp.Func1[B, fp.Option[R]]] {
	return func(a A) fp.Func1[B, fp.Option[R]] {
		return Flap(Ap(tfab, Some(a)))
	}
}

// (a -> b -> r) -> m a -> b -> m r
// Map 호출 후에 Flap 을 호출 한 것
//
// https://hoogle.haskell.org/?hoogle=%28+a+-%3E+b+-%3E++r+%29+-%3E+m+a+-%3E++b+-%3E+m+r+&scope=set%3Astackage
// liftOp 라는 이름으로 정의된 것이 있음
func FlapMap[A, B, R any](tfab func(A, B) R, a fp.Option[A]) func(B) fp.Option[R] {
	return Flap(Map(a, as.Curried2(tfab)))
}

// ( a -> b -> m r ) -> m a -> b -> m r
//
//	Flatten . FlapMap
//
// https://hoogle.haskell.org/?hoogle=(%20a%20-%3E%20b%20-%3E%20m%20r%20)%20-%3E%20m%20a%20-%3E%20%20b%20-%3E%20m%20r%20
// om , ==<<  이름으로 정의된 것이 있음
func FlatFlapMap[A, B, R any](fab func(A, B) fp.Option[R], ta fp.Option[A]) func(B) fp.Option[R] {
	return fp.Compose(FlapMap(fab, ta), Flatten)
}

// FlatMap 과는 아규먼트 순서가 다른 함수로
// Go 나 Java 에서는 메소드 레퍼런스를 이용하여,  객체내의 메소드를 리턴 타입만 lift 된 형태로 리턴하게 할 수 있음.
// Method 라는 이름보다  Ap 와 비슷한 이름이 좋을 거 같은데
// Ap와 비슷한 이름으로 하기에는 Ap 와 타입이 너무 다름.
func Method1[A, B, R any](ta fp.Option[A], fab func(a A, b B) R) func(B) fp.Option[R] {
	return FlapMap(fab, ta)
}

func FlatMethod1[A, B, R any](ta fp.Option[A], fab func(a A, b B) fp.Option[R]) func(B) fp.Option[R] {
	return FlatFlapMap(fab, ta)
}

func Method2[A, B, C, R any](ta fp.Option[A], fabc func(a A, b B, c C) R) func(B, C) fp.Option[R] {

	return curried.Revert2(Flap2(Map(ta, as.Curried3(fabc))))
	// return func(b B, c C) fp.Option[R] {
	// 	return Map(a, func(a A) R {
	// 		return cf(a, b, c)
	// 	})
	// }
}

func FlatMethod2[A, B, C, R any](ta fp.Option[A], fabc func(a A, b B, c C) fp.Option[R]) func(B, C) fp.Option[R] {

	return curried.Revert2(curried.Compose2(Flap2(Map(ta, as.Curried3(fabc))), Flatten))

	// return func(b B, c C) fp.Option[R] {
	// 	return FlatMap(ta, func(a A) fp.Option[R] {
	// 		return cf(a, b, c)
	// 	})
	// }
}

func Zip[A, B any](c1 fp.Option[A], c2 fp.Option[B]) fp.Option[fp.Tuple2[A, B]] {
	return Map2(c1, c2, product.Tuple2)
}

func Zip3[A, B, C any](c1 fp.Option[A], c2 fp.Option[B], c3 fp.Option[C]) fp.Option[fp.Tuple3[A, B, C]] {
	return LiftA3(as.Tuple3[A, B, C])(c1, c2, c3)
}

func SequenceIterator[T any](optItr fp.Iterator[fp.Option[T]]) fp.Option[fp.Iterator[T]] {
	return iterator.Fold(optItr, Some(iterator.Empty[T]()), LiftA2(fp.Iterator[T].Appended))
}

func Traverse[T, U any](itr fp.Iterator[T], fn func(T) fp.Option[U]) fp.Option[fp.Iterator[U]] {
	return iterator.FoldOption(itr, iterator.Empty[U](), func(acc fp.Iterator[U], v T) fp.Option[fp.Iterator[U]] {
		return Map(fn(v), acc.Appended)
	})
}

func TraverseSeq[T, U any](seq fp.Seq[T], fn func(T) fp.Option[U]) fp.Option[fp.Seq[U]] {
	return Map(TraverseSlice(seq, fn), as.Seq)
}

func TraverseSlice[T, U any](seq []T, fn func(T) fp.Option[U]) fp.Option[[]U] {
	return Map(Traverse(fp.IteratorOfSeq(seq), fn), fp.Iterator[U].ToSeq)
}

func TraverseFunc[A, R any](far func(A) fp.Option[R]) func(fp.Iterator[A]) fp.Option[fp.Iterator[R]] {
	return func(iterA fp.Iterator[A]) fp.Option[fp.Iterator[R]] {
		return Traverse(iterA, far)
	}
}

func TraverseSeqFunc[A, R any](far func(A) fp.Option[R]) func(fp.Seq[A]) fp.Option[fp.Seq[R]] {
	return func(seqA fp.Seq[A]) fp.Option[fp.Seq[R]] {
		return TraverseSeq(seqA, far)
	}
}

func TraverseSliceFunc[A, R any](far func(A) fp.Option[R]) func([]A) fp.Option[[]R] {
	return func(seqA []A) fp.Option[[]R] {
		return TraverseSlice(seqA, far)
	}
}

func Sequence[T any](optSeq []fp.Option[T]) fp.Option[[]T] {
	return Map(SequenceIterator(fp.IteratorOfSeq(optSeq)), fp.Iterator[T].ToSeq)
}

func Fold[A, B any](s fp.Option[A], zero B, f func(B, A) B) B {
	if s.IsEmpty() {
		return zero
	}

	return f(zero, s.Get())
}

func FoldRight[A, B any](s fp.Option[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if s.IsEmpty() {
		return lazy.Done(zero)
	}

	return f(s.Get(), lazy.Done(zero))
}

func ToSeq[T any](r fp.Option[T]) fp.Seq[T] {
	if r.IsDefined() {
		return fp.Seq[T]{r.Get()}
	}
	return nil
}

func Iterator[T any](r fp.Option[T]) fp.Iterator[T] {
	return fp.IteratorOfSeq(ToSeq(r))
}

func Deref[R any, T fp.Deref[R]](opt fp.Option[T]) fp.Option[R] {
	return Map(opt, T.Deref)
}

type MonadChain1[H hlist.Header[HT], HT, A, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A, R]]
}

func (r MonadChain1[H, HT, A, R]) Map(a func(HT) A) fp.Option[R] {
	return r.FlatMap(func(h HT) fp.Option[A] {
		return Some(a(h))
	})
}

func (r MonadChain1[H, HT, A, R]) HListMap(a func(H) A) fp.Option[R] {
	return r.HListFlatMap(func(h H) fp.Option[A] {
		return Some(a(h))
	})
}

func (r MonadChain1[H, HT, A, R]) HListFlatMap(a func(H) fp.Option[A]) fp.Option[R] {
	av := FlatMap(r.h, func(v H) fp.Option[A] {
		return a(v)
	})

	return r.ApOption(av)
}

func (r MonadChain1[H, HT, A, R]) FlatMap(a func(HT) fp.Option[A]) fp.Option[R] {
	av := FlatMap(r.h, func(v H) fp.Option[A] {
		return a(v.Head())
	})

	return r.ApOption(av)
}

func (r MonadChain1[H, HT, A, R]) ApOption(a fp.Option[A]) fp.Option[R] {
	return Ap(r.fn, a)
}

func (r MonadChain1[H, HT, A, R]) Ap(a A) fp.Option[R] {
	return r.ApOption(Some(a))
}

func (r MonadChain1[H, HT, A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Option[R] {

	av := FlatMap(r.h, func(v H) fp.Option[A] {
		return a()
	})
	return r.ApOption(av)
}
func (r MonadChain1[H, HT, A, R]) ApFunc(a func() A) fp.Option[R] {

	av := Map(r.h, func(v H) A {
		return a()
	})
	return r.ApOption(av)
}

func Chain1[A, R any](fn fp.Func1[A, R]) MonadChain1[hlist.Nil, hlist.Nil, A, R] {
	return MonadChain1[hlist.Nil, hlist.Nil, A, R]{Some(hlist.Empty()), Some(fn)}
}

type ApplicativeFunctor1[A, R any] struct {
	fn fp.Option[fp.Func1[A, R]]
}

func (r ApplicativeFunctor1[A, R]) ApOption(a fp.Option[A]) fp.Option[R] {
	return Ap(r.fn, a)
}

func (r ApplicativeFunctor1[A, R]) Ap(a A) fp.Option[R] {
	return r.ApOption(Some(a))
}

func (r ApplicativeFunctor1[A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Option[R] {

	return ApFunc(r.fn, a)
}
func (r ApplicativeFunctor1[A, R]) ApFunc(a func() A) fp.Option[R] {
	return ApFunc(r.fn, func() fp.Option[A] {
		return Some(a())
	})
}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[A, R] {
	return ApplicativeFunctor1[A, R]{fn: Some(fn)}
}

func Pure0[R any](f func() R) fp.Func1[fp.Unit, fp.Option[R]] {
	return func(fp.Unit) fp.Option[R] {
		return Some(f())
	}
}

func Pure1[A, R any](f func(A) R) fp.Func1[A, fp.Option[R]] {
	return func(a A) fp.Option[R] {
		return Some(f(a))
	}
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "applicative_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/curried", Name: "curried"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
{{define "Receiver"}}func (r MonadChain{{.N}}[H, HT, {{TypeArgs 1 .N}}, R]){{end}}
{{define "Next"}}MonadChain{{dec .N}}[hlist.Cons[A1, H], {{TypeArgs 1 .N}}, R]{{end}}


type MonadChain{{.N}}[H hlist.Header[HT], HT, {{TypeArgs 1 .N}}, R any] struct {
	h  fp.Option[H]
	fn fp.Option[{{CurriedFunc 1 .N "R"}}]
}

{{template "Receiver" .}} FlatMap(a func(HT) fp.Option[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v.Head())
	})
	return r.ApOption(av)
}
{{template "Receiver" .}} Map(a func(HT) A1) {{template "Next" .}} {

	return r.FlatMap(func(h HT) fp.Option[A1] {
		return Some(a(h))
	})
}
{{template "Receiver" .}} HListMap(a func(H) A1) {{template "Next" .}} {

	return r.HListFlatMap(func(h H) fp.Option[A1] {
		return Some(a(h))
	})
}
{{template "Receiver" .}} HListFlatMap(a func(H) fp.Option[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a(v)
	})

	return r.ApOption(av)
}

{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return {{template "Next" .}}{nh, Ap(r.fn, a)}
}

{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {

	return r.ApOption(Some(a))

}

{{template "Receiver" .}} ApOptionFunc(a func() fp.Option[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Option[A1] {
		return a()
	})
	return r.ApOption(av)
}

{{template "Receiver" .}} ApFunc(a func() A1) {{template "Next" .}} {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApOption(av)
}

func Chain{{.N}}[{{TypeArgs 1 .N}}, R any](fn fp.Func{{.N}}[{{TypeArgs 1 .N}}, R]) MonadChain{{.N}}[hlist.Nil, hlist.Nil, {{TypeArgs 1 .N}}, R] {
	return MonadChain{{.N}}[hlist.Nil, hlist.Nil, {{TypeArgs 1 .N}}, R]{Some(hlist.Empty()), Some(curried.Func{{.N}}(fn))}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "applicative_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/curried", Name: "curried"},
		{Package: "github.com/csgura/fp/hlist", Name: "hlist"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
{{define "Receiver"}}func (r ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]){{end}}
{{define "Next"}}ApplicativeFunctor{{dec .N}}[{{TypeArgs 2 .N}}, R]{{end}}

type ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R any] struct {
	fn fp.Option[{{CurriedFunc 1 .N "R"}}]
}


{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {

	return {{template "Next" .}}{Ap(r.fn, a)}
}

{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {

	return r.ApOption(Some(a))

}

{{template "Receiver" .}} ApOptionFunc(a func() fp.Option[A1]) {{template "Next" .}} {

	return {{template "Next" .}}{ApFunc(r.fn, a)}

}

{{template "Receiver" .}} ApFunc(a func() A1) {{template "Next" .}} {

	return r.ApOptionFunc(func() fp.Option[A1] {
		return Some(a())
	})
}
func Applicative{{.N}}[{{TypeArgs 1 .N}}, R any](fn fp.Func{{.N}}[{{TypeArgs 1 .N}}, R]) ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R] {
	return ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]{Some(curried.Func{{.N}}(fn))}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  3,
	Until: genfp.MaxFunc,
	Template: `
func LiftA{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{DeclArgs 1 .N}}) R) func({{TypeClassArgs 1 .N "fp.Option"}}) fp.Option[R] {
	return func({{DeclTypeClassArgs 1 .N "fp.Option"}}) fp.Option[R] {

		return FlatMap(ins1, func(a1 A1) fp.Option[R] {
			return LiftA{{dec .N}}(func({{DeclArgs 2 .N}}) R {
				return f({{CallArgs 1 .N}})
			})({{CallArgs 2 .N "ins"}})
		})
	}
}

func LiftM{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{DeclArgs 1 .N}}) fp.Option[R]) func({{TypeClassArgs 1 .N "fp.Option"}}) fp.Option[R] {
	return func({{DeclTypeClassArgs 1 .N "fp.Option"}}) fp.Option[R] {

		return FlatMap(ins1, func(a1 A1) fp.Option[R] {
			return LiftM{{dec .N}}(func({{DeclArgs 2 .N}}) fp.Option[R] {
				return f({{CallArgs 1 .N}})
			})({{CallArgs 2 .N "ins"}})
		})
	}
}

func Flap{{.N}}[{{TypeArgs 1 .N}}, R any](tf fp.Option[{{CurriedFunc 1 .N "R"}}]) {{CurriedFunc 1 .N "fp.Option[R]"}} {
	return func(a1 A1) {{CurriedFunc 2 .N "fp.Option[R]"}} {
		return Flap{{dec .N}}(Ap(tf, Some(a1)))
	}
}

func Method{{.N}}[{{TypeArgs 1 .N}}, R any](ta1 fp.Option[A1], fa1 func({{DeclArgs 1 .N}}) R) func({{TypeArgs 2 .N}}) fp.Option[R] {
	return func({{DeclArgs 2 .N}}) fp.Option[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1({{CallArgs 1 .N}})
		})
	}
}

func FlatMethod{{.N}}[{{TypeArgs 1 .N}}, R any](ta1 fp.Option[A1], fa1 func({{DeclArgs 1 .N}}) fp.Option[R]) func({{TypeArgs 2 .N}}) fp.Option[R] {
	return func({{DeclArgs 2 .N}}) fp.Option[R] {
		return FlatMap(ta1, func(a1 A1) fp.Option[R] {
			return fa1({{CallArgs 1 .N}})
		})
	}
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "func_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
	},
	From:  3,
	Until: genfp.MaxCompose,
	Template: `
func Compose{{.N}}[{{TypeArgs 1 .N}}, R any]({{(Monad "fp.Option").FuncChain 1 .N}}) fp.Func1[A1, fp.Option[R]] {
	return Compose2(f1, Compose{{dec .N}}({{CallArgs 2 .N "f"}}))
}
	`,
}
