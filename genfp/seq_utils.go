package genfp

import (
	"bytes"
	"fmt"
)

func asPtr[T any](v T) *T {
	return &v
}

func seqFlatMap[T, U any](opt []T, fn func(v T) []U) []U {
	ret := make([]U, 0, len(opt))

	for _, v := range opt {
		ret = append(ret, fn(v)...)
	}

	return ret
}

func seqMap[T, U any](opt []T, fn func(v T) U) []U {
	ret := make([]U, len(opt))

	for i, v := range opt {
		ret[i] = fn(v)
	}

	return ret
}

func seqFilter[T any](r []T, p func(v T) bool) []T {
	ret := make([]T, 0, len(r))
	for _, v := range r {
		if p(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func seqLast[T any](r []T) (T, bool) {
	if len(r) > 0 {
		return r[len(r)-1], true
	} else {
		var zero T
		return zero, false
	}
}

func seqMakeString[T any](r []T, sep string) string {
	buf := &bytes.Buffer{}

	for i, v := range r {
		if i != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(fmt.Sprint(v))
	}
	return buf.String()
}

type tuple2[A, B any] struct {
	I1 A
	I2 B
}

func (r tuple2[A, B]) Unapply() (A, B) {
	return r.I1, r.I2
}

func seqToGoMap[K comparable, V any](s []tuple2[K, V]) map[K]V {
	ret := map[K]V{}
	for _, e := range s {
		k, v := e.Unapply()
		ret[k] = v
	}
	return ret
}

func asTuple[A, B any](a A, b B) tuple2[A, B] {
	return tuple2[A, B]{a, b}
}
