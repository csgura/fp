package slice

import (
	"sort"

	"github.com/csgura/fp"
	"github.com/csgura/fp/immutable"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/product"
)

func Iterator[T any](r fp.Slice[T]) fp.Iterator[T] {
	return fp.IteratorOfSeq(r)
}

func Collect[T any](r fp.Iterator[T]) fp.Slice[T] {
	ret := fp.Slice[T]{}
	for r.HasNext() {
		ret = append(ret, r.Next())
	}
	return ret
}

func Empty[T any]() fp.Slice[T] {
	return nil
}

func Pure[T any](v T) fp.Slice[T] {
	return fp.Slice[T]{v}
}

func Of[T any](list ...T) fp.Slice[T] {
	return list
}

func MapKey[KA, KB, V any](s fp.Slice[fp.Tuple2[KA, V]], f func(KA) KB) fp.Slice[fp.Tuple2[KB, V]] {
	return Map(s, func(v fp.Tuple2[KA, V]) fp.Tuple2[KB, V] {
		return fp.Tuple2[KB, V]{
			I1: f(v.I1),
			I2: v.I2,
		}
	})
}

func FilterMapKey[KA, KB, V any](s fp.Slice[fp.Tuple2[KA, V]], f func(KA) fp.Option[KB]) fp.Slice[fp.Tuple2[KB, V]] {
	return FilterMap(s, func(v fp.Tuple2[KA, V]) fp.Option[fp.Tuple2[KB, V]] {
		return option.Zip(f(v.I1), option.Some(v.I2))
	})
}

func MapValue[K, VA, VB any](s fp.Slice[fp.Tuple2[K, VA]], f func(VA) VB) fp.Slice[fp.Tuple2[K, VB]] {
	return Map(s, func(v fp.Tuple2[K, VA]) fp.Tuple2[K, VB] {
		return fp.Tuple2[K, VB]{
			I1: v.I1,
			I2: f(v.I2),
		}
	})
}

func FilterMapValue[K, VA, VB any](s fp.Slice[fp.Tuple2[K, VA]], f func(VA) fp.Option[VB]) fp.Slice[fp.Tuple2[K, VB]] {
	return FilterMap(s, func(v fp.Tuple2[K, VA]) fp.Option[fp.Tuple2[K, VB]] {
		return option.Zip(option.Some(v.I1), f(v.I2))
	})
}

func FromMap[K comparable, V any](m map[K]V) fp.Slice[fp.Tuple2[K, V]] {
	seq := make([]fp.Tuple2[K, V], 0, len(m))
	for k, v := range m {
		seq = append(seq, fp.Tuple2[K, V]{I1: k, I2: v})
	}
	return seq
}

func FromMapValues[K comparable, V any](m map[K]V) fp.Slice[V] {
	seq := make([]V, 0, len(m))
	for _, v := range m {
		seq = append(seq, v)
	}
	return seq
}

func FromMapKeys[K comparable, V any](m map[K]V) fp.Slice[K] {
	seq := make([]K, 0, len(m))
	for k := range m {
		seq = append(seq, k)
	}
	return seq
}

func Ap[T, U any](t fp.Slice[fp.Func1[T, U]], a fp.Slice[T]) fp.Slice[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Slice[U] {
		return Map(a, f)
	})
}

func MapWithIndex[T, U any](opt fp.Slice[T], fn func(i int, v T) U) fp.Slice[U] {
	ret := make(fp.Slice[U], len(opt))

	for i, v := range opt {
		ret[i] = fn(i, v)
	}

	return ret
}

func Map2[A, B, U any](a fp.Slice[A], b fp.Slice[B], f func(A, B) U) fp.Slice[U] {
	return FlatMap(a, func(v1 A) fp.Slice[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		})
	})
}

func FilterMap[T, U any](opt fp.Slice[T], fn func(v T) fp.Option[U]) fp.Slice[U] {
	return FlatMap(opt, fp.Compose(fn, option.ToSlice))
}

func FilterNil[T any](opt fp.Slice[*T]) fp.Slice[T] {
	return FilterMap(opt, option.Ptr)
}

func Lift[T, U any](f func(v T) U) func(fp.Slice[T]) fp.Slice[U] {
	return func(opt fp.Slice[T]) fp.Slice[U] {
		return Map(opt, f)
	}
}

func LiftM[T, U any](f func(v T) fp.Slice[U]) func(fp.Slice[T]) fp.Slice[U] {
	return func(opt fp.Slice[T]) fp.Slice[U] {
		return FlatMap(opt, f)
	}
}

func Compose[A, B, C any](f1 func(A) fp.Slice[B], f2 func(B) fp.Slice[C]) func(A) fp.Slice[C] {
	return func(a A) fp.Slice[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposePure[A, B any](fab func(A) B) func(A) fp.Slice[B] {
	return func(a A) fp.Slice[B] {
		return Of(fab(a))
	}
}

func Flatten[T any](opt fp.Slice[fp.Slice[T]]) fp.Slice[T] {
	return FlatMap(opt, func(v fp.Slice[T]) fp.Slice[T] {
		return v
	})
}

func Prepend[T any](head T, tail fp.Slice[T]) fp.Slice[T] {
	return Concat(Of(head), tail)
}

func Zip[A, B any](s1 fp.Slice[A], s2 fp.Slice[B]) fp.Slice[fp.Tuple2[A, B]] {
	minSize := fp.Min(Size(s1), Size(s2))

	ret := make(fp.Slice[fp.Tuple2[A, B]], minSize)
	for i := 0; i < minSize; i++ {
		ret[i] = product.Tuple2(s1[i], s2[i])
	}
	return ret
}

func ZipWithIndex[A any](s1 fp.Slice[A]) fp.Slice[fp.Tuple2[int, A]] {

	ret := make(fp.Slice[fp.Tuple2[int, A]], Size(s1))
	for i := 0; i < Size(s1); i++ {
		ret[i] = product.Tuple2(i, s1[i])
	}
	return ret
}

func Reduce[T any](r fp.Slice[T], m fp.Monoid[T]) T {
	if Size(r) == 0 {
		return m.Empty()
	}

	reduce := m.Empty()
	for i := 0; i < len(r); i++ {
		reduce = m.Combine(reduce, r[i])
	}

	return reduce
}

func Fold[A, B any](s fp.Slice[A], zero B, f func(B, A) B) B {
	sum := zero
	for _, v := range s {
		sum = f(sum, v)
	}
	return sum
}

// FoldTry 는 foldM 의 Try 버젼
// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldTry[A, B any](s fp.Slice[A], zero B, f func(B, A) fp.Try[B]) fp.Try[B] {
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

func FoldFuture[A, B any](itr fp.Slice[A], zero B, fn func(B, A) fp.Future[B], ctx ...fp.Executor) fp.Future[B] {
	p := fp.NewPromise[B]()
	p.Success(zero)
	return Fold(itr, p.Future(), func(acc fp.Future[B], v A) fp.Future[B] {
		return acc.FlatMap(func(acc B) fp.Future[B] {
			return fn(acc, v)
		}, ctx...)
	})
}

// FoldOption 는 foldM 의 Option 버젼
// foldM : (b -> a -> m b ) -> b -> t a -> m b
func FoldOption[A, B any](s fp.Slice[A], zero B, f func(B, A) fp.Option[B]) fp.Option[B] {
	sum := zero
	for _, v := range s {
		t := f(sum, v)
		if t.IsDefined() {
			sum = t.Get()
		} else {
			return t
		}
	}
	return fp.Some(sum)
}

// FoldError 는  FoldTry[A,fp.Unit]와 같은 함수인데
// 하스켈에서 동일한 기능을 하는 함수를 찾아 보면 traverse_ 혹은 mapM_ 과 같은 함수
// 하스켈에서 _ 가 붙어 있는 함수들은 결과를 discard 해서  m() 를 리턴함.
func FoldError[A any](s fp.Slice[A], f func(A) error) error {
	for _, v := range s {
		err := f(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func FoldMap[A, M any](s fp.Slice[A], m fp.Monoid[M], f func(A) M) M {
	return Fold(s, m.Empty(), func(b M, a A) M {
		return m.Combine(b, f(a))
	})
}

func FoldRight[A, B any](s fp.Slice[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if IsEmpty(s) {
		return lazy.Done(zero)
	}

	head, tail := Unapply(s)
	v := lazy.TailCall(func() lazy.Eval[B] {
		return FoldRight(tail, zero, f)
	})

	return f(head.Get(), v)
}

// func FoldRightImplementUsingFoldMap[A, B any](s fp.Slice[A], zero B, f func(A, B) B) B {

// 	m := monoid.Endo[B]()

// 	c := as.Curried2(f)

// 	ret := FoldMap(s, m, func(a A) fp.Endo[B] {
// 		return fp.Endo[B](c(a))
// 	})

// 	return ret(zero)
// }

func Scan[A, B any](s fp.Slice[A], zero B, f func(B, A) B) fp.Slice[B] {

	if IsEmpty(s) {
		return Of(zero)
	}

	ret := make(fp.Slice[B], Size(s)+1)
	sum := zero
	ret[0] = sum
	for i, v := range s {
		sum = f(sum, v)
		ret[i+1] = sum
	}
	return ret

}

func GroupBy[A any, K comparable](s fp.Slice[A], keyFunc func(A) K) map[K]fp.Slice[A] {

	ret := map[K]fp.Slice[A]{}

	return Fold(s, ret, func(b map[K]fp.Slice[A], a A) map[K]fp.Slice[A] {
		k := keyFunc(a)
		b[k] = Append(b[k], a)
		return b
	})
}

func Distinct[V comparable](s fp.Slice[V]) fp.Slice[V] {

	dupcheck := map[V]bool{}

	return Fold(s, make(fp.Slice[V], 0, Size(s)), func(acc fp.Slice[V], a V) fp.Slice[V] {
		if dupcheck[a] {
			return acc
		}

		dupcheck[a] = true
		return append(acc, a)
	})
}

func ToMap[K, V any](s fp.Slice[fp.Tuple2[K, V]], hasher fp.Hashable[K]) fp.Map[K, V] {
	ret := immutable.MapBuilder[K, V](hasher)

	for _, e := range s {
		k, v := e.Unapply()
		ret = ret.Add(k, v)
	}

	return ret.Build()
}

func ToGoMap[K comparable, V any](s fp.Slice[fp.Tuple2[K, V]]) map[K]V {
	ret := map[K]V{}
	for _, e := range s {
		k, v := e.Unapply()
		ret[k] = v
	}
	return ret
}

func ToSet[V any](s fp.Slice[V], hasher fp.Hashable[V]) fp.Set[V] {
	ret := immutable.SetBuilder(hasher)

	for _, e := range s {
		ret = ret.Add(e)
	}

	return ret.Build()
}

func ToGoSet[V comparable](s fp.Slice[V]) mutable.Set[V] {
	ret := map[V]bool{}
	for _, e := range s {
		ret[e] = true
	}
	return ret
}

type seqSorter[T any] struct {
	seq fp.Slice[T]
	ord fp.Ord[T]
}

func (p *seqSorter[T]) Len() int           { return len(p.seq) }
func (p *seqSorter[T]) Less(i, j int) bool { return p.ord.Less(p.seq[i], p.seq[j]) }
func (p *seqSorter[T]) Swap(i, j int)      { p.seq[i], p.seq[j] = p.seq[j], p.seq[i] }

func Sort[T any](r fp.Slice[T], ord fp.Ord[T]) fp.Slice[T] {
	ns := Concat(r, nil)
	sort.Sort(&seqSorter[T]{ns, ord})
	return ns
}

func Min[T any](r fp.Slice[T], ord fp.Ord[T]) fp.Option[T] {
	return Fold(r, fp.Option[T]{}, func(min fp.Option[T], v T) fp.Option[T] {
		if min.IsDefined() && ord.Less(min.Get(), v) {
			return min
		}
		return fp.Some[T](v)
	})
}

func Max[T any](r fp.Slice[T], ord fp.Ord[T]) fp.Option[T] {
	return Fold(r, fp.Option[T]{}, func(max fp.Option[T], v T) fp.Option[T] {
		if max.IsDefined() && ord.Less(v, max.Get()) {
			return max
		}
		return fp.Some[T](v)
	})
}

func Span[T any](r fp.Slice[T], p func(T) bool) (fp.Slice[T], fp.Slice[T]) {
	left := fp.Slice[T]{}
	right := fp.Slice[T]{}

	span := false
	for _, v := range r {
		if span {
			right = append(right, v)
		} else {
			if p(v) {
				left = append(left, v)
			} else {
				span = true
				right = append(right, v)
			}
		}
	}

	return left, right

}

func Partition[T any](r fp.Slice[T], p func(T) bool) (fp.Slice[T], fp.Slice[T]) {
	left := fp.Slice[T]{}
	right := fp.Slice[T]{}

	for _, v := range r {
		if p(v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return left, right
}

func PartitionEithers[L, R any](r fp.Slice[fp.Either[L, R]]) (fp.Slice[L], fp.Slice[R]) {
	left := fp.Slice[L]{}
	right := fp.Slice[R]{}

	for _, v := range r {
		if v.IsLeft() {
			left = append(left, v.Left())
		} else {
			right = append(right, v.Get())
		}
	}
	return left, right
}
