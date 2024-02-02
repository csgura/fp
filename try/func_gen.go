// Code generated by template_gen, DO NOT EDIT.
package try

import (
	"github.com/csgura/fp"
)

func Func1[A1, R any](f func(A1) (R, error)) fp.Func1[A1, fp.Try[R]] {
	return func(a1 A1) fp.Try[R] {
		ret, err := f(a1)
		return Apply(ret, err)
	}
}

func Pure1[A1, R any](f func(A1) R) fp.Func1[A1, fp.Try[R]] {
	return func(a1 A1) fp.Try[R] {
		return Success(f(a1))
	}
}

func Unit1[A1 any](f func(A1) error) fp.Func1[A1, fp.Try[fp.Unit]] {
	return func(a1 A1) fp.Try[fp.Unit] {
		err := f(a1)
		return Apply(fp.Unit{}, err)
	}
}

func Ptr1[A1, R any](f func(A1) (*R, error)) fp.Func1[A1, fp.Try[R]] {
	return func(a1 A1) fp.Try[R] {
		ret, err := f(a1)
		return FlatMap(Apply(ret, err), FromPtr)
	}
}

func Func2[A1, A2, R any](f func(A1, A2) (R, error)) fp.Func2[A1, A2, fp.Try[R]] {
	return func(a1 A1, a2 A2) fp.Try[R] {
		ret, err := f(a1, a2)
		return Apply(ret, err)
	}
}

func Pure2[A1, A2, R any](f func(A1, A2) R) fp.Func2[A1, A2, fp.Try[R]] {
	return func(a1 A1, a2 A2) fp.Try[R] {
		return Success(f(a1, a2))
	}
}

func Unit2[A1, A2 any](f func(A1, A2) error) fp.Func2[A1, A2, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2) fp.Try[fp.Unit] {
		err := f(a1, a2)
		return Apply(fp.Unit{}, err)
	}
}

func Ptr2[A1, A2, R any](f func(A1, A2) (*R, error)) fp.Func2[A1, A2, fp.Try[R]] {
	return func(a1 A1, a2 A2) fp.Try[R] {
		ret, err := f(a1, a2)
		return FlatMap(Apply(ret, err), FromPtr)
	}
}

func Func3[A1, A2, A3, R any](f func(A1, A2, A3) (R, error)) fp.Func3[A1, A2, A3, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Try[R] {
		ret, err := f(a1, a2, a3)
		return Apply(ret, err)
	}
}

func Pure3[A1, A2, A3, R any](f func(A1, A2, A3) R) fp.Func3[A1, A2, A3, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Try[R] {
		return Success(f(a1, a2, a3))
	}
}

func Unit3[A1, A2, A3 any](f func(A1, A2, A3) error) fp.Func3[A1, A2, A3, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Try[fp.Unit] {
		err := f(a1, a2, a3)
		return Apply(fp.Unit{}, err)
	}
}

func Ptr3[A1, A2, A3, R any](f func(A1, A2, A3) (*R, error)) fp.Func3[A1, A2, A3, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Try[R] {
		ret, err := f(a1, a2, a3)
		return FlatMap(Apply(ret, err), FromPtr)
	}
}

func Func4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) (R, error)) fp.Func4[A1, A2, A3, A4, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4)
		return Apply(ret, err)
	}
}

func Pure4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) R) fp.Func4[A1, A2, A3, A4, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R] {
		return Success(f(a1, a2, a3, a4))
	}
}

func Unit4[A1, A2, A3, A4 any](f func(A1, A2, A3, A4) error) fp.Func4[A1, A2, A3, A4, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4)
		return Apply(fp.Unit{}, err)
	}
}

func Ptr4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) (*R, error)) fp.Func4[A1, A2, A3, A4, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4)
		return FlatMap(Apply(ret, err), FromPtr)
	}
}

func Func5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) (R, error)) fp.Func5[A1, A2, A3, A4, A5, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5)
		return Apply(ret, err)
	}
}

func Pure5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) R) fp.Func5[A1, A2, A3, A4, A5, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
		return Success(f(a1, a2, a3, a4, a5))
	}
}

func Unit5[A1, A2, A3, A4, A5 any](f func(A1, A2, A3, A4, A5) error) fp.Func5[A1, A2, A3, A4, A5, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5)
		return Apply(fp.Unit{}, err)
	}
}

func Ptr5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) (*R, error)) fp.Func5[A1, A2, A3, A4, A5, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5)
		return FlatMap(Apply(ret, err), FromPtr)
	}
}

func Func6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) (R, error)) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6)
		return Apply(ret, err)
	}
}

func Pure6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) R) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
		return Success(f(a1, a2, a3, a4, a5, a6))
	}
}

func Unit6[A1, A2, A3, A4, A5, A6 any](f func(A1, A2, A3, A4, A5, A6) error) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5, a6)
		return Apply(fp.Unit{}, err)
	}
}

func Ptr6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) (*R, error)) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6)
		return FlatMap(Apply(ret, err), FromPtr)
	}
}

func Func7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) (R, error)) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6, a7)
		return Apply(ret, err)
	}
}

func Pure7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) R) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
		return Success(f(a1, a2, a3, a4, a5, a6, a7))
	}
}

func Unit7[A1, A2, A3, A4, A5, A6, A7 any](f func(A1, A2, A3, A4, A5, A6, A7) error) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5, a6, a7)
		return Apply(fp.Unit{}, err)
	}
}

func Ptr7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) (*R, error)) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6, a7)
		return FlatMap(Apply(ret, err), FromPtr)
	}
}

func Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R, error)) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
		return Apply(ret, err)
	}
}

func Pure8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) R) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
		return Success(f(a1, a2, a3, a4, a5, a6, a7, a8))
	}
}

func Unit8[A1, A2, A3, A4, A5, A6, A7, A8 any](f func(A1, A2, A3, A4, A5, A6, A7, A8) error) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5, a6, a7, a8)
		return Apply(fp.Unit{}, err)
	}
}

func Ptr8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) (*R, error)) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
		return FlatMap(Apply(ret, err), FromPtr)
	}
}

func Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R, error)) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		return Apply(ret, err)
	}
}

func Pure9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) R) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
		return Success(f(a1, a2, a3, a4, a5, a6, a7, a8, a9))
	}
}

func Unit9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) error) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		return Apply(fp.Unit{}, err)
	}
}

func Ptr9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (*R, error)) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		return FlatMap(Apply(ret, err), FromPtr)
	}
}

func Compose3[A1, A2, A3, R any](f1 fp.Func1[A1, fp.Try[A2]], f2 fp.Func1[A2, fp.Try[A3]], f3 fp.Func1[A3, fp.Try[R]]) fp.Func1[A1, fp.Try[R]] {
	return Compose2(f1, Compose2(f2, f3))
}

func Compose4[A1, A2, A3, A4, R any](f1 fp.Func1[A1, fp.Try[A2]], f2 fp.Func1[A2, fp.Try[A3]], f3 fp.Func1[A3, fp.Try[A4]], f4 fp.Func1[A4, fp.Try[R]]) fp.Func1[A1, fp.Try[R]] {
	return Compose2(f1, Compose3(f2, f3, f4))
}

func Compose5[A1, A2, A3, A4, A5, R any](f1 fp.Func1[A1, fp.Try[A2]], f2 fp.Func1[A2, fp.Try[A3]], f3 fp.Func1[A3, fp.Try[A4]], f4 fp.Func1[A4, fp.Try[A5]], f5 fp.Func1[A5, fp.Try[R]]) fp.Func1[A1, fp.Try[R]] {
	return Compose2(f1, Compose4(f2, f3, f4, f5))
}
