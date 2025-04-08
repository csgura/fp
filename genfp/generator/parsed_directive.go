package generator

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/types"
	"path"
	"strconv"
	"strings"

	"github.com/csgura/fp/genfp"
	"golang.org/x/tools/go/packages"
)

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
			//fmt.Printf("eval const : %s, %T\n", v.String(), v)

			return constant.StringVal(v), nil
		}
	case *ast.SelectorExpr:
		v, _ := evalConst(p, e)
		if v.Kind() == constant.String {
			return constant.StringVal(v), nil
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

type TypeReference struct {
	Expr       ast.Expr
	StringExpr string
	Type       types.Type
	Imports    []genfp.ImportPackage
}

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

func asFuncRef(p *packages.Package, expr ast.Expr, name string) (FuncReference, error) {
	tr := evalTypeReference(p, expr)
	tp, err := extractTypeParam(p, expr)
	if err != nil {
		return FuncReference{}, err
	}
	return FuncReference{
		TypeReference: tr,
		TypeParams:    tp,
		Name:          name,
	}, nil
}

func getFuncName(p *packages.Package, e ast.Expr) (string, error) {
	switch t := e.(type) {
	case *ast.SelectorExpr:
		switch t.X.(type) {
		case *ast.SelectorExpr:
			return t.Sel.Name, nil
		// p, err := evalSelectorExpr(x)
		// if err != nil {
		// 	return "", err
		// }
		// if strings.HasSuffix(p, tname) {
		// 	return t.Sel.Name, nil
		// }
		// return "", fmt.Errorf("invalid method reference : %s", p)
		case *ast.IndexExpr:
			return t.Sel.Name, nil
		case *ast.Ident:
			return t.Sel.Name, nil
			// if x.Name == tname {
			// 	return t.Sel.Name, nil
			// }
			// return "", fmt.Errorf("invalid method reference : %s", x.Name)
		}
		return "", fmt.Errorf("can't eval %T as func reference. '%s'", t.X, types.ExprString(t.X))
	case *ast.Ident:
		return t.Name, nil
	case *ast.IndexListExpr:
		return getFuncName(p, t.X)
	case *ast.IndexExpr:
		return getFuncName(p, t.X)
	}
	return "", fmt.Errorf("can't eval %T as func reference. '%s'", e, types.ExprString(e))
}

func evalFunctionRef() func(p *packages.Package, e ast.Expr) (FuncReference, error) {
	return func(p *packages.Package, e ast.Expr) (FuncReference, error) {
		name, err := getFuncName(p, e)
		if err != nil {
			return FuncReference{}, err
		}
		return asFuncRef(p, e, name)
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
			return MethodReference{}, fmt.Errorf("can't eval %T as method reference. '%s'", t.X, types.ExprString(t.X))
		case *ast.IndexListExpr:
			return evalMethodRef(tname)(p, t.X)
		case *ast.IndexExpr:
			return evalMethodRef(tname)(p, t.X)
		}
		return MethodReference{}, fmt.Errorf("can't eval %T as method reference. '%s'", e, types.ExprString(e))
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
			i, ok := constant.Int64Val(v)
			if !ok {
				return 0, fmt.Errorf("can't parseInt %s", v.String())
			}
			return int(i), nil
		}
	case *ast.SelectorExpr:
		v, _ := evalConst(p, e)
		if v.Kind() == constant.Int {
			i, ok := constant.Int64Val(v)
			if !ok {
				return 0, fmt.Errorf("can't parseInt %s", v.String())
			}
			return int(i), nil
		}
	}
	return 0, fmt.Errorf("can't eval %T as int", e)
}

func evalImport(p *packages.Package, e ast.Expr) ([]genfp.ImportPackage, error) {
	if lt, ok := e.(*ast.CompositeLit); ok {
		ret := genfp.ImportPackage{}
		names := []string{"Package", "Name"}
		for idx, e := range lt.Elts {
			if idx >= len(names) {
				return nil, fmt.Errorf("invalid number of literals")
			}

			name := names[idx]
			name, value := asKeyValue(e, name)

			switch name {
			case "Package":
				v, err := evalStringValue(p, value)
				if err != nil {
					return nil, err
				}
				ret.Package = v
			case "Name":
				v, err := evalStringValue(p, value)
				if err != nil {
					return nil, err
				}
				ret.Name = v
			}
		}
		return []genfp.ImportPackage{ret}, nil
	} else if ce, ok := e.(*ast.CallExpr); ok {
		if matchFuncName(ce, "genfp.Import") && len(ce.Args) == 1 {
			p, err := evalStringValue(p, ce.Args[0])
			if err != nil {
				return nil, err
			}
			return []genfp.ImportPackage{
				{
					Package: p,
					Name:    path.Base(p),
				},
			}, nil
		} else if matchFuncName(ce, "genfp.PackageOfType") {
			arr, err := extractTypeParam(p, ce.Fun)
			if err != nil {
				return nil, err
			}
			if len(arr) != 1 {
				return nil, fmt.Errorf("not allowed expression. use genfp.PackageOfType[T]")
			}
			return arr[0].Imports, nil
		}
	}
	return nil, fmt.Errorf("expr is not composite expr : %T", e)

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

	} else if ce, ok := e.(*ast.CallExpr); ok {
		if matchFuncName(ce, "seq.Of") || matchFuncName(ce, "slice.Of") {
			var ret []T

			for _, e := range ce.Args {
				v, err := f(p, e)
				if err != nil {
					return nil, err
				}
				ret = append(ret, v)
			}
			return ret, nil
		}
	}
	return nil, fmt.Errorf("expr is not array expr : %T", e)
}

func flatEvalArray[T any](p *packages.Package, e ast.Expr, f func(*packages.Package, ast.Expr) ([]T, error)) ([]T, error) {

	if lt, ok := e.(*ast.CompositeLit); ok {
		var ret []T

		for _, e := range lt.Elts {
			v, err := f(p, e)
			if err != nil {
				return nil, err
			}
			ret = append(ret, v...)
		}
		return ret, nil

	} else if ce, ok := e.(*ast.CallExpr); ok {
		if matchFuncName(ce, "seq.Of") || matchFuncName(ce, "slice.Of") {
			var ret []T

			for _, e := range ce.Args {
				v, err := f(p, e)
				if err != nil {
					return nil, err
				}
				ret = append(ret, v...)
			}
			return ret, nil
		}
	}
	return nil, fmt.Errorf("expr is not array expr : %T", e)
}

func evalImports(p *packages.Package, e ast.Expr) ([]genfp.ImportPackage, error) {

	if ce, ok := e.(*ast.CallExpr); ok {
		if matchFuncName(ce, "genfp.Imports") {
			var ret []genfp.ImportPackage

			for _, e := range ce.Args {
				v, err := evalStringValue(p, e)
				if err != nil {
					return nil, err
				}
				ret = append(ret, genfp.ImportPackage{
					Package: v,
					Name:    path.Base(v),
				})
			}
			return ret, nil
		}
	}
	return flatEvalArray(p, e, evalImport)
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

type GenerateFromUntil struct {
	Decl           TaggedLit
	File           string
	Imports        []genfp.ImportPackage
	From           int
	Until          int
	Parameters     map[string]string
	TypeReferences map[string]TypeReference
	Template       string
}

func ParseGenerateFromUntil(tagged TaggedLit) (GenerateFromUntil, error) {

	lit := tagged.Lit
	ret := GenerateFromUntil{Decl: tagged}

	names := []string{"File", "Imports", "From", "Until", "Parameters", "Template"}
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
			v, err := evalImports(tagged.Package, value)
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
		case "Parameters":
			v, err := evalMap(tagged.Package, value, evalStringValue, evalStringValue)
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.Parameters = v
		case "TypeReferences":
			v, err := evalMap(tagged.Package, value, evalStringValue, evalTypeOf(tagged.Package))
			if err != nil {
				return GenerateFromUntil{}, err
			}
			ret.TypeReferences = v
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

type GenerateFromList struct {
	Decl           TaggedLit
	File           string
	Imports        []genfp.ImportPackage
	List           []string
	Parameters     map[string]string
	TypeReferences map[string]TypeReference
	Template       string
}

func ParseGenerateFromList(tagged TaggedLit) (GenerateFromList, error) {

	lit := tagged.Lit
	ret := GenerateFromList{Decl: tagged}

	names := []string{"File", "Imports", "List", "Parameters", "TypeReferences", "Template"}
	for idx, e := range lit.Elts {
		if idx >= len(names) {
			return GenerateFromList{}, fmt.Errorf("invalid number of literals")
		}

		name := names[idx]
		name, value := asKeyValue(e, name)
		switch name {
		case "File":
			v, err := evalStringValue(tagged.Package, value)
			if err != nil {
				return GenerateFromList{}, err
			}
			ret.File = v
		case "Imports":
			v, err := evalImports(tagged.Package, value)
			if err != nil {
				return GenerateFromList{}, err
			}
			ret.Imports = v
		case "List":
			v, err := evalArray(tagged.Package, value, evalStringValue)
			if err != nil {
				return GenerateFromList{}, err
			}
			ret.List = v
		case "Parameters":
			v, err := evalMap(tagged.Package, value, evalStringValue, evalStringValue)
			if err != nil {
				return GenerateFromList{}, err
			}
			ret.Parameters = v
		case "TypeReferences":
			v, err := evalMap(tagged.Package, value, evalStringValue, evalTypeOf(tagged.Package))
			if err != nil {
				return GenerateFromList{}, err
			}
			ret.TypeReferences = v
		case "Template":
			v, err := evalStringValue(tagged.Package, value)
			if err != nil {
				return GenerateFromList{}, err
			}
			ret.Template = v
		}
	}
	return ret, nil
}

type DelegateDirective struct {
	TypeOf TypeReference
	Field  string
}

type GenerateAdaptorDirective struct {
	Decl                TaggedLit
	Package             genfp.WorkingPackage
	TargetAlias         *types.Alias
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
		return nil, fmt.Errorf("not expected expression : %T, %s", exp, types.ExprString(exp))
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
		Decl:    lit,
		Package: lit.WorkingPackage(),
		Methods: map[string]ImplOptionDirective{},
	}

	if lit.Type.TypeArgs().Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	aliasType, ok := lit.Type.TypeArgs().At(0).(*types.Alias)
	if ok {
		ret.TargetAlias = aliasType
	}

	argType, ok := types.Unalias(lit.Type.TypeArgs().At(0)).(*types.Named)
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
	DefaultImplImports      []genfp.ImportPackage
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

func evalConst(pk *packages.Package, constExpr ast.Expr) (constant.Value, []genfp.ImportPackage) {
	info := &types.Info{
		Types:     make(map[ast.Expr]types.TypeAndValue),
		Instances: map[*ast.Ident]types.Instance{},
		Defs:      map[*ast.Ident]types.Object{},
		Uses:      map[*ast.Ident]types.Object{},
	}
	types.CheckExpr(pk.Fset, pk.Types, constExpr.End(), constExpr, info)

	var imports []genfp.ImportPackage
	for k, v := range info.Uses {
		if pk, ok := v.(*types.PkgName); ok {
			imports = append(imports, genfp.ImportPackage{
				Package: pk.Imported().Path(),
				Name:    k.Name,
			})

		}
	}

	ti := info.Types[constExpr]
	return ti.Value, imports
}

func evalFuncLit(pk *packages.Package, typeExpr ast.Expr) (types.Type, []genfp.ImportPackage) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Uses:  map[*ast.Ident]types.Object{},
	}
	err := types.CheckExpr(pk.Fset, pk.Types, typeExpr.End(), typeExpr, info)
	if err != nil {
		// TODO: Hello.World 같은 경우 receiver type 을  eval 하는데
		// seq.Sort 처럼  package 이름이 있는 경우,  여기서 에러가 출력됨.
		fmt.Printf("check expr '%s' err = %s\n", types.ExprString(typeExpr), err)
	}

	var imports []genfp.ImportPackage
	for k, v := range info.Uses {
		if pk, ok := v.(*types.PkgName); ok {
			imports = append(imports, genfp.ImportPackage{
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

type GenerateMonadFunctionsDirective struct {
	Decl       TaggedLit
	Package    genfp.WorkingPackage
	TargetType GenericType
	// 생성될 file 이름
	File     string
	TypeParm *types.TypeParam
}

func ParseGenerateMonadFunctions(lit TaggedLit) (GenerateMonadFunctionsDirective, error) {
	ret := GenerateMonadFunctionsDirective{
		Decl:    lit,
		Package: lit.WorkingPackage(),
	}

	if lit.Type.TypeArgs().Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	aliasType, ok := lit.Type.TypeArgs().At(0).(*types.Alias)
	if ok {
		ret.TargetType = aliasType
	} else {

		argType, ok := types.Unalias(lit.Type.TypeArgs().At(0)).(*types.Named)
		if !ok {
			return ret, fmt.Errorf("target type is not named type : %s", lit.Type.TypeArgs().At(0))
		}
		ret.TargetType = argType

	}

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

type MonadFunctions struct {
	Pure    TypeReference
	FlatMap TypeReference
}

type GenerateMonadTransformerDirective struct {
	Decl              TaggedLit
	Name              string
	Package           genfp.WorkingPackage
	TargetAlias       *types.Alias
	TargetType        *types.Named
	ExposureMonadType GenericType
	// 생성될 file 이름
	File          string
	TypeParm      *types.TypeParam
	GivenMonad    MonadFunctions
	ExposureMonad MonadFunctions
	Sequence      TypeReference
	Transform     []FuncReference
}

func ParseGenerateMonadTransformer(lit TaggedLit) (GenerateMonadTransformerDirective, error) {
	ret := GenerateMonadTransformerDirective{
		Decl:    lit,
		Package: genfp.NewWorkingPackage(lit.Package.Types, lit.Package.Fset, lit.Package.Syntax),
	}

	if lit.Type.TypeArgs().Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	aliasType, ok := lit.Type.TypeArgs().At(0).(*types.Alias)
	if ok {
		ret.TargetAlias = aliasType
	}

	argType, ok := types.Unalias(lit.Type.TypeArgs().At(0)).(*types.Named)
	if !ok {
		return ret, fmt.Errorf("target type is not named type : %s", lit.Type.TypeArgs().At(0))
	}

	typeArgs := argType.TypeArgs()
	if typeArgs.Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	if monadType, ok := typeArgs.At(0).(GenericType); ok {
		ret.ExposureMonadType = monadType
	} else {
		return ret, fmt.Errorf("inside type is not named type : %s", typeArgs.At(0))
	}

	names := []string{"Name", "File", "TypeParm", "GivenMonad", "ExposureMonad", "Sequence", "Transform"}
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
		case "GivenMonad":
			v, err := evalMonadFunctions(lit.Package, value)
			if err != nil {
				return ret, err
			}
			ret.GivenMonad = v
		case "ExposureMonad":
			v, err := evalMonadFunctions(lit.Package, value)
			if err != nil {
				return ret, err
			}
			ret.ExposureMonad = v
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
		ret.Name = ret.ExposureMonadType.Obj().Name() + ret.TargetType.Obj().Name()[:1]
	}

	return ret, nil
}

func evalMonadFunctions(p *packages.Package, e ast.Expr) (MonadFunctions, error) {
	if lt, ok := e.(*ast.CompositeLit); ok {
		ret := MonadFunctions{}
		names := []string{"Pure", "FlatMap"}
		for idx, e := range lt.Elts {
			if idx >= len(names) {
				return MonadFunctions{}, fmt.Errorf("invalid number of literals")
			}

			name := names[idx]
			name, value := asKeyValue(e, name)

			switch name {
			case "Pure":
				found := evalTypeReference(p, value)
				ret.Pure = found

			case "FlatMap":
				ret.FlatMap = evalTypeReference(p, value)
			}
		}
		return ret, nil
	}
	return MonadFunctions{}, fmt.Errorf("expr is not composite expr : %T", e)

}

type GenerateFromStructs struct {
	Decl           TaggedLit
	File           string
	Imports        []genfp.ImportPackage
	List           []TypeReference
	Recursive      bool
	Parameters     map[string]string
	TypeReferences map[string]TypeReference
	Template       string
}

func ParseGenerateFromStructs(tagged TaggedLit) (GenerateFromStructs, error) {

	lit := tagged.Lit
	ret := GenerateFromStructs{Decl: tagged}

	names := []string{"File", "Imports", "List", "Recursive", "Parameters", "TypeReferences", "Template"}
	for idx, e := range lit.Elts {
		if idx >= len(names) {
			return GenerateFromStructs{}, fmt.Errorf("invalid number of literals")
		}

		name := names[idx]
		name, value := asKeyValue(e, name)
		switch name {
		case "File":
			v, err := evalStringValue(tagged.Package, value)
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.File = v
		case "Imports":
			v, err := evalImports(tagged.Package, value)
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.Imports = v
		case "List":
			v, err := evalArray(tagged.Package, value, evalTypeOf(tagged.Package))
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.List = v
		case "Recursive":
			v, err := evalBoolValue(value)
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.Recursive = v
		case "Parameters":
			v, err := evalMap(tagged.Package, value, evalStringValue, evalStringValue)
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.Parameters = v
		case "TypeReferences":
			v, err := evalMap(tagged.Package, value, evalStringValue, evalTypeOf(tagged.Package))
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.TypeReferences = v
		case "Template":
			v, err := evalStringValue(tagged.Package, value)
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.Template = v
		}
	}

	return ret, nil

}

func ParseGenerateFromInterfaces(tagged TaggedLit) (GenerateFromStructs, error) {

	lit := tagged.Lit
	ret := GenerateFromStructs{Decl: tagged}

	names := []string{"File", "Imports", "List", "Parameters", "TypeReferences", "Template"}
	for idx, e := range lit.Elts {
		if idx >= len(names) {
			return GenerateFromStructs{}, fmt.Errorf("invalid number of literals")
		}

		name := names[idx]
		name, value := asKeyValue(e, name)
		switch name {
		case "File":
			v, err := evalStringValue(tagged.Package, value)
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.File = v
		case "Imports":
			v, err := evalImports(tagged.Package, value)
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.Imports = v
		case "List":
			v, err := evalArray(tagged.Package, value, evalTypeOf(tagged.Package))
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.List = v
		case "Parameters":
			v, err := evalMap(tagged.Package, value, evalStringValue, evalStringValue)
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.Parameters = v
		case "TypeReferences":
			v, err := evalMap(tagged.Package, value, evalStringValue, evalTypeOf(tagged.Package))
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.TypeReferences = v
		case "Template":
			v, err := evalStringValue(tagged.Package, value)
			if err != nil {
				return GenerateFromStructs{}, err
			}
			ret.Template = v
		}
	}

	return ret, nil

}

type GenerateApplicative struct {
	Decl        TaggedLit
	Package     genfp.WorkingPackage
	TargetAlias *types.Alias
	TargetType  *types.Named
	// 생성될 file 이름
	File      string
	TypeParam *types.TypeParam
	Mapper    []Mapping
}

type Mapping struct {
	Prefix string
	Mapper FuncReference
}

func ParseGenerateApplicative(lit TaggedLit) (GenerateApplicative, error) {
	ret := GenerateApplicative{
		Decl:    lit,
		Package: genfp.NewWorkingPackage(lit.Package.Types, lit.Package.Fset, lit.Package.Syntax),
	}

	if lit.Type.TypeArgs().Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	aliasType, ok := lit.Type.TypeArgs().At(0).(*types.Alias)
	if ok {
		ret.TargetAlias = aliasType
	}

	argType, ok := types.Unalias(lit.Type.TypeArgs().At(0)).(*types.Named)
	if !ok {
		return ret, fmt.Errorf("target type is not named type : %s", lit.Type.TypeArgs().At(0))
	}

	typeArgs := argType.TypeArgs()
	if typeArgs.Len() != 1 {
		return ret, fmt.Errorf("invalid number of type argument")
	}

	names := []string{"File", "TypeParm", "Mapper"}
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
				ret.TypeParam = tp
			} else {
				return ret, fmt.Errorf("invalid TypeParam. %s is not type param", v.Type)
			}
		case "Mapper":
			v, err := evalArray(lit.Package, value, evalMapping(lit.Package))
			if err != nil {
				return ret, err
			}
			ret.Mapper = v
		}
	}
	ctx := types.NewContext()

	ins, err := types.Instantiate(ctx, argType.Origin(), []types.Type{ret.TypeParam}, false)

	if err != nil {
		return ret, err
	}

	ret.TargetType = ins.(*types.Named)

	return ret, nil
}

func evalMapping(pk *packages.Package) func(p *packages.Package, e ast.Expr) (Mapping, error) {
	return func(p *packages.Package, e ast.Expr) (Mapping, error) {
		if lt, ok := e.(*ast.CompositeLit); ok {
			ret := Mapping{}
			names := []string{"Prefix", "Mapper"}
			for idx, e := range lt.Elts {
				if idx >= len(names) {
					return ret, fmt.Errorf("invalid number of literals")
				}

				name := names[idx]
				name, value := asKeyValue(e, name)

				switch name {

				case "Prefix":
					v, err := evalStringValue(p, value)
					if err != nil {
						return ret, err
					}
					ret.Prefix = v
				case "Mapper":
					m, err := evalFunctionRef()(pk, value)
					if err != nil {
						return ret, err
					}
					ret.Mapper = m
				}
			}
			return ret, nil
		}
		return Mapping{}, fmt.Errorf("expr is not composite expr : %T", e)
	}

}
