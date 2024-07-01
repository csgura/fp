package fp

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ShowOption struct {
	Indent string

	OmitEmpty bool

	// true 인 경우, type name 생략
	OmitTypeName bool

	// true 인 경우  1, 2, 3
	// false 인 경우 1,2,3
	SpaceAfterComma bool

	// {} 나 [] 가 여러줄로 출력될 경우, 마지막 , 를 붙일지 여부
	TrailingComma bool

	// true 인 경우  a: 10,b: 20
	// false인 경우 a:10,b:20
	SpaceAfterColon bool

	// type name + {} 형태에 적용되는 옵션
	// true 인 경우  Hello {}
	// false 인 경우 Hello{}
	// a: {} 의 경우 SpaceAfterColon 옵션을 사용.
	SpaceBeforeBrace bool

	// true 인 경우 { 1,2,3 } 혹은 [ 1,2,3 ]
	// false인 경우 {1,2,3} 혹은 [1,2,3]
	SpaceWithinBrace bool

	// true 이면 array 의 경우 [] 사용
	// false 이면 {} 사용
	SquareBracketForArray bool

	// true 이면 nil 을 null 로 출력
	// false 이면 nil 사용
	NullForNil bool

	// true 면  "name": value 로 출력
	// false 면  name: value 로 출력
	QuoteNames bool

	currentIndent string
}

func (r ShowOption) CurrentIndent() string {
	return r.currentIndent
}

func (r ShowOption) IncreaseIndent() ShowOption {
	r.currentIndent = r.currentIndent + r.Indent
	return r
}

func (r ShowOption) WithIndent(indent string) ShowOption {
	r.Indent = indent
	return r
}

func (r ShowOption) WithOmitEmpty(b bool) ShowOption {
	r.OmitEmpty = b
	return r
}

func (r ShowOption) WithOmitTypeName(b bool) ShowOption {
	r.OmitTypeName = b
	return r
}

func (r ShowOption) WithSpaceAfterComma(b bool) ShowOption {
	r.SpaceAfterComma = b
	return r
}

func (r ShowOption) WithSpaceAfterColon(b bool) ShowOption {
	r.SpaceAfterColon = b
	return r
}

func (r ShowOption) WithSpaceBeforeBrace(b bool) ShowOption {
	r.SpaceBeforeBrace = b
	return r
}

func (r ShowOption) WithSpaceWithinBrace(b bool) ShowOption {
	r.SpaceWithinBrace = b
	return r
}

func (r ShowOption) WithSquareBracketForArray(b bool) ShowOption {
	r.SquareBracketForArray = b
	return r
}

func (r ShowOption) WithNullForNil(b bool) ShowOption {
	r.NullForNil = b
	return r
}

func (r ShowOption) WithQuoteNames(b bool) ShowOption {
	r.QuoteNames = b
	return r
}

func (r ShowOption) WithTrailingComma(b bool) ShowOption {
	r.TrailingComma = b
	return r
}

func (r ShowOption) WithSpace() ShowOption {
	r.SpaceAfterColon = true
	r.SpaceAfterComma = true
	r.SpaceBeforeBrace = true
	r.SpaceWithinBrace = true
	return r
}

type Show[T any] interface {
	Show(t T) string
	ShowIndent(t T, option ShowOption) string
	Append(buf []string, t T, option ShowOption) []string
	Stringer(t T, option ShowOption) fmt.Stringer
}

type ShowIndentFunc[T any] func(t T, option ShowOption) string

func (r ShowIndentFunc[T]) Show(t T) string {
	return r(t, ShowOption{})
}

func (r ShowIndentFunc[T]) ShowIndent(t T, opt ShowOption) string {
	return r(t, opt)
}

func (r ShowIndentFunc[T]) Append(buf []string, t T, option ShowOption) []string {
	return append(buf, r(t, option))
}

type StringerFunc func() string

func (r StringerFunc) String() string {
	return r()
}

func (r StringerFunc) MarshalJSON() ([]byte, error) {
	s := r()
	return json.Marshal(s)
}

func (r ShowIndentFunc[T]) Stringer(t T, option ShowOption) fmt.Stringer {
	return StringerFunc(func() string {
		return r.ShowIndent(t, option)
	})
}

type ShowFunc[T any] func(T) string

func Sprint[T any]() Show[T] {
	return ShowFunc[T](func(t T) string {
		return fmt.Sprint(t)
	})
}

func (r ShowFunc[T]) Show(t T) string {
	return r(t)
}

func (r ShowFunc[T]) ShowIndent(t T, opt ShowOption) string {
	return r(t)
}

func (r ShowFunc[T]) Append(buf []string, t T, option ShowOption) []string {
	return append(buf, r(t))
}

func (r ShowFunc[T]) Stringer(t T, option ShowOption) fmt.Stringer {
	return StringerFunc(func() string {
		return r.Show(t)
	})
}

type ShowAppendFunc[T any] func(buf []string, t T, option ShowOption) []string

func (r ShowAppendFunc[T]) Show(t T) string {
	return r.ShowIndent(t, ShowOption{})
}

func (r ShowAppendFunc[T]) ShowIndent(t T, opt ShowOption) string {
	ret := r(nil, t, opt)
	return strings.Join(ret, "")
}

func (r ShowAppendFunc[T]) Append(buf []string, t T, option ShowOption) []string {
	return r(buf, t, option)
}

func (r ShowAppendFunc[T]) Stringer(t T, option ShowOption) fmt.Stringer {
	return StringerFunc(func() string {
		return r.ShowIndent(t, option)
	})
}

type Eq[T any] interface {
	Eqv(a T, b T) bool
}

type EqFunc[T any] func(a, b T) bool

func (r EqFunc[T]) Eqv(a, b T) bool {
	return r(a, b)
}

func EqGiven[T comparable]() Eq[T] {
	return EqFunc[T](func(a, b T) bool {
		return a == b
	})
}

type Hashable[T any] interface {
	Eq[T]
	Hash(T) uint32
}

type Ord[T any] interface {
	Eq[T]

	/*
		-1 if x is less than y,
		 0 if x equals y,
		+1 if x is greater than y.
	*/
	Compare(a, b T) int
	Less(a T, b T) bool
	ThenComparing(other Ord[T]) Ord[T]
	Reversed() Ord[T]
	Max(a T, b T) T
	Min(a T, b T) T
}

type CompareFunc[T any] func(a, b T) int

func (r CompareFunc[T]) Eqv(a, b T) bool {
	return r(a, b) == 0
}

func (r CompareFunc[T]) Compare(a, b T) int {
	return r(a, b)
}

func (r CompareFunc[T]) Less(a, b T) bool {
	return r(a, b) < 0
}

func (r CompareFunc[T]) Max(a, b T) T {
	if r(a, b) < 0 {
		return b
	}
	return a
}
func (r CompareFunc[T]) Min(a, b T) T {
	if r(a, b) < 0 {
		return a
	}
	return b
}

func (r CompareFunc[T]) ThenComparing(other Ord[T]) Ord[T] {
	return CompareFunc[T](func(a, b T) int {
		res := r.Compare(a, b)
		if res == 0 {
			return other.Compare(a, b)
		}
		return res
	})
}

func (r CompareFunc[T]) Reversed() Ord[T] {
	return CompareFunc[T](func(a, b T) int {
		return -r.Compare(a, b)
	})
}

type LessFunc[T any] func(a, b T) bool

func (r LessFunc[T]) Compare(a, b T) int {
	if r(a, b) {
		return -1
	}
	if r(b, a) {
		return 1
	}
	return 0
}

func (r LessFunc[T]) Eqv(a, b T) bool {
	return r.Compare(a, b) == 0
}

func (r LessFunc[T]) Less(a, b T) bool {
	return r(a, b)
}

func (r LessFunc[T]) Max(a, b T) T {
	if r(a, b) {
		return b
	}
	return a
}

func (r LessFunc[T]) Min(a, b T) T {
	if r(a, b) {
		return a
	}
	return b
}

func (r LessFunc[T]) ThenComparing(other Ord[T]) Ord[T] {
	return CompareFunc[T](func(a, b T) int {
		res := r.Compare(a, b)
		if res == 0 {
			return other.Compare(a, b)
		}
		return res
	})
}

func (r LessFunc[T]) Reversed() Ord[T] {
	return CompareFunc[T](func(a, b T) int {
		return -r.Compare(a, b)
	})
}

func LessGiven[T ImplicitOrd]() Ord[T] {
	return CompareFunc[T](func(a, b T) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
}

type Clone[T any] interface {
	Clone(t T) T
}

type CloneFunc[T any] func(t T) T

func (r CloneFunc[T]) Clone(t T) T {
	return r(t)
}
