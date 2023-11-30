package list

import (
	"sort"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/immutable"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/monoid"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
)

type Nil[T any] struct {
}

func (r Nil[T]) IsEmpty() bool {
	return true
}

func (r Nil[T]) NonEmpty() bool {
	return false
}

func (r Nil[T]) Head() T {
	panic("List.empty")
}

func (r Nil[T]) Tail() fp.List[T] {
	return r
}

func (r Nil[T]) Unapply() (T, fp.List[T]) {
	return r.Head(), r
}

func (r Nil[T]) Foreach(f func(v T)) {
}

func (r Nil[T]) ToSeq() []T {
	return nil
}

type Cons[T any] struct {
	head T
	tail fp.List[T]
}

func (r Cons[T]) IsEmpty() bool {
	return false
}

func (r Cons[T]) NonEmpty() bool {
	return true
}

func (r Cons[T]) Head() T {
	return r.head
}

func (r Cons[T]) Tail() fp.List[T] {
	return r.tail
}

func (r Cons[T]) Unapply() (T, fp.List[T]) {
	return r.Head(), r.Tail()
}

func (r Cons[T]) Foreach(f func(v T)) {
	f(r.head)
	r.tail.Foreach(f)
}

func (r Cons[T]) ToSeq() []T {
	ret := []T{}
	r.Foreach(func(v T) {
		ret = append(ret, v)
	})
	return ret
}

func Empty[T any]() fp.List[T] {
	return Nil[T]{}
}

func Generate[T any](generator func(index int) fp.Option[T]) fp.List[T] {
	return GenerateFrom(0, generator)
}

func GenerateFrom[T any](startIndex int, generator func(index int) fp.Option[T]) fp.List[T] {
	return fp.MakeList(
		func() fp.Option[T] {
			return generator(startIndex)
		},
		func() fp.List[T] {
			return GenerateFrom(startIndex+1, generator)
		},
	)
}

func Recurrence1[T any](a1 T, relation func(n1 T) T) fp.List[T] {
	return fp.MakeList(
		func() fp.Option[T] {
			return option.Some(a1)
		},
		func() fp.List[T] {
			return Recurrence1(relation(a1), relation)
		},
	)
}

func Recurrence2[T any](a1 T, a2 T, relation func(n1, n2 T) T) fp.List[T] {
	return fp.MakeList(
		func() fp.Option[T] {
			return option.Some(a1)
		},
		func() fp.List[T] {
			return Recurrence2(a2, relation(a1, a2), relation)
		},
	)
}

func Head[T any](l fp.List[T]) fp.Option[T] {
	if l.IsEmpty() {
		return fp.Option[T]{}
	}
	return fp.Some(l.Head())
}

type Seq[T any] []T

var _ fp.List[int] = Seq[int]{}

func (r Seq[T]) IsEmpty() bool {
	return len(r) == 0
}

func (r Seq[T]) NonEmpty() bool {
	return len(r) > 0
}
func (r Seq[T]) Head() T {
	if r.IsEmpty() {
		panic("List.Empty")
	}
	return r[0]
}

func (r Seq[T]) Tail() fp.List[T] {
	if r.IsEmpty() {
		return Empty[T]()
	}
	return Seq[T](r[1:])
}

func (r Seq[T]) Unapply() (T, fp.List[T]) {
	return r.Head(), r.Tail()
}

func (r Seq[T]) Foreach(f func(v T)) {
	for _, v := range r {
		f(v)
	}
}

func (r Seq[T]) ToSeq() []T {
	if len(r) > 0 {
		ret := make([]T, len(r))
		copy(ret, r)
		return ret
	}
	return nil
}

func Map[T, U any](opt fp.List[T], fn func(v T) U) fp.List[U] {
	return fp.MakeList(
		func() fp.Option[U] {
			return option.Map(Head(opt), fn)
		},
		func() fp.List[U] {
			return Map(opt.Tail(), fn)
		},
	)
}

func Map2[A, B, U any](a fp.List[A], b fp.List[B], f func(A, B) U) fp.List[U] {
	return FlatMap(a, func(v1 A) fp.List[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		})
	})
}

func FilterMap[T, U any](opt fp.List[T], fn func(v T) fp.Option[U]) fp.List[U] {
	return FlatMap(opt, fp.Compose(fn, FromOption))
}

func FlatMap[T, U any](opt fp.List[T], fn func(v T) fp.List[U]) fp.List[U] {

	if opt.IsEmpty() {
		return Empty[U]()
	}

	mappedHeadLazy := lazy.Call(func() fp.List[U] {
		return fn(opt.Head())
	})

	tail := opt.Tail()

	return fp.MakeList(
		func() fp.Option[U] {
			headList := mappedHeadLazy.Get()

			if headList.IsEmpty() {
				return Head(FlatMap(tail, fn))
			}

			return fp.Some(headList.Head())
		},
		func() fp.List[U] {
			headList := mappedHeadLazy.Get()

			if headList.IsEmpty() {
				return FlatMap(tail, fn).Tail()
			}

			return Combine(headList.Tail(), FlatMap(tail, fn))
		},
	)

}

func Apply[T any](head T, tail fp.List[T]) fp.List[T] {
	return Cons[T]{head, tail}
}

func Of[T any](e ...T) fp.List[T] {
	return Seq[T](e)
}

func FromSeq[T any](seq fp.Seq[T]) fp.List[T] {
	return Seq[T](seq)
}

func FromSlice[T any](seq []T) fp.List[T] {
	return Seq[T](seq)
}
func ReverseSeq[T any](seq fp.Seq[T]) fp.List[T] {
	return fp.MakeList(
		func() fp.Option[T] {
			return fp.Seq[T](seq).Last()
		},
		func() fp.List[T] {
			return ReverseSeq(fp.Seq[T](seq).Init())
		},
	)
}
func ReverseSlice[T any](seq []T) fp.List[T] {
	return ReverseSeq(seq)
}

func FromPtr[T any](ptr *T) fp.List[T] {
	if ptr == nil {
		return Empty[T]()
	}
	return Of(*ptr)
}

func FromOption[T any](opt fp.Option[T]) fp.List[T] {
	if opt.IsDefined() {
		return Of(opt.Get())
	}
	return Empty[T]()
}

func FromMap[K comparable, V any](m map[K]V) fp.List[fp.Tuple2[K, V]] {
	return Collect(fp.IteratorOfGoMap(m))
}

func FromMapKey[K comparable, V any](m map[K]V) fp.List[K] {
	return Collect(mutable.MapOf(m).Keys())
}

func FromMapValue[K comparable, V any](m map[K]V) fp.List[V] {
	return Collect(mutable.MapOf(m).Values())
}

func Collect[T any](itr fp.Iterator[T]) fp.List[T] {
	head := itr.NextOption()

	return fp.MakeList(func() fp.Option[T] {
		return head
	}, func() fp.List[T] {

		return Collect(itr)
	})
}

func Concat[T any](head T, tail fp.List[T]) fp.List[T] {
	return Apply(head, tail)
}

func Combine[T any](l1 fp.List[T], l2 fp.List[T]) fp.List[T] {

	if l1.IsEmpty() {
		return l2
	}

	return fp.MakeList(
		func() fp.Option[T] {
			return fp.Some(l1.Head())
		},
		func() fp.List[T] {
			l1Tail := l1.Tail()
			if l1Tail.NonEmpty() {
				return Combine(l1Tail, l2)
			}
			return l2
		},
	)
}

func Ap[T, U any](t fp.List[fp.Func1[T, U]], a fp.List[T]) fp.List[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.List[U] {
		return Map(a, f)
	})
}

func Lift[T, U any](f func(v T) U) func(fp.List[T]) fp.List[U] {
	return func(opt fp.List[T]) fp.List[U] {
		return Map(opt, f)
	}
}

func Compose[A, B, C any](f1 func(A) fp.List[B], f2 func(B) fp.List[C]) func(A) fp.List[C] {
	return func(a A) fp.List[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposePure[A, B any](fab func(A) B) func(A) fp.List[B] {
	return func(a A) fp.List[B] {
		return Of(fab(a))
	}
}

func Flatten[T any](opt fp.List[fp.List[T]]) fp.List[T] {
	return FlatMap(opt, func(v fp.List[T]) fp.List[T] {
		return v
	})
}

func Flap[A, R any](tfa fp.List[fp.Func1[A, R]]) func(A) fp.List[R] {
	return func(a A) fp.List[R] {
		return Ap(tfa, Of(a))
	}
}

func Flap2[A, B, R any](tfab fp.List[fp.Func1[A, fp.Func1[B, R]]]) fp.Func1[A, fp.Func1[B, fp.List[R]]] {
	return func(a A) fp.Func1[B, fp.List[R]] {
		return Flap(Ap(tfab, Of(a)))
	}
}

func FlapMap[A, B, R any](tfab func(A, B) R, a fp.List[A]) func(B) fp.List[R] {
	return Flap(Map(a, as.Curried2(tfab)))
}

func Method1[A, B, R any](ta fp.List[A], fab func(a A, b B) R) func(B) fp.List[R] {
	return FlapMap(fab, ta)
}

func Method2[A, B, C, R any](ta fp.List[A], fabc func(a A, b B, c C) R) func(B, C) fp.List[R] {
	return curried.Revert2(Flap2(Map(ta, as.Curried3(fabc))))
}

func ToMap[K, V any](list fp.List[fp.Tuple2[K, V]], hasher fp.Hashable[K]) fp.Map[K, V] {
	ret := immutable.MapBuilder[K, V](hasher)

	cursor := list
	for !cursor.IsEmpty() {
		k, v := cursor.Head().Unapply()
		ret = ret.Add(k, v)
		cursor = cursor.Tail()
	}

	return ret.Build()
}

func ToGoMap[K comparable, V any](list fp.List[fp.Tuple2[K, V]]) map[K]V {
	ret := map[K]V{}
	cursor := list
	for !cursor.IsEmpty() {
		k, v := cursor.Head().Unapply()
		ret[k] = v
		cursor = cursor.Tail()
	}
	return ret
}

func ToSet[V any](list fp.List[V], hasher fp.Hashable[V]) fp.Set[V] {
	ret := immutable.SetBuilder(hasher)

	cursor := list
	for !cursor.IsEmpty() {
		v := cursor.Head()
		ret = ret.Add(v)
		cursor = cursor.Tail()
	}

	return ret.Build()
}

func ToGoSet[V comparable](list fp.List[V]) mutable.Set[V] {
	ret := map[V]bool{}
	cursor := list
	for !cursor.IsEmpty() {
		k := cursor.Head()
		ret[k] = true
		cursor = cursor.Tail()
	}
	return ret
}

func Zip[T, U any](a fp.List[T], b fp.List[U]) fp.List[fp.Tuple2[T, U]] {
	return fp.MakeList(
		func() fp.Option[fp.Tuple2[T, U]] {
			return option.LiftA2(as.Tuple[T, U])(Head(a), Head(b))
		},
		func() fp.List[fp.Tuple2[T, U]] {
			return Zip(a.Tail(), b.Tail())
		},
	)
}

func ZipWithIndex[A any](s1 fp.List[A]) fp.List[fp.Tuple2[int, A]] {
	idxList := Generate(func(index int) fp.Option[int] {
		return fp.Some(index)
	})
	return Zip(idxList, s1)
}

func Zip3[A, B, C any](a fp.List[A], b fp.List[B], c fp.List[C]) fp.List[fp.Tuple3[A, B, C]] {
	return fp.MakeList(
		func() fp.Option[fp.Tuple3[A, B, C]] {
			return option.LiftA3(as.Tuple3[A, B, C])(Head(a), Head(b), Head(c))

		},
		func() fp.List[fp.Tuple3[A, B, C]] {
			return Zip3(a.Tail(), b.Tail(), c.Tail())
		},
	)
}

func Reduce[A any](s fp.List[A], m fp.Monoid[A]) A {
	return FoldRight(s, m.Empty(), func(a A, b lazy.Eval[A]) lazy.Eval[A] {
		return b.Map(func(v A) A {
			return m.Combine(a, v)
		})
	}).Get()
}

func Fold[A, B any](s fp.List[A], zero B, f func(B, A) B) B {
	sum := zero

	cursor := s

	for !cursor.IsEmpty() {
		sum = f(sum, cursor.Head())
		cursor = cursor.Tail()
	}
	return sum
}

func FoldLeft[A, B any](s fp.List[A], zero B, f func(B, A) B) B {
	cf := curried.Flip(as.Curried2(f))
	ret := FoldRight[A, fp.Endo[B]](s, fp.Id, func(a A, endo lazy.Eval[fp.Endo[B]]) lazy.Eval[fp.Endo[B]] {
		ef := endo.Get().AsFunc()
		return lazy.Done(fp.Endo[B](fp.Compose(cf(a), ef)))
	})
	return ret.Get()(zero)
}

// FoldTry 는 foldM 의 Try 버젼
// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldTry[A, B any](s fp.List[A], zero B, f func(B, A) fp.Try[B]) fp.Try[B] {
	sum := zero
	cursor := s
	for cursor.NonEmpty() {
		t := f(sum, cursor.Head())
		if t.IsSuccess() {
			sum = t.Get()
		} else {
			return t
		}
		cursor = cursor.Tail()
	}
	return fp.Success(sum)
}

func FoldFuture[A, B any](s fp.List[A], zero B, fn func(B, A) fp.Future[B], ctx ...fp.Executor) fp.Future[B] {
	p := fp.NewPromise[B]()
	p.Success(zero)
	return Fold(s, p.Future(), func(acc fp.Future[B], v A) fp.Future[B] {
		return acc.FlatMap(func(acc B) fp.Future[B] {
			return fn(acc, v)
		}, ctx...)
	})
}

// FoldError 는  FoldTry[A,fp.Unit]와 같은 함수인데
// 하스켈에서 동일한 기능을 하는 함수를 찾아 보면 traverse_ 혹은 mapM_ 과 같은 함수
// 하스켈에서 _ 가 붙어 있는 함수들은 결과를 discard 해서  m() 를 리턴함.
func FoldError[A any](s fp.List[A], f func(A) error) error {
	cursor := s

	for cursor.NonEmpty() {
		err := f(cursor.Head())
		if err != nil {
			return err
		}
		cursor = cursor.Tail()
	}
	return nil
}

// FoldOption 는 foldM 의 Option 버젼
// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldOption[A, B any](s fp.List[A], zero B, f func(B, A) fp.Option[B]) fp.Option[B] {
	sum := zero
	cursor := s

	for cursor.NonEmpty() {
		t := f(sum, cursor.Head())
		if t.IsDefined() {
			sum = t.Get()
		} else {
			return t
		}
	}
	return fp.Some(sum)
}

func FoldMap[A, B any](s fp.List[A], m fp.Monoid[B], f func(A) B) B {
	ret := FoldRight(s, m.Empty(), func(a A, b lazy.Eval[B]) lazy.Eval[B] {
		ab := f(a)

		return b.Map(func(t B) B {
			return m.Combine(ab, t)
		})

	})

	return ret.Get()
}

func FoldRight[A, B any](s fp.List[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if s.IsEmpty() {
		return lazy.Done(zero)
	}

	v := lazy.TailCall(func() lazy.Eval[B] {
		return FoldRight(s.Tail(), zero, f)
	})
	return f(s.Head(), v)
}

func FoldLeftUsingMap[A, B any](s fp.List[A], zero B, f func(B, A) B) B {
	cf := curried.Flip(as.Curried2(f))
	m := monoid.Dual(monoid.Endo[B]())

	f2 := func(a A) fp.Dual[fp.Endo[B]] {
		return as.Dual(as.Endo(cf(a)))
	}

	ret := FoldMap(s, m, f2)
	return ret.GetDual(zero)
}

func FoldRightUsingMap[A, B any](s fp.List[A], zero B, f func(A, B) B) B {
	cf := as.Curried2(f)
	m := monoid.Endo[B]()

	f2 := func(a A) fp.Endo[B] {
		return as.Endo(cf(a))
	}

	ret := FoldMap(s, m, f2)
	return ret(zero)
}

func Scan[A, B any](s fp.List[A], zero B, f func(B, A) B) fp.List[B] {

	cf := as.Curried2(f)
	return fp.MakeList(
		func() fp.Option[B] {
			return option.Some(zero)
		},
		func() fp.List[B] {
			z := option.Map(Head(s), cf(zero))
			if z.IsDefined() {
				return Scan(s.Tail(), z.Get(), f)
			}
			return Empty[B]()
		},
	)
}

func GroupBy[A any, K comparable](s fp.List[A], keyFunc func(A) K) map[K]fp.Seq[A] {

	ret := map[K]fp.Seq[A]{}

	return Fold(s, ret, func(b map[K]fp.Seq[A], a A) map[K]fp.Seq[A] {
		k := keyFunc(a)
		b[k] = b[k].Append(a)
		return b
	})
}

func Range(from, exclusive int) fp.List[int] {
	return GenerateFrom(from, func(index int) fp.Option[int] {
		if index < exclusive {
			return option.Some(index)
		}
		return option.None[int]()
	})
}

func RangeClosed(from, inclusive int) fp.List[int] {
	return GenerateFrom(from, func(index int) fp.Option[int] {
		if index <= inclusive {
			return option.Some(index)
		}
		return option.None[int]()
	})
}

type seqSorter[T any] struct {
	seq fp.Seq[T]
	ord fp.Ord[T]
}

func (p *seqSorter[T]) Len() int           { return len(p.seq) }
func (p *seqSorter[T]) Less(i, j int) bool { return p.ord.Less(p.seq[i], p.seq[j]) }
func (p *seqSorter[T]) Swap(i, j int)      { p.seq[i], p.seq[j] = p.seq[j], p.seq[i] }

func Sort[T any](r fp.List[T], ord fp.Ord[T]) fp.Seq[T] {
	s := r.ToSeq()
	sort.Sort(&seqSorter[T]{s, ord})
	return s
}

func Min[T any](r fp.List[T], ord fp.Ord[T]) fp.Option[T] {
	return Fold(r, fp.Option[T]{}, func(min fp.Option[T], v T) fp.Option[T] {
		if min.IsDefined() && ord.Less(min.Get(), v) {
			return min
		}
		return fp.Some[T](v)
	})
}

func Max[T any](r fp.List[T], ord fp.Ord[T]) fp.Option[T] {
	return Fold(r, fp.Option[T]{}, func(max fp.Option[T], v T) fp.Option[T] {
		if max.IsDefined() && ord.Less(v, max.Get()) {
			return max
		}
		return fp.Some[T](v)
	})
}
