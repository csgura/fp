package fp

import (
	"fmt"
	"github.com/csgura/fp/hlist"
)

type Tuple2[T1, T2 any] struct {
	I1 T1
	I2 T2
}

func (r Tuple2[T1, T2]) Head() T1 {
	return r.I1
}

func (r Tuple2[T1, T2]) Tail() Tuple1[T2] {
	return Tuple1[T2]{r.I2}
}

func (r Tuple2[T1, T2]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Nil]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple2[T1, T2]) String() string {
	return fmt.Sprintf("(%v,%v)", r.I1, r.I2)
}

func (r Tuple2[T1, T2]) Unapply() (T1, T2) {
	return r.I1, r.I2
}

type Tuple3[T1, T2, T3 any] struct {
	I1 T1
	I2 T2
	I3 T3
}

func (r Tuple3[T1, T2, T3]) Head() T1 {
	return r.I1
}

func (r Tuple3[T1, T2, T3]) Tail() Tuple2[T2, T3] {
	return Tuple2[T2, T3]{r.I2, r.I3}
}

func (r Tuple3[T1, T2, T3]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Nil]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple3[T1, T2, T3]) String() string {
	return fmt.Sprintf("(%v,%v,%v)", r.I1, r.I2, r.I3)
}

func (r Tuple3[T1, T2, T3]) Unapply() (T1, T2, T3) {
	return r.I1, r.I2, r.I3
}

type Tuple4[T1, T2, T3, T4 any] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
}

func (r Tuple4[T1, T2, T3, T4]) Head() T1 {
	return r.I1
}

func (r Tuple4[T1, T2, T3, T4]) Tail() Tuple3[T2, T3, T4] {
	return Tuple3[T2, T3, T4]{r.I2, r.I3, r.I4}
}

func (r Tuple4[T1, T2, T3, T4]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Nil]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple4[T1, T2, T3, T4]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4)
}

func (r Tuple4[T1, T2, T3, T4]) Unapply() (T1, T2, T3, T4) {
	return r.I1, r.I2, r.I3, r.I4
}

type Tuple5[T1, T2, T3, T4, T5 any] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
}

func (r Tuple5[T1, T2, T3, T4, T5]) Head() T1 {
	return r.I1
}

func (r Tuple5[T1, T2, T3, T4, T5]) Tail() Tuple4[T2, T3, T4, T5] {
	return Tuple4[T2, T3, T4, T5]{r.I2, r.I3, r.I4, r.I5}
}

func (r Tuple5[T1, T2, T3, T4, T5]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Nil]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple5[T1, T2, T3, T4, T5]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5)
}

func (r Tuple5[T1, T2, T3, T4, T5]) Unapply() (T1, T2, T3, T4, T5) {
	return r.I1, r.I2, r.I3, r.I4, r.I5
}
