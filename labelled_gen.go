package fp

import (
	"fmt"
)

type Labelled2[T1, T2 any] struct {
	I1 Field[T1]
	I2 Field[T2]
}

func (r Labelled2[T1, T2]) Head() Field[T1] {
	return r.I1
}

func (r Labelled2[T1, T2]) Tail() Labelled1[T2] {
	return Labelled1[T2]{r.I2}
}

func (r Labelled2[T1, T2]) String() string {
	return fmt.Sprintf("(%v,%v)", r.I1, r.I2)
}

func (r Labelled2[T1, T2]) Unapply() (Field[T1], Field[T2]) {
	return r.I1, r.I2
}

type Labelled3[T1, T2, T3 any] struct {
	I1 Field[T1]
	I2 Field[T2]
	I3 Field[T3]
}

func (r Labelled3[T1, T2, T3]) Head() Field[T1] {
	return r.I1
}

func (r Labelled3[T1, T2, T3]) Tail() Labelled2[T2, T3] {
	return Labelled2[T2, T3]{r.I2, r.I3}
}

func (r Labelled3[T1, T2, T3]) String() string {
	return fmt.Sprintf("(%v,%v,%v)", r.I1, r.I2, r.I3)
}

func (r Labelled3[T1, T2, T3]) Unapply() (Field[T1], Field[T2], Field[T3]) {
	return r.I1, r.I2, r.I3
}

type Labelled4[T1, T2, T3, T4 any] struct {
	I1 Field[T1]
	I2 Field[T2]
	I3 Field[T3]
	I4 Field[T4]
}

func (r Labelled4[T1, T2, T3, T4]) Head() Field[T1] {
	return r.I1
}

func (r Labelled4[T1, T2, T3, T4]) Tail() Labelled3[T2, T3, T4] {
	return Labelled3[T2, T3, T4]{r.I2, r.I3, r.I4}
}

func (r Labelled4[T1, T2, T3, T4]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4)
}

func (r Labelled4[T1, T2, T3, T4]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4]) {
	return r.I1, r.I2, r.I3, r.I4
}

type Labelled5[T1, T2, T3, T4, T5 any] struct {
	I1 Field[T1]
	I2 Field[T2]
	I3 Field[T3]
	I4 Field[T4]
	I5 Field[T5]
}

func (r Labelled5[T1, T2, T3, T4, T5]) Head() Field[T1] {
	return r.I1
}

func (r Labelled5[T1, T2, T3, T4, T5]) Tail() Labelled4[T2, T3, T4, T5] {
	return Labelled4[T2, T3, T4, T5]{r.I2, r.I3, r.I4, r.I5}
}

func (r Labelled5[T1, T2, T3, T4, T5]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5)
}

func (r Labelled5[T1, T2, T3, T4, T5]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5
}

type Labelled6[T1, T2, T3, T4, T5, T6 any] struct {
	I1 Field[T1]
	I2 Field[T2]
	I3 Field[T3]
	I4 Field[T4]
	I5 Field[T5]
	I6 Field[T6]
}

func (r Labelled6[T1, T2, T3, T4, T5, T6]) Head() Field[T1] {
	return r.I1
}

func (r Labelled6[T1, T2, T3, T4, T5, T6]) Tail() Labelled5[T2, T3, T4, T5, T6] {
	return Labelled5[T2, T3, T4, T5, T6]{r.I2, r.I3, r.I4, r.I5, r.I6}
}

func (r Labelled6[T1, T2, T3, T4, T5, T6]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6)
}

func (r Labelled6[T1, T2, T3, T4, T5, T6]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6
}

type Labelled7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	I1 Field[T1]
	I2 Field[T2]
	I3 Field[T3]
	I4 Field[T4]
	I5 Field[T5]
	I6 Field[T6]
	I7 Field[T7]
}

func (r Labelled7[T1, T2, T3, T4, T5, T6, T7]) Head() Field[T1] {
	return r.I1
}

func (r Labelled7[T1, T2, T3, T4, T5, T6, T7]) Tail() Labelled6[T2, T3, T4, T5, T6, T7] {
	return Labelled6[T2, T3, T4, T5, T6, T7]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7}
}

func (r Labelled7[T1, T2, T3, T4, T5, T6, T7]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7)
}

func (r Labelled7[T1, T2, T3, T4, T5, T6, T7]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7
}

type Labelled8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	I1 Field[T1]
	I2 Field[T2]
	I3 Field[T3]
	I4 Field[T4]
	I5 Field[T5]
	I6 Field[T6]
	I7 Field[T7]
	I8 Field[T8]
}

func (r Labelled8[T1, T2, T3, T4, T5, T6, T7, T8]) Head() Field[T1] {
	return r.I1
}

func (r Labelled8[T1, T2, T3, T4, T5, T6, T7, T8]) Tail() Labelled7[T2, T3, T4, T5, T6, T7, T8] {
	return Labelled7[T2, T3, T4, T5, T6, T7, T8]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8}
}

func (r Labelled8[T1, T2, T3, T4, T5, T6, T7, T8]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8)
}

func (r Labelled8[T1, T2, T3, T4, T5, T6, T7, T8]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8
}

type Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	I1 Field[T1]
	I2 Field[T2]
	I3 Field[T3]
	I4 Field[T4]
	I5 Field[T5]
	I6 Field[T6]
	I7 Field[T7]
	I8 Field[T8]
	I9 Field[T9]
}

func (r Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Head() Field[T1] {
	return r.I1
}

func (r Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Tail() Labelled8[T2, T3, T4, T5, T6, T7, T8, T9] {
	return Labelled8[T2, T3, T4, T5, T6, T7, T8, T9]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9}
}

func (r Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9)
}

func (r Labelled9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9
}

type Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
}

func (r Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Head() Field[T1] {
	return r.I1
}

func (r Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Tail() Labelled9[T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Labelled9[T2, T3, T4, T5, T6, T7, T8, T9, T10]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10}
}

func (r Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10)
}

func (r Labelled10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10
}

type Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
}

func (r Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Head() Field[T1] {
	return r.I1
}

func (r Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Tail() Labelled10[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11] {
	return Labelled10[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11}
}

func (r Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11)
}

func (r Labelled11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11
}

type Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
}

func (r Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Head() Field[T1] {
	return r.I1
}

func (r Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Tail() Labelled11[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12] {
	return Labelled11[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12}
}

func (r Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12)
}

func (r Labelled12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12
}

type Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
	I13 Field[T13]
}

func (r Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Head() Field[T1] {
	return r.I1
}

func (r Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Tail() Labelled12[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13] {
	return Labelled12[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13}
}

func (r Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13)
}

func (r Labelled13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12], Field[T13]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13
}

type Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
	I13 Field[T13]
	I14 Field[T14]
}

func (r Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Head() Field[T1] {
	return r.I1
}

func (r Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Tail() Labelled13[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14] {
	return Labelled13[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14}
}

func (r Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14)
}

func (r Labelled14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12], Field[T13], Field[T14]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14
}

type Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
	I13 Field[T13]
	I14 Field[T14]
	I15 Field[T15]
}

func (r Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Head() Field[T1] {
	return r.I1
}

func (r Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Tail() Labelled14[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15] {
	return Labelled14[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15}
}

func (r Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15)
}

func (r Labelled15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12], Field[T13], Field[T14], Field[T15]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15
}

type Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
	I13 Field[T13]
	I14 Field[T14]
	I15 Field[T15]
	I16 Field[T16]
}

func (r Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Head() Field[T1] {
	return r.I1
}

func (r Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Tail() Labelled15[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16] {
	return Labelled15[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16}
}

func (r Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16)
}

func (r Labelled16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12], Field[T13], Field[T14], Field[T15], Field[T16]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16
}

type Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
	I13 Field[T13]
	I14 Field[T14]
	I15 Field[T15]
	I16 Field[T16]
	I17 Field[T17]
}

func (r Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) Head() Field[T1] {
	return r.I1
}

func (r Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) Tail() Labelled16[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17] {
	return Labelled16[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17}
}

func (r Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17)
}

func (r Labelled17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12], Field[T13], Field[T14], Field[T15], Field[T16], Field[T17]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17
}

type Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
	I13 Field[T13]
	I14 Field[T14]
	I15 Field[T15]
	I16 Field[T16]
	I17 Field[T17]
	I18 Field[T18]
}

func (r Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) Head() Field[T1] {
	return r.I1
}

func (r Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) Tail() Labelled17[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18] {
	return Labelled17[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18}
}

func (r Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18)
}

func (r Labelled18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12], Field[T13], Field[T14], Field[T15], Field[T16], Field[T17], Field[T18]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18
}

type Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
	I13 Field[T13]
	I14 Field[T14]
	I15 Field[T15]
	I16 Field[T16]
	I17 Field[T17]
	I18 Field[T18]
	I19 Field[T19]
}

func (r Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) Head() Field[T1] {
	return r.I1
}

func (r Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) Tail() Labelled18[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19] {
	return Labelled18[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19}
}

func (r Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19)
}

func (r Labelled19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12], Field[T13], Field[T14], Field[T15], Field[T16], Field[T17], Field[T18], Field[T19]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19
}

type Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
	I13 Field[T13]
	I14 Field[T14]
	I15 Field[T15]
	I16 Field[T16]
	I17 Field[T17]
	I18 Field[T18]
	I19 Field[T19]
	I20 Field[T20]
}

func (r Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) Head() Field[T1] {
	return r.I1
}

func (r Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) Tail() Labelled19[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20] {
	return Labelled19[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20}
}

func (r Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20)
}

func (r Labelled20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12], Field[T13], Field[T14], Field[T15], Field[T16], Field[T17], Field[T18], Field[T19], Field[T20]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20
}

type Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21 any] struct {
	I1  Field[T1]
	I2  Field[T2]
	I3  Field[T3]
	I4  Field[T4]
	I5  Field[T5]
	I6  Field[T6]
	I7  Field[T7]
	I8  Field[T8]
	I9  Field[T9]
	I10 Field[T10]
	I11 Field[T11]
	I12 Field[T12]
	I13 Field[T13]
	I14 Field[T14]
	I15 Field[T15]
	I16 Field[T16]
	I17 Field[T17]
	I18 Field[T18]
	I19 Field[T19]
	I20 Field[T20]
	I21 Field[T21]
}

func (r Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) Head() Field[T1] {
	return r.I1
}

func (r Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) Tail() Labelled20[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21] {
	return Labelled20[T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]{r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20, r.I21}
}

func (r Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20, r.I21)
}

func (r Labelled21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]) Unapply() (Field[T1], Field[T2], Field[T3], Field[T4], Field[T5], Field[T6], Field[T7], Field[T8], Field[T9], Field[T10], Field[T11], Field[T12], Field[T13], Field[T14], Field[T15], Field[T16], Field[T17], Field[T18], Field[T19], Field[T20], Field[T21]) {
	return r.I1, r.I2, r.I3, r.I4, r.I5, r.I6, r.I7, r.I8, r.I9, r.I10, r.I11, r.I12, r.I13, r.I14, r.I15, r.I16, r.I17, r.I18, r.I19, r.I20, r.I21
}
