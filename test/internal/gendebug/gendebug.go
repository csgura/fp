package gendebug

import (
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/mshow"
)

//go:generate go run github.com/csgura/fp/cmd/gombok

type ProtocolIESingleContainerGlobalRANNodeIDExtIEs struct {
}

type HasTuple struct {
	HList hlist.Cons[string, hlist.Nil]
}

// @fp.Derive
var _ mshow.Derives[mshow.Show[ProtocolIESingleContainerGlobalRANNodeIDExtIEs]]
