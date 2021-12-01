package seq

import "github.com/csgura/fp"

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
