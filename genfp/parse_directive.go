package genfp

import (
	"fmt"
	"go/ast"
	"go/types"
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

type TypeReference struct {
	Expr       ast.Expr
	StringExpr string
	Type       types.Type
	Imports    []ImportPackage
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

type MethodReference struct {
	Receiver TypeReference
	Name     string
}

func (r MethodReference) GetName() string {
	return r.Name
}

type GenerateMonadTransformer[T any] struct {
	Name string
	// 생성될 file 이름
	File      string
	TypeParm  TypeTag
	Pure      any
	FlatMap   any
	Sequence  any
	Transform []any
}
