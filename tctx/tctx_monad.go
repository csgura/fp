// Code generated by monad_gen, DO NOT EDIT.
package tctx

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/product"
	"github.com/csgura/fp/xtr"
)

func Flatten[A any](tta State[State[A]]) State[A] {
	return FlatMap(tta, func(v State[A]) State[A] {
		return v
	})
}

func Map[A any, R any](m State[A], f func(A) R) State[R] {
	return FlatMap(m, fp.Compose2(f, Pure[R]))
}

// haskell 의 <$
// map . const 와 같은 함수
func Replace[A any, R any](s State[A], b R) State[R] {
	return Map(s, fp.Const[A](b))
}

func Map2[A any, B, R any](first State[A], second State[B], fab func(A, B) R) State[R] {
	return FlatMap(first, func(a A) State[R] {
		return Map(second, func(b B) R {
			return fab(a, b)
		})
	})
}

func Zip[A any, B any](first State[A], second State[B]) State[fp.Tuple2[A, B]] {
	return Map2(first, second, product.Tuple2)
}

func Ap[A any, B any](tfab State[fp.Func1[A, B]], ta State[A]) State[B] {
	return FlatMap(tfab, func(fab fp.Func1[A, B]) State[B] {
		return Map(ta, fab)
	})
}

func Compose[A any, B, C any](f1 func(A) State[B], f2 func(B) State[C]) func(A) State[C] {
	return func(a A) State[C] {
		return FlatMap(f1(a), f2)
	}
}

func Compose2[A any, B, C any](f1 func(A) State[B], f2 func(B) State[C]) func(A) State[C] {
	return func(a A) State[C] {
		return FlatMap(f1(a), f2)
	}
}

func ApFunc[A any, B any](tfab State[fp.Func1[A, B]], ta func() State[A]) State[B] {
	return FlatMap(tfab, func(fab fp.Func1[A, B]) State[B] {
		return Map(ta(), fab)
	})
}

// Map(ta , seq.Lift(f)) 와 동일
func MapSeqLift[A any, B any](ta State[fp.Seq[A]], f func(v A) B) State[fp.Seq[B]] {

	return Map(ta, func(a fp.Seq[A]) fp.Seq[B] {
		return iterator.Map(iterator.FromSeq(a), f).ToSeq()
	})
}

// Map(ta , seq.Lift(f)) 와 동일
func MapSliceLift[A any, B any](ta State[[]A], f func(v A) B) State[[]B] {

	return Map(ta, func(a []A) []B {
		return iterator.Map(iterator.FromSeq(a), f).ToSeq()
	})
}

func Lift[A any, R any](fa func(v A) R) func(State[A]) State[R] {
	return func(ta State[A]) State[R] {
		return Map(ta, fa)
	}
}

func LiftA2[A any, B, R any](fab func(A, B) R) func(State[A], State[B]) State[R] {
	return func(a State[A], b State[B]) State[R] {
		return Map2(a, b, fab)
	}
}

func LiftM[A any, R any](fa func(v A) State[R]) func(State[A]) State[R] {
	return func(ta State[A]) State[R] {
		return Flatten(Map(ta, fa))
	}
}

// (a -> b -> m r) -> m a -> m b -> m r
// 하스켈에서는  liftM2 와 liftA2 는 같은 함수이고
// 위와 같은 함수는 존재하지 않음.
// hoogle 에서 검색해 보면 , liftJoin2 , bindM2 등의 이름으로 정의된 것이 있음.
// 하지만 ,  fp 패키지에서도   LiftA2 와 LiftM2 를 동일하게 하는 것은 낭비이고
// M 은 Monad 라는 뜻인데, Monad는 Flatten, FlatMap 의 의미가 있으니까
// LiftM2 를 다음과 같이 정의함.
func LiftM2[A any, B, R any](fab func(A, B) State[R]) func(State[A], State[B]) State[R] {
	return func(a State[A], b State[B]) State[R] {
		return Flatten(Map2(a, b, fab))
	}
}

// 하스켈 : m( a -> r ) -> a -> m r
// 스칼라 : M[ A => r ] => A => M[R]
// 하스켈이나 스칼라의 기본 패키지에는 이런 기능을 하는 함수가 없는데,
// hoogle 에서 검색해 보면
// https://hoogle.haskell.org/?hoogle=m%20(%20a%20-%3E%20b)%20-%3E%20a%20-%3E%20m%20b
// ?? 혹은 flap 이라는 이름으로 정의된 함수가 있음
func Flap[A any, R any](tfa State[fp.Func1[A, R]]) func(A) State[R] {
	return func(a A) State[R] {
		return Ap(tfa, Pure(a))
	}
}

// 하스켈 : m( a -> b -> r ) -> a -> b -> m r
func Flap2[A any, B, R any](tfab State[fp.Func1[A, fp.Func1[B, R]]]) fp.Func1[A, fp.Func1[B, State[R]]] {
	return func(a A) fp.Func1[B, State[R]] {
		return Flap(Ap(tfab, Pure(a)))
	}
}

// (a -> b -> r) -> m a -> b -> m r
// Map 호출 후에 Flap 을 호출 한 것
//
// https://hoogle.haskell.org/?hoogle=%28+a+-%3E+b+-%3E++r+%29+-%3E+m+a+-%3E++b+-%3E+m+r+&scope=set%3Astackage
// liftOp 라는 이름으로 정의된 것이 있음
func FlapMap[A any, B, R any](tfab func(A, B) R, a State[A]) func(B) State[R] {
	return Flap(Map(a, curried.Func2(tfab)))
}

// ( a -> b -> m r) -> m a -> b -> m r
//
//	Flatten . FlapMap
//
// https://hoogle.haskell.org/?hoogle=(%20a%20-%3E%20b%20-%3E%20m%20r%20)%20-%3E%20m%20a%20-%3E%20%20b%20-%3E%20m%20r%20
// om , ==<<  이름으로 정의된 것이 있음
func FlatFlapMap[A any, B, R any](fab func(A, B) State[R], ta State[A]) func(B) State[R] {
	return fp.Compose(FlapMap(fab, ta), Flatten[R])
}

// FlatMap 과는 아규먼트 순서가 다른 함수로
// Go 나 Java 에서는 메소드 레퍼런스를 이용하여,  객체내의 메소드를 리턴 타입만 lift 된 형태로 리턴하게 할 수 있음.
// Method 라는 이름보다  Ap 와 비슷한 이름이 좋을 거 같은데
// Ap와 비슷한 이름으로 하기에는 Ap 와 타입이 너무 다름.
func Method1[A any, B, R any](ta State[A], fab func(a A, b B) R) func(B) State[R] {
	return FlapMap(fab, ta)
}

func FlatMethod1[A any, B, R any](ta State[A], fab func(a A, b B) State[R]) func(B) State[R] {
	return FlatFlapMap(fab, ta)
}

func Method2[A any, B, C, R any](ta State[A], fabc func(a A, b B, c C) R) func(B, C) State[R] {

	return curried.Revert2(Flap2(Map(ta, curried.Func3(fabc))))
	// return func(b B, c C) State[R] {
	// 	return Map(a, func(a A) R {
	// 		return cf(a, b, c)
	// 	})
	// }
}

func FlatMethod2[A any, B, C, R any](ta State[A], fabc func(a A, b B, c C) State[R]) func(B, C) State[R] {

	return curried.Revert2(curried.Compose2(Flap2(Map(ta, curried.Func3(fabc))), Flatten[R]))

	// return func(b B, c C) State[R] {
	// 	return FlatMap(ta, func(a A) State[R] {
	// 		return cf(a, b, c)
	// 	})
	// }
}

func UnZip[A any, B any](t State[fp.Tuple2[A, B]]) (State[A], State[B]) {
	return Map(t, xtr.Head), Map(t, xtr.Tail)
}

func Zip3[A any, B, C any](ta State[A], tb State[B], tc State[C]) State[fp.Tuple3[A, B, C]] {
	return LiftA3(product.Tuple3[A, B, C])(ta, tb, tc)
}

// fp.With 의 try 버젼
// fp.With 가 Flip 과 사실상 같은 것처럼
// FlapMap 의 Flip 버젼과 동일
// var b fp.Try[B]
// a := try.Sucesss(A{})
// a.FlatMap( try.With(A.WithB, b))
// 형태로 코딩 가능
func With[A any, B any](withf func(A, B) A, v State[B]) func(A) State[A] {
	return Flap(Map(v, fp.Flip2(withf)))
}

func LiftA3[A1 any, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) R) func(State[A1], State[A2], State[A3]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftA2(func(a2 A2, a3 A3) R {
				return f(a1, a2, a3)
			})(ins2, ins3)
		})
	}
}

func LiftM3[A1 any, A2, A3, R any](f func(a1 A1, a2 A2, a3 A3) State[R]) func(State[A1], State[A2], State[A3]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftM2(func(a2 A2, a3 A3) State[R] {
				return f(a1, a2, a3)
			})(ins2, ins3)
		})
	}
}

func Flap3[A1 any, A2, A3, R any](tf State[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, State[R]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, State[R]]] {
		return Flap2(Ap(tf, Pure(a1)))
	}
}

func Method3[A1 any, A2, A3, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3) R) func(A2, A3) State[R] {
	return func(a2 A2, a3 A3) State[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3)
		})
	}
}

func FlatMethod3[A1 any, A2, A3, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3) State[R]) func(A2, A3) State[R] {
	return func(a2 A2, a3 A3) State[R] {
		return FlatMap(ta1, func(a1 A1) State[R] {
			return fa1(a1, a2, a3)
		})
	}
}

func LiftA4[A1 any, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) R) func(State[A1], State[A2], State[A3], State[A4]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftA3(func(a2 A2, a3 A3, a4 A4) R {
				return f(a1, a2, a3, a4)
			})(ins2, ins3, ins4)
		})
	}
}

func LiftM4[A1 any, A2, A3, A4, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4) State[R]) func(State[A1], State[A2], State[A3], State[A4]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftM3(func(a2 A2, a3 A3, a4 A4) State[R] {
				return f(a1, a2, a3, a4)
			})(ins2, ins3, ins4)
		})
	}
}

func Flap4[A1 any, A2, A3, A4, R any](tf State[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, State[R]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, State[R]]]] {
		return Flap3(Ap(tf, Pure(a1)))
	}
}

func Method4[A1 any, A2, A3, A4, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4) R) func(A2, A3, A4) State[R] {
	return func(a2 A2, a3 A3, a4 A4) State[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4)
		})
	}
}

func FlatMethod4[A1 any, A2, A3, A4, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4) State[R]) func(A2, A3, A4) State[R] {
	return func(a2 A2, a3 A3, a4 A4) State[R] {
		return FlatMap(ta1, func(a1 A1) State[R] {
			return fa1(a1, a2, a3, a4)
		})
	}
}

func LiftA5[A1 any, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) func(State[A1], State[A2], State[A3], State[A4], State[A5]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftA4(func(a2 A2, a3 A3, a4 A4, a5 A5) R {
				return f(a1, a2, a3, a4, a5)
			})(ins2, ins3, ins4, ins5)
		})
	}
}

func LiftM5[A1 any, A2, A3, A4, A5, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) State[R]) func(State[A1], State[A2], State[A3], State[A4], State[A5]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftM4(func(a2 A2, a3 A3, a4 A4, a5 A5) State[R] {
				return f(a1, a2, a3, a4, a5)
			})(ins2, ins3, ins4, ins5)
		})
	}
}

func Flap5[A1 any, A2, A3, A4, A5, R any](tf State[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, State[R]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, State[R]]]]] {
		return Flap4(Ap(tf, Pure(a1)))
	}
}

func Method5[A1 any, A2, A3, A4, A5, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R) func(A2, A3, A4, A5) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) State[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5)
		})
	}
}

func FlatMethod5[A1 any, A2, A3, A4, A5, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) State[R]) func(A2, A3, A4, A5) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) State[R] {
		return FlatMap(ta1, func(a1 A1) State[R] {
			return fa1(a1, a2, a3, a4, a5)
		})
	}
}

func LiftA6[A1 any, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) func(State[A1], State[A2], State[A3], State[A4], State[A5], State[A6]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5], ins6 State[A6]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftA5(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
				return f(a1, a2, a3, a4, a5, a6)
			})(ins2, ins3, ins4, ins5, ins6)
		})
	}
}

func LiftM6[A1 any, A2, A3, A4, A5, A6, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) State[R]) func(State[A1], State[A2], State[A3], State[A4], State[A5], State[A6]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5], ins6 State[A6]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftM5(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) State[R] {
				return f(a1, a2, a3, a4, a5, a6)
			})(ins2, ins3, ins4, ins5, ins6)
		})
	}
}

func Flap6[A1 any, A2, A3, A4, A5, A6, R any](tf State[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, State[R]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, State[R]]]]]] {
		return Flap5(Ap(tf, Pure(a1)))
	}
}

func Method6[A1 any, A2, A3, A4, A5, A6, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R) func(A2, A3, A4, A5, A6) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) State[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6)
		})
	}
}

func FlatMethod6[A1 any, A2, A3, A4, A5, A6, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) State[R]) func(A2, A3, A4, A5, A6) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) State[R] {
		return FlatMap(ta1, func(a1 A1) State[R] {
			return fa1(a1, a2, a3, a4, a5, a6)
		})
	}
}

func LiftA7[A1 any, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) func(State[A1], State[A2], State[A3], State[A4], State[A5], State[A6], State[A7]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5], ins6 State[A6], ins7 State[A7]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftA6(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
				return f(a1, a2, a3, a4, a5, a6, a7)
			})(ins2, ins3, ins4, ins5, ins6, ins7)
		})
	}
}

func LiftM7[A1 any, A2, A3, A4, A5, A6, A7, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) State[R]) func(State[A1], State[A2], State[A3], State[A4], State[A5], State[A6], State[A7]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5], ins6 State[A6], ins7 State[A7]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftM6(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) State[R] {
				return f(a1, a2, a3, a4, a5, a6, a7)
			})(ins2, ins3, ins4, ins5, ins6, ins7)
		})
	}
}

func Flap7[A1 any, A2, A3, A4, A5, A6, A7, R any](tf State[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, State[R]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, State[R]]]]]]] {
		return Flap6(Ap(tf, Pure(a1)))
	}
}

func Method7[A1 any, A2, A3, A4, A5, A6, A7, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R) func(A2, A3, A4, A5, A6, A7) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) State[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6, a7)
		})
	}
}

func FlatMethod7[A1 any, A2, A3, A4, A5, A6, A7, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) State[R]) func(A2, A3, A4, A5, A6, A7) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) State[R] {
		return FlatMap(ta1, func(a1 A1) State[R] {
			return fa1(a1, a2, a3, a4, a5, a6, a7)
		})
	}
}

func LiftA8[A1 any, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) func(State[A1], State[A2], State[A3], State[A4], State[A5], State[A6], State[A7], State[A8]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5], ins6 State[A6], ins7 State[A7], ins8 State[A8]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftA7(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
				return f(a1, a2, a3, a4, a5, a6, a7, a8)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8)
		})
	}
}

func LiftM8[A1 any, A2, A3, A4, A5, A6, A7, A8, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) State[R]) func(State[A1], State[A2], State[A3], State[A4], State[A5], State[A6], State[A7], State[A8]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5], ins6 State[A6], ins7 State[A7], ins8 State[A8]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftM7(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) State[R] {
				return f(a1, a2, a3, a4, a5, a6, a7, a8)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8)
		})
	}
}

func Flap8[A1 any, A2, A3, A4, A5, A6, A7, A8, R any](tf State[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, State[R]]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, State[R]]]]]]]] {
		return Flap7(Ap(tf, Pure(a1)))
	}
}

func Method8[A1 any, A2, A3, A4, A5, A6, A7, A8, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R) func(A2, A3, A4, A5, A6, A7, A8) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) State[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8)
		})
	}
}

func FlatMethod8[A1 any, A2, A3, A4, A5, A6, A7, A8, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) State[R]) func(A2, A3, A4, A5, A6, A7, A8) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) State[R] {
		return FlatMap(ta1, func(a1 A1) State[R] {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8)
		})
	}
}

func LiftA9[A1 any, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) func(State[A1], State[A2], State[A3], State[A4], State[A5], State[A6], State[A7], State[A8], State[A9]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5], ins6 State[A6], ins7 State[A7], ins8 State[A8], ins9 State[A9]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftA8(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
				return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)
		})
	}
}

func LiftM9[A1 any, A2, A3, A4, A5, A6, A7, A8, A9, R any](f func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) State[R]) func(State[A1], State[A2], State[A3], State[A4], State[A5], State[A6], State[A7], State[A8], State[A9]) State[R] {
	return func(ins1 State[A1], ins2 State[A2], ins3 State[A3], ins4 State[A4], ins5 State[A5], ins6 State[A6], ins7 State[A7], ins8 State[A8], ins9 State[A9]) State[R] {

		return FlatMap(ins1, func(a1 A1) State[R] {
			return LiftM8(func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) State[R] {
				return f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			})(ins2, ins3, ins4, ins5, ins6, ins7, ins8, ins9)
		})
	}
}

func Flap9[A1 any, A2, A3, A4, A5, A6, A7, A8, A9, R any](tf State[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]]) fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, State[R]]]]]]]]]] {
	return func(a1 A1) fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, State[R]]]]]]]]] {
		return Flap8(Ap(tf, Pure(a1)))
	}
}

func Method9[A1 any, A2, A3, A4, A5, A6, A7, A8, A9, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R) func(A2, A3, A4, A5, A6, A7, A8, A9) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) State[R] {
		return Map(ta1, func(a1 A1) R {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		})
	}
}

func FlatMethod9[A1 any, A2, A3, A4, A5, A6, A7, A8, A9, R any](ta1 State[A1], fa1 func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) State[R]) func(A2, A3, A4, A5, A6, A7, A8, A9) State[R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) State[R] {
		return FlatMap(ta1, func(a1 A1) State[R] {
			return fa1(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		})
	}
}

func Compose3[A1 any, A2, A3, R any](f1 fp.Func1[A1, State[A2]], f2 fp.Func1[A2, State[A3]], f3 fp.Func1[A3, State[R]]) fp.Func1[A1, State[R]] {
	return Compose2(f1, Compose2(f2, f3))
}

func Compose4[A1 any, A2, A3, A4, R any](f1 fp.Func1[A1, State[A2]], f2 fp.Func1[A2, State[A3]], f3 fp.Func1[A3, State[A4]], f4 fp.Func1[A4, State[R]]) fp.Func1[A1, State[R]] {
	return Compose2(f1, Compose3(f2, f3, f4))
}

func Compose5[A1 any, A2, A3, A4, A5, R any](f1 fp.Func1[A1, State[A2]], f2 fp.Func1[A2, State[A3]], f3 fp.Func1[A3, State[A4]], f4 fp.Func1[A4, State[A5]], f5 fp.Func1[A5, State[R]]) fp.Func1[A1, State[R]] {
	return Compose2(f1, Compose4(f2, f3, f4, f5))
}
