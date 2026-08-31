package gendebug

//go:generate go run github.com/csgura/fp/cmd/gombok

type serverSetter struct {
	message string
}

func (r serverSetter) Pre(a string) serverSetter {

	return r
}

// @fp.WithPubField
type genericSetter[REQ, RES any] struct {
	serverSetter
}

// @fp.WithPubField
type validatedSetter[REQ, RES any] struct {
	genericSetter[REQ, RES]
}
