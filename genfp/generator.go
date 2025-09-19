package genfp

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/scanner"
	"go/token"
	"go/types"
	"io"
	"log"
	"maps"
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

func TypeArg(prefix string, idx int) string {
	return fmt.Sprintf("%s%d", prefix, idx)
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

func FuncTypeArg(idx int, prefix ...string) string {
	if len(prefix) > 0 {
		return TypeArg(prefix[0], idx)
	}
	return TypeArg("A", idx)
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
	BuildFlag            string
	Package              string
	Filename             string
	CmdName              string
	PackageCommentBuffer *bytes.Buffer
	Buffer               *bytes.Buffer
	numSaved             int
	importSet
}

func (r *writer) PackageCommentWriter() io.Writer {
	return r.PackageCommentBuffer
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
		if param["N"] == nil {
			param["N"] = 1
		}
		err := tpl.Execute(r, param)
		if err != nil {
			lines := strings.Split(templateStr, "\n")
			for i := range lines {
				lines[i] = fmt.Sprintf("%d: %s", i, lines[i])
			}
			log.Print(strings.Join(lines, "\n"))
			log.Fatal("template error ", err)

		}

	} else {
		lines := strings.Split(templateStr, "\n")
		for i := range lines {
			lines[i] = fmt.Sprintf("%d: %s", i, lines[i])
		}
		log.Print(strings.Join(lines, "\n"))
		log.Fatal("template error ", err)
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

func (r *importSet) All() func(yield func(PackageId) bool) {
	return func(yield func(PackageId) bool) {
		for _, v := range r.PathToName {
			if !yield(NewImportPackage(v.path, v.alias)) {
				return
			}
		}
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

func (r *writer) CheckMaxFileSize(maxSize int) {
	if r.Buffer.Len() > maxSize {
		r.saveFile()
	}
}

type FormattingError struct {
	Lines       []string
	FormatError error
}

func (r *FormattingError) LineNumbered() string {
	ret := make([]string, len(r.Lines))
	for i := range r.Lines {
		ret[i] = fmt.Sprintf("%d: %s", i+1, r.Lines[i])
	}
	return strings.Join(ret, "\n")
}

func (r *FormattingError) Error() string {
	el := r.ErrorLine()

	return strings.Join([]string{
		r.NumberedLineAt(el - 1),
		r.NumberedLineAt(el),
		r.NumberedLineAt(el + 1),
		r.FormatError.Error(),
	}, "\n")

}

func (r *FormattingError) NumberedLineAt(n int) string {
	if n >= 0 && n <= len(r.Lines) {
		return fmt.Sprintf("%d: %s", n, r.Lines[n-1])
	}
	return ""
}

func (r *FormattingError) LineAt(n int) string {
	if n > 0 && n <= len(r.Lines) {
		return r.Lines[n-1]
	}
	return ""
}

func (r *FormattingError) ErrorLine() int {
	if se, ok := r.FormatError.(scanner.ErrorList); ok {
		if len(se) > 0 {
			return se[0].Pos.Line
		}
	}
	return 0
}

func (r *writer) makeFile() (string, *FormattingError) {
	wholeSource := &bytes.Buffer{}
	if r.BuildFlag != "" {
		fmt.Fprintf(wholeSource, "//go:build %s\n", r.BuildFlag)
	}
	fmt.Fprintf(wholeSource, "// Code generated by %s, DO NOT EDIT.\n", r.CmdName)

	if r.PackageCommentBuffer.Len() > 0 {
		fmt.Fprintf(wholeSource, "\n%s", r.PackageCommentBuffer.String())
	}
	fmt.Fprintf(wholeSource, "package %s\n\n", r.Package)

	importList := r.ImportList()
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

	wholeSource.Write(r.Buffer.Bytes())

	formatted, err := format.Source(wholeSource.Bytes())
	if err != nil {
		lines := strings.Split(wholeSource.String(), "\n")
		return "", &FormattingError{
			Lines:       lines,
			FormatError: err,
		}
	}

	return string(formatted), nil

}

func (r *writer) saveFile() (ImportSet, error) {
	if r.Buffer.Len() == 0 {
		return &r.importSet, nil
	}

	err := func() error {
		formatted, ferr := r.makeFile()
		if ferr != nil {
			log.Println(ferr.LineNumbered())
			log.Println(ferr.FormatError)
			return ferr
		}

		if r.Filename != "" {
			filename := r.Filename
			if r.numSaved > 0 {
				ext := path.Ext(r.Filename)
				base := r.Filename[0 : len(r.Filename)-len(ext)]
				filename = fmt.Sprintf("%s%d%s", base, r.numSaved, ext)
			}
			d := path.Dir(filename)
			if d != "." && d != "/" {
				os.MkdirAll(d, 0755)
			}
			err := os.WriteFile(filename, []byte(formatted), 0644)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("%s", formatted)
			fmt.Println()
		}
		return nil
	}()

	ret := &importSet{
		PathToName: r.importSet.PathToName,
		NameToPath: r.importSet.NameToPath,
	}
	r.importSet = importSet{PathToName: map[string]importAlias{}, NameToPath: map[string]importAlias{}}
	r.numSaved = r.numSaved + 1
	r.Buffer = &bytes.Buffer{}
	return ret, err
}

func (r *importSet) Clonse() ImportSet {
	return &importSet{
		PathToName: maps.Clone(r.PathToName),
		NameToPath: maps.Clone(r.NameToPath),
	}
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
	case *types.Alias:
		if _, ok := realtp.Underlying().(*types.Interface); ok {
			return "nil"
		}
		return fmt.Sprintf("%s{}", r.TypeName(pk, tpe))
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

		if realtp.TypeParams().Len() > 0 && realtp.TypeArgs() != nil {
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
	case *types.Alias:
		tpname := realtp.Obj().Name()

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

	}

	return tpe.String()
}

type ImportPackage struct {
	Package string
	Name    string
}

func (r ImportPackage) String() string {
	return r.Alias()
}

func (r ImportPackage) Path() string {
	return r.Package
}

func (r ImportPackage) Alias() string {
	if r.Name == "" {
		return path.Base(r.Package)
	}
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
	FindNode(pos token.Pos) ast.Node
	EvalTypeExpr(typeExpr ast.Expr) (types.Type, []ImportPackage)
}

func NewWorkingPackage(pk *types.Package, fset *token.FileSet, syntax []*ast.File) WorkingPackage {

	return &workingPackage{
		currentPackage: pk,
		fset:           fset,
		syntax:         syntax,
	}
}

type findPosVisitor struct {
	pos   token.Pos
	found ast.Node
}

func (r *findPosVisitor) Visit(node ast.Node) (w ast.Visitor) {

	if node != nil && node.Pos() == r.pos {
		r.found = node
		return nil
	}

	return r
}

type workingPackage struct {
	fset           *token.FileSet
	syntax         []*ast.File
	currentPackage *types.Package
}

func (w *workingPackage) EvalTypeExpr(typeExpr ast.Expr) (types.Type, []ImportPackage) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Uses:  map[*ast.Ident]types.Object{},
	}
	err := types.CheckExpr(w.fset, w.currentPackage, typeExpr.End(), typeExpr, info)
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

	// for _, v := range info.Uses {
	// 	if tn, ok := v.(*types.TypeName); ok {
	// 		if tn.Type() == ti.Type {
	// 			fmt.Printf("type aliased\n")
	// 		}
	// 	}
	// }

	return ti.Type, imports
}

func (w *workingPackage) FindNode(pos token.Pos) ast.Node {
	for _, f := range w.syntax {
		if pos >= f.Pos() && pos <= f.End() {
			for _, d := range f.Decls {
				if pos >= d.Pos() && pos <= d.End() {

					v := &findPosVisitor{pos: pos}
					ast.Walk(v, d)
					return v.found

				}
			}
		}
	}
	return nil
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
	All() func(yield func(PackageId) bool)
}

type Writer interface {
	io.Writer
	ImportSet
	PackageCommentWriter() io.Writer
	Iteration(start, end int) Range
	Render(template string, funcs map[string]any, param map[string]any)
	CheckMaxFileSize(maxSize int)
}

func GenerateString(packname string, writeFunc func(w Writer)) (string, error) {

	cmdName := path.Base(os.Args[0])

	f := &writer{
		"",
		packname,
		"",
		cmdName,
		&bytes.Buffer{},
		&bytes.Buffer{},
		0,
		importSet{PathToName: map[string]importAlias{}, NameToPath: map[string]importAlias{}}}

	writeFunc(f)

	ret, err := f.makeFile()
	if err != nil {
		return "", err
	}
	return ret, nil
}

func Generate(packname string, filename string, writeFunc func(w Writer)) (ImportSet, error) {

	cmdName := path.Base(os.Args[0])
	if filename != "" {
		fmt.Printf("%s generate %s", cmdName, filename)
		fmt.Println()
		os.Remove(filename)
	}

	f := &writer{
		"",
		packname,
		filename,
		cmdName,
		&bytes.Buffer{},
		&bytes.Buffer{},
		0,
		importSet{PathToName: map[string]importAlias{}, NameToPath: map[string]importAlias{}}}

	writeFunc(f)

	return f.saveFile()

}

func GenerateWithBuildFlag(buildFlag string, packname string, filename string, writeFunc func(w Writer)) (ImportSet, error) {

	cmdName := path.Base(os.Args[0])
	if filename != "" {
		fmt.Printf("%s generate %s", cmdName, filename)
		fmt.Println()
		os.Remove(filename)
	}

	f := &writer{
		buildFlag,
		packname,
		filename,
		cmdName,
		&bytes.Buffer{},
		&bytes.Buffer{},
		0,
		importSet{PathToName: map[string]importAlias{}, NameToPath: map[string]importAlias{}}}

	writeFunc(f)

	return f.saveFile()

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
