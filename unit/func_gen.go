package unit

import (
	"github.com/csgura/fp"
)

func Func1[A1 any](f func(A1)) fp.Func1[A1, fp.Unit] {
	return func(a1 A1) fp.Unit {
		f(a1)
		return fp.Unit{}
	}
}

func Func2[A1, A2 any](f func(A1, A2)) fp.Func2[A1, A2, fp.Unit] {
	return func(a1 A1, a2 A2) fp.Unit {
		f(a1, a2)
		return fp.Unit{}
	}
}

func Func3[A1, A2, A3 any](f func(A1, A2, A3)) fp.Func3[A1, A2, A3, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3) fp.Unit {
		f(a1, a2, a3)
		return fp.Unit{}
	}
}

func Func4[A1, A2, A3, A4 any](f func(A1, A2, A3, A4)) fp.Func4[A1, A2, A3, A4, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Unit {
		f(a1, a2, a3, a4)
		return fp.Unit{}
	}
}

func Func5[A1, A2, A3, A4, A5 any](f func(A1, A2, A3, A4, A5)) fp.Func5[A1, A2, A3, A4, A5, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Unit {
		f(a1, a2, a3, a4, a5)
		return fp.Unit{}
	}
}
