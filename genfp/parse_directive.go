package genfp

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
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

type GenerateFromUntil struct {
	File     string
	Imports  []ImportPackage
	From     int
	Until    int
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

func asKeyValue(e ast.Expr, defName string) (string, ast.Expr) {
	if kv, ok := e.(*ast.KeyValueExpr); ok {
		if id, ok := kv.Key.(*ast.Ident); ok {
			return id.Name, kv.Value
		}
	}
	return defName, e
}

func evalStringValue(p *packages.Package, e ast.Expr) (string, error) {
	switch t := e.(type) {
	case *ast.BasicLit:
		if strings.HasPrefix(t.Value, `"`) && strings.HasSuffix(t.Value, `"`) {
			return t.Value[1 : len(t.Value)-1], nil
		} else if strings.HasPrefix(t.Value, "`") && strings.HasSuffix(t.Value, "`") {
			return t.Value[1 : len(t.Value)-1], nil
		}
	case *ast.Ident:
		v, _ := evalConst(p, e)
		if v.Kind() == constant.String {
			return v.String(), nil
		}
	case *ast.SelectorExpr:
		v, _ := evalConst(p, e)
		if v.Kind() == constant.String {
			return v.String(), nil
		}
	}
	return "", fmt.Errorf("can't eval %T as string", e)
}

// func evalSelectorExpr(e ast.Expr) (string, error) {
// 	switch t := e.(type) {
// 	case *ast.Ident:
// 		return t.Name, nil
// 	case *ast.SelectorExpr:
// 		switch x := t.X.(type) {
// 		case *ast.SelectorExpr:
// 			p, err := evalSelectorExpr(x)
// 			if err != nil {
// 				return "", err
// 			}
// 			return fmt.Sprintf("%s.%s", p, t.Sel.Name), nil
// 		case *ast.Ident:
// 			return fmt.Sprintf("%s.%s", x.Name, t.Sel.Name), nil
// 		}
// 	}
// 	return "", fmt.Errorf("can't eval %T as selector expr", e)
// }

type MethodReference struct {
	Receiver TypeReference
	Name     string
}

func (r MethodReference) GetName() string {
	return r.Name
}

func asMethodRef(p *packages.Package, receiver ast.Expr, name string) MethodReference {
	return MethodReference{
		Receiver: evalTypeReference(p, receiver),
		Name:     name,
	}
}

func evalMethodRef(tname string) func(p *packages.Package, e ast.Expr) (MethodReference, error) {
	return func(p *packages.Package, e ast.Expr) (MethodReference, error) {
		switch t := e.(type) {
		case *ast.SelectorExpr:
			switch t.X.(type) {
			case *ast.SelectorExpr:
				return asMethodRef(p, t.X, t.Sel.Name), nil
			// p, err := evalSelectorExpr(x)
			// if err != nil {
			// 	return "", err
			// }
			// if strings.HasSuffix(p, tname) {
			// 	return t.Sel.Name, nil
			// }
			// return "", fmt.Errorf("invalid method reference : %s", p)
			case *ast.IndexExpr:
				return asMethodRef(p, t.X, t.Sel.Name), nil
			case *ast.Ident:
				return asMethodRef(p, t.X, t.Sel.Name), nil
				// if x.Name == tname {
				// 	return t.Sel.Name, nil
				// }
				// return "", fmt.Errorf("invalid method reference : %s", x.Name)
			}
			return MethodReference{}, fmt.Errorf("can't eval %T as method reference", t.X)
		case *ast.IndexListExpr:
			return evalMethodRef(tname)(p, t.X)
		}
		return MethodReference{}, fmt.Errorf("can't eval %T as method reference", e)
	}

}

func evalBoolValue(e ast.Expr) (bool, error) {
	switch t := e.(type) {
	case *ast.Ident:
		if t.Name == "true" {
			return true, nil
		} else if t.Name == "false" {
			return false, nil
		}
	}
	return false, fmt.Errorf("can't eval %T as bool", e)
}

func evalIntValue(p *packages.Package, e ast.Expr) (int, error) {
	switch t := e.(type) {
	case *ast.BasicLit:
		i, err := strconv.ParseInt(t.Value, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("can't parseInt %s", t.Value)
		}
		return int(i), nil
	case *ast.Ident:
		v, _ := evalConst(p, e)
		if v.Kind() == constant.Int {
			i, err := strconv.ParseInt(v.String(), 10, 64)
			if err != nil {
				return 0, fmt.Errorf("can't parseInt %s", v.String())
			}
			return int(i), nil
		}
	case *ast.SelectorExpr:
		v, _ := evalConst(p, e)
		if v.Kind() == constant.Int {
			i, err := strconv.ParseInt(v.String(), 10, 64)
			if err != nil {
				return 0, fmt.Errorf("can't parseInt %s", v.String())
			}
			return int(i), nil
		}
	}
	return 0, fmt.Errorf("can't eval %T as int", e)
}

func evalImport(p *packages.Package, e ast.Expr) (ImportPackage, error) {
	if lt, ok := e.(*ast.CompositeLit); ok {
		ret := ImportPackage{}
		names := []string{"Package", "Name"}
		for idx, e := range lt.Elts {
			if idx >= len(names) {
				return ImportPackage{}, fmt.Errorf("invalid number of literals")
			}

			name := names[idx]
			name, value := asKeyValue(e, name)

			switch name {
			case "Package":
				v, err := evalStringValue(p, value)
				if err != nil {
					return ImportPackage{}, err
				}
				ret.Package = v
			case "Name":
				v, err := evalStringValue(p, value)
				if err != nil {
					return ImportPackage{}, err
				}
				ret.Name = v
			}
		}
		return ret, nil
	}
	return ImportPackage{}, fmt.Errorf("expr is not composite expr : %T", e)

}

func matchSelExpr(sel *ast.SelectorExpr, exp []string) bool {
	if len(exp) == 0 {
		return false
	}
	init := exp[0 : len(exp)-1]
	last := exp[len(exp)-1]
	if sel.Sel.Name == last {
		switch t := sel.X.(type) {
		case *ast.SelectorExpr:
			return matchSelExpr(t, init)
		case *ast.Ident:
			if len(init) == 1 {
				return t.Name == init[0]
			}
		}
	}
	return false
}

func evalArray[T any](p *packages.Package, e ast.Expr, f func(*packages.Package, ast.Expr) (T, error)) ([]T, error) {

	if lt, ok := e.(*ast.CompositeLit); ok {
		var ret []T

		for _, e := range lt.Elts {
			v, err := f(p, e)
			if err != nil {
				return nil, err
			}
			ret = append(ret, v)
		}
		return ret, nil

	}
	return nil, fmt.Errorf("expr is not array expr : %T", e)
}

func evalMap[K comparable, V any](p *packages.Package, e ast.Expr, kf func(*packages.Package, ast.Expr) (K, error), vf func(*packages.Package, ast.Expr) (V, error)) (map[K]V, error) {

	if lt, ok := e.(*ast.CompositeLit); ok {
		ret := map[K]V{}

		for _, e := range lt.Elts {
			if kve, ok := e.(*ast.KeyValueExpr); ok {
				kv, err := kf(p, kve.Key)
				if err != nil {
					return nil, err
				}

				vv, err := vf(p, kve.Value)
				if err != nil {
					return nil, err
				}
				ret[kv] = vv

			} else {
				return nil, fmt.Errorf("expr is not map expr : %T", e)
			}
		}
		return ret, nil

	}
	return nil, fmt.Errorf("expr is not map expr : %T", e)
}

func ParseGenerateFromUntil(tagged TaggedLit) (GenerateFromUntil, error) {

	lit := tagged.Lit
	ret := GenerateFromUntil{}

	names := []string{"File", "Imports", "From", "Until", "Template"}
	for idx, e := range lit.Elts {
		if idx >= len(names) {
			return GenerateFromUntil{}, fmt.Errorf("invalid number of literals")
		}

		name := names[idx]
		name, value := asKeyValue(e, name)
		switch name {
		case "File":
			v, err := evalStringValue(tagged.Package, value)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.File = v
		case "Imports":
			v, err := evalArray(tagged.Package, value, evalImport)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.Imports = v
		case "From":
			v, err := evalIntValue(tagged.Package, value)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.From = v
		case "Until":
			v, err := evalIntValue(tagged.Package, value)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.Until = v
		case "Template":
			v, err := evalStringValue(tagged.Package, value)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.Template = v
		}
	}

	return ret, nil

}

type GenerateAdaptorDirective struct {
	Package             *packages.Package
	Interface           *types.Named
	File                string
	Name                string
	Extends             bool
	Self                bool
	ExtendsSelfCheck    bool
	ImplementsWith      []TypeReference
	ExtendsWith         map[string]TypeReference
	Embedding           []TypeReference
	EmbeddingInterface  []TypeReference
	ExtendsByEmbedding  bool
	Delegate            []DelegateDirective
	Getter              []string
	EventHandler        []string
	ValOverride         []string
	ValOverrideUsingPtr []string
	ZeroReturn          []string
	Methods             map[string]ImplOptionDirective
}

type TypeReference struct {
	Expr       ast.Expr
	StringExpr string
	Type       types.Type
	Imports    []ImportPackage
}

type FuncReference struct {
	Name          string
	TypeParams    []TypeReference
	TypeReference TypeReference
}

func evalTypeReference(pk *packages.Package, exp ast.Expr) TypeReference {
	t, imp := evalFuncLit(pk, exp)

	return TypeReference{
		Expr:       exp,
		StringExpr: types.ExprString(exp),
		Type:       t,
		Imports:    imp,
	}
}
func extractTypeParam(pk *packages.Package, exp ast.Expr) ([]TypeReference, error) {
	switch e := exp.(type) {
	case *ast.SelectorExpr:
		return extractTypeParam(pk, e.X)
	case *ast.IndexExpr:
		return []TypeReference{evalTypeReference(pk, e.Index)}, nil
	case *ast.IndexListExpr:

		var ret []TypeReference
		for _, v := range e.Indices {
			ret = append(ret, evalTypeReference(pk, v))
		}
		return ret, nil

	default:
		return nil, fmt.Errorf("not expected expression : %s", types.ExprString(exp))
	}
}

func extractFuncName(exp *ast.CallExpr) (string, error) {
	switch v := exp.Fun.(type) {
	case *ast.IndexExpr:
		return types.ExprString(v.X), nil
	case *ast.IndexListExpr:
		return types.ExprString(v.X), nil
	case *ast.Ident:
		return v.Name, nil
	case *ast.SelectorExpr:
		return types.ExprString(v), nil
	}
	return "", fmt.Errorf("unexpeced func type %T", exp.Fun)
}

func matchFuncName(exp *ast.CallExpr, fnname string) bool {
	n, err := extractFuncName(exp)
	if err != nil {
		return false
	}
	return n == fnname
}

func evalDelegatedBy(pk *packages.Package) func(exp ast.Expr) (DelegateDirective, error) {
	var zero DelegateDirective
	return func(exp ast.Expr) (DelegateDirective, error) {
		switch v := exp.(type) {
		case *ast.CallExpr:
			if !matchFuncName(v, "genfp.DelegatedBy") {
				return zero, fmt.Errorf("not allowed expression. use genfp.DelegatedBy[T](fieldName)")
			}
			arr, err := extractTypeParam(pk, v.Fun)
			if err != nil {
				return zero, err
			}
			if len(arr) != 1 {
				return zero, fmt.Errorf("not allowed expression. use genfp.DelegatedBy[T](fieldName)")
			}

			fn, err := evalStringValue(pk, v.Args[0])
			if err != nil {
				return zero, err
			}
			return DelegateDirective{
				TypeOf: arr[0],
				Field:  fn,
			}, nil

		default:
			return zero, fmt.Errorf("not allowed expression. use genfp.DelegatedBy[T](fieldName)")

		}
	}

}

func evalTypeOf(pk *packages.Package) func(p *packages.Package, exp ast.Expr) (TypeReference, error) {
	var zero TypeReference
	return func(p *packages.Package, exp ast.Expr) (TypeReference, error) {
		switch v := exp.(type) {
		case *ast.CallExpr:
			if !matchFuncName(v, "genfp.TypeOf") {
				return zero, fmt.Errorf("not allowed expression. use genfp.TypeOf[T]")
			}
			arr, err := extractTypeParam(pk, v.Fun)
			if err != nil {
				return zero, err
			}
			if len(arr) != 1 {
				return zero, fmt.Errorf("not allowed expression. use genfp.TypeOf[T]")
			}
			return arr[0], nil

		default:
			return zero, fmt.Errorf("not allowed expression. use genfp.TypeOf[T]")

		}
	}

}

func ParseGenerateAdaptor(lit TaggedLit) (GenerateAdaptorDirective, error) {
	ret := GenerateAdaptorDirective{
		Package: lit.Package,
		Methods: map[string]ImplOptionDirective{},
	}

	if lit.Type.TypeArgs().Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	argType, ok := lit.Type.TypeArgs().At(0).(*types.Named)
	if !ok {
		return ret, fmt.Errorf("target type is not named type : %s", lit.Type.TypeArgs().At(0))
	}

	if _, ok := argType.Underlying().(*types.Interface); !ok {
		return ret, fmt.Errorf("target type is not interface type : %s", lit.Type.TypeArgs().At(0))
	}

	ret.Interface = argType
	intfname := argType.Obj().Name()

	names := []string{"File", "Name", "Extends", "Self", "ImplementsWith", "ExtendsWith", "Embedding", "EmbeddingInterface", "ExtendsByEmbedding", "Delegate", "Getter", "EventHandler", "ValOverride", "ValOverrideUsingPtr", "ZeroReturn", "Options"}
	for idx, e := range lit.Lit.Elts {
		if idx >= len(names) {
			return ret, fmt.Errorf("invalid number of literals")
		}

		name := names[idx]
		name, value := asKeyValue(e, name)
		switch name {
		case "File":
			v, err := evalStringValue(lit.Package, value)
			if err != nil {
				return ret, err
			}
			ret.File = v
		case "Name":
			v, err := evalStringValue(lit.Package, value)
			if err != nil {
				return ret, err
			}
			ret.Name = v
		case "Extends":
			v, err := evalBoolValue(value)
			if err != nil {
				return ret, err
			}
			ret.Extends = v
		case "Self":
			v, err := evalBoolValue(value)
			if err != nil {
				return ret, err
			}
			ret.Self = v
		case "ExtendsSelfCheck":
			v, err := evalBoolValue(value)
			if err != nil {
				return ret, err
			}
			ret.ExtendsSelfCheck = v
		case "ImplementsWith":
			v, err := evalArray(lit.Package, value, evalTypeOf(lit.Package))
			if err != nil {
				return ret, err
			}
			ret.ImplementsWith = v
		case "ExtendsWith":
			v, err := evalMap(lit.Package, value, evalStringValue, evalTypeOf(lit.Package))
			if err != nil {
				return ret, err
			}
			ret.ExtendsWith = v
		case "Embedding":
			v, err := evalArray(lit.Package, value, evalTypeOf(lit.Package))
			if err != nil {
				return ret, err
			}
			ret.Embedding = v
		case "EmbeddingInterface":
			v, err := evalArray(lit.Package, value, evalTypeOf(lit.Package))
			if err != nil {
				return ret, err
			}
			ret.EmbeddingInterface = v
		case "ExtendsByEmbedding":
			v, err := evalBoolValue(value)
			if err != nil {
				return ret, err
			}
			ret.ExtendsByEmbedding = v
		case "Delegate":
			v, err := evalArray(lit.Package, value, evalDelegate(lit.Package))
			if err != nil {
				return ret, err
			}
			ret.Delegate = v
		case "Getter":
			v, err := evalArray(lit.Package, value, evalMethodRef(intfname))
			if err != nil {
				return ret, err
			}
			ret.Getter = seqMap(v, MethodReference.GetName)
		case "EventHandler":
			v, err := evalArray(lit.Package, value, evalMethodRef(intfname))
			if err != nil {
				return ret, err
			}
			ret.EventHandler = seqMap(v, MethodReference.GetName)
		case "ValOverride":
			v, err := evalArray(lit.Package, value, evalMethodRef(intfname))
			if err != nil {
				return ret, err
			}
			ret.ValOverride = seqMap(v, MethodReference.GetName)
		case "ValOverrideUsingPtr":
			v, err := evalArray(lit.Package, value, evalMethodRef(intfname))
			if err != nil {
				return ret, err
			}
			ret.ValOverrideUsingPtr = seqMap(v, MethodReference.GetName)
		case "ZeroReturn":
			v, err := evalArray(lit.Package, value, evalMethodRef(intfname))
			if err != nil {
				return ret, err
			}
			ret.ZeroReturn = seqMap(v, MethodReference.GetName)
		case "Options":
			v, err := evalArray(lit.Package, value, evalImplOption(lit.Package, intfname))
			if err != nil {
				return ret, err
			}
			for _, impl := range v {
				ret.Methods[impl.Method] = impl
			}
		}
	}

	if ret.Self {
		ret.ExtendsSelfCheck = true
	}

	return ret, nil
}

type ImplOptionDirective struct {
	Method                  string
	ReceiverType            TypeReference
	Prefix                  string
	Name                    string
	Private                 bool
	ValOverride             bool
	ValOverrideUsingPtr     bool
	OmitGetterIfValOverride bool
	DefaultImplExpr         ast.Expr
	DefaultImplSignature    *types.Signature
	DefaultImplImports      []ImportPackage
	Delegate                *DelegateDirective

	Type      *types.Func
	Signature *types.Signature
}

// func lookupIdent(pk *packages.Package, typeExpr ast.Expr, pos token.Pos) types.Type {
// 	info := &types.Info{
// 		Types: make(map[ast.Expr]types.TypeAndValue),
// 		Uses:  map[*ast.Ident]types.Object{},
// 	}
// 	types.CheckExpr(pk.Fset, pk.Types, pos, typeExpr, info)

// 	// for k, v := range info.Uses {
// 	// 	fmt.Printf("use = %s, %s\n", k.Name, v.Name())
// 	// }

// 	ti := info.Types[typeExpr]
// 	return ti.Type
// }

func evalConst(pk *packages.Package, constExpr ast.Expr) (constant.Value, []ImportPackage) {
	info := &types.Info{
		Types:     make(map[ast.Expr]types.TypeAndValue),
		Instances: map[*ast.Ident]types.Instance{},
		Defs:      map[*ast.Ident]types.Object{},
		Uses:      map[*ast.Ident]types.Object{},
	}
	types.CheckExpr(pk.Fset, pk.Types, constExpr.End(), constExpr, info)

	var imports []ImportPackage
	for k, v := range info.Uses {
		if pk, ok := v.(*types.PkgName); ok {
			imports = append(imports, ImportPackage{
				Package: pk.Imported().Path(),
				Name:    k.Name,
			})

		}
	}

	ti := info.Types[constExpr]
	return ti.Value, imports
}

func evalFuncLit(pk *packages.Package, typeExpr ast.Expr) (types.Type, []ImportPackage) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Uses:  map[*ast.Ident]types.Object{},
	}
	err := types.CheckExpr(pk.Fset, pk.Types, typeExpr.End(), typeExpr, info)
	if err != nil {
		fmt.Printf("check expr err = %s\n", err)
	}

	var imports []ImportPackage
	for k, v := range info.Uses {
		if pk, ok := v.(*types.PkgName); ok {
			imports = append(imports, ImportPackage{
				Package: pk.Imported().Path(),
				Name:    k.Name,
			})

		}
	}

	ti := info.Types[typeExpr]
	return ti.Type, imports
}

func evalImplOption(pk *packages.Package, intfname string) func(p *packages.Package, e ast.Expr) (ImplOptionDirective, error) {
	return func(p *packages.Package, e ast.Expr) (ImplOptionDirective, error) {
		if lt, ok := e.(*ast.CompositeLit); ok {
			ret := ImplOptionDirective{}
			names := []string{"Method", "Prefix", "Name", "Private", "ValOverride", "OmitGetterIfValOverride", "Delegate", "DefaultImpl"}
			for idx, e := range lt.Elts {
				if idx >= len(names) {
					return ret, fmt.Errorf("invalid number of literals")
				}

				name := names[idx]
				name, value := asKeyValue(e, name)

				switch name {
				case "Method":
					v, err := evalMethodRef(intfname)(p, value)
					if err != nil {
						return ret, err
					}
					ret.Method = v.Name

					ret.ReceiverType = v.Receiver
				case "Prefix":
					v, err := evalStringValue(p, value)
					if err != nil {
						return ret, err
					}
					ret.Prefix = v
				case "Name":
					v, err := evalStringValue(p, value)
					if err != nil {
						return ret, err
					}
					ret.Name = v
				case "Private":
					v, err := evalBoolValue(value)
					if err != nil {
						return ret, err
					}
					ret.Private = v
				case "ValOverride":
					v, err := evalBoolValue(value)
					if err != nil {
						return ret, err
					}
					ret.ValOverride = v
				case "ValOverrideUsingPtr":
					v, err := evalBoolValue(value)
					if err != nil {
						return ret, err
					}
					ret.ValOverrideUsingPtr = v
				case "OmitGetterIfValOverride":
					v, err := evalBoolValue(value)
					if err != nil {
						return ret, err
					}
					ret.OmitGetterIfValOverride = v
				case "Delegate":
					v, err := evalDelegate(p)(p, value)
					if err != nil {
						return ret, err
					}
					ret.Delegate = &v
				case "DefaultImpl":

					found, imports := evalFuncLit(pk, value)
					ret.DefaultImplImports = imports

					ret.DefaultImplExpr = value
					if sig, ok := found.(*types.Signature); ok {
						ret.DefaultImplSignature = sig
					}

				}
			}
			return ret, nil
		}
		return ImplOptionDirective{}, fmt.Errorf("expr is not composite expr : %T", e)
	}

}

type DelegateDirective struct {
	TypeOf TypeReference
	Field  string
}

func evalDelegate(pk *packages.Package) func(p *packages.Package, e ast.Expr) (DelegateDirective, error) {
	return func(p *packages.Package, e ast.Expr) (DelegateDirective, error) {
		var zero DelegateDirective
		if lt, ok := e.(*ast.CompositeLit); ok {
			ret := DelegateDirective{}
			names := []string{"TypeOf", "Field"}
			for idx, e := range lt.Elts {
				if idx >= len(names) {
					return ret, fmt.Errorf("invalid number of literals")
				}

				name := names[idx]
				name, value := asKeyValue(e, name)

				switch name {
				case "TypeOf":
					v, err := evalTypeOf(pk)(pk, value)
					if err != nil {
						return ret, err
					}
					ret.TypeOf = v
				case "Field":
					v, err := evalStringValue(pk, value)
					if err != nil {
						return ret, err
					}
					ret.Field = v

				}
			}
			return ret, nil
		} else if ct, ok := e.(*ast.CallExpr); ok {
			return evalDelegatedBy(pk)(ct)
		}
		return zero, fmt.Errorf("expr is not composite expr : %T", e)
	}

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

type GenerateMonadFunctionsDirective struct {
	Package    *packages.Package
	TargetType *types.Named
	// 생성될 file 이름
	File     string
	TypeParm *types.TypeParam
}

func ParseGenerateMonadFunctions(lit TaggedLit) (GenerateMonadFunctionsDirective, error) {
	ret := GenerateMonadFunctionsDirective{
		Package: lit.Package,
	}

	if lit.Type.TypeArgs().Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	argType, ok := lit.Type.TypeArgs().At(0).(*types.Named)
	if !ok {
		return ret, fmt.Errorf("target type is not named type : %s", lit.Type.TypeArgs().At(0))
	}

	ret.TargetType = argType

	names := []string{"File", "TypeParm"}
	for idx, e := range lit.Lit.Elts {
		if idx >= len(names) {
			return ret, fmt.Errorf("invalid number of literals")
		}
		name := names[idx]
		name, value := asKeyValue(e, name)
		switch name {
		case "File":
			v, err := evalStringValue(lit.Package, value)
			if err != nil {
				return ret, err
			}
			ret.File = v
		case "TypeParm":
			v, err := evalTypeOf(lit.Package)(lit.Package, value)
			if err != nil {
				return ret, err
			}
			if tp, ok := v.Type.(*types.TypeParam); ok {
				ret.TypeParm = tp
			} else {
				return ret, fmt.Errorf("invalid TypeParam. %s is not type param", v.Type)
			}
		}
	}
	return ret, nil
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

type GenerateMonadTransformerDirective struct {
	Name       string
	Package    *packages.Package
	TargetType *types.Named
	MonadType  *types.Named
	// 생성될 file 이름
	File      string
	TypeParm  *types.TypeParam
	Pure      TypeReference
	FlatMap   TypeReference
	Sequence  TypeReference
	Transform []FuncReference
}

func ParseGenerateMonadTransformer(lit TaggedLit) (GenerateMonadTransformerDirective, error) {
	ret := GenerateMonadTransformerDirective{
		Package: lit.Package,
	}

	if lit.Type.TypeArgs().Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	argType, ok := lit.Type.TypeArgs().At(0).(*types.Named)
	if !ok {
		return ret, fmt.Errorf("target type is not named type : %s", lit.Type.TypeArgs().At(0))
	}

	typeArgs := argType.TypeArgs()
	if typeArgs.Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	if monadType, ok := typeArgs.At(0).(*types.Named); ok {
		ret.MonadType = monadType
	} else {
		return ret, fmt.Errorf("target type is not named type : %s", typeArgs.At(0))
	}

	names := []string{"Name", "File", "TypeParm", "Pure", "FlatMap", "Sequence", "Transform"}
	for idx, e := range lit.Lit.Elts {
		if idx >= len(names) {
			return ret, fmt.Errorf("invalid number of literals")
		}
		name := names[idx]
		name, value := asKeyValue(e, name)
		switch name {
		case "Name":
			v, err := evalStringValue(lit.Package, value)
			if err != nil {
				return ret, err
			}
			ret.Name = v
		case "File":
			v, err := evalStringValue(lit.Package, value)
			if err != nil {
				return ret, err
			}
			ret.File = v
		case "TypeParm":
			v, err := evalTypeOf(lit.Package)(lit.Package, value)
			if err != nil {
				return ret, err
			}
			if tp, ok := v.Type.(*types.TypeParam); ok {
				ret.TypeParm = tp
			} else {
				return ret, fmt.Errorf("invalid TypeParam. %s is not type param", v.Type)
			}
		case "Pure":
			found := evalTypeReference(lit.Package, value)
			ret.Pure = found

		case "FlatMap":
			ret.FlatMap = evalTypeReference(lit.Package, value)
		case "Sequence":
			ret.Sequence = evalTypeReference(lit.Package, value)
		case "Transform":
			v, err := evalArray(lit.Package, value, func(p *packages.Package, e ast.Expr) (FuncReference, error) {
				name, err := evalMethodRef("")(p, e)
				if err != nil {
					return FuncReference{}, err
				}
				ret := evalTypeReference(p, e)

				tp, err := extractTypeParam(p, e)
				if err != nil {
					return FuncReference{}, err
				}

				return FuncReference{Name: name.Name, TypeReference: ret, TypeParams: tp}, nil

			})
			if err != nil {
				return ret, err
			}
			ret.Transform = v
		}
	}
	ctx := types.NewContext()

	ins, err := types.Instantiate(ctx, argType.Origin(), []types.Type{ret.TypeParm}, false)

	if err != nil {
		return ret, err
	}

	ret.TargetType = ins.(*types.Named)

	if ret.Name == "" {
		ret.Name = ret.MonadType.Obj().Name() + "T"
	}

	return ret, nil
}
