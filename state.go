package fp

type State[S, A any] func(S) Tuple2[A, S]

func (r State[S, A]) Run(s S) (A, S) {
	return r(s).Unapply()
}

func (r State[S, A]) Exec(s S) S {
	return r(s).Tail()
}

func (r State[S, A]) Eval(s S) A {
	return r(s).Head()
}

func (r State[S, A]) Widen() func(S) Tuple2[A, S] {
	return r
}

type StateT[S, A any] func(S) Try[Tuple2[A, S]]

func (r StateT[S, A]) Run(s S) (Try[A], Try[S]) {
	ret := r(s)
	if ret.IsSuccess() {
		a, s := ret.Get().Unapply()
		return Success(a), Success(s)
	}
	return Failure[A](ret.Failed().Get()), Failure[S](ret.Failed().Get())
}

func (r StateT[S, A]) Exec(s S) Try[S] {
	_, state := r.Run(s)
	return state
}

func (r StateT[S, A]) Eval(s S) Try[A] {
	result, _ := r.Run(s)
	return result
}

func (r StateT[S, A]) Widen() func(S) Try[Tuple2[A, S]] {
	return r
}
