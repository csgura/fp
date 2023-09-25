package genfp

import (
	"fmt"
	"go/ast"
	"go/token"
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

func DelegateOf[T any]() Delegate {
	return Delegate{
		TypeOf: TypeOf[T](),
		Field:  "",
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
	File           string
	Name           string
	Extends        bool
	Self           bool
	ImplementsWith []TypeTag
	ExtendsWith    map[TypeTag]string
	Embedding      []TypeTag
	Delegate       []TypeTag
	Getter         []any
	EventHandler   []any
	ValOverride    []any
	ZeroReturn     []any
	Options        []ImplOption
}

type AdaptorMethods []ImplOption

type ImplOption struct {
	Method                  any
	Prefix                  string
	Name                    string
	Private                 bool
	ValOverride             bool
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

func evalStringValue(e ast.Expr) (string, error) {
	switch t := e.(type) {
	case *ast.BasicLit:
		if strings.HasPrefix(t.Value, `"`) && strings.HasSuffix(t.Value, `"`) {
			return t.Value[1 : len(t.Value)-1], nil
		} else if strings.HasPrefix(t.Value, "`") && strings.HasSuffix(t.Value, "`") {
			return t.Value[1 : len(t.Value)-1], nil
		}
	}
	return "", fmt.Errorf("can't eval %T as string", e)
}

func evalSelectorExpr(e ast.Expr) (string, error) {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name, nil
	case *ast.SelectorExpr:
		switch x := t.X.(type) {
		case *ast.SelectorExpr:
			p, err := evalSelectorExpr(x)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("%s.%s", p, t.Sel.Name), nil
		case *ast.Ident:
			return fmt.Sprintf("%s.%s", x.Name, t.Sel.Name), nil
		}
	}
	return "", fmt.Errorf("can't eval %T as method reference", e)
}

func evalMethodRef(tname string) func(e ast.Expr) (string, error) {
	return func(e ast.Expr) (string, error) {
		switch t := e.(type) {
		case *ast.SelectorExpr:
			switch x := t.X.(type) {
			case *ast.SelectorExpr:
				p, err := evalSelectorExpr(x)
				if err != nil {
					return "", err
				}
				if strings.HasSuffix(p, tname) {
					return t.Sel.Name, nil
				}
				return "", fmt.Errorf("invalid method reference : %s", p)
			case *ast.Ident:
				if x.Name == tname {
					return t.Sel.Name, nil
				}
				return "", fmt.Errorf("invalid method reference : %s", x.Name)
			}
		}
		return "", fmt.Errorf("can't eval %T as method reference", e)
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

func evalIntValue(e ast.Expr) (int, error) {
	switch t := e.(type) {
	case *ast.BasicLit:
		i, err := strconv.ParseInt(t.Value, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("can't parseInt %s", t.Value)
		}
		return int(i), nil
	case *ast.SelectorExpr:
		if matchSelExpr(t, []string{"genfp", "MaxProduct"}) {
			return MaxProduct, nil
		}
		if matchSelExpr(t, []string{"genfp", "MaxFunc"}) {
			return MaxFunc, nil
		}

		if matchSelExpr(t, []string{"genfp", "MaxCompose"}) {
			return MaxCompose, nil
		}

		if matchSelExpr(t, []string{"genfp", "MaxShift"}) {
			return MaxShift, nil
		}

		if matchSelExpr(t, []string{"genfp", "MaxFlip"}) {
			return MaxFlip, nil
		}
	}
	return 0, fmt.Errorf("can't eval %T as int", e)
}

func evalImport(e ast.Expr) (ImportPackage, error) {
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
				v, err := evalStringValue(value)
				if err != nil {
					return ImportPackage{}, err
				}
				ret.Package = v
			case "Name":
				v, err := evalStringValue(value)
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

func evalArray[T any](e ast.Expr, f func(ast.Expr) (T, error)) ([]T, error) {

	if lt, ok := e.(*ast.CompositeLit); ok {
		var ret []T

		for _, e := range lt.Elts {
			v, err := f(e)
			if err != nil {
				return nil, err
			}
			ret = append(ret, v)
		}
		return ret, nil

	}
	return nil, fmt.Errorf("expr is not array expr : %T", e)
}

func evalMap[K comparable, V any](e ast.Expr, kf func(ast.Expr) (K, error), vf func(ast.Expr) (V, error)) (map[K]V, error) {

	if lt, ok := e.(*ast.CompositeLit); ok {
		ret := map[K]V{}

		for _, e := range lt.Elts {
			if kve, ok := e.(*ast.KeyValueExpr); ok {
				kv, err := kf(kve.Key)
				if err != nil {
					return nil, err
				}

				vv, err := vf(kve.Value)
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

func ParseGenerateFromUntil(lit *ast.CompositeLit) (GenerateFromUntil, error) {

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
			v, err := evalStringValue(value)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.File = v
		case "Imports":
			v, err := evalArray(value, evalImport)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.Imports = v
		case "From":
			v, err := evalIntValue(value)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.From = v
		case "Until":
			v, err := evalIntValue(value)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.Until = v
		case "Template":
			v, err := evalStringValue(value)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.Template = v
		}
	}

	return ret, nil

}

type GenerateAdaptorDirective struct {
	Package        *packages.Package
	Interface      *types.Named
	File           string
	Name           string
	Extends        bool
	Self           bool
	ImplementsWith []TypeReference
	Embedding      []TypeReference
	Delegate       []TypeReference
	Getter         []string
	EventHandler   []string
	ValOverride    []string
	ZeroReturn     []string
	Methods        map[string]ImplOptionDirective
}

type TypeReference struct {
	Expr       ast.Expr
	StringExpr string
	Type       types.Type
	Imports    []ImportPackage
}

func evalTypeReference(pk *packages.Package, exp ast.Expr) TypeReference {
	t, imp := evalFuncLit(pk, exp, exp.Pos())

	return TypeReference{
		Expr:       exp,
		StringExpr: types.ExprString(exp),
		Type:       t,
		Imports:    imp,
	}
}
func extractTypeParam(pk *packages.Package, exp ast.Expr) ([]TypeReference, error) {
	switch e := exp.(type) {
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

func evalTypeOf(pk *packages.Package) func(exp ast.Expr) (TypeReference, error) {
	var zero TypeReference
	return func(exp ast.Expr) (TypeReference, error) {
		switch v := exp.(type) {
		case *ast.CallExpr:
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

	names := []string{"File", "Name", "Extends", "Self", "ImplementsWith", "Embedding", "Delegate", "Getter", "EventHandler", "ValOverride", "ZeroReturn", "Options"}
	for idx, e := range lit.Lit.Elts {
		if idx >= len(names) {
			return ret, fmt.Errorf("invalid number of literals")
		}

		name := names[idx]
		name, value := asKeyValue(e, name)
		switch name {
		case "File":
			v, err := evalStringValue(value)
			if err != nil {
				return ret, err
			}
			ret.File = v
		case "Name":
			v, err := evalStringValue(value)
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
		case "ImplementsWith":
			v, err := evalArray(value, evalTypeOf(lit.Package))
			if err != nil {
				return ret, err
			}
			ret.ImplementsWith = v
		case "Embedding":
			v, err := evalArray(value, evalTypeOf(lit.Package))
			if err != nil {
				return ret, err
			}
			ret.Embedding = v
		case "Delegate":
			v, err := evalArray(value, evalTypeOf(lit.Package))
			if err != nil {
				return ret, err
			}
			ret.Delegate = v
		case "Getter":
			v, err := evalArray(value, evalMethodRef(intfname))
			if err != nil {
				return ret, err
			}
			ret.Getter = v
		case "EventHandler":
			v, err := evalArray(value, evalMethodRef(intfname))
			if err != nil {
				return ret, err
			}
			ret.EventHandler = v
		case "ValOverride":
			v, err := evalArray(value, evalMethodRef(intfname))
			if err != nil {
				return ret, err
			}
			ret.ValOverride = v
		case "ZeroReturn":
			v, err := evalArray(value, evalMethodRef(intfname))
			if err != nil {
				return ret, err
			}
			ret.ZeroReturn = v
		case "Options":
			v, err := evalArray(value, evalImplOption(lit.Package, intfname))
			if err != nil {
				return ret, err
			}
			for _, impl := range v {
				ret.Methods[impl.Method] = impl
			}
		}
	}

	return ret, nil
}

type ImplOptionDirective struct {
	Method                  string
	Prefix                  string
	Name                    string
	Private                 bool
	ValOverride             bool
	OmitGetterIfValOverride bool
	DefaultImplExpr         ast.Expr
	DefaultImplSignature    *types.Signature
	DefaultImplImports      []ImportPackage

	Type      *types.Func
	Signature *types.Signature
}

func lookupIdent(pk *packages.Package, typeExpr ast.Expr, pos token.Pos) types.Type {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Uses:  map[*ast.Ident]types.Object{},
	}
	types.CheckExpr(pk.Fset, pk.Types, pos, typeExpr, info)

	// for k, v := range info.Uses {
	// 	fmt.Printf("use = %s, %s\n", k.Name, v.Name())
	// }

	ti := info.Types[typeExpr]
	return ti.Type
}

func evalFuncLit(pk *packages.Package, typeExpr ast.Expr, pos token.Pos) (types.Type, []ImportPackage) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Uses:  map[*ast.Ident]types.Object{},
	}
	types.CheckExpr(pk.Fset, pk.Types, pos, typeExpr, info)

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

func evalImplOption(pk *packages.Package, intfname string) func(e ast.Expr) (ImplOptionDirective, error) {
	return func(e ast.Expr) (ImplOptionDirective, error) {
		if lt, ok := e.(*ast.CompositeLit); ok {
			ret := ImplOptionDirective{}
			names := []string{"Method", "Prefix", "Name", "Private", "ValOverride", "OmitGetterIfValOverride", "DefaultImpl"}
			for idx, e := range lt.Elts {
				if idx >= len(names) {
					return ret, fmt.Errorf("invalid number of literals")
				}

				name := names[idx]
				name, value := asKeyValue(e, name)

				switch name {
				case "Method":
					v, err := evalMethodRef(intfname)(value)
					if err != nil {
						return ret, err
					}
					ret.Method = v
				case "Prefix":
					v, err := evalStringValue(value)
					if err != nil {
						return ret, err
					}
					ret.Prefix = v
				case "Name":
					v, err := evalStringValue(value)
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
				case "OmitGetterIfValOverride":
					v, err := evalBoolValue(value)
					if err != nil {
						return ret, err
					}
					ret.OmitGetterIfValOverride = v
				case "DefaultImpl":

					found, imports := evalFuncLit(pk, value, value.Pos())
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
