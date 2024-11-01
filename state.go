package fp

// haskell 에서 StateT 타입은   s -> m ( a,s ) 타입이지만
// s 까지  m 안에 있는 것은 불편할 때가 더 많기 때문에
// ( 에러 발생시  에러에  대한 것을 s안에 기록하기 힘듬 )
// s -> (m a , s )  형태로 정의
type StateT[S, A any] func(S) (Try[A], S)

func (r StateT[S, A]) Run(s S) (Try[A], S) {
	return r(s)
}

func (r StateT[S, A]) Exec(s S) Try[S] {
	res, state := r.Run(s)
	if res.IsSuccess() {
		return Success(state)
	}
	return Failure[S](res.Failed().Get())
}

func (r StateT[S, A]) Eval(s S) Try[A] {
	result, _ := r.Run(s)
	return result
}

func (r StateT[S, A]) Recover(f func(err error) A) StateT[S, A] {
	return func(s S) (Try[A], S) {
		at, ns := r.Run(s)
		if at.IsFailure() {
			ra := f(at.Failed().Get())
			return Success(ra), ns
		}
		return at, ns
	}
}

func (r StateT[S, A]) RecoverT(f func(err error) Try[A]) StateT[S, A] {
	return func(s S) (Try[A], S) {
		at, ns := r.Run(s)
		if at.IsFailure() {
			rt := f(at.Failed().Get())
			return rt, ns
		}
		return at, ns
	}
}

func (r StateT[S, A]) RecoverWithState(f func(s S, err error) A) StateT[S, A] {
	return func(s S) (Try[A], S) {
		at, ns := r.Run(s)
		if at.IsFailure() {
			ra := f(ns, at.Failed().Get())
			return Success(ra), ns
		}
		return at, ns
	}
}

func (r StateT[S, A]) RecoverWithStateT(f func(s S, err error) Try[A]) StateT[S, A] {
	return func(s S) (Try[A], S) {
		at, ns := r.Run(s)
		if at.IsFailure() {
			rt := f(s, at.Failed().Get())
			return rt, ns
		}
		return at, ns
	}
}

func (r StateT[S, A]) RecoverWith(f func(err error) StateT[S, A]) StateT[S, A] {
	return func(s S) (Try[A], S) {
		at, ns := r.Run(s)
		if at.IsFailure() {
			rat, nns := f(at.Failed().Get()).Run(ns)
			return rat, nns
		}
		return at, ns
	}
}

func (r StateT[S, A]) RecoverCase(isDefinedAt func(error) bool, then func(error) A) StateT[S, A] {
	return func(s S) (Try[A], S) {
		at, ns := r.Run(s)

		if at.IsSuccess() {
			return at, ns
		}

		if isDefinedAt(at.Failed().Get()) {
			return Success(then(at.Failed().Get())), ns
		}

		return at, ns
	}
}

func (r StateT[S, A]) RecoverCaseT(isDefinedAt func(error) bool, then func(error) Try[A]) StateT[S, A] {
	return func(s S) (Try[A], S) {
		at, ns := r.Run(s)

		if at.IsSuccess() {
			return at, ns
		}

		if isDefinedAt(at.Failed().Get()) {
			return then(at.Failed().Get()), ns
		}

		return at, ns
	}
}

func (r StateT[S, A]) RecoverCaseWith(isDefinedAt func(error) bool, then func(error) StateT[S, A]) StateT[S, A] {
	return func(s S) (Try[A], S) {
		at, ns := r.Run(s)

		if at.IsSuccess() {
			return at, ns
		}

		if isDefinedAt(at.Failed().Get()) {
			return then(at.Failed().Get())(ns)
		}

		return at, ns
	}
}
