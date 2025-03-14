package gendebug

import "github.com/csgura/fp/mshow"

//go:generate go run github.com/csgura/fp/cmd/gombok

type TLV16 struct {
	Value []byte
}

type PayloadContainer = TLV16

type Container struct {
	TLV *PayloadContainer
}

// @fp.Derive(recursive=true)
var _ mshow.Derives[mshow.Show[Container]]
