package fn1

// first :: a b c -> a (b, d) (c, d)
func First[D, B, C any](f func(B) C) func(B, D) (C, D) {
	return func(b B, d D) (C, D) {
		return f(b), d
	}
}

// second :: a b c -> a (d, b) (d, c)
func Second[D, B, C any](f func(B) C) func(D, B) (D, C) {
	return func(d D, b B) (D, C) {
		return d, f(b)
	}
}

// (***) :: a b c -> a b' c' -> a (b, b') (c, c')
// haskell 에서는 기호외에 따로 이름은 없는데
// scala cats 에서 split 이라는 이름으로 정의함
func Split[A, B, C, D any](f1 func(A) B, f2 func(C) D) func(A, C) (B, D) {
	return func(a A, c C) (B, D) {
		return f1(a), f2(c)
	}
}

// (&&&) :: a b c -> a b c' -> a b (c, c')
// haskell 에서는 기호외에 따로 이름은 없는데
// scala cats 에서 merge 라는 이름으로 정의함.
func Merge[A, B, C any](f1 func(A) B, f2 func(A) C) func(A) (B, C) {
	return func(a A) (B, C) {
		return f1(a), f2(a)
	}
}
