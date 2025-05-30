//go:generate go run github.com/csgura/fp/internal/generator/template_gen
package iterator

import (
	"iter"
	"sort"
	"sync"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/genfp"
	"github.com/csgura/fp/immutable"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/xtr"
)

func Pull[T any](seq iter.Seq[T]) fp.Iterator[T] {
	return fp.MakePullIterator(seq)
}

func Empty[T any]() fp.Iterator[T] {
	return fp.MakeIterator(func() bool {
		return false
	}, func() T {
		panic("next on empty iterator")
	})
}

func List[T any](list fp.List[T]) fp.Iterator[T] {
	return FromList(list)
}

func FromOption[T any](opt fp.Option[T]) fp.Iterator[T] {
	return fp.IteratorOfOption(opt)
}

func FromList[T any](list fp.List[T]) fp.Iterator[T] {
	current := list

	return fp.MakeIterator(
		func() bool {
			return current.NonEmpty()
		},
		func() T {
			if current.NonEmpty() {
				ret := current.Head()
				current = current.Tail()
				return ret
			}
			panic("next on empty iterator")
		},
	)
}

func Of[T any](list ...T) fp.Iterator[T] {
	return fp.IteratorOfSeq(list)
}

func FromSeq[T any](seq fp.Seq[T]) fp.Iterator[T] {
	return fp.IteratorOfSeq(seq)
}

func FromSlice[T any](seq []T) fp.Iterator[T] {
	return fp.IteratorOfSeq(seq)
}

func ReverseSeq[T any](seq []T) fp.Iterator[T] {
	idx := len(seq)

	return fp.MakeIterator(
		func() bool {
			return idx > 0
		},
		func() T {
			if idx > 0 {
				ret := seq[idx-1]
				idx--
				return ret
			}
			panic("next on empty iterator")
		},
	)
}

func ReverseSlice[T any](seq []T) fp.Iterator[T] {
	return ReverseSeq(seq)
}

func FromPtr[T any](ptr *T) fp.Iterator[T] {
	if ptr == nil {
		return Empty[T]()
	} else {
		return Of(*ptr)
	}
}

func FromMap[K comparable, V any](m map[K]V) fp.Iterator[fp.Tuple2[K, V]] {
	return fp.IteratorOfGoMap(m)
}

func FromMapKey[K comparable, V any](m map[K]V) fp.Iterator[K] {
	return mutable.MapOf(m).Keys()
}

func FromMapValue[K comparable, V any](m map[K]V) fp.Iterator[V] {
	return mutable.MapOf(m).Values()
}

func Ap[T, U any](t fp.Iterator[fp.Func1[T, U]], a fp.Iterator[T]) fp.Iterator[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Iterator[U] {
		return Map(a, f)
	})
}

func Lift[T, U any](f func(v T) U) func(fp.Iterator[T]) fp.Iterator[U] {
	return func(opt fp.Iterator[T]) fp.Iterator[U] {
		return Map(opt, f)
	}
}

func Compose[A, B, C any](f1 func(A) fp.Iterator[B], f2 func(B) fp.Iterator[C]) func(A) fp.Iterator[C] {
	return func(a A) fp.Iterator[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposePure[A, B any](fab func(A) B) func(A) fp.Iterator[B] {
	return func(a A) fp.Iterator[B] {
		return Of(fab(a))
	}
}

func Flatten[T any](opt fp.Iterator[fp.Iterator[T]]) fp.Iterator[T] {
	return FlatMap(opt, func(v fp.Iterator[T]) fp.Iterator[T] {
		return v
	})
}

func Concat[T any](head T, tail fp.Iterator[T]) fp.Iterator[T] {
	return Of(head).Concat(tail)
}

func Map[T, U any](opt fp.Iterator[T], fn func(v T) U) fp.Iterator[U] {
	return fp.MakeIterator(
		func() bool {
			return opt.HasNext()
		},
		func() U {
			return fn(opt.Next())
		},
	)
}

func MapKey[KA, KB, V any](s fp.Iterator[fp.Tuple2[KA, V]], f func(KA) KB) fp.Iterator[fp.Tuple2[KB, V]] {
	return Map(s, func(v fp.Tuple2[KA, V]) fp.Tuple2[KB, V] {
		return fp.Tuple2[KB, V]{
			I1: f(v.I1),
			I2: v.I2,
		}
	})
}

func FilterMapKey[KA, KB, V any](s fp.Iterator[fp.Tuple2[KA, V]], f func(KA) fp.Option[KB]) fp.Iterator[fp.Tuple2[KB, V]] {
	return FilterMap(s, func(v fp.Tuple2[KA, V]) fp.Option[fp.Tuple2[KB, V]] {
		ov := f(v.I1)
		if ov.IsDefined() {
			return fp.Some(fp.Tuple2[KB, V]{
				I1: ov.Get(),
				I2: v.I2,
			})
		}
		return fp.None[fp.Tuple2[KB, V]]()
	})
}

func MapValue[K, VA, VB any](s fp.Iterator[fp.Tuple2[K, VA]], f func(VA) VB) fp.Iterator[fp.Tuple2[K, VB]] {
	return Map(s, func(v fp.Tuple2[K, VA]) fp.Tuple2[K, VB] {
		return fp.Tuple2[K, VB]{
			I1: v.I1,
			I2: f(v.I2),
		}
	})
}

func FilterMapValue[K, VA, VB any](s fp.Iterator[fp.Tuple2[K, VA]], f func(VA) fp.Option[VB]) fp.Iterator[fp.Tuple2[K, VB]] {
	return FilterMap(s, func(v fp.Tuple2[K, VA]) fp.Option[fp.Tuple2[K, VB]] {
		ov := f(v.I2)
		if ov.IsDefined() {
			return fp.Some(fp.Tuple2[K, VB]{
				I1: v.I1,
				I2: ov.Get(),
			})
		}
		return fp.None[fp.Tuple2[K, VB]]()
	})
}

func Map2[A, B, U any](a fp.Iterator[A], b fp.Iterator[B], f func(A, B) U) fp.Iterator[U] {
	return FlatMap(a, func(v1 A) fp.Iterator[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		})
	})
}

func FilterMap[T, U any](opt fp.Iterator[T], fn func(v T) fp.Option[U]) fp.Iterator[U] {
	return FlatMap(opt, fp.Compose(fn, fp.IteratorOfOption))
}

func FlatMap[T, U any](opt fp.Iterator[T], fn func(v T) fp.Iterator[U]) fp.Iterator[U] {

	current := fp.None[fp.Iterator[U]]()

	hasNext := func() bool {
		if current.IsDefined() && current.Get().HasNext() {
			return true
		}

		for opt.HasNext() {
			nextItr := fn(opt.Next())
			current = fp.Some(nextItr)
			if nextItr.HasNext() {
				return true
			}
		}

		return false
	}

	return fp.MakeIterator(
		hasNext,
		func() U {
			if hasNext() {
				return current.Get().Next()
			}
			panic("next on empty iterator")
		},
	)
}

func Flap[A, R any](tfa fp.Iterator[fp.Func1[A, R]]) func(A) fp.Iterator[R] {
	return func(a A) fp.Iterator[R] {
		return Ap(tfa, Of(a))
	}
}

func Flap2[A, B, R any](tfab fp.Iterator[fp.Func1[A, fp.Func1[B, R]]]) fp.Func1[A, fp.Func1[B, fp.Iterator[R]]] {
	return func(a A) fp.Func1[B, fp.Iterator[R]] {
		return Flap(Ap(tfab, Of(a)))
	}
}

func FlapMap[A, B, R any](tfab func(A, B) R, a fp.Iterator[A]) func(B) fp.Iterator[R] {
	return Flap(Map(a, as.Curried2(tfab)))
}

func Method1[A, B, R any](ta fp.Iterator[A], fab func(a A, b B) R) func(B) fp.Iterator[R] {
	return FlapMap(fab, ta)
}

func Method2[A, B, C, R any](ta fp.Iterator[A], fabc func(a A, b B, c C) R) func(B, C) fp.Iterator[R] {
	return curried.Revert2(Flap2(Map(ta, as.Curried3(fabc))))
}

func ToMap[K, V any](itr fp.Iterator[fp.Tuple2[K, V]], hasher fp.Hashable[K]) fp.Map[K, V] {
	ret := immutable.MapBuilder[K, V](hasher)

	for itr.HasNext() {
		k, v := itr.Next().Unapply()
		ret = ret.Add(k, v)
	}

	return ret.Build()
}

func ToGoMap[K comparable, V any](itr fp.Iterator[fp.Tuple2[K, V]]) map[K]V {
	ret := map[K]V{}
	for itr.HasNext() {
		k, v := itr.Next().Unapply()
		ret[k] = v
	}
	return ret
}

func ToSlice[V any](itr fp.Iterator[V]) []V {
	return itr.ToSeq()
}

func ToSeq[V any](itr fp.Iterator[V]) fp.Seq[V] {
	return itr.ToSeq()
}

func ToSet[V any](itr fp.Iterator[V], hasher fp.Hashable[V]) fp.Set[V] {
	ret := immutable.SetBuilder(hasher)

	for itr.HasNext() {
		v := itr.Next()
		ret = ret.Add(v)
	}

	return ret.Build()
}

func ToGoSet[V comparable](itr fp.Iterator[V]) mutable.Set[V] {
	ret := map[V]bool{}
	for itr.HasNext() {
		k := itr.Next()
		ret[k] = true
	}
	return ret
}

func ToList[V any](itr fp.Iterator[V]) fp.List[V] {
	head := itr.NextOption()

	return fp.MakeList(func() fp.Option[V] {
		return head
	}, func() fp.List[V] {

		return ToList(itr)
	})
}

func Zip[T, U any](a fp.Iterator[T], b fp.Iterator[U]) fp.Iterator[fp.Tuple2[T, U]] {
	return fp.MakeIterator(
		func() bool {
			return a.HasNext() && b.HasNext()
		},
		func() fp.Tuple2[T, U] {
			return as.Tuple(a.Next(), b.Next())
		},
	)
}

func ZipWithIndex[A any](s1 fp.Iterator[A]) fp.Iterator[fp.Tuple2[int, A]] {
	idx := 0
	idxList := Generate(func() int {
		ret := idx
		idx++
		return ret
	})
	return Zip(idxList, s1)
}

func Zip3[A, B, C any](a fp.Iterator[A], b fp.Iterator[B], c fp.Iterator[C]) fp.Iterator[fp.Tuple3[A, B, C]] {
	return fp.MakeIterator(
		func() bool {
			return a.HasNext() && b.HasNext() && c.HasNext()
		},
		func() fp.Tuple3[A, B, C] {
			return as.Tuple3(a.Next(), b.Next(), c.Next())
		},
	)
}

func Reduce[T any](r fp.Iterator[T], m fp.Monoid[T]) T {
	ret := m.Empty()
	for r.HasNext() {
		v := r.Next()
		ret = m.Combine(ret, v)
	}
	return ret
}

func Fold[A, B any](s fp.Iterator[A], zero B, f func(B, A) B) B {
	sum := zero
	for s.HasNext() {
		sum = f(sum, s.Next())
	}
	return sum
}

// FoldTry 는 foldM 의 Try 버젼
// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldTry[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.Try[B]) fp.Try[B] {
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

func FoldFuture[A, B any](itr fp.Iterator[A], zero B, fn func(B, A) fp.Future[B], ctx ...fp.Executor) fp.Future[B] {
	p := fp.NewPromise[B]()
	p.Success(zero)
	return Fold(itr, p.Future(), func(acc fp.Future[B], v A) fp.Future[B] {
		return acc.FlatMap(func(acc B) fp.Future[B] {
			return fn(acc, v)
		}, ctx...)
	})
}

// FoldError 는  FoldTry[A,fp.Unit]와 같은 함수인데
// 하스켈에서 동일한 기능을 하는 함수를 찾아 보면 traverse_ 혹은 mapM_ 과 같은 함수
// 하스켈에서 _ 가 붙어 있는 함수들은 결과를 discard 해서  m() 를 리턴함.
func FoldError[A any](s fp.Iterator[A], f func(A) error) error {
	for s.HasNext() {
		err := f(s.Next())
		if err != nil {
			return err
		}
	}
	return nil
}

// FoldOption 는 foldM 의 Option 버젼
// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldOption[A, B any](s fp.Iterator[A], zero B, f func(B, A) fp.Option[B]) fp.Option[B] {
	sum := zero
	for s.HasNext() {
		t := f(sum, s.Next())
		if t.IsDefined() {
			sum = t.Get()
		} else {
			return t
		}
	}
	return fp.Some(sum)
}

func FoldRight[A, B any](s fp.Iterator[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if s.IsEmpty() {
		return lazy.Done(zero)
	}

	head := s.Next()

	v := lazy.TailCall(func() lazy.Eval[B] {
		return FoldRight(s, zero, f)
	})

	return f(head, v)

}

func GroupBy[A any, K comparable](s fp.Iterator[A], keyFunc func(A) K) map[K]fp.Seq[A] {

	ret := map[K]fp.Seq[A]{}

	return Fold(s, ret, func(b map[K]fp.Seq[A], a A) map[K]fp.Seq[A] {
		k := keyFunc(a)
		b[k] = b[k].Append(a)
		return b
	})
}

func Scan[A, B any](s fp.Iterator[A], zero B, f func(B, A) B) fp.Iterator[B] {

	first := true
	sum := zero
	hasNext := func() bool {
		if first {
			return true
		}
		return s.HasNext()
	}
	return fp.MakeIterator(
		hasNext,
		func() B {
			if hasNext() {
				if first {
					first = false
					return sum
				}
				sum = f(sum, s.Next())
				return sum
			}
			panic("next on empty iterator")
		},
	)
}

func Generate[T any](generator func() T) fp.Iterator[T] {
	return fp.MakeIterator(func() bool {
		return true
	}, func() T {
		return generator()
	})
}

func Range(from, exclusive int) fp.Iterator[int] {
	i := from
	return fp.MakeIterator(
		func() bool {
			return i < exclusive
		},
		func() int {
			if i < exclusive {
				ret := i
				i++
				return ret
			}
			panic("next on empty iterator")
		},
	)
}

func RangeClosed(from, inclusive int) fp.Iterator[int] {
	i := from
	return fp.MakeIterator(
		func() bool {
			return i <= inclusive
		},
		func() int {
			if i <= inclusive {
				ret := i
				i++
				return ret
			}
			panic("next on empty iterator")
		},
	)
}

func Duplicate[T any](r fp.Iterator[T]) (fp.Iterator[T], fp.Iterator[T]) {
	lock := sync.Mutex{}

	queue := []T{}

	leftAhead := true

	unseq := func(r []T) (fp.Option[T], []T) {
		if len(r) > 0 {
			return fp.Some(r[0]), r[1:]
		} else {
			return fp.None[T](), nil
		}
	}

	left := fp.MakeIterator(
		func() bool {
			lock.Lock()
			defer lock.Unlock()

			if leftAhead || len(queue) == 0 {
				return r.HasNext()
			}
			// queue not empty
			return true
		},
		func() T {
			lock.Lock()
			defer lock.Unlock()

			if len(queue) == 0 {
				leftAhead = true
			}
			if leftAhead {
				ret := r.Next()
				queue = append(queue, ret)
				return ret
			}

			// leftAhead == false means queue not empty
			head, tail := unseq(queue)
			queue = tail
			return head.Get()
		},
	)

	right := fp.MakeIterator(
		func() bool {
			lock.Lock()
			defer lock.Unlock()

			if !leftAhead || len(queue) == 0 {
				return r.HasNext()
			}
			// queue not empty
			return true
		},
		func() T {
			lock.Lock()
			defer lock.Unlock()

			if len(queue) == 0 {
				// right ahead
				leftAhead = false
			}
			if !leftAhead {
				ret := r.Next()
				queue = append(queue, ret)
				return ret
			}
			// rightAhead means queue not empty
			head, tail := unseq(queue)
			queue = tail
			return head.Get()
		},
	)

	return left, right
}

func Span[T any](r fp.Iterator[T], p func(T) bool) (fp.Iterator[T], fp.Iterator[T]) {
	left, right := Duplicate(r)

	return left.TakeWhile(p), right.DropWhile(p)

}

func Partition[T any](r fp.Iterator[T], p func(T) bool) (fp.Iterator[T], fp.Iterator[T]) {
	left, right := Duplicate(r)

	return left.Filter(p), right.FilterNot(p)
}

func PartitionEithers[L, R any](r fp.Iterator[fp.Either[L, R]]) (fp.Iterator[L], fp.Iterator[R]) {
	left, right := Duplicate(r)
	return Map(left.Filter(xtr.IsLeft), xtr.Left), Map(right.Filter(xtr.IsRight), xtr.Get)
}

type seqSorter[T any] struct {
	seq fp.Seq[T]
	ord fp.Ord[T]
}

func (p *seqSorter[T]) Len() int           { return len(p.seq) }
func (p *seqSorter[T]) Less(i, j int) bool { return p.ord.Less(p.seq[i], p.seq[j]) }
func (p *seqSorter[T]) Swap(i, j int)      { p.seq[i], p.seq[j] = p.seq[j], p.seq[i] }

func Sort[T any](r fp.Iterator[T], ord fp.Ord[T]) fp.Seq[T] {
	s := r.ToSeq()
	sort.Sort(&seqSorter[T]{s, ord})
	return s
}

func Min[T any](r fp.Iterator[T], ord fp.Ord[T]) fp.Option[T] {
	return Fold(r, fp.Option[T]{}, func(min fp.Option[T], v T) fp.Option[T] {
		if min.IsDefined() && ord.Less(min.Get(), v) {
			return min
		}
		return fp.Some[T](v)
	})
}

func Max[T any](r fp.Iterator[T], ord fp.Ord[T]) fp.Option[T] {
	return Fold(r, fp.Option[T]{}, func(max fp.Option[T], v T) fp.Option[T] {
		if max.IsDefined() && ord.Less(v, max.Get()) {
			return max
		}
		return fp.Some[T](v)
	})
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
func Flap{{.N}}[{{TypeArgs 1 .N}}, R any](tf fp.Iterator[{{CurriedFunc 1 .N "R"}}]) {{CurriedFunc 1 .N "fp.Iterator[R]"}} {
	return func(a1 A1) {{CurriedFunc 2 .N "fp.Iterator[R]"}} {
		return Flap{{dec .N}}(Ap(tf, Of(a1)))
	}
}

func Method{{.N}}[{{TypeArgs 1 .N}}, R any](ta1 fp.Iterator[A1], fa1 func({{TypeArgs 1 .N}}) R) func({{TypeArgs 2 .N}}) fp.Iterator[R] {
	return func({{DeclArgs 2 .N}}) fp.Iterator[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1({{CallArgs 1 .N}})
		})
	}
}
	`,
}
