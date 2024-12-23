package findfp

import "github.com/csgura/fp/test/internal/ngap"

type Derives[T any] interface {
	Target() T
}

type NgapType struct {
	Present int
	First   *int     `aper:"id=20"`
	Second  *string  `aper:"id=30"`
	Third   *float64 `aper:"id=40"`
}

//go:generate go run github.com/csgura/fp/cmd/gombok

// @fp.Derive
var _ ngap.Derives[ngap.Split[NgapType]]
