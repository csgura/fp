package genfp

import (
	"fmt"
	"path"
	"reflect"
)

type TypeTag interface {
	Type() reflect.Type
}

type TypeTagOf[T any] struct {
}

func (r TypeTagOf[T]) Type() reflect.Type {
	var pt *T
	return reflect.TypeOf(pt).Elem()
}

func TypeOf[T any]() TypeTag {
	return TypeTagOf[T]{}
}

func Import(p string) ImportPackage {
	return ImportPackage{
		Package: p,
		Name:    path.Base(p),
	}
}

type Delegate struct {
	TypeOf TypeTag
	Field  string
}

func DelegatedBy[T any](fieldName string) Delegate {
	return Delegate{
		TypeOf: TypeOf[T](),
		Field:  fieldName,
	}
}

var defaultFunc = map[string]any{
	"FormatStr":         FormatStr,
	"FuncChain":         FuncChain,
	"ConsType":          ConsType,
	"ReversConsType":    ReversConsType,
	"TypeArg":           FuncTypeArg,
	"TypeArgs":          FuncTypeArgs,
	"DeclArgs":          FuncDeclArgs,
	"CallArgs":          FuncCallArgs,
	"ReverseCallArgs":   ReverseFuncCallArgs,
	"DeclTypeClassArgs": FuncDeclTypeClassArgs,
	"CurriedCallArgs":   CurriedCallArgs,
	"TypeClassArgs":     TypeClassArgs,
	"CurriedFunc":       CurriedType,
	"RecursiveType":     RecursiveType,
	"EmptyMap": func() map[string]any {
		return map[string]any{}
	},
	"CloneMap": func(m map[string]any) map[string]any {
		ret := map[string]any{}
		for k, v := range m {
			ret[k] = v
		}
		return ret
	},
	"DeleteMap": func(delkey string, target map[string]any) map[string]any {
		delete(target, delkey)
		return target
	},
	"PutMap": func(addk string, addv any, target map[string]any) map[string]any {
		target[addk] = addv
		return target
	},
	"Range": func(start, until int) []int {
		var ret = make([]int, until-start+1)
		for i := start; i <= until; i++ {
			ret[i-start] = i
		}
		return ret
	},
	"Monad": func(s string) Monad {
		return Monad(s)
	},
	"Args": func(s string, start, until int) ArgsRange {
		return ArgsRange{s, start, until}
	},
	"TupleType": func(n int) string {
		return fmt.Sprintf("Tuple%d[%s]", n, FuncTypeArgs(1, n))
	},
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
	"dec": func(n int) int {
		return n - 1
	},
	"inc": func(n int) int {
		return n + 1
	},
}

type GenerateFromUntil struct {
	File     string
	Imports  []ImportPackage
	From     int
	Until    int
	Template string
}

type GenerateFromList struct {
	File     string
	Imports  []ImportPackage
	List     []string
	Template string
}

type TypeName struct {
	// fp.Option[string] 처럼 완전한 type 이름.
	Complete string

	// package
	Package ImportPackage

	// Option 과 같이 package 와 type arg 없는 이름.
	Name string

	// pointer 인지 여부
	IsPtr bool

	IsStruct bool

	// nilable 타입인지 여부
	IsNilable bool

	// zero 값
	ZeroExpr string
}

func (r TypeName) String() string {
	return r.Complete
}

type StructFieldDef struct {
	// field 이름
	Name string

	// field type
	Type TypeName

	// field tag
	Tag string

	// []T, *T, Option[T] 같은 타입인 경우 T 타입
	ElemType TypeName

	// *T 의 경우 T, 아니면 Type과 동일한 값.
	IndirectType TypeName

	// 참조 가능한 field 인지
	IsVisible bool

	// 대문자로 시작하는지
	IsPublic bool
}

type StructDef struct {
	Name      string
	Type      TypeName
	Fields    []StructFieldDef
	AllFields []StructFieldDef
}

func (r StructDef) String() string {
	return r.Name
}

type GenerateFromStructs struct {
	File      string
	Imports   []ImportPackage
	List      []TypeTag
	Recursive bool
	// StructDef 가 .N 에 들어 있음.
	Template string
}

type GenerateAdaptor[T any] struct {
	// 생성될 file 이름
	File string
	// adaptor type 이름
	Name string

	// Extends field 추가 여부
	Extends bool

	// callback 함수에 self 변수 추가 여부
	Self bool

	// Extends 호출 할때 , Self아규먼트 있는지 체크 여부
	ExtendsSelfCheck bool

	// T 이외에 추가로 구현할  interface 목록
	ImplementsWith []TypeTag

	// adaptor struct 에 추가로 포함시킬  field
	ExtendsWith         map[string]TypeTag
	Embedding           []TypeTag
	EmbeddingInterface  []TypeTag
	ExtendsByEmbedding  bool
	Delegate            []Delegate
	Getter              []any
	EventHandler        []any
	ValOverride         []any
	ValOverrideUsingPtr []any
	ZeroReturn          []any
	Options             []ImplOption
}

type AdaptorMethods []ImplOption

type ImplOption struct {
	Method                  any
	Prefix                  string
	Name                    string
	Private                 bool
	ValOverride             bool
	ValOverrideUsingPtr     bool
	OmitGetterIfValOverride bool
	Delegate                Delegate
	DefaultImpl             any
}

func ZeroReturn() {

}

type GenerateMonadFunctions[T any] struct {
	// 생성될 file 이름
	File     string
	TypeParm TypeTag
}

type GenerateTraverseFunctions[T any] struct {
	// 생성될 file 이름
	File     string
	TypeParm TypeTag
}

type MonadFunctions struct {
	Pure    any
	FlatMap any
}

type GenerateMonadTransformer[T any] struct {
	Name string
	// 생성될 file 이름
	File string

	// sequence 를 구현할 때는 T 외에 U 같은 타입 파라미터를 더 쓸 수 있기 때문에
	// 어느 파라미터가  타입 파라미터인지 표시
	TypeParm TypeTag

	// monad transformer 의 아규먼트  ( fp.Try[fp.Option[T]] 의 경우 fp.Try )
	GivenMonad MonadFunctions

	// API 를 노출 하는 monad ( fp.Try[fp.Option[T]] 의 경우 fp.Option )
	ExposureMonad MonadFunctions

	// sequence 함수  fp.Option[fp.Try[T]]  => fp.Try[fp.Option[T]]
	Sequence any

	// ExposureMonad 를 첫번째 아규먼트로 받는 함수들을 지정하면
	// Monad transformer 를 아규먼트로 받는 코드를 생성해 줌
	Transform []any
}
