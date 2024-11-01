package fp

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
