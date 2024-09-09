package genfp

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/csgura/fp/internal/max"
)

const MaxProduct = max.Product

const MaxFunc = max.Func

const MaxCompose = max.Compose

const MaxShift = max.Shift
const MaxFlip = max.Flip

var OrdinalName = []string{
	"Zero",
	"First",
	"Second",
	"Third",
	"Fourth",
	"Fifth",
	"Sixth",
	"Seventh",
	"Eighth",
	"Ninth",
	"Tenth",
}

var defaultFunc = map[string]any{
	"FormatStr":         FormatStr,
	"FuncChain":         FuncChain,
	"ConsType":          ConsType,
	"ReversConsType":    ReversConsType,
	"TypeArgs":          FuncTypeArgs,
	"DeclArgs":          FuncDeclArgs,
	"CallArgs":          FuncCallArgs,
	"ReverseCallArgs":   ReverseFuncCallArgs,
	"DeclTypeClassArgs": FuncDeclTypeClassArgs,
	"CurriedCallArgs":   CurriedCallArgs,
	"TypeClassArgs":     TypeClassArgs,
	"CurriedFunc":       CurriedType,
	"RecursiveType":     RecursiveType,

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
	"dec": func(n int) int {
		return n - 1
	},
	"inc": func(n int) int {
		return n + 1
	},
}

type ArgsRange struct {
	prefix string
	start  int
	until  int
}

func (r ArgsRange) Dot(expr string) string {
	return Args(r.prefix).Dot(r.start, r.until, expr)
}

func FuncDecl(prefix string, start, until int, ret string) string {
	return fmt.Sprintf("func(%s) %s", TypeArgs(prefix, start, until), ret)
}

func FormatStr(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ",")
		}
		fmt.Fprintf(f, "%s", "%v")
	}
	return f.String()
}

func TypeArgs(prefix string, start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s%d", prefix, j)
	}
	return f.String()
}

func FuncTypeArgs(start, until int, prefix ...string) string {
	if len(prefix) > 0 {
		return TypeArgs(prefix[0], start, until)
	}
	return TypeArgs("A", start, until)
}

func FuncChain(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		if j == until {
			fmt.Fprintf(f, "f%d Func1[A%d,R]", j, j)
		} else {
			fmt.Fprintf(f, "f%d Func1[A%d,A%d]", j, j, j+1)
		}
	}
	return f.String()
}

type Monad string

func (r Monad) ConsType(start, until int, last string) string {
	ret := last
	for j := until; j >= start; j-- {
		ret = fmt.Sprintf("hlist.Cons[%s[A%d], %s]", r, j, ret)
	}
	return ret
}

func (r Monad) TypeDeclArgs(start, until int, prefixOpt ...string) string {

	prefix := "A"
	if len(prefixOpt) > 0 {
		prefix = prefixOpt[0]
	}

	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s[%s%d]", r, prefix, j)
	}
	return f.String()
}

func (r Monad) FuncChain(start, until int, funcType ...string) string {
	ft := "fp.Func1"
	if len(funcType) > 0 {
		ft = funcType[0]
	}

	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		if j == until {
			fmt.Fprintf(f, "f%d %s[A%d,%s[R]]", j, ft, j, r)
		} else {
			fmt.Fprintf(f, "f%d %s[A%d,%s[A%d]]", j, ft, j, r, j+1)
		}
	}
	return f.String()
}

type Args string

func (r Args) Call(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s%d", r, j)
	}
	return f.String()
}

func (r Args) Dot(start, until int, expr string) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s%d.%s", r, j, expr)
	}
	return f.String()
}

type importSet struct {
	PathToName map[string]importAlias
	NameToPath map[string]importAlias
}

type writer struct {
	Package string
	Buffer  *bytes.Buffer
	importSet
}

func (r *writer) Write(b []byte) (int, error) {
	return r.Buffer.Write(b)
}

func (r *writer) Render(templateStr string, funcs map[string]any, param map[string]any) {
	if param == nil {
		param = map[string]any{}
	}

	tpl, err := template.New("write").Funcs(defaultFunc).Funcs(funcs).Parse(templateStr)
	if err == nil {
		param["N"] = 1
		err := tpl.Execute(r, param)
		if err != nil {
			fmt.Printf("template = %s\n", templateStr)
			panic(err)
		}

	} else {
		fmt.Printf("template = %s\n", templateStr)
		panic(err)
	}
}

func (r *writer) Iteration(start, end int) Range {
	return Range{r, start, end}
}

type importAlias struct {
	path    string
	alias   string
	isalias bool
}

func (r *importSet) AddImport(p PackageId) bool {
	alias := p.Alias()

	_, ok := r.NameToPath[alias]
	if ok {
		return false
	}

	parr := strings.Split(p.Path(), "/")
	plast := parr[len(parr)-1]

	ia := importAlias{
		path:    p.Path(),
		alias:   alias,
		isalias: plast != alias,
	}

	r.PathToName[p.Path()] = ia
	r.NameToPath[alias] = ia

	return true
}

func (r *importSet) GetImportedName(p PackageId) string {
	ret, ok := r.PathToName[p.Path()]
	if ok {
		return ret.alias
	}

	i := 1
	alias := p.Alias()

	for {
		added := r.AddImport(p)
		if added {
			return alias
		}

		alias = fmt.Sprintf("%s%d", p.Alias(), i)
		p = NewImportPackage(p.Path(), alias)
		i++
	}
}

func NewImportSet() ImportSet {
	return &importSet{PathToName: map[string]importAlias{}, NameToPath: map[string]importAlias{}}
}

func (r *writer) ImportList() []string {
	var ret = []string{}

	for _, v := range r.NameToPath {
		if !v.isalias {
			ret = append(ret, fmt.Sprintf(`"%s"`, v.path))

		} else {
			ret = append(ret, fmt.Sprintf(`%s "%s"`, v.alias, v.path))
		}
	}

	return ret
}

func (r *importSet) ZeroExpr(pk WorkingPackage, tpe types.Type) string {
	switch realtp := tpe.(type) {
	case *types.Named:
		if _, ok := realtp.Underlying().(*types.Interface); ok {
			return "nil"
		}
		return fmt.Sprintf("%s{}", r.TypeName(pk, tpe))
	case *types.Array:
		return fmt.Sprintf("%s{}", r.TypeName(pk, tpe))
	case *types.Map:
		return fmt.Sprintf("%s{}", r.TypeName(pk, tpe))
	case *types.Slice:
		return "nil"
	case *types.Pointer:
		return "nil"
	case *types.Chan:
		return "nil"
	case *types.Signature:
		return "nil"
	case *types.Struct:
		return fmt.Sprintf("%s{}", r.TypeName(pk, tpe))
	case *types.Interface:
		return "nil"
	case *types.Basic:

		switch realtp.Info() {
		case types.IsBoolean:
			return "false"
		case types.IsInteger:
			return "0"
		case types.IsUnsigned:
			return "0"
		case types.IsFloat:
			return "0"
		case types.IsComplex:
			return "0"
		case types.IsString:
			return `""`
		}
	}

	return tpe.String()
}

func (r *importSet) TypeName(pk WorkingPackage, tpe types.Type) string {
	//fmt.Printf("type %s %T\n", tpe.String(), tpe)
	switch realtp := tpe.(type) {
	case *types.Basic:
		if realtp.Kind() == types.UnsafePointer {
			r.AddImport(NewImportPackage("unsafe", "unsafe"))
		}
		return tpe.String()
	case *types.Named:
		tpname := realtp.Origin().Obj().Name()
		nameWithPkg := tpname
		if realtp.Obj().Pkg() != nil && realtp.Obj().Pkg().Path() != pk.Path() {
			alias := r.GetImportedName(FromTypesPackage(realtp.Obj().Pkg()))

			nameWithPkg = fmt.Sprintf("%s.%s", alias, tpname)
		}

		if realtp.TypeArgs() != nil {
			args := []string{}
			for i := 0; i < realtp.TypeArgs().Len(); i++ {
				args = append(args, r.TypeName(pk, realtp.TypeArgs().At(i)))
			}

			argsstr := strings.Join(args, ",")

			return fmt.Sprintf("%s[%s]", nameWithPkg, argsstr)
		} else {

			return nameWithPkg

		}

	case *types.Array:
		elemType := r.TypeName(pk, realtp.Elem())
		return fmt.Sprintf("[%d]%s", realtp.Len(), elemType)

	case *types.Map:
		keyType := r.TypeName(pk, realtp.Key())

		elemType := r.TypeName(pk, realtp.Elem())
		return fmt.Sprintf("map[%s]%s", keyType, elemType)
	case *types.Slice:
		elemType := r.TypeName(pk, realtp.Elem())
		return "[]" + elemType
	case *types.Pointer:
		elemType := r.TypeName(pk, realtp.Elem())
		return "*" + elemType
	case *types.Chan:
		elemType := r.TypeName(pk, realtp.Elem())
		switch realtp.Dir() {
		case types.RecvOnly:
			return "<-chan " + elemType
		case types.SendOnly:
			return "chan<- " + elemType
		default:
			return "chan " + elemType

		}
	case *types.Signature:
		argsstr := iterate(realtp.Params().Len(), realtp.Params().At, func(idx int, v *types.Var) string {
			return v.Name() + " " + r.TypeName(pk, v.Type())
		})

		resultstr := iterate(realtp.Results().Len(), realtp.Results().At, func(idx int, v *types.Var) string {
			return v.Name() + " " + r.TypeName(pk, v.Type())
		})

		return fmt.Sprintf("func (%s) (%s)", strings.Join(argsstr, ","), strings.Join(resultstr, ","))
	case *types.Struct:
		fields := iterate(realtp.NumFields(), realtp.Field, func(idx int, v *types.Var) string {
			if v.Embedded() {
				return fmt.Sprintf("%s %s",
					r.TypeName(pk, v.Type()),
					realtp.Tag(idx),
				)
			}
			return fmt.Sprintf("%s %s %s",
				v.Name(),
				r.TypeName(pk, v.Type()),
				realtp.Tag(idx),
			)
		})
		return fmt.Sprintf(`struct {
			%s
		}`, strings.Join(fields, "\n"))
	case *types.Interface:
		if realtp.NumMethods() == 0 {
			return "any"
		}
		embeded := iterate(realtp.NumEmbeddeds(), realtp.EmbeddedType, func(idx int, v types.Type) string {
			return r.TypeName(pk, realtp.EmbeddedType(idx))
		})

		fields := iterate(realtp.NumExplicitMethods(), realtp.ExplicitMethod, func(idx int, v *types.Func) string {
			m := realtp.ExplicitMethod(idx)

			return fmt.Sprintf("%s%s", m.Name(), r.TypeName(pk, m.Type())[4:])

		})
		return fmt.Sprintf(`interface {
			%s
			%s
		}`, strings.Join(embeded, "\n"), strings.Join(fields, "\n"))
	}

	return tpe.String()
}

type ImportPackage struct {
	Package string
	Name    string
}

func (r ImportPackage) Path() string {
	return r.Package
}

func (r ImportPackage) Alias() string {
	return r.Name
}

func FromTypesPackage(pk *types.Package) ImportPackage {
	if pk == nil {
		return ImportPackage{}
	}
	return ImportPackage{
		Package: pk.Path(),
		Name:    pk.Name(),
	}
}

func NewImportPackage(pkg string, name string) ImportPackage {
	return ImportPackage{
		Package: pkg,
		Name:    name,
	}
}

type WorkingPackage interface {
	Path() string
	Alias() string
	Scope() *types.Scope
	Package() *types.Package
}

func NewWorkingPackage(pk *types.Package, fset *token.FileSet, syntax []*ast.File) WorkingPackage {

	fmt.Printf("new working package fset = \n")

	return &workingPackage{
		currentPackage: pk,
		fset:           fset,
		syntax:         syntax,
	}
}

type workingPackage struct {
	fset           *token.FileSet
	syntax         []*ast.File
	currentPackage *types.Package
}

// Scope implements WorkingPackage.
func (w *workingPackage) Scope() *types.Scope {
	if w == nil || w.currentPackage == nil {
		return nil
	}

	return w.currentPackage.Scope()
}

// Name implements WorkingPackage.
func (w *workingPackage) Alias() string {
	if w == nil || w.currentPackage == nil {
		return ""
	}
	return w.currentPackage.Name()
}

// Path implements WorkingPackage.
func (w *workingPackage) Path() string {
	if w == nil || w.currentPackage == nil {
		return ""
	}
	return w.currentPackage.Path()
}

func (w *workingPackage) Package() *types.Package {
	if w == nil || w.currentPackage == nil {
		return nil
	}
	return w.currentPackage
}

var _ WorkingPackage = &workingPackage{}

type PackageId interface {
	Path() string
	Alias() string
}
type ImportSet interface {
	AddImport(p PackageId) bool
	GetImportedName(p PackageId) string
	TypeName(pk WorkingPackage, tpe types.Type) string
	ZeroExpr(pk WorkingPackage, tpe types.Type) string
}

type Writer interface {
	io.Writer
	ImportSet
	Iteration(start, end int) Range
	Render(template string, funcs map[string]any, param map[string]any)
}

func Generate(packname string, filename string, writeFunc func(w Writer)) {

	cmdName := path.Base(os.Args[0])
	if filename != "" {
		fmt.Printf("%s generate %s", cmdName, filename)
		fmt.Println()
		os.Remove(filename)
	}

	f := &writer{packname, &bytes.Buffer{}, importSet{PathToName: map[string]importAlias{}, NameToPath: map[string]importAlias{}}}

	writeFunc(f)

	if f.Buffer.Len() == 0 {
		return
	}
	wholeSource := &bytes.Buffer{}
	fmt.Fprintf(wholeSource, "// Code generated by %s, DO NOT EDIT.\n", cmdName)
	fmt.Fprintf(wholeSource, "package %s\n\n", packname)

	importList := f.ImportList()
	if len(importList) > 0 {
		fmt.Fprintf(wholeSource, `
			import (
		`)
		for _, v := range importList {
			fmt.Fprintf(wholeSource, "%s\n", v)

		}
		fmt.Fprintf(wholeSource, `
			)
		`)
	}

	wholeSource.Write(f.Buffer.Bytes())

	formatted, err := format.Source(wholeSource.Bytes())
	if err != nil {
		lines := strings.Split(wholeSource.String(), "\n")
		for i := range lines {
			lines[i] = fmt.Sprintf("%d: %s", i, lines[i])
		}
		log.Print(strings.Join(lines, "\n"))
		log.Fatal("format error ", err)

		return
	}

	if filename != "" {
		err = os.WriteFile(filename, formatted, 0644)
		if err != nil {
			return
		}
	} else {
		fmt.Printf("%s", formatted)
		fmt.Println()
	}
}

func ConsType(start, until int, last string) string {
	return RecursiveType("hlist.Cons", start, until, last)
}

func ReversConsType(start, until int) string {
	ret := "hlist.Nil"
	for j := start; j <= until; j++ {
		ret = fmt.Sprintf("hlist.Cons[A%d, %s]", j, ret)
	}
	return ret
}

func FuncDeclArgs(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "a%d A%d", j, j)
	}
	return f.String()
}

func FuncCallArgs(start, until int, prefixOpt ...string) string {

	prefix := "a"
	if len(prefixOpt) > 0 {
		prefix = prefixOpt[0]
	}
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s%d", prefix, j)
	}
	return f.String()
}

func ReverseFuncCallArgs(start, until int, prefixOpt ...string) string {
	prefix := "a"
	if len(prefixOpt) > 0 {
		prefix = prefixOpt[0]
	}

	f := &bytes.Buffer{}
	for j := until; j >= start; j-- {
		if j != until {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s%d", prefix, j)
	}
	return f.String()
}

func CurriedCallArgs(start, until int, prefixOpt ...string) string {
	prefix := "a"
	if len(prefixOpt) > 0 {
		prefix = prefixOpt[0]
	}

	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {

		fmt.Fprintf(f, "(%s%d)", prefix, j)
	}
	return f.String()
}

func FuncDeclTypeClassArgs(start, until int, typeClass string) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "ins%d %s[A%d]", j, typeClass, j)
	}
	return f.String()
}

func TypeClassArgs(start, until int, typeClass string) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%s[A%d]", typeClass, j)
	}
	return f.String()
}

func CurriedType(start, until int, rtype string) string {
	return RecursiveType("fp.Func1", start, until, rtype)
}

func RecursiveType(tp string, start, until int, rtype string) string {
	if start > until {
		f := &bytes.Buffer{}
		endBracket := ""
		for j := start; j >= until; j-- {
			fmt.Fprintf(f, "%s[A%d, ", tp, j)
			endBracket = endBracket + "]"
		}
		fmt.Fprintf(f, "%s%s", rtype, endBracket)

		return f.String()
	} else {
		f := &bytes.Buffer{}
		endBracket := ""
		for j := start; j <= until; j++ {
			fmt.Fprintf(f, "%s[A%d, ", tp, j)
			endBracket = endBracket + "]"
		}
		fmt.Fprintf(f, "%s%s", rtype, endBracket)

		return f.String()
	}

}

type Range struct {
	writer *writer
	start  int
	end    int
}

func (r Range) Write(txt string, param map[string]any) {

	if param == nil {
		param = map[string]any{}
	}

	tpl, err := template.New("write").Funcs(defaultFunc).Parse(txt)
	if err == nil {
		for i := r.start; i < r.end; i++ {
			param["N"] = i
			err := tpl.Execute(r.writer, param)
			if err != nil {
				fmt.Printf("template = %s\n", txt)
				panic(err)
			}
		}
	} else {
		fmt.Printf("template = %s\n", txt)
		panic(err)
	}
}

func (r Range) Render(txt string, funcs map[string]any, param map[string]any) {

	if param == nil {
		param = map[string]any{}
	}

	tpl, err := template.New("write").Funcs(defaultFunc).Funcs(funcs).Parse(txt)
	if err == nil {
		for i := r.start; i < r.end; i++ {
			param["N"] = i
			err := tpl.Execute(r.writer, param)
			if err != nil {
				fmt.Printf("template = %s\n", txt)
				panic(err)
			}
		}
	} else {
		fmt.Printf("template = %s\n", txt)
		panic(err)
	}
}
