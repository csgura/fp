package slice

import (
	"bytes"
	"fmt"

	"github.com/csgura/fp"
)

// fp.Seq[T] 의 method 로 있던것

func Size[T any](r fp.Slice[T]) int {
	return len(r)
}

func IsEmpty[T any](r fp.Slice[T]) bool {
	return Size(r) == 0
}

func NonEmpty[T any](r fp.Slice[T]) bool {
	return Size(r) > 0
}

func Get[T any](r fp.Slice[T], idx int) fp.Option[T] {
	if Size(r) > idx {
		return fp.Some(r[idx])
	} else {
		return fp.None[T]()
	}
}

func Head[T any](r fp.Slice[T]) fp.Option[T] {
	if Size(r) > 0 {
		return fp.Some(r[0])
	} else {
		return fp.None[T]()
	}
}

func Init[T any](r fp.Slice[T]) fp.Slice[T] {
	if Size(r) > 1 {
		return r[:Size(r)-1]
	} else {
		return nil
	}
}

func Last[T any](r fp.Slice[T]) fp.Option[T] {
	if Size(r) > 0 {
		return fp.Some(r[Size(r)-1])
	} else {
		return fp.None[T]()
	}
}

func Tail[T any](r fp.Slice[T]) fp.Slice[T] {
	if Size(r) > 0 {
		return r[1:]
	} else {
		return nil
	}
}

func Unapply[T any](r fp.Slice[T]) (fp.Option[T], fp.Slice[T]) {
	if Size(r) > 0 {
		return Head(r), r[1:]
	} else {
		return Head(r), nil
	}
}

func Take[T any](r fp.Slice[T], n int) fp.Slice[T] {
	if len(r) < n {
		return r
	}
	return r[0:n]
}

func Drop[T any](r fp.Slice[T], n int) fp.Slice[T] {
	if len(r) < n {
		return nil
	}
	return r[n:]
}

func Foreach[T any](r fp.Slice[T], f func(v T)) {
	for _, v := range r {
		f(v)
	}
}

func Filter[T any](r fp.Slice[T], p func(v T) bool) fp.Slice[T] {
	ret := make([]T, 0, len(r))
	for _, v := range r {
		if p(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func FilterNot[T any](r fp.Slice[T], p func(v T) bool) fp.Slice[T] {
	return Filter(r, func(t T) bool {
		return !p(t)
	})
}

func Exists[T any](r fp.Slice[T], p func(v T) bool) bool {
	for _, v := range r {
		if p(v) {
			return true
		}
	}
	return false
}

func ForAll[T any](r fp.Slice[T], p func(v T) bool) bool {
	for _, v := range r {
		if !p(v) {
			return false
		}
	}
	return true
}

func Find[T any](r fp.Slice[T], p func(v T) bool) fp.Option[T] {
	for _, v := range r {
		if p(v) {
			return fp.Some(v)
		}
	}
	return fp.None[T]()
}

func Add[T any](r fp.Slice[T], item T) fp.Slice[T] {
	return Append(r, item)
}

func Append[T any](r fp.Slice[T], items ...T) fp.Slice[T] {
	if len(items) > 0 {
		tail := fp.Slice[T](items)
		ret := make(fp.Slice[T], Size(r)+Size(tail))

		copy(ret, r)

		for i := range tail {
			ret[i+Size(r)] = tail[i]
		}

		return ret
	}
	return r
}

func Concat[T any](r fp.Slice[T], tail fp.Slice[T]) fp.Slice[T] {
	if len(tail) > 0 {
		ret := make(fp.Slice[T], Size(r)+Size(tail))

		copy(ret, r)

		for i := range tail {
			ret[i+Size(r)] = tail[i]
		}

		return ret
	}
	return r
}

func Reverse[T any](r fp.Slice[T]) fp.Slice[T] {
	ret := make(fp.Slice[T], Size(r))

	for i := range r {
		ret[Size(r)-i-1] = r[i]
	}

	return ret
}

func MakeString[T any](r fp.Slice[T], sep string) string {
	buf := &bytes.Buffer{}

	for i, v := range r {
		if i != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(fmt.Sprint(v))
	}
	return buf.String()
}
