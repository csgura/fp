package assert

func True(b bool) {
	if !b {
		panic("assert fail")
	}
}

func False(b bool) {
	if b {
		panic("assert fail")
	}
}

func Equal[T comparable](a, b T) {
	if a != b {
		panic("assert fail")
	}
}
