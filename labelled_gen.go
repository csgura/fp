// Code generated by fp_gen, DO NOT EDIT.
package fp

import (
	"fmt"
)

type Labelled2[T1, T2 Named] struct {
	I1 T1
	I2 T2
}

func (r Labelled2[T1, T2]) Head() T1 {
	return r.I1
}

func (r Labelled2[T1, T2]) Tail() Labelled1[T2] {
	return Labelled1[T2]{r.I2}
}

func (r Labelled2[T1, T2]) String() string {
	return fmt.Sprintf("(%v,%v)", r.I1, r.I2)
}

func (r Labelled2[T1, T2]) Unapply() (T1, T2) {
	return r.I1, r.I2
}

type Labelled3[T1, T2, T3 Named] struct {
	I1 T1
	I2 T2
	I3 T3
}

func (r Labelled3[T1, T2, T3]) Head() T1 {
	return r.I1
}

func (r Labelled3[T1, T2, T3]) Tail() Labelled2[T2, T3] {
	return Labelled2[T2, T3]{r.I2, r.I3}
}

func (r Labelled3[T1, T2, T3]) String() string {
	return fmt.Sprintf("(%v,%v,%v)", r.I1, r.I2, r.I3)
}

func (r Labelled3[T1, T2, T3]) Unapply() (T1, T2, T3) {
	return r.I1, r.I2, r.I3
}

type Labelled4[T1, T2, T3, T4 Named] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
}

func (r Labelled4[T1, T2, T3, T4]) Head() T1 {
	return r.I1
}

func (r Labelled4[T1, T2, T3, T4]) Tail() Labelled3[T2, T3, T4] {
	return Labelled3[T2, T3, T4]{r.I2, r.I3, r.I4}
}

func (r Labelled4[T1, T2, T3, T4]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4)
}

func (r Labelled4[T1, T2, T3, T4]) Unapply() (T1, T2, T3, T4) {
	return r.I1, r.I2, r.I3, r.I4
}

type Labelled5[T1, T2, T3, T4, T5 Named] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
}

func (r Labelled5[T1, T2, T3, T4, T5]) Head() T1 {
	return r.I1
}

func (r Labelled5[T1, T2, T3, T4, T5]) Tail() Labelled4[T2, T3, T4, T5] {
	return Labelled4[T2, T3, T4, T5]{r.I2, r.I3, r.I4, r.I5}
}

func (r Labelled5[T1, T2, T3, T4, T5]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5)
}

func (r Labelled5[T1, T2, T3, T4, T5]) Unapply() (T1, T2, T3, T4, T5) {
	return r.I1, r.I2, r.I3, r.I4, r.I5
}

type Labelled6[T1, T2, T3, T4, T5, T6 Named] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
	I6 T6
}

func (r Labelled6[T1, T2, T3, T4, T5, T6]) Head() T1 {
	return r.I1
}

func (r Labelled6[T1, T2, T3, T4, T5, T6]) Tail() Labelled5[T2, T3, T4, T5, T6] {
	return Labelled5[T2, T3, T4, T5, T6]{r.I2, r.I3, r.I4, r.I5, r.I6}
}

func (r Labelled6[T1, T2, T3, T4, T5, T6]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6)
}

func (r Labelled6[T1, T2, T3, T4, T5, T6]) Unapply() (T1, T2, T3, T4, T5, T6) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6
}

type Labelled7[T1, T2, T3, T4, T5, T6, T7 Named] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
	I6 T6
	I7 T7
}

func (r Labelled7[T1, T2, T3, T4, T5, T6, T7]) Head() T1 {
	return r.I1
}

func (r Labelled7[T1, T2, T3, T4, T5, T6, T7]) Tail() Labelled6[T2, T3, T4, T5, T6, T7] {
	return Labelled6[T2, T3, T4, T5, T6, T7]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7}
}

func (r Labelled7[T1, T2, T3, T4, T5, T6, T7]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7)
}

func (r Labelled7[T1, T2, T3, T4, T5, T6, T7]) Unapply() (T1, T2, T3, T4, T5, T6, T7) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7
}

type Labelled8[T1, T2, T3, T4, T5, T6, T7, T8 Named] struct {
	I1 T1
	I2 T2
	I3 T3
	I4 T4
	I5 T5
	I6 T6
	I7 T7
	I8 T8
}

func (r Labelled8[T1, T2, T3, T4, T5, T6, T7, T8]) Head() T1 {
	return r.I1
}

func (r Labelled8[T1, T2, T3, T4, T5, T6, T7, T8]) Tail() Labelled7[T2, T3, T4, T5, T6, T7, T8] {
	return Labelled7[T2, T3, T4, T5, T6, T7, T8]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8}
}

func (r Labelled8[T1, T2, T3, T4, T5, T6, T7, T8]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8)
}

func (r Labelled8[T1, T2, T3, T4, T5, T6, T7, T8]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8
}

type Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9 Named] struct {
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

func (r Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Head() T1 {
	return r.I1
}

func (r Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Tail() Labelled8[T2, T3, T4, T5, T6, T7, T8, T9] {
	return Labelled8[T2, T3, T4, T5, T6, T7, T8, T9]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9}
}

func (r Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9)
}

func (r Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9
}

type Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 Named] struct {
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

func (r Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Head() T1 {
	return r.I1
}

func (r Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Tail() Labelled9[T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Labelled9[T2, T3, T4, T5, T6, T7, T8, T9, T10]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10}
}

func (r Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10)
}

func (r Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10
}

type Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 Named] struct {
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

func (r Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Head() T1 {
	return r.I1
}

func (r Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Tail() Labelled10[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11] {
	return Labelled10[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11}
}

func (r Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11)
}

func (r Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11
}

type Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 Named] struct {
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

func (r Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Head() T1 {
	return r.I1
}

func (r Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Tail() Labelled11[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12] {
	return Labelled11[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12}
}

func (r Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12)
}

func (r Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12
}

type Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 Named] struct {
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

func (r Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Head() T1 {
	return r.I1
}

func (r Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Tail() Labelled12[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13] {
	return Labelled12[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13}
}

func (r Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13)
}

func (r Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13
}

type Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 Named] struct {
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

func (r Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Head() T1 {
	return r.I1
}

func (r Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Tail() Labelled13[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14] {
	return Labelled13[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14}
}

func (r Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14)
}

func (r Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14
}

type Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 Named] struct {
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

func (r Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Head() T1 {
	return r.I1
}

func (r Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Tail() Labelled14[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15] {
	return Labelled14[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15}
}

func (r Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15)
}

func (r Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15
}

type Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 Named] struct {
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

func (r Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Head() T1 {
	return r.I1
}

func (r Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Tail() Labelled15[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16] {
	return Labelled15[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16}
}

func (r Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16)
}

func (r Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16
}

type Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17 Named] struct {
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

func (r Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) Head() T1 {
	return r.I1
}

func (r Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) Tail() Labelled16[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17] {
	return Labelled16[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17}
}

func (r Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17)
}

func (r Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17
}

type Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18 Named] struct {
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

func (r Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) Head() T1 {
	return r.I1
}

func (r Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) Tail() Labelled17[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18] {
	return Labelled17[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18}
}

func (r Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18)
}

func (r Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18
}

type Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19 Named] struct {
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

func (r Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) Head() T1 {
	return r.I1
}

func (r Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) Tail() Labelled18[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19] {
	return Labelled18[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19}
}

func (r Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19)
}

func (r Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19
}

type Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20 Named] struct {
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

func (r Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) Head() T1 {
	return r.I1
}

func (r Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) Tail() Labelled19[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20] {
	return Labelled19[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20}
}

func (r Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20)
}

func (r Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20
}

type Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21 Named] struct {
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

func (r Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) Head() T1 {
	return r.I1
}

func (r Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) Tail() Labelled20[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21] {
	return Labelled20[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20, r.I21}
}

func (r Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20, r.I21)
}

func (r Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) Unapply() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20, r.I21
}