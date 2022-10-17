package future

import (
	"github.com/csgura/fp"
)

func LiftA3[A1, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) R, exec ...fp.Executor) fp.Func3[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftA2(func(a2 A2, a3 A3) R {
				return f(a1, a2, a3)
			}, exec...)(ins2, ins3)
		}, exec...)
	}
}

func LiftM3[A1, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) fp.Future[R], exec ...fp.Executor) fp.Func3[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftM2(func(a2 A2, a3 A3) fp.Future[R] {
				return f(a1, a2, a3)
			}, exec...)(ins2, ins3)
		}, exec...)
	}
}

func Flap3[A1, A2, A3, R any](tf fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]], exec ...fp.Executor) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Future[R]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Future[R]]] {
		return Flap2(Ap(tf, Successful(a1)), exec...)
	}
}

func Method3[A1, A2, A3, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3) R, exec ...fp.Executor) fp.Func2[A2, A3, fp.Future[R]] {
	return func(a2 A2, a3 A3) fp.Future[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3)
		}, exec...)
	}
}

func FlatMethod3[A1, A2, A3, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3) fp.Future[R], exec ...fp.Executor) fp.Func2[A2, A3, fp.Future[R]] {
	return func(a2 A2, a3 A3) fp.Future[R] {
		return FlatMap(ta1, func(a1 A1) fp.Future[R] {
			return fa1(a1, a2, a3)
		}, exec...)
	}
}

func LiftA4[A1, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) R, exec ...fp.Executor) fp.Func4[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftA3(func(a2 A2, a3 A3, a4 A4) R {
				return f(a1, a2, a3, a4)
			}, exec...)(ins2, ins3, ins4)
		}, exec...)
	}
}

func LiftM4[A1, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Future[R], exec ...fp.Executor) fp.Func4[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftM3(func(a2 A2, a3 A3, a4 A4) fp.Future[R] {
				return f(a1, a2, a3, a4)
			}, exec...)(ins2, ins3, ins4)
		}, exec...)
	}
}

func Flap4[A1, A2, A3, A4, R any](tf fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]], exec ...fp.Executor) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Future[R]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Future[R]]]] {
		return Flap3(Ap(tf, Successful(a1)), exec...)
	}
}

func Method4[A1, A2, A3, A4, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4) R, exec ...fp.Executor) fp.Func3[A2, A3, A4, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4) fp.Future[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4)
		}, exec...)
	}
}

func FlatMethod4[A1, A2, A3, A4, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Future[R], exec ...fp.Executor) fp.Func3[A2, A3, A4, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4) fp.Future[R] {
		return FlatMap(ta1, func(a1 A1) fp.Future[R] {
			return fa1(a1, a2, a3, a4)
		}, exec...)
	}
}

func LiftA5[A1, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R, exec ...fp.Executor) fp.Func5[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftA4(func(a2 A2, a3 A3, a4 A4, a5 A5) R {
				return f(a1, a2, a3, a4, a5)
			}, exec...)(ins2, ins3, ins4, ins5)
		}, exec...)
	}
}

func LiftM5[A1, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[R], exec ...fp.Executor) fp.Func5[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftM4(func(a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[R] {
				return f(a1, a2, a3, a4, a5)
			}, exec...)(ins2, ins3, ins4, ins5)
		}, exec...)
	}
}

func Flap5[A1, A2, A3, A4, A5, R any](tf fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]], exec ...fp.Executor) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Future[R]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Future[R]]]]] {
		return Flap4(Ap(tf, Successful(a1)), exec...)
	}
}

func Method5[A1, A2, A3, A4, A5, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R, exec ...fp.Executor) fp.Func4[A2, A3, A4, A5, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5)
		}, exec...)
	}
}

func FlatMethod5[A1, A2, A3, A4, A5, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[R], exec ...fp.Executor) fp.Func4[A2, A3, A4, A5, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[R] {
		return FlatMap(ta1, func(a1 A1) fp.Future[R] {
			return fa1(a1, a2, a3, a4, a5)
		}, exec...)
	}
}

func LiftA6[A1, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R, exec ...fp.Executor) fp.Func6[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[A6], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5], ins6 fp.Future[A6]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftA5(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
				return f(a1, a2, a3, a4, a5, a6)
			}, exec...)(ins2, ins3, ins4, ins5, ins6)
		}, exec...)
	}
}

func LiftM6[A1, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[R], exec ...fp.Executor) fp.Func6[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[A6], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5], ins6 fp.Future[A6]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftM5(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[R] {
				return f(a1, a2, a3, a4, a5, a6)
			}, exec...)(ins2, ins3, ins4, ins5, ins6)
		}, exec...)
	}
}

func Flap6[A1, A2, A3, A4, A5, A6, R any](tf fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]], exec ...fp.Executor) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Future[R]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Future[R]]]]]] {
		return Flap5(Ap(tf, Successful(a1)), exec...)
	}
}

func Method6[A1, A2, A3, A4, A5, A6, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R, exec ...fp.Executor) fp.Func5[A2, A3, A4, A5, A6, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6)
		}, exec...)
	}
}

func FlatMethod6[A1, A2, A3, A4, A5, A6, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[R], exec ...fp.Executor) fp.Func5[A2, A3, A4, A5, A6, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[R] {
		return FlatMap(ta1, func(a1 A1) fp.Future[R] {
			return fa1(a1, a2, a3, a4, a5, a6)
		}, exec...)
	}
}

func LiftA7[A1, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R, exec ...fp.Executor) fp.Func7[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[A6], fp.Future[A7], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5], ins6 fp.Future[A6], ins7 fp.Future[A7]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftA6(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
				return f(a1, a2, a3, a4, a5, a6, a7)
			}, exec...)(ins2, ins3, ins4, ins5, ins6, ins7)
		}, exec...)
	}
}

func LiftM7[A1, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[R], exec ...fp.Executor) fp.Func7[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[A6], fp.Future[A7], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5], ins6 fp.Future[A6], ins7 fp.Future[A7]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftM6(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[R] {
				return f(a1, a2, a3, a4, a5, a6, a7)
			}, exec...)(ins2, ins3, ins4, ins5, ins6, ins7)
		}, exec...)
	}
}

func Flap7[A1, A2, A3, A4, A5, A6, A7, R any](tf fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]], exec ...fp.Executor) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Future[R]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Future[R]]]]]]] {
		return Flap6(Ap(tf, Successful(a1)), exec...)
	}
}

func Method7[A1, A2, A3, A4, A5, A6, A7, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R, exec ...fp.Executor) fp.Func6[A2, A3, A4, A5, A6, A7, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6, a7)
		}, exec...)
	}
}

func FlatMethod7[A1, A2, A3, A4, A5, A6, A7, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[R], exec ...fp.Executor) fp.Func6[A2, A3, A4, A5, A6, A7, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[R] {
		return FlatMap(ta1, func(a1 A1) fp.Future[R] {
			return fa1(a1, a2, a3, a4, a5, a6, a7)
		}, exec...)
	}
}

func LiftA8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R, exec ...fp.Executor) fp.Func8[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[A6], fp.Future[A7], fp.Future[A8], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5], ins6 fp.Future[A6], ins7 fp.Future[A7], ins8 fp.Future[A8]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftA7(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
				return f(a1, a2, a3, a4, a5, a6, a7, a8)
			}, exec...)(ins2, ins3, ins4, ins5, ins6, ins7, ins8)
		}, exec...)
	}
}

func LiftM8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[R], exec ...fp.Executor) fp.Func8[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[A6], fp.Future[A7], fp.Future[A8], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5], ins6 fp.Future[A6], ins7 fp.Future[A7], ins8 fp.Future[A8]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftM7(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[R] {
				return f(a1, a2, a3, a4, a5, a6, a7, a8)
			}, exec...)(ins2, ins3, ins4, ins5, ins6, ins7, ins8)
		}, exec...)
	}
}

func Flap8[A1, A2, A3, A4, A5, A6, A7, A8, R any](tf fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]], exec ...fp.Executor) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Future[R]]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Future[R]]]]]]]] {
		return Flap7(Ap(tf, Successful(a1)), exec...)
	}
}

func Method8[A1, A2, A3, A4, A5, A6, A7, A8, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R, exec ...fp.Executor) fp.Func7[A2, A3, A4, A5, A6, A7, A8, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8)
		}, exec...)
	}
}

func FlatMethod8[A1, A2, A3, A4, A5, A6, A7, A8, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[R], exec ...fp.Executor) fp.Func7[A2, A3, A4, A5, A6, A7, A8, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[R] {
		return FlatMap(ta1, func(a1 A1) fp.Future[R] {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8)
		}, exec...)
	}
}

func LiftA9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R, exec ...fp.Executor) fp.Func9[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[A6], fp.Future[A7], fp.Future[A8], fp.Future[A9], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5], ins6 fp.Future[A6], ins7 fp.Future[A7], ins8 fp.Future[A8], ins9 fp.Future[A9]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftA8(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
				return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			}, exec...)(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)
		}, exec...)
	}
}

func LiftM9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[R], exec ...fp.Executor) fp.Func9[fp.Future[A1], fp.Future[A2], fp.Future[A3], fp.Future[A4], fp.Future[A5], fp.Future[A6], fp.Future[A7], fp.Future[A8], fp.Future[A9], fp.Future[R]] {
	return func(ins1 fp.Future[A1], ins2 fp.Future[A2], ins3 fp.Future[A3], ins4 fp.Future[A4], ins5 fp.Future[A5], ins6 fp.Future[A6], ins7 fp.Future[A7], ins8 fp.Future[A8], ins9 fp.Future[A9]) fp.Future[R] {

		return FlatMap(ins1, func(a1 A1) fp.Future[R] {
			return LiftM8(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[R] {
				return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			}, exec...)(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)
		}, exec...)
	}
}

func Flap9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](tf fp.Future[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]], exec ...fp.Executor) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Future[R]]]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Future[R]]]]]]]]] {
		return Flap8(Ap(tf, Successful(a1)), exec...)
	}
}

func Method9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R, exec ...fp.Executor) fp.Func8[A2, A3, A4, A5, A6, A7, A8, A9, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}, exec...)
	}
}

func FlatMethod9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](ta1 fp.Future[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[R], exec ...fp.Executor) fp.Func8[A2, A3, A4, A5, A6, A7, A8, A9, fp.Future[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[R] {
		return FlatMap(ta1, func(a1 A1) fp.Future[R] {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}, exec...)
	}
}

func Func1[A1, R any](f func(A1) (R, error), exec ...fp.Executor) fp.Func1[A1, fp.Future[R]] {
	return func(a1 A1) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1)
		})
	}
}

func Unit1[A1 any](f func(A1) error, exec ...fp.Executor) fp.Func1[A1, fp.Future[fp.Unit]] {
	return func(a1 A1) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1)
			return fp.Unit{}, err
		})
	}
}

func Func2[A1, A2, R any](f func(A1, A2) (R, error), exec ...fp.Executor) fp.Func2[A1, A2, fp.Future[R]] {
	return func(a1 A1, a2 A2) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2)
		})
	}
}

func Unit2[A1, A2 any](f func(A1, A2) error, exec ...fp.Executor) fp.Func2[A1, A2, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2)
			return fp.Unit{}, err
		})
	}
}

func Func3[A1, A2, A3, R any](f func(A1, A2, A3) (R, error), exec ...fp.Executor) fp.Func3[A1, A2, A3, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3)
		})
	}
}

func Unit3[A1, A2, A3 any](f func(A1, A2, A3) error, exec ...fp.Executor) fp.Func3[A1, A2, A3, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3)
			return fp.Unit{}, err
		})
	}
}

func Func4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) (R, error), exec ...fp.Executor) fp.Func4[A1, A2, A3, A4, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4)
		})
	}
}

func Unit4[A1, A2, A3, A4 any](f func(A1, A2, A3, A4) error, exec ...fp.Executor) fp.Func4[A1, A2, A3, A4, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4)
			return fp.Unit{}, err
		})
	}
}

func Func5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) (R, error), exec ...fp.Executor) fp.Func5[A1, A2, A3, A4, A5, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5)
		})
	}
}

func Unit5[A1, A2, A3, A4, A5 any](f func(A1, A2, A3, A4, A5) error, exec ...fp.Executor) fp.Func5[A1, A2, A3, A4, A5, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5)
			return fp.Unit{}, err
		})
	}
}

func Func6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) (R, error), exec ...fp.Executor) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6)
		})
	}
}

func Unit6[A1, A2, A3, A4, A5, A6 any](f func(A1, A2, A3, A4, A5, A6) error, exec ...fp.Executor) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5, a6)
			return fp.Unit{}, err
		})
	}
}

func Func7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) (R, error), exec ...fp.Executor) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7)
		})
	}
}

func Unit7[A1, A2, A3, A4, A5, A6, A7 any](f func(A1, A2, A3, A4, A5, A6, A7) error, exec ...fp.Executor) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5, a6, a7)
			return fp.Unit{}, err
		})
	}
}

func Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R, error), exec ...fp.Executor) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8)
		})
	}
}

func Unit8[A1, A2, A3, A4, A5, A6, A7, A8 any](f func(A1, A2, A3, A4, A5, A6, A7, A8) error, exec ...fp.Executor) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5, a6, a7, a8)
			return fp.Unit{}, err
		})
	}
}

func Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R, error), exec ...fp.Executor) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Future[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[R] {
		return Apply2(func() (R, error) {
			return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		})
	}
}

func Unit9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) error, exec ...fp.Executor) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Future[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Future[fp.Unit] {
		return Apply2(func() (fp.Unit, error) {
			err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			return fp.Unit{}, err
		})
	}
}

func Compose3[A1, A2, A3, R any](f1 fp.Func1[A1, fp.Future[A2]], f2 fp.Func1[A2, fp.Future[A3]], f3 fp.Func1[A3, fp.Future[R]], exec ...fp.Executor) fp.Func1[A1, fp.Future[R]] {
	return Compose2(f1, Compose2(f2, f3, exec...), exec...)
}

func Compose4[A1, A2, A3, A4, R any](f1 fp.Func1[A1, fp.Future[A2]], f2 fp.Func1[A2, fp.Future[A3]], f3 fp.Func1[A3, fp.Future[A4]], f4 fp.Func1[A4, fp.Future[R]], exec ...fp.Executor) fp.Func1[A1, fp.Future[R]] {
	return Compose2(f1, Compose3(f2, f3, f4, exec...), exec...)
}

func Compose5[A1, A2, A3, A4, A5, R any](f1 fp.Func1[A1, fp.Future[A2]], f2 fp.Func1[A2, fp.Future[A3]], f3 fp.Func1[A3, fp.Future[A4]], f4 fp.Func1[A4, fp.Future[A5]], f5 fp.Func1[A5, fp.Future[R]], exec ...fp.Executor) fp.Func1[A1, fp.Future[R]] {
	return Compose2(f1, Compose4(f2, f3, f4, f5, exec...), exec...)
}
