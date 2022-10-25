package value

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/eq"
)

var EqPerson = eq.ContraMap(eq.Tuple8(eq.String, eq.Given[int](), EqFloat64, eq.Option(eq.String), eq.Slice(eq.String), eq.HCons(eq.String, eq.HCons(eq.Given[int](), eq.HNil)), EqFpSeq(EqFloat64), eq.Bytes), Person.AsTuple)

var EqWallet = eq.ContraMap(eq.Tuple2(EqPerson, eq.Given[int64]()), Wallet.AsTuple)

func EqEntry[A any, B any](eqA fp.Eq[A], eqB fp.Eq[B]) fp.Eq[Entry[A, B]] {
	return eq.ContraMap(eq.Tuple3(eq.String, eqA, eq.Tuple2(eqA, eqB)), Entry[A, B].AsTuple)
}
