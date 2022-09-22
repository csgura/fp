package seq

import (
	"sort"

	"github.com/csgura/fp"
	"github.com/csgura/fp/immutable"
	"github.com/csgura/fp/lazy"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/product"
)

func Collect[T any](r fp.Iterator[T]) fp.Seq[T] {
	ret := fp.Seq[T]{}
	for r.HasNext() {
		ret = append(ret, r.Next())
	}
	return ret
}

func Empty[T any]() fp.Seq[T] {
	return nil
}

func Of[T any](list ...T) fp.Seq[T] {
	return list
}

func FromMap[K comparable, V any](m map[K]V) fp.Seq[fp.Tuple2[K, V]] {
	seq := make([]fp.Tuple2[K, V], 0, len(m))
	for k, v := range m {
		seq = append(seq, fp.Tuple2[K, V]{I1: k, I2: v})
	}
	return seq
}

func FromMapValues[K comparable, V any](m map[K]V) fp.Seq[V] {
	seq := make([]V, 0, len(m))
	for _, v := range m {
		seq = append(seq, v)
	}
	return seq
}

func FromMapKeys[K comparable, V any](m map[K]V) fp.Seq[K] {
	seq := make([]K, 0, len(m))
	for k := range m {
		seq = append(seq, k)
	}
	return seq
}

func Ap[T, U any](t fp.Seq[fp.Func1[T, U]], a fp.Seq[T]) fp.Seq[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Seq[U] {
		return Map(a, f)
	})
}

func Map[T, U any](opt fp.Seq[T], fn func(v T) U) fp.Seq[U] {
	ret := make(fp.Seq[U], len(opt))

	for i, v := range opt {
		ret[i] = fn(v)
	}

	return ret
}

func Map2[A, B, U any](a fp.Seq[A], b fp.Seq[B], f func(A, B) U) fp.Seq[U] {
	return FlatMap(a, func(v1 A) fp.Seq[U] {
		return Map(b, func(v2 B) U {
			return f(v1, v2)
		})
	})
}

func Lift[T, U any](f func(v T) U) fp.Func1[fp.Seq[T], fp.Seq[U]] {
	return func(opt fp.Seq[T]) fp.Seq[U] {
		return Map(opt, f)
	}
}

func Compose[A, B, C any](f1 fp.Func1[A, fp.Seq[B]], f2 fp.Func1[B, fp.Seq[C]]) fp.Func1[A, fp.Seq[C]] {
	return func(a A) fp.Seq[C] {
		return FlatMap(f1(a), f2)
	}
}

func ComposePure[A, B, C any](f1 fp.Func1[A, fp.Seq[B]], f2 fp.Func1[B, C]) fp.Func1[A, fp.Seq[C]] {
	return func(a A) fp.Seq[C] {
		return Map(f1(a), f2)
	}
}

func FlatMap[T, U any](opt fp.Seq[T], fn func(v T) fp.Seq[U]) fp.Seq[U] {
	ret := make(fp.Seq[U], 0, len(opt))

	for _, v := range opt {
		ret = append(ret, fn(v)...)
	}

	return ret
}

func Flatten[T any](opt fp.Seq[fp.Seq[T]]) fp.Seq[T] {
	return FlatMap(opt, func(v fp.Seq[T]) fp.Seq[T] {
		return v
	})
}

func Concat[T any](head T, tail fp.Seq[T]) fp.Seq[T] {
	return Of(head).Concat(tail)
}

func Zip[A, B any](s1 fp.Seq[A], s2 fp.Seq[B]) fp.Seq[fp.Tuple2[A, B]] {
	minSize := fp.Min(s1.Size(), s2.Size())

	ret := make(fp.Seq[fp.Tuple2[A, B]], minSize)
	for i := 0; i < minSize; i++ {
		ret[i] = product.Tuple2(s1[i], s2[i])
	}
	return ret
}

func Fold[A, B any](s fp.Seq[A], zero B, f func(B, A) B) B {
	sum := zero
	for _, v := range s {
		sum = f(sum, v)
	}
	return sum
}

func FoldMap[A, M any](s fp.Seq[A], m fp.Monoid[M], f func(A) M) M {
	return Fold(s, m.Empty(), func(b M, a A) M {
		return m.Combine(b, f(a))
	})
}

func FoldRight[A, B any](s fp.Seq[A], zero B, f func(A, lazy.Eval[B]) lazy.Eval[B]) lazy.Eval[B] {
	if s.IsEmpty() {
		return lazy.Done(zero)
	}

	head, tail := s.UnSeq()
	v := lazy.TailCall(func() lazy.Eval[B] {
		return FoldRight(tail, zero, f)
	})

	return f(head.Get(), v)
}

// func FoldRightImplementUsingFoldMap[A, B any](s fp.Seq[A], zero B, f func(A, B) B) B {

// 	m := monoid.Endo[B]()

// 	c := as.Curried2(f)

// 	ret := FoldMap(s, m, func(a A) fp.Endo[B] {
// 		return fp.Endo[B](c(a))
// 	})

// 	return ret(zero)
// }

func Scan[A, B any](s fp.Seq[A], zero B, f func(B, A) B) fp.Seq[B] {

	if s.IsEmpty() {
		return Of(zero)
	}

	ret := make(fp.Seq[B], s.Size()+1)
	sum := zero
	ret[0] = sum
	for i, v := range s {
		sum = f(sum, v)
		ret[i+1] = sum
	}
	return ret

}

func GroupBy[A any, K any](s fp.Seq[A], keyFunc func(A) K) mutable.Map[K, fp.Seq[A]] {

	ret := map[any]fp.Seq[A]{}

	return Fold(s, ret, func(b map[any]fp.Seq[A], a A) map[any]fp.Seq[A] {
		k := keyFunc(a)
		b[k] = b[k].Append(a)
		return b
	})
}

func ToMap[K, V any](s fp.Seq[fp.Tuple2[K, V]], hasher fp.Hashable[K]) fp.Map[K, V] {
	ret := immutable.MapBuilder[K, V](hasher)

	for _, e := range s {
		k, v := e.Unapply()
		ret = ret.Add(k, v)
	}

	return fp.MakeMap(ret.Build())
}

func ToGoMap[K comparable, V any](s fp.Seq[fp.Tuple2[K, V]]) map[K]V {
	ret := map[K]V{}
	for _, e := range s {
		k, v := e.Unapply()
		ret[k] = v
	}
	return ret
}

func ToSet[V any](s fp.Seq[V], hasher fp.Hashable[V]) fp.Set[V] {
	ret := immutable.SetBuilder(hasher)

	for _, e := range s {
		ret = ret.Add(e)
	}

	return ret.Build()
}

type seqSorter[T any] struct {
	seq fp.Seq[T]
	ord fp.Ord[T]
}

func (p *seqSorter[T]) Len() int           { return len(p.seq) }
func (p *seqSorter[T]) Less(i, j int) bool { return p.ord.Less(p.seq[i], p.seq[j]) }
func (p *seqSorter[T]) Swap(i, j int)      { p.seq[i], p.seq[j] = p.seq[j], p.seq[i] }

func Sort[T any](r fp.Seq[T], ord fp.Ord[T]) fp.Seq[T] {
	ns := r.Concat(nil)
	sort.Sort(&seqSorter[T]{ns, ord})
	return ns
}
