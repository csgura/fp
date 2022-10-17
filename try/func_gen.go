package try

import (
	"github.com/csgura/fp"
)

func LiftA3[A1, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) R) fp.Func3[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftA2(func(a2 A2, a3 A3) R {
				return f(a1, a2, a3)
			})(ins2, ins3)
		})
	}
}

func LiftM3[A1, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) fp.Try[R]) fp.Func3[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftM2(func(a2 A2, a3 A3) fp.Try[R] {
				return f(a1, a2, a3)
			})(ins2, ins3)
		})
	}
}

func Flap3[A1, A2, A3, R any](tf fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Try[R]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Try[R]]] {
		return Flap2(Ap(tf, Success(a1)))
	}
}

func Method3[A1, A2, A3, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3) R) fp.Func2[A2, A3, fp.Try[R]] {
	return func(a2 A2, a3 A3) fp.Try[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3)
		})
	}
}

func FlatMethod3[A1, A2, A3, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3) fp.Try[R]) fp.Func2[A2, A3, fp.Try[R]] {
	return func(a2 A2, a3 A3) fp.Try[R] {
		return FlatMap(ta1, func(a1 A1) fp.Try[R] {
			return fa1(a1, a2, a3)
		})
	}
}

func LiftA4[A1, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) R) fp.Func4[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftA3(func(a2 A2, a3 A3, a4 A4) R {
				return f(a1, a2, a3, a4)
			})(ins2, ins3, ins4)
		})
	}
}

func LiftM4[A1, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R]) fp.Func4[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftM3(func(a2 A2, a3 A3, a4 A4) fp.Try[R] {
				return f(a1, a2, a3, a4)
			})(ins2, ins3, ins4)
		})
	}
}

func Flap4[A1, A2, A3, A4, R any](tf fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Try[R]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Try[R]]]] {
		return Flap3(Ap(tf, Success(a1)))
	}
}

func Method4[A1, A2, A3, A4, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4) R) fp.Func3[A2, A3, A4, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4) fp.Try[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4)
		})
	}
}

func FlatMethod4[A1, A2, A3, A4, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R]) fp.Func3[A2, A3, A4, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4) fp.Try[R] {
		return FlatMap(ta1, func(a1 A1) fp.Try[R] {
			return fa1(a1, a2, a3, a4)
		})
	}
}

func LiftA5[A1, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) fp.Func5[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftA4(func(a2 A2, a3 A3, a4 A4, a5 A5) R {
				return f(a1, a2, a3, a4, a5)
			})(ins2, ins3, ins4, ins5)
		})
	}
}

func LiftM5[A1, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R]) fp.Func5[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftM4(func(a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
				return f(a1, a2, a3, a4, a5)
			})(ins2, ins3, ins4, ins5)
		})
	}
}

func Flap5[A1, A2, A3, A4, A5, R any](tf fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Try[R]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Try[R]]]]] {
		return Flap4(Ap(tf, Success(a1)))
	}
}

func Method5[A1, A2, A3, A4, A5, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) fp.Func4[A2, A3, A4, A5, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5)
		})
	}
}

func FlatMethod5[A1, A2, A3, A4, A5, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R]) fp.Func4[A2, A3, A4, A5, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
		return FlatMap(ta1, func(a1 A1) fp.Try[R] {
			return fa1(a1, a2, a3, a4, a5)
		})
	}
}

func LiftA6[A1, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) fp.Func6[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[A6], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftA5(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
				return f(a1, a2, a3, a4, a5, a6)
			})(ins2, ins3, ins4, ins5, ins6)
		})
	}
}

func LiftM6[A1, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R]) fp.Func6[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[A6], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftM5(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
				return f(a1, a2, a3, a4, a5, a6)
			})(ins2, ins3, ins4, ins5, ins6)
		})
	}
}

func Flap6[A1, A2, A3, A4, A5, A6, R any](tf fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Try[R]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Try[R]]]]]] {
		return Flap5(Ap(tf, Success(a1)))
	}
}

func Method6[A1, A2, A3, A4, A5, A6, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) fp.Func5[A2, A3, A4, A5, A6, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6)
		})
	}
}

func FlatMethod6[A1, A2, A3, A4, A5, A6, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R]) fp.Func5[A2, A3, A4, A5, A6, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
		return FlatMap(ta1, func(a1 A1) fp.Try[R] {
			return fa1(a1, a2, a3, a4, a5, a6)
		})
	}
}

func LiftA7[A1, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) fp.Func7[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[A6], fp.Try[A7], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6], ins7 fp.Try[A7]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftA6(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
				return f(a1, a2, a3, a4, a5, a6, a7)
			})(ins2, ins3, ins4, ins5, ins6, ins7)
		})
	}
}

func LiftM7[A1, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R]) fp.Func7[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[A6], fp.Try[A7], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6], ins7 fp.Try[A7]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftM6(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
				return f(a1, a2, a3, a4, a5, a6, a7)
			})(ins2, ins3, ins4, ins5, ins6, ins7)
		})
	}
}

func Flap7[A1, A2, A3, A4, A5, A6, A7, R any](tf fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Try[R]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Try[R]]]]]]] {
		return Flap6(Ap(tf, Success(a1)))
	}
}

func Method7[A1, A2, A3, A4, A5, A6, A7, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) fp.Func6[A2, A3, A4, A5, A6, A7, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6, a7)
		})
	}
}

func FlatMethod7[A1, A2, A3, A4, A5, A6, A7, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R]) fp.Func6[A2, A3, A4, A5, A6, A7, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
		return FlatMap(ta1, func(a1 A1) fp.Try[R] {
			return fa1(a1, a2, a3, a4, a5, a6, a7)
		})
	}
}

func LiftA8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) fp.Func8[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[A6], fp.Try[A7], fp.Try[A8], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6], ins7 fp.Try[A7], ins8 fp.Try[A8]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftA7(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
				return f(a1, a2, a3, a4, a5, a6, a7, a8)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8)
		})
	}
}

func LiftM8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R]) fp.Func8[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[A6], fp.Try[A7], fp.Try[A8], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6], ins7 fp.Try[A7], ins8 fp.Try[A8]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftM7(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
				return f(a1, a2, a3, a4, a5, a6, a7, a8)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8)
		})
	}
}

func Flap8[A1, A2, A3, A4, A5, A6, A7, A8, R any](tf fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Try[R]]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Try[R]]]]]]]] {
		return Flap7(Ap(tf, Success(a1)))
	}
}

func Method8[A1, A2, A3, A4, A5, A6, A7, A8, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) fp.Func7[A2, A3, A4, A5, A6, A7, A8, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8)
		})
	}
}

func FlatMethod8[A1, A2, A3, A4, A5, A6, A7, A8, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R]) fp.Func7[A2, A3, A4, A5, A6, A7, A8, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
		return FlatMap(ta1, func(a1 A1) fp.Try[R] {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8)
		})
	}
}

func LiftA9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) fp.Func9[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[A6], fp.Try[A7], fp.Try[A8], fp.Try[A9], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6], ins7 fp.Try[A7], ins8 fp.Try[A8], ins9 fp.Try[A9]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftA8(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
				return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)
		})
	}
}

func LiftM9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R]) fp.Func9[fp.Try[A1], fp.Try[A2], fp.Try[A3], fp.Try[A4], fp.Try[A5], fp.Try[A6], fp.Try[A7], fp.Try[A8], fp.Try[A9], fp.Try[R]] {
	return func(ins1 fp.Try[A1], ins2 fp.Try[A2], ins3 fp.Try[A3], ins4 fp.Try[A4], ins5 fp.Try[A5], ins6 fp.Try[A6], ins7 fp.Try[A7], ins8 fp.Try[A8], ins9 fp.Try[A9]) fp.Try[R] {

		return FlatMap(ins1, func(a1 A1) fp.Try[R] {
			return LiftM8(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
				return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)
		})
	}
}

func Flap9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](tf fp.Try[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Try[R]]]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Try[R]]]]]]]]] {
		return Flap8(Ap(tf, Success(a1)))
	}
}

func Method9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) fp.Func8[A2, A3, A4, A5, A6, A7, A8, A9, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		})
	}
}

func FlatMethod9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](ta1 fp.Try[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R]) fp.Func8[A2, A3, A4, A5, A6, A7, A8, A9, fp.Try[R]] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
		return FlatMap(ta1, func(a1 A1) fp.Try[R] {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		})
	}
}

func Func1[A1, R any](f func(A1) (R, error)) fp.Func1[A1, fp.Try[R]] {
	return func(a1 A1) fp.Try[R] {
		ret, err := f(a1)
		return Apply(ret, err)
	}
}

func Unit1[A1 any](f func(A1) error) fp.Func1[A1, fp.Try[fp.Unit]] {
	return func(a1 A1) fp.Try[fp.Unit] {
		err := f(a1)
		return Apply(fp.Unit{}, err)
	}
}

func Func2[A1, A2, R any](f func(A1, A2) (R, error)) fp.Func2[A1, A2, fp.Try[R]] {
	return func(a1 A1, a2 A2) fp.Try[R] {
		ret, err := f(a1, a2)
		return Apply(ret, err)
	}
}

func Unit2[A1, A2 any](f func(A1, A2) error) fp.Func2[A1, A2, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2) fp.Try[fp.Unit] {
		err := f(a1, a2)
		return Apply(fp.Unit{}, err)
	}
}

func Func3[A1, A2, A3, R any](f func(A1, A2, A3) (R, error)) fp.Func3[A1, A2, A3, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Try[R] {
		ret, err := f(a1, a2, a3)
		return Apply(ret, err)
	}
}

func Unit3[A1, A2, A3 any](f func(A1, A2, A3) error) fp.Func3[A1, A2, A3, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3) fp.Try[fp.Unit] {
		err := f(a1, a2, a3)
		return Apply(fp.Unit{}, err)
	}
}

func Func4[A1, A2, A3, A4, R any](f func(A1, A2, A3, A4) (R, error)) fp.Func4[A1, A2, A3, A4, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4)
		return Apply(ret, err)
	}
}

func Unit4[A1, A2, A3, A4 any](f func(A1, A2, A3, A4) error) fp.Func4[A1, A2, A3, A4, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4)
		return Apply(fp.Unit{}, err)
	}
}

func Func5[A1, A2, A3, A4, A5, R any](f func(A1, A2, A3, A4, A5) (R, error)) fp.Func5[A1, A2, A3, A4, A5, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5)
		return Apply(ret, err)
	}
}

func Unit5[A1, A2, A3, A4, A5 any](f func(A1, A2, A3, A4, A5) error) fp.Func5[A1, A2, A3, A4, A5, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5)
		return Apply(fp.Unit{}, err)
	}
}

func Func6[A1, A2, A3, A4, A5, A6, R any](f func(A1, A2, A3, A4, A5, A6) (R, error)) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6)
		return Apply(ret, err)
	}
}

func Unit6[A1, A2, A3, A4, A5, A6 any](f func(A1, A2, A3, A4, A5, A6) error) fp.Func6[A1, A2, A3, A4, A5, A6, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5, a6)
		return Apply(fp.Unit{}, err)
	}
}

func Func7[A1, A2, A3, A4, A5, A6, A7, R any](f func(A1, A2, A3, A4, A5, A6, A7) (R, error)) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6, a7)
		return Apply(ret, err)
	}
}

func Unit7[A1, A2, A3, A4, A5, A6, A7 any](f func(A1, A2, A3, A4, A5, A6, A7) error) fp.Func7[A1, A2, A3, A4, A5, A6, A7, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5, a6, a7)
		return Apply(fp.Unit{}, err)
	}
}

func Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R, error)) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
		return Apply(ret, err)
	}
}

func Unit8[A1, A2, A3, A4, A5, A6, A7, A8 any](f func(A1, A2, A3, A4, A5, A6, A7, A8) error) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5, a6, a7, a8)
		return Apply(fp.Unit{}, err)
	}
}

func Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R, error)) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Try[R]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[R] {
		ret, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		return Apply(ret, err)
	}
}

func Unit9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) error) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, fp.Try[fp.Unit]] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) fp.Try[fp.Unit] {
		err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		return Apply(fp.Unit{}, err)
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
