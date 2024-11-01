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
