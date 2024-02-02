//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package try

import (
	"fmt"
	"runtime/debug"

	"github.com/csgura/fp"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/seq"
)

func Pure[T any](t T) fp.Try[T] {
	return fp.Success(t)
}

func Success[T any](t T) fp.Try[T] {
	return fp.Success(t)
}

func Failure[T any](err error) fp.Try[T] {
	return fp.Failure[T](err)
}

func FromOption[T any](v fp.Option[T]) fp.Try[T] {
	if v.IsDefined() {
		return Success(v.Get())
	} else {
		return Failure[T](fp.ErrOptionEmpty)
	}
}

func FromPtr[T any](v *T) fp.Try[T] {
	if v != nil {
		return Success(*v)
	} else {
		return Failure[T](fp.ErrOptionEmpty)
	}
}

type Panic interface {
	error
	Panic() any
	Stack() []byte
}

type panicError struct {
	cause any
	stack []byte
}

func (r *panicError) Error() string {
	return fmt.Sprintf("%v %v", r.cause, string(r.stack))
}

func (r *panicError) Stack() []byte {
	return r.stack
}

func (r *panicError) Panic() any {
	return r.cause
}

func Of[T any](f func() T) (ret fp.Try[T]) {
	defer func() {
		if p := recover(); p != nil {
			ret = Failure[T](&panicError{p, debug.Stack()})
		}
	}()

	ret = Success(f())
	return
}

func Call[T any](f func() (T, error)) (ret fp.Try[T]) {
	defer func() {
		if p := recover(); p != nil {
			ret = Failure[T](&panicError{p, debug.Stack()})
		}
	}()

	ret = Apply(f())
	return
}

func CallUnit(f func() error) (ret fp.Try[fp.Unit]) {
	defer func() {
		if p := recover(); p != nil {
			ret = Failure[fp.Unit](&panicError{p, debug.Stack()})
		}
	}()

	err := f()
	ret = Apply(fp.Unit{}, err)
	return
}

func Apply[T any](v T, err error) fp.Try[T] {
	if err != nil {
		return Failure[T](err)
	}
	return Success(v)
}

func ComposeOption[A, B, C any](f1 func(A) fp.Option[B], f2 func(B) fp.Try[C]) func(A) fp.Try[C] {
	return func(a A) fp.Try[C] {
		return FlatMap(FromOption(f1(a)), f2)
	}
}

func ComposePure[A, B any](fab func(A) B) func(A) fp.Try[B] {
	return fp.Compose(fab, Success)
}

var Unit fp.Try[fp.Unit] = Success(fp.Unit{})

// func Map[T, U any](opt fp.Try[T], f func(v T) U) fp.Try[U] {
// 	return Ap(Success(as.Func1(f)), opt)
// }

func FlatMap[A, B any](ta fp.Try[A], fn func(v A) fp.Try[B]) fp.Try[B] {
	if ta.IsSuccess() {
		return fn(ta.Get())
	}
	return Failure[B](ta.Failed().Get())
}

func FlatMapTraverseSeq[A, B any](ta fp.Try[fp.Seq[A]], f func(v A) fp.Try[B]) fp.Try[fp.Seq[B]] {
	return FlatMap(ta, TraverseSeqFunc(f))
}

func FlatMapTraverseSlice[A, B any](ta fp.Try[[]A], f func(v A) fp.Try[B]) fp.Try[[]B] {
	return FlatMap(ta, TraverseSliceFunc(f))
}

func SequenceIterator[A any](ita fp.Iterator[fp.Try[A]]) fp.Try[fp.Iterator[A]] {
	ret := iterator.Fold(ita, Success(seq.Empty[A]()), LiftA2(fp.Seq[A].Add))
	return Map(ret, iterator.FromSeq)

}

func traverse[A, R any](ia fp.Iterator[A], fn func(A) fp.Try[R]) fp.Try[fp.Seq[R]] {
	return iterator.FoldTry(ia, seq.Empty[R](), func(acc fp.Seq[R], a A) fp.Try[fp.Seq[R]] {
		return Map(fn(a), acc.Add)
	})

}

func Traverse[A, R any](ia fp.Iterator[A], fn func(A) fp.Try[R]) fp.Try[fp.Iterator[R]] {
	ret := traverse(ia, fn)
	return Map(ret, iterator.FromSeq)
}

// Traverse_ 가  Traverse 와 다른점은,  result 를 무시 한다는 것
// 다음은 하스켈의 traverse_ 함수
// traverse_ :: (Foldable t, Applicative f) => (a -> f b) -> t a -> f ()
func Traverse_[A, R any](ia fp.Iterator[A], fn func(A) fp.Try[R]) error {
	return iterator.FoldError(ia, func(a A) error {
		_, err := fn(a).Unapply()
		return err
	})
}

func TraverseOption[A, R any](opta fp.Option[A], fa func(A) fp.Try[R]) fp.Try[fp.Option[R]] {
	return Map(Traverse(fp.IteratorOfOption(opta), fa), fp.Iterator[R].NextOption)
}

func TraverseSeq[A, R any](sa fp.Seq[A], fa func(A) fp.Try[R]) fp.Try[fp.Seq[R]] {
	return traverse(fp.IteratorOfSeq(sa), fa)
}

func TraverseSlice[A, R any](sa []A, fa func(A) fp.Try[R]) fp.Try[[]R] {
	return Map(traverse(fp.IteratorOfSeq(sa), fa), fp.Seq[R].Widen)
}

func TraverseFunc[A, R any](far func(A) fp.Try[R]) func(fp.Iterator[A]) fp.Try[fp.Iterator[R]] {
	return func(iterA fp.Iterator[A]) fp.Try[fp.Iterator[R]] {
		return Traverse(iterA, far)
	}
}

func TraverseSeqFunc[A, R any](far func(A) fp.Try[R]) func(fp.Seq[A]) fp.Try[fp.Seq[R]] {
	return func(seqA fp.Seq[A]) fp.Try[fp.Seq[R]] {
		return TraverseSeq(seqA, far)
	}
}

func TraverseSliceFunc[A, R any](far func(A) fp.Try[R]) func([]A) fp.Try[[]R] {
	return func(seqA []A) fp.Try[[]R] {
		return TraverseSlice(seqA, far)
	}
}
func Sequence[A any](tsa []fp.Try[A]) fp.Try[[]A] {
	ret := iterator.Fold(iterator.FromSeq(tsa), Success(seq.Empty[A]()), LiftA2(fp.Seq[A].Add))
	return Map(ret, fp.Seq[A].Widen)
}

func Fold[A, B any](ta fp.Try[A], bzero B, fba func(B, A) B) B {
	if ta.IsFailure() {
		return bzero
	}

	return fba(bzero, ta.Get())
}

func FoldRight[A, B any](ta fp.Try[A], bzero B, fab func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if ta.IsFailure() {
		return lazy.Done(bzero)
	}

	return fab(ta.Get(), lazy.Done(bzero))
}

func ToSeq[A any](ta fp.Try[A]) fp.Seq[A] {
	if ta.IsSuccess() {
		return fp.Seq[A]{ta.Get()}
	}
	return nil
}

func Iterator[A any](ta fp.Try[A]) fp.Iterator[A] {
	return fp.IteratorOfSeq(ToSeq(ta))
}

type MonadChain1[H hlist.Header[HT], HT, A, R any] struct {
	h  fp.Try[H]
	fn fp.Try[fp.Func1[A, R]]
}

func (r MonadChain1[H, HT, A, R]) Map(a func(HT) A) fp.Try[R] {
	return r.FlatMap(func(h HT) fp.Try[A] {
		return Success(a(h))
	})
}

func (r MonadChain1[H, HT, A, R]) HListMap(a func(H) A) fp.Try[R] {
	return r.HListFlatMap(func(h H) fp.Try[A] {
		return Success(a(h))
	})
}

func (r MonadChain1[H, HT, A, R]) HListFlatMap(a func(H) fp.Try[A]) fp.Try[R] {
	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return a(v)
	})

	return r.ApTry(av)
}

func (r MonadChain1[H, HT, A, R]) FlatMap(a func(HT) fp.Try[A]) fp.Try[R] {
	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return a(v.Head())
	})

	return r.ApTry(av)
}

func (r MonadChain1[H, HT, A, R]) ApOption(a fp.Option[A]) fp.Try[R] {
	return r.ApTry(FromOption(a))
}

func (r MonadChain1[H, HT, A, R]) ApTry(a fp.Try[A]) fp.Try[R] {
	return Ap(r.fn, a)
}

func (r MonadChain1[H, HT, A, R]) Ap(a A) fp.Try[R] {
	return r.ApTry(Success(a))
}

func (r MonadChain1[H, HT, A, R]) ApTryFunc(a func() fp.Try[A]) fp.Try[R] {

	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return a()
	})
	return r.ApTry(av)
}
func (r MonadChain1[H, HT, A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Try[R] {

	av := FlatMap(r.h, func(v H) fp.Try[A] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
func (r MonadChain1[H, HT, A, R]) ApFunc(a func() A) fp.Try[R] {

	av := Map(r.h, func(v H) A {
		return a()
	})
	return r.ApTry(av)
}

func Chain1[A, R any](fn fp.Func1[A, R]) MonadChain1[hlist.Nil, hlist.Nil, A, R] {
	return MonadChain1[hlist.Nil, hlist.Nil, A, R]{Success(hlist.Empty()), Success(fn)}
}

type ApplicativeFunctor1[A, R any] struct {
	fn fp.Try[fp.Func1[A, R]]
}

func (r ApplicativeFunctor1[A, R]) ApOption(a fp.Option[A]) fp.Try[R] {
	return r.ApTry(FromOption(a))
}

func (r ApplicativeFunctor1[A, R]) ApTry(a fp.Try[A]) fp.Try[R] {
	return Ap(r.fn, a)
}

func (r ApplicativeFunctor1[A, R]) Ap(a A) fp.Try[R] {
	return r.ApTry(Success(a))
}

func (r ApplicativeFunctor1[A, R]) ApTryFunc(a func() fp.Try[A]) fp.Try[R] {

	return ApFunc(r.fn, a)

}
func (r ApplicativeFunctor1[A, R]) ApOptionFunc(a func() fp.Option[A]) fp.Try[R] {
	return r.ApTryFunc(func() fp.Try[A] {
		return FromOption(a())
	})
}
func (r ApplicativeFunctor1[A, R]) ApFunc(a func() A) fp.Try[R] {

	return r.ApTryFunc(func() fp.Try[A] {
		return Success(a())
	})

}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[A, R] {
	return ApplicativeFunctor1[A, R]{Success(fn)}
}

func Pure0[R any](f func() R) fp.Func1[fp.Unit, fp.Try[R]] {
	return func(fp.Unit) fp.Try[R] {
		return Success(f())
	}
}

func Func0[R any](f func() (R, error)) fp.Func1[fp.Unit, fp.Try[R]] {
	return func(fp.Unit) fp.Try[R] {
		ret, err := f()
		return Apply(ret, err)
	}
}

func Unit0(f func() error) fp.Func1[fp.Unit, fp.Try[fp.Unit]] {
	return func(fp.Unit) fp.Try[fp.Unit] {
		err := f()
		return Apply(fp.Unit{}, err)
	}
}

//go:generate go run github.com/csgura/fp/internal/generator/monad_gen

// @internal.Generate
func _[A any]() genfp.GenerateMonadFunctions[fp.Try[A]] {
	return genfp.GenerateMonadFunctions[fp.Try[A]]{
		File:     "try_monad.go",
		TypeParm: genfp.TypeOf[A](),
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
	h  fp.Try[H]
	fn fp.Try[{{CurriedFunc 1 .N "R"}}]
}

{{template "Receiver" .}} FlatMap(a func(HT) fp.Try[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v.Head())
	})
	return r.ApTry(av)
}
{{template "Receiver" .}} Map(a func(HT) A1) {{template "Next" .}} {

	return r.FlatMap(func(h HT) fp.Try[A1] {
		return Success(a(h))
	})
}
{{template "Receiver" .}} HListMap(a func(H) A1) {{template "Next" .}} {

	return r.HListFlatMap(func(h H) fp.Try[A1] {
		return Success(a(h))
	})
}
{{template "Receiver" .}} HListFlatMap(a func(H) fp.Try[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a(v)
	})

	return r.ApTry(av)
}
{{template "Receiver" .}} ApTry(a fp.Try[A1]) {{template "Next" .}} {

	nh := Map2(a, r.h, hlist.Concat[A1, H])

	return {{template "Next" .}}{nh, Ap(r.fn, a)}
}
{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {

	return r.ApTry(FromOption(a))
}
{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {

	return r.ApTry(Success(a))

}
{{template "Receiver" .}} ApTryFunc(a func() fp.Try[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return a()
	})
	return r.ApTry(av)
}
{{template "Receiver" .}} ApOptionFunc(a func() fp.Option[A1]) {{template "Next" .}} {

	av := FlatMap(r.h, func(v H) fp.Try[A1] {
		return FromOption(a())
	})
	return r.ApTry(av)
}
{{template "Receiver" .}} ApFunc(a func() A1) {{template "Next" .}} {

	av := Map(r.h, func(v H) A1 {
		return a()
	})
	return r.ApTry(av)
}

func Chain{{.N}}[{{TypeArgs 1 .N}}, R any](fn fp.Func{{.N}}[{{TypeArgs 1 .N}}, R]) MonadChain{{.N}}[hlist.Nil, hlist.Nil, {{TypeArgs 1 .N}}, R] {
	return MonadChain{{.N}}[hlist.Nil, hlist.Nil, {{TypeArgs 1 .N}}, R]{Success(hlist.Empty()), Success(curried.Func{{.N}}(fn))}
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
	fn fp.Try[{{CurriedFunc 1 .N "R"}}]
}

{{template "Receiver" .}} ApTry(a fp.Try[A1]) {{template "Next" .}} {

	return {{template "Next" .}}{Ap(r.fn, a)}
}
{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {

	return r.ApTry(FromOption(a))
}
{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {

	return r.ApTry(Success(a))

}
{{template "Receiver" .}} ApTryFunc(a func() fp.Try[A1]) {{template "Next" .}} {

	return {{template "Next" .}}{ApFunc(r.fn, a)}

}
{{template "Receiver" .}} ApOptionFunc(a func() fp.Option[A1]) {{template "Next" .}} {

	return r.ApTryFunc(func() fp.Try[A1] {
		return FromOption(a())
	})
}
{{template "Receiver" .}} ApFunc(a func() A1) {{template "Next" .}} {

	return r.ApTryFunc(func() fp.Try[A1] {
		return Success(a())
	})
}
func Applicative{{.N}}[{{TypeArgs 1 .N}}, R any](fn fp.Func{{.N}}[{{TypeArgs 1 .N}}, R]) ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R] {
	return ApplicativeFunctor{{.N}}[{{TypeArgs 1 .N}}, R]{Success(curried.Func{{.N}}(fn))}
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

func LiftM{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{DeclArgs 1 .N}}) fp.Try[R]) func({{TypeClassArgs 1 .N "fp.Try"}}) fp.Try[R] {
	return func({{DeclTypeClassArgs 1 .N "fp.Try"}}) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftM{{dec .N}}(func({{DeclArgs 2 .N}}) fp.Try[R] {
				return f({{CallArgs 1 .N}})
			})({{CallArgs 2 .N "ins"}})
		})
	}
}

func Flap{{.N}}[{{TypeArgs 1 .N}}, R any](tf fp.Try[{{CurriedFunc 1 .N "R"}}]) {{CurriedFunc 1 .N "fp.Try[R]"}} {
	return func(a1 A1) {{CurriedFunc 2 .N "fp.Try[R]"}} {
		return Flap{{dec .N}}(Ap(tf, Success(a1)))
	}
}

func Method{{.N}}[{{TypeArgs 1 .N}}, R any](ta1 fp.Try[A1], fa1 func({{DeclArgs 1 .N}}) R) func({{TypeArgs 2 .N}}) fp.Try[R] {
	return func({{DeclArgs 2 .N}}) fp.Try[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1({{CallArgs 1 .N}})
		})
	}
}

func FlatMethod{{.N}}[{{TypeArgs 1 .N}}, R any](ta1 fp.Try[A1], fa1 func({{DeclArgs 1 .N}}) fp.Try[R]) func({{TypeArgs 2 .N}}) fp.Try[R] {
	return func({{DeclArgs 2 .N}}) fp.Try[R] {
		return FlatMap(ta1, func(a1 A1) fp.Try[R] {
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
	From:  1,
	Until: genfp.MaxFunc,
	Template: `
func Func{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) (R, error)) fp.Func{{.N}}[{{TypeArgs 1 .N}}, fp.Try[R]] {
	return func({{DeclArgs 1 .N}}) fp.Try[R] {
		ret, err := f({{CallArgs 1 .N}})
		return Apply(ret, err)
	}
}

func Pure{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R) fp.Func{{.N}}[{{TypeArgs 1 .N}}, fp.Try[R]] {
	return func({{DeclArgs 1 .N}}) fp.Try[R] {
		return Success(f({{CallArgs 1 .N}}))
	}
}

func Unit{{.N}}[{{TypeArgs 1 .N}} any](f func({{TypeArgs 1 .N}}) error) fp.Func{{.N}}[{{TypeArgs 1 .N}}, fp.Try[fp.Unit]] {
	return func({{DeclArgs 1 .N}}) fp.Try[fp.Unit] {
		err := f({{CallArgs 1 .N}})
		return Apply(fp.Unit{}, err)
	}
}

func Ptr{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) (*R, error)) fp.Func{{.N}}[{{TypeArgs 1 .N}}, fp.Try[R]] {
	return func({{DeclArgs 1 .N}}) fp.Try[R] {
		ret, err := f({{CallArgs 1 .N}})
		return FlatMap(Apply(ret, err), FromPtr)
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
func Compose{{.N}}[{{TypeArgs 1 .N}}, R any]({{(Monad "fp.Try").FuncChain 1 .N}}) fp.Func1[A1, fp.Try[R]] {
	return Compose2(f1, Compose{{dec .N}}({{CallArgs 2 .N "f"}}))
}
	`,
}

// @internal.Generate
var _ = genfp.GenerateFromUntil{
	File: "curried_gen.go",
	Imports: []genfp.ImportPackage{
		{Package: "github.com/csgura/fp", Name: "fp"},
		{Package: "github.com/csgura/fp/as", Name: "as"},
	},
	From:  2,
	Until: genfp.MaxFunc,
	Template: `
func Curried{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) (R,error)) {{CurriedFunc 1 .N "fp.Try[R]"}} {
	return as.Curried{{.N}}(func({{DeclArgs 1 .N}}) fp.Try[R] {
		return Apply(f({{CallArgs 1 .N}}))
	})	
}

func CurriedPure{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) R) {{CurriedFunc 1 .N "fp.Try[R]"}} {
	return as.Curried{{.N}}(func({{DeclArgs 1 .N}}) fp.Try[R] {
		return Success(f({{CallArgs 1 .N}}))
	})	
}

func CurriedUnit{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) error) {{CurriedFunc 1 .N "fp.Try[fp.Unit]"}} {
	return as.Curried{{.N}}(func({{DeclArgs 1 .N}}) fp.Try[fp.Unit] {
		return Apply(fp.Unit{}, f({{CallArgs 1 .N}}))
	})	
}

func CurriedPtr{{.N}}[{{TypeArgs 1 .N}}, R any](f func({{TypeArgs 1 .N}}) (*R,error)) {{CurriedFunc 1 .N "fp.Try[R]"}} {
	return as.Curried{{.N}}(func({{DeclArgs 1 .N}}) fp.Try[R] {
		return FlatMap(Apply(f({{CallArgs 1 .N}})),FromPtr)
	})	
}
`,
}
