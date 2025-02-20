package slice

import "github.com/csgura/fp"

func Map[T, U any](opt fp.Slice[T], fn func(v T) U) fp.Seq[U] {
	ret := make(fp.Seq[U], len(opt))

	for i, v := range opt {
		ret[i] = fn(v)
	}

	return ret
}

func FlatMap[T, U any](opt fp.Slice[T], fn func(v T) fp.Slice[U]) fp.Slice[U] {
	ret := make(fp.Slice[U], 0, len(opt))

	for _, v := range opt {
		ret = append(ret, fn(v)...)
	}

	return ret
}
