//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package try

import (
	"fmt"
	"runtime/debug"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
)

func Constructors[T fp.Try[V], V any]() (failure func(error) fp.Try[V], success func(V) fp.Try[V]) {
	return Failure[V], Success[V]
}

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

func FromOptionT[T any](v fp.OptionT[T]) fp.Try[T] {
	return FlatMap(v, FromOption)
}

func FromPtr[T any](v fp.Ptr[T]) fp.Try[T] {
	if v != nil {
		return Success(*v)
	} else {
		return Failure[T](fp.ErrOptionEmpty)
	}
}

func FromPtrT[T any](v fp.PtrT[T]) fp.Try[T] {
	return FlatMap(v, FromPtr)
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

func Map[T, U any](tt fp.Try[T], f func(v T) U) fp.Try[U] {
	if tt.IsSuccess() {
		return Success(f(tt.Get()))
	}
	return Failure[U](tt.Failed().Get())
}

func FlatMap[A, B any](ta fp.Try[A], fn func(v A) fp.Try[B]) fp.Try[B] {
	if ta.IsSuccess() {
		return fn(ta.Get())
	}
	return Failure[B](ta.Failed().Get())
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

func TraverseOption[A, R any](opta fp.Option[A], fa func(A) fp.Try[R]) fp.OptionT[R] {

	if opta.IsEmpty() {
		return Success(fp.None[R]())
	}

	return Map(fa(opta.Get()), func(v R) fp.Option[R] {
		return fp.Some(v)
	})
}

func TraversePtr[A, R any](p fp.Ptr[A], fa func(A) fp.Try[R]) fp.PtrT[R] {
	if p == nil {
		return Success[fp.Ptr[R]](nil)
	}

	return Map(fa(*p), func(v R) fp.Ptr[R] {
		return &v
	})
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
	return fp.IteratorOfSeq(ta.ToSeq())
}

// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldM[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.Try[B]) fp.Try[B] {
	sum := zero
	for s.HasNext() {
		t := f(sum, s.Next())
		if t.IsSuccess() {
			sum = t.Get()
		} else {
			return t
		}
	}
	return fp.Success(sum)
}

func FoldSliceM[A, B any](s fp.Slice[A], zero B, f func(B, A) fp.Try[B]) fp.Try[B] {
	sum := zero
	for _, v := range s {
		t := f(sum, v)
		if t.IsSuccess() {
			sum = t.Get()
		} else {
			return t
		}
	}
	return fp.Success(sum)
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
func _[A any]() genfp.GenerateTraverseFunctions[fp.Try[A]] {
	return genfp.GenerateTraverseFunctions[fp.Try[A]]{
		File:     "try_traverse.go",
		TypeParm: genfp.TypeOf[A](),
	}
}

// OptionT[F,A] 는 F[Option[A]]  와 같은 타입
// SeqT[Try,A] 는  Try[Seq[A]] 타입임
// PureSeqT
// MapSeqT
// FlatMapSeqT

// @internal.Generate
func _[A, B any]() genfp.GenerateMonadTransformer[fp.Try[fp.Seq[A]]] {
	return genfp.GenerateMonadTransformer[fp.Try[fp.Seq[A]]]{
		File:     "try_seqt.go",
		TypeParm: genfp.TypeOf[A](),
		ExposureMonad: genfp.MonadFunctions{
			Pure:    seq.Pure[A],
			FlatMap: seq.FlatMap[A, B],
		},
		Sequence: func(v fp.Seq[fp.Try[A]]) fp.Try[fp.Seq[A]] {
			return Map(Sequence(v), as.Seq)
		},
		Transform: []any{
			fp.Seq[A].Filter,
			fp.Seq[A].Add,
			fp.Seq[A].Append,
			fp.Seq[A].Concat,
			fp.Seq[A].Drop,
			fp.Seq[A].Exists,
			fp.Seq[A].FilterNot,
			fp.Seq[A].Find,
			fp.Seq[A].ForAll,
			fp.Seq[A].Foreach,
			fp.Seq[A].Get,
			fp.Seq[A].Head,
			fp.Seq[A].Tail,
			fp.Seq[A].Init,
			fp.Seq[A].IsEmpty,
			fp.Seq[A].Last,
			fp.Seq[A].MakeString,
			fp.Seq[A].NonEmpty,
			fp.Seq[A].Reverse,
			fp.Seq[A].Size,
			fp.Seq[A].Take,
			seq.Fold[A, B],
			seq.Scan[A, B],
			seq.Sort[A],
			seq.Min[A],
			seq.Max[A],

			// TODO: SPAN (multi value return)
			// ToGoMAP :  comparable constraints
			// FoldTry : -> flatten
		},
	}
}

// try.OptionT[A] =  fp.Try[fp.Option[A]]

// @internal.Generate
func _[T, U any]() genfp.GenerateMonadTransformer[fp.Try[fp.Option[T]]] {
	return genfp.GenerateMonadTransformer[fp.Try[fp.Option[T]]]{
		File:     "try_optiont.go",
		TypeParm: genfp.TypeOf[T](),
		ExposureMonad: genfp.MonadFunctions{
			Pure:    option.Pure[T],
			FlatMap: option.FlatMap[T, U],
		},
		Sequence: func(v fp.Option[fp.Try[T]]) fp.Try[fp.Option[T]] {
			if v.IsDefined() {
				return Map(v.Get(), option.Some)
			}
			return Success(fp.Option[T]{})
		},
		Transform: []any{
			fp.Option[T].Filter,
			fp.Option[T].OrElse,
			fp.Option[T].OrZero,
			fp.Option[T].OrElseGet,
			fp.Option[T].Or,
			fp.Option[T].OrOption,
			fp.Option[T].OrPtr,
			fp.Option[T].Recover,
			option.Fold[T, U],
		},
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

{{template "Receiver" .}} ApTryAll({{DeclTypeClassArgs 1 .N "fp.Try"}}) fp.Try[R] {
	return r.
	{{- range (dec .N) -}}
		ApTry(ins{{inc .}}).
	{{- end -}}
		ApTry(ins{{.N}})
}

{{template "Receiver" .}} ApOption(a fp.Option[A1]) {{template "Next" .}} {
	return r.ApTry(FromOption(a))
}

{{template "Receiver" .}} ApOptionAll({{DeclTypeClassArgs 1 .N "fp.Option"}}) fp.Try[R] {
	return r.
	{{- range (dec .N) -}}
		ApOption(ins{{inc .}}).
	{{- end -}}
		ApOption(ins{{.N}})
}

{{template "Receiver" .}} Ap(a A1) {{template "Next" .}} {
	return r.ApTry(Success(a))
}

{{template "Receiver" .}} ApAll({{DeclArgs 1 .N}}) fp.Try[R] {
	return r.
	{{- range (dec .N) -}}
		Ap(a{{inc .}}).
	{{- end -}}
		Ap(a{{.N}})
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
