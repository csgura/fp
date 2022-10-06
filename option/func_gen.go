package option

import (
	"github.com/csgura/fp"
)

func LiftA3[A1, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) R) fp.Func3[fp.Option[A1], fp.Option[A2], fp.Option[A3], fp.Option[R]] {
	return func(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3]) fp.Option[R] {

		return FlatMap(ins1, func(a1 A1) fp.Option[R] {
			return LiftA2(func(a2 A2, a3 A3) R {
				return f(a1, a2, a3)
			})(ins2, ins3)
		})
	}
}

func LiftA4[A1, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) R) fp.Func4[fp.Option[A1], fp.Option[A2], fp.Option[A3], fp.Option[A4], fp.Option[R]] {
	return func(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4]) fp.Option[R] {

		return FlatMap(ins1, func(a1 A1) fp.Option[R] {
			return LiftA3(func(a2 A2, a3 A3, a4 A4) R {
				return f(a1, a2, a3, a4)
			})(ins2, ins3, ins4)
		})
	}
}

func LiftA5[A1, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) fp.Func5[fp.Option[A1], fp.Option[A2], fp.Option[A3], fp.Option[A4], fp.Option[A5], fp.Option[R]] {
	return func(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5]) fp.Option[R] {

		return FlatMap(ins1, func(a1 A1) fp.Option[R] {
			return LiftA4(func(a2 A2, a3 A3, a4 A4, a5 A5) R {
				return f(a1, a2, a3, a4, a5)
			})(ins2, ins3, ins4, ins5)
		})
	}
}

func LiftA6[A1, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) fp.Func6[fp.Option[A1], fp.Option[A2], fp.Option[A3], fp.Option[A4], fp.Option[A5], fp.Option[A6], fp.Option[R]] {
	return func(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5], ins6 fp.Option[A6]) fp.Option[R] {

		return FlatMap(ins1, func(a1 A1) fp.Option[R] {
			return LiftA5(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
				return f(a1, a2, a3, a4, a5, a6)
			})(ins2, ins3, ins4, ins5, ins6)
		})
	}
}

func LiftA7[A1, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) fp.Func7[fp.Option[A1], fp.Option[A2], fp.Option[A3], fp.Option[A4], fp.Option[A5], fp.Option[A6], fp.Option[A7], fp.Option[R]] {
	return func(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5], ins6 fp.Option[A6], ins7 fp.Option[A7]) fp.Option[R] {

		return FlatMap(ins1, func(a1 A1) fp.Option[R] {
			return LiftA6(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
				return f(a1, a2, a3, a4, a5, a6, a7)
			})(ins2, ins3, ins4, ins5, ins6, ins7)
		})
	}
}

func LiftA8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) fp.Func8[fp.Option[A1], fp.Option[A2], fp.Option[A3], fp.Option[A4], fp.Option[A5], fp.Option[A6], fp.Option[A7], fp.Option[A8], fp.Option[R]] {
	return func(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5], ins6 fp.Option[A6], ins7 fp.Option[A7], ins8 fp.Option[A8]) fp.Option[R] {

		return FlatMap(ins1, func(a1 A1) fp.Option[R] {
			return LiftA7(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
				return f(a1, a2, a3, a4, a5, a6, a7, a8)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8)
		})
	}
}

func LiftA9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) fp.Func9[fp.Option[A1], fp.Option[A2], fp.Option[A3], fp.Option[A4], fp.Option[A5], fp.Option[A6], fp.Option[A7], fp.Option[A8], fp.Option[A9], fp.Option[R]] {
	return func(ins1 fp.Option[A1], ins2 fp.Option[A2], ins3 fp.Option[A3], ins4 fp.Option[A4], ins5 fp.Option[A5], ins6 fp.Option[A6], ins7 fp.Option[A7], ins8 fp.Option[A8], ins9 fp.Option[A9]) fp.Option[R] {

		return FlatMap(ins1, func(a1 A1) fp.Option[R] {
			return LiftA8(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
				return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)
		})
	}
}

func Compose3[A1, A2, A3, R any](f1 fp.Func1[A1, fp.Option[A2]], f2 fp.Func1[A2, fp.Option[A3]], f3 fp.Func1[A3, fp.Option[R]]) fp.Func1[A1, fp.Option[R]] {
	return Compose2(f1, Compose2(f2, f3))
}

func Compose4[A1, A2, A3, A4, R any](f1 fp.Func1[A1, fp.Option[A2]], f2 fp.Func1[A2, fp.Option[A3]], f3 fp.Func1[A3, fp.Option[A4]], f4 fp.Func1[A4, fp.Option[R]]) fp.Func1[A1, fp.Option[R]] {
	return Compose2(f1, Compose3(f2, f3, f4))
}

func Compose5[A1, A2, A3, A4, A5, R any](f1 fp.Func1[A1, fp.Option[A2]], f2 fp.Func1[A2, fp.Option[A3]], f3 fp.Func1[A3, fp.Option[A4]], f4 fp.Func1[A4, fp.Option[A5]], f5 fp.Func1[A5, fp.Option[R]]) fp.Func1[A1, fp.Option[R]] {
	return Compose2(f1, Compose4(f2, f3, f4, f5))
}
