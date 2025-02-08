package genfp

import (
	"fmt"
	"go/types"
	"path"
	"reflect"
	"strings"
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

func Imports(p ...string) []ImportPackage {
	var ret []ImportPackage
	for _, v := range p {
		ret = append(ret, Import(v))
	}
	return ret
}

func PackageOfType[T any]() ImportPackage {
	tp := TypeOf[T]()
	return Import(tp.Type().PkgPath())
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
	"PublicName":        PublicName,
	"PrivateName":       PrivateName,
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

func PublicName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func PrivateName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

type GenerateFromUntil struct {
	File      string
	Imports   []ImportPackage
	From      int
	Until     int
	Variables map[string]string
	Template  string
}

type GenerateFromList struct {
	File      string
	Imports   []ImportPackage
	List      []string
	Variables map[string]string
	Template  string
}

// TypeDecl 을 사용하여 type string 으로 변환 가능
type TypeInfo struct {
	// golang types.Type
	Type types.Type

	// fp.Option[string] 처럼 완전한 type 이름.
	// TypeDecl과 다른 점은 ,  import 를 하지 않음.
	// "" 안에서 사용할 경우 이걸 사용
	// import가 필요하면 TypeDecl 사용
	Complete string

	// package
	Package ImportPackage

	IsCurrentPackage bool

	// "Option" 형태로 package 와 type arg 없는 이름.
	Name string

	// nilable 타입인지 여부
	IsNilable bool

	IsBasic bool

	// pointer 인지 여부
	IsPtr bool

	IsNumber bool
	IsBool   bool
	IsString bool
	IsSlice  bool
	IsMap    bool
	IsFunc   bool

	IsStruct     bool
	IsInterface  bool
	IsError      bool
	IsComparable bool
	IsAny        bool

	IsOption bool
	IsTry    bool

	// zero 값
	ZeroExpr string

	TypeArgs []TypeInfo
}

func (r TypeInfo) String() string {
	return r.Complete
}

func (r TypeInfo) TypeArgAt(i int) *TypeInfo {
	if i < len(r.TypeArgs) {
		return &r.TypeArgs[i]
	}
	return nil
}

type StructFieldInfo struct {
	// field 이름
	Name string

	// field type
	Type TypeInfo

	// field tag
	Tag string

	// []T, *T, Option[T] 같은 타입인 경우 T 타입
	ElemType TypeInfo

	// *T 의 경우 T, 아니면 Type과 동일한 값.
	IndirectType TypeInfo

	// 참조 가능한 field 인지
	IsVisible bool

	// 대문자로 시작하는지
	IsPublic bool
}

type StructInfo struct {
	Package          ImportPackage
	IsCurrentPackage bool
	Name             string
	Type             TypeInfo
	// 참조 가능한 field목록
	Fields []StructFieldInfo
	// 모든 field 목록
	AllFields []StructFieldInfo
}

func (r StructInfo) FieldAt(at int) *StructFieldInfo {
	if at < len(r.Fields) {
		return &r.Fields[at]
	}
	return nil
}

func (r StructInfo) String() string {
	return r.Name
}

type GenerateFromStructs struct {
	File      string
	Imports   []ImportPackage
	List      []TypeTag
	Recursive bool
	Variables map[string]string
	// StructInfo 가 .N 에 들어 있음.
	Template string
}

// VarDecl 함수를 사용하면 name type 형태로 변환 가능
// TypeDecl 함수를 사용하면 type 만 리턴
type VarInfo struct {
	Index int
	// 선언에 변수이름 없으면 ""
	Name string
	Type TypeInfo
}

func (r VarInfo) String() string {
	return r.Type.String()
}

type InterfaceMethodInfo struct {
	// method 이름
	Name string

	// arg type
	// VarDecl .Args 하면  a type, b type 형태로 리턴
	// TypeDecl .Args 하면 type, type 형태로 리턴
	Args []VarInfo

	// return type
	// Args와 동일하게 VarDecl , TypeDecl 사용 가능
	Returns []VarInfo
}

// TypeDecl .Args와 다른 점은 import하지 않음.
func (r InterfaceMethodInfo) ArgsDef() string {
	return seqMakeString(seqMap(r.Args, func(v VarInfo) string {
		return fmt.Sprintf("%s %s", v.Name, v.Type.Complete)
	}), ",")
}

// a,b,c 형태로 이름만 리턴
func (r InterfaceMethodInfo) ArgsCall() string {
	return seqMakeString(seqMap(r.Args, func(v VarInfo) string {
		return fmt.Sprintf("%s", v.Name)
	}), ",")
}

// TypeDecl .Returns와 다른 점은 import하지 않음.
func (r InterfaceMethodInfo) ReturnsDef() string {
	return seqMakeString(seqMap(r.Returns, func(v VarInfo) string {
		return fmt.Sprintf("%s", v.Type.Complete)
	}), ",")
}

func (r InterfaceMethodInfo) ArgAt(i int) *VarInfo {
	if i < len(r.Args) {
		return &r.Args[i]
	}
	return nil
}

func (r InterfaceMethodInfo) ReturnAt(i int) *VarInfo {
	if i < len(r.Returns) {
		return &r.Returns[i]
	}
	return nil
}

type InterfaceInfo struct {
	Package          ImportPackage
	IsCurrentPackage bool
	Name             string
	Type             TypeInfo
	Methods          []InterfaceMethodInfo
}

type GenerateFromInterfaces struct {
	File      string
	Imports   []ImportPackage
	List      []TypeTag
	Variables map[string]string
	// InterfaceInfo 가 .N 에 들어 있음.
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
