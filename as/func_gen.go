package as

import (
	"github.com/csgura/fp"
)

func Func1[A1, R any](f func(A1) R) fp.Func1[A1, R] {
	return fp.Func1[A1, R](f)
}

func Func2[A1, A2, R any](f func(A1, A2) R) fp.Func2[A1, A2, R] {
	return fp.Func2[A1, A2, R](f)
}

func Func3[A1, A2, A3, R any](f func(A1, A2, A3) R) fp.Func3[A1, A2, A3, R] {
	return fp.Func3[A1, A2, A3, R](f)
}

func Func4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) R) fp.Func4[A1, A2, A3, A4, R] {
	return fp.Func4[A1, A2, A3, A4, R](f)
}

func Func5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) R) fp.Func5[A1, A2, A3, A4, A5, R] {
	return fp.Func5[A1, A2, A3, A4, A5, R](f)
}

func Func6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) R) fp.Func6[A1, A2, A3, A4, A5, A6, R] {
	return fp.Func6[A1, A2, A3, A4, A5, A6, R](f)
}

func Func7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) R) fp.Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return fp.Func7[A1, A2, A3, A4, A5, A6, A7, R](f)
}

func Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) R) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R](f)
}

func Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) R) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R](f)
}

func Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) R) fp.Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	return fp.Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R](f)
}

func Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) R) fp.Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	return fp.Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R](f)
}

func Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) R) fp.Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	return fp.Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R](f)
}

func Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) R) fp.Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	return fp.Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R](f)
}

func Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) R) fp.Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	return fp.Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R](f)
}

func Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) R) fp.Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	return fp.Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R](f)
}

func Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) R) fp.Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	return fp.Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R](f)
}

func Func17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) R) fp.Func17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R] {
	return fp.Func17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R](f)
}

func Func18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) R) fp.Func18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R] {
	return fp.Func18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R](f)
}

func Func19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) R) fp.Func19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R] {
	return fp.Func19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R](f)
}

func Func20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) R) fp.Func20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R] {
	return fp.Func20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R](f)
}

func Func21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) R) fp.Func21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R] {
	return fp.Func21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R](f)
}

func Func22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) R) fp.Func22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R] {
	return fp.Func22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R](f)
}
