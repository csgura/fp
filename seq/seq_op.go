package seq

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/product"
)

func Of[T any](list ...T) fp.Seq[T] {
	return list
}

func Map[T, U any](opt fp.Seq[T], fn func(v T) U) fp.Seq[U] {
	ret := make(fp.Seq[U], len(opt))

	for i, v := range opt {
		ret[i] = fn(v)
	}

	return ret
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

func Concact[T any](head T, tail fp.Seq[T]) fp.Seq[T] {
	return Of(head).Concact(tail)
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

func GroupBy[A any, K comparable](s fp.Seq[A], keyFunc func(A) K) map[K]fp.Seq[A] {

	ret := map[K]fp.Seq[A]{}

	return Fold(s, ret, func(b map[K]fp.Seq[A], a A) map[K]fp.Seq[A] {
		k := keyFunc(a)
		b[k] = b[k].Append(a)
		return b
	})
}
