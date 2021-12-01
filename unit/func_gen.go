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

func Func6[A1, A2, A3, A4, A5, A6 any](f func(A1, A2, A3, A4, A5, A6)) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Unit {
		f(a1, a2, a3, a4, a5, a6)
		return fp.Unit{}
	}
}

func Func7[A1, A2, A3, A4, A5, A6, A7 any](f func(A1, A2, A3, A4, A5, A6, A7)) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7)
		return fp.Unit{}
	}
}

func Func8[A1, A2, A3, A4, A5, A6, A7, A8 any](f func(A1, A2, A3, A4, A5, A6, A7, A8)) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8)
		return fp.Unit{}
	}
}

func Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9)) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		return fp.Unit{}
	}
}

func Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10)) fp.Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
		return fp.Unit{}
	}
}

func Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11)) fp.Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
		return fp.Unit{}
	}
}

func Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12)) fp.Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
		return fp.Unit{}
	}
}

func Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13)) fp.Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
		return fp.Unit{}
	}
}

func Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14)) fp.Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
		return fp.Unit{}
	}
}

func Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15)) fp.Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
		return fp.Unit{}
	}
}

func Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16)) fp.Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
		return fp.Unit{}
	}
}

func Func17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17)) fp.Func17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
		return fp.Unit{}
	}
}

func Func18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18)) fp.Func18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
		return fp.Unit{}
	}
}

func Func19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19)) fp.Func19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
		return fp.Unit{}
	}
}

func Func20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20)) fp.Func20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
		return fp.Unit{}
	}
}

func Func21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21)) fp.Func21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
		return fp.Unit{}
	}
}

func Func22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22)) fp.Func22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, fp.Unit] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) fp.Unit {
		f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
		return fp.Unit{}
	}
}
