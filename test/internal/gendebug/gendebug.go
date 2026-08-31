package gendebug

//go:generate go run github.com/csgura/fp/cmd/gombok

type serverSetter struct {
	message string
}

type genericSetter[REQ, RES any] struct {
	serverSetter
}

func (r genericSetter[REQ, RES]) WithGeneric[S any](s S) genericSetter[REQ, RES] {
	return r
}

// @fp.WithPubField
type genericStatSetter[REQ, RES, S any] struct {
	genericSetter[REQ, RES]
}
