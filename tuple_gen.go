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

type Tuple6[T1, T2, T3, T4, T5, T6 any] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
	I6 T6
}

func (r Tuple6[T1, T2, T3, T4, T5, T6]) Head() T1 {
	return r.I1
}

func (r Tuple6[T1, T2, T3, T4, T5, T6]) Tail() Tuple5[T2, T3, T4, T5, T6] {
	return Tuple5[T2, T3, T4, T5, T6]{r.I2, r.I3, r.I4, r.I5, r.I6}
}

func (r Tuple6[T1, T2, T3, T4, T5, T6]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Nil]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple6[T1, T2, T3, T4, T5, T6]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6)
}

func (r Tuple6[T1, T2, T3, T4, T5, T6]) Unapply() (T1, T2, T3, T4, T5, T6) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6
}

type Tuple7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
	I6 T6
	I7 T7
}

func (r Tuple7[T1, T2, T3, T4, T5, T6, T7]) Head() T1 {
	return r.I1
}

func (r Tuple7[T1, T2, T3, T4, T5, T6, T7]) Tail() Tuple6[T2, T3, T4, T5, T6, T7] {
	return Tuple6[T2, T3, T4, T5, T6, T7]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7}
}

func (r Tuple7[T1, T2, T3, T4, T5, T6, T7]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Nil]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple7[T1, T2, T3, T4, T5, T6, T7]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7)
}

func (r Tuple7[T1, T2, T3, T4, T5, T6, T7]) Unapply() (T1, T2, T3, T4, T5, T6, T7) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7
}

type Tuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
	I6 T6
	I7 T7
	I8 T8
}

func (r Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Head() T1 {
	return r.I1
}

func (r Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Tail() Tuple7[T2, T3, T4, T5, T6, T7, T8] {
	return Tuple7[T2, T3, T4, T5, T6, T7, T8]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8}
}

func (r Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Nil]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8)
}

func (r Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8
}

type Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
	I6 T6
	I7 T7
	I8 T8
	I9 T9
}

func (r Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Head() T1 {
	return r.I1
}

func (r Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Tail() Tuple8[T2, T3, T4, T5, T6, T7, T8, T9] {
	return Tuple8[T2, T3, T4, T5, T6, T7, T8, T9]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9}
}

func (r Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Nil]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9)
}

func (r Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9
}

type Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
}

func (r Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Head() T1 {
	return r.I1
}

func (r Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Tail() Tuple9[T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Tuple9[T2, T3, T4, T5, T6, T7, T8, T9, T10]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10}
}

func (r Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Nil]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10)
}

func (r Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10
}

type Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
}

func (r Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Head() T1 {
	return r.I1
}

func (r Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Tail() Tuple10[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11] {
	return Tuple10[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11}
}

func (r Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Nil]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11)
}

func (r Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11
}

type Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
}

func (r Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Head() T1 {
	return r.I1
}

func (r Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Tail() Tuple11[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12] {
	return Tuple11[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12}
}

func (r Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Nil]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12)
}

func (r Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12
}

type Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
	I13 T13
}

func (r Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Head() T1 {
	return r.I1
}

func (r Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Tail() Tuple12[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13] {
	return Tuple12[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13}
}

func (r Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Cons[T13, hlist.Nil]]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13)
}

func (r Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13
}

type Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
	I13 T13
	I14 T14
}

func (r Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Head() T1 {
	return r.I1
}

func (r Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Tail() Tuple13[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14] {
	return Tuple13[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14}
}

func (r Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Cons[T13, hlist.Cons[T14, hlist.Nil]]]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14)
}

func (r Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14
}

type Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
	I13 T13
	I14 T14
	I15 T15
}

func (r Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Head() T1 {
	return r.I1
}

func (r Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Tail() Tuple14[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15] {
	return Tuple14[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15}
}

func (r Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Cons[T13, hlist.Cons[T14, hlist.Cons[T15, hlist.Nil]]]]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15)
}

func (r Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15
}

type Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
	I13 T13
	I14 T14
	I15 T15
	I16 T16
}

func (r Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Head() T1 {
	return r.I1
}

func (r Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Tail() Tuple15[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16] {
	return Tuple15[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16}
}

func (r Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Cons[T13, hlist.Cons[T14, hlist.Cons[T15, hlist.Cons[T16, hlist.Nil]]]]]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16)
}

func (r Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16
}

type Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
	I13 T13
	I14 T14
	I15 T15
	I16 T16
	I17 T17
}

func (r Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) Head() T1 {
	return r.I1
}

func (r Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) Tail() Tuple16[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17] {
	return Tuple16[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17}
}

func (r Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Cons[T13, hlist.Cons[T14, hlist.Cons[T15, hlist.Cons[T16, hlist.Cons[T17, hlist.Nil]]]]]]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17)
}

func (r Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17
}

type Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
	I13 T13
	I14 T14
	I15 T15
	I16 T16
	I17 T17
	I18 T18
}

func (r Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) Head() T1 {
	return r.I1
}

func (r Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) Tail() Tuple17[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18] {
	return Tuple17[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18}
}

func (r Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Cons[T13, hlist.Cons[T14, hlist.Cons[T15, hlist.Cons[T16, hlist.Cons[T17, hlist.Cons[T18, hlist.Nil]]]]]]]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18)
}

func (r Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18
}

type Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
	I13 T13
	I14 T14
	I15 T15
	I16 T16
	I17 T17
	I18 T18
	I19 T19
}

func (r Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) Head() T1 {
	return r.I1
}

func (r Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) Tail() Tuple18[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19] {
	return Tuple18[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19}
}

func (r Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Cons[T13, hlist.Cons[T14, hlist.Cons[T15, hlist.Cons[T16, hlist.Cons[T17, hlist.Cons[T18, hlist.Cons[T19, hlist.Nil]]]]]]]]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19)
}

func (r Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19
}

type Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
	I13 T13
	I14 T14
	I15 T15
	I16 T16
	I17 T17
	I18 T18
	I19 T19
	I20 T20
}

func (r Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) Head() T1 {
	return r.I1
}

func (r Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) Tail() Tuple19[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20] {
	return Tuple19[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20}
}

func (r Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Cons[T13, hlist.Cons[T14, hlist.Cons[T15, hlist.Cons[T16, hlist.Cons[T17, hlist.Cons[T18, hlist.Cons[T19, hlist.Cons[T20, hlist.Nil]]]]]]]]]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20)
}

func (r Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20
}

type Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21 any] struct {
	I1  T1
	I2  T2
	I3  T3
	I4  T4
	I5  T5
	I6  T6
	I7  T7
	I8  T8
	I9  T9
	I10 T10
	I11 T11
	I12 T12
	I13 T13
	I14 T14
	I15 T15
	I16 T16
	I17 T17
	I18 T18
	I19 T19
	I20 T20
	I21 T21
}

func (r Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) Head() T1 {
	return r.I1
}

func (r Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) Tail() Tuple20[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21] {
	return Tuple20[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20, r.I21}
}

func (r Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) ToHList() hlist.Cons[T1, hlist.Cons[T2, hlist.Cons[T3, hlist.Cons[T4, hlist.Cons[T5, hlist.Cons[T6, hlist.Cons[T7, hlist.Cons[T8, hlist.Cons[T9, hlist.Cons[T10, hlist.Cons[T11, hlist.Cons[T12, hlist.Cons[T13, hlist.Cons[T14, hlist.Cons[T15, hlist.Cons[T16, hlist.Cons[T17, hlist.Cons[T18, hlist.Cons[T19, hlist.Cons[T20, hlist.Cons[T21, hlist.Nil]]]]]]]]]]]]]]]]]]]]] {
	return hlist.Concat(r.Head(), r.Tail().ToHList())
}

func (r Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20, r.I21)
}

func (r Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20, r.I21
}
