package metafp

import (
	"bytes"
	"fmt"
	"go/format"
	"go/types"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

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

func FuncTypeArgs(start, until int) string {
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

func (r Monad) FuncChain(start, until int) string {
	f := &bytes.Buffer{}
	for j := start; j <= until; j++ {
		if j != start {
			fmt.Fprintf(f, ", ")
		}
		if j == until {
			fmt.Fprintf(f, "f%d fp.Func1[A%d,%s[R]]", j, j, r)
		} else {
			fmt.Fprintf(f, "f%d fp.Func1[A%d,%s[A%d]]", j, j, r, j+1)
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

type writer struct {
	Package    string
	Buffer     *bytes.Buffer
	PathToName map[string]importAlias
	NameToPath map[string]string
}

func (r *writer) Write(b []byte) (int, error) {
	return r.Buffer.Write(b)
}

func (r *writer) Iteration(start, end int) Range {
	return Range{r, start, end}
}

type importAlias struct {
	alias   string
	isalias bool
}

func (r *writer) AddImport(p *types.Package, alias string) bool {
	_, ok := r.PathToName[p.Path()]
	if ok {
		return false
	}

	_, ok = r.NameToPath[alias]
	if ok {
		return false
	}

	r.PathToName[p.Path()] = importAlias{
		alias:   alias,
		isalias: p.Name() == alias,
	}
	r.NameToPath[alias] = p.Path()

	return true
}

func (r *writer) GetImportedName(p *types.Package) string {
	ret, ok := r.PathToName[p.Path()]
	if ok {
		return ret.alias
	}

	i := 1
	alias := p.Name()

	for {
		added := r.AddImport(p, alias)
		if added {
			return alias
		}

		alias = fmt.Sprintf("%s%d", p.Name(), i)
	}
}

func (r *writer) ImportList() []string {
	var ret = []string{}

	for k, v := range r.PathToName {
		if v.isalias {
			ret = append(ret, fmt.Sprintf(`"%s"`, k))

		} else {
			ret = append(ret, fmt.Sprintf(`%s "%s"`, v.alias, k))
		}
	}

	return ret
}

func (r *writer) TypeName(pk *types.Package, tpe types.Type) string {
	switch realtp := tpe.(type) {
	case *types.Named:
		tpname := realtp.Origin().Obj().Name()
		nameWithPkg := tpname
		if realtp.Obj().Pkg().Path() != pk.Path() {
			alias := r.GetImportedName(realtp.Obj().Pkg())

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
	}

	return tpe.String()
}

type ImportSet interface {
	GetImportedName(p *types.Package) string
	TypeName(pk *types.Package, tpe types.Type) string
}

type Writer interface {
	io.Writer
	ImportSet
	Iteration(start, end int) Range
}

func Generate(packname string, filename string, writeFunc func(w Writer)) {
	os.Remove(filename)

	f := &writer{packname, &bytes.Buffer{}, map[string]importAlias{}, map[string]string{}}

	writeFunc(f)

	wholeSource := &bytes.Buffer{}
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

	err = ioutil.WriteFile(filename, formatted, 0644)
	if err != nil {
		return
	}
}

func ConsType(start, until int, last string) string {
	ret := last
	for j := until; j >= start; j-- {
		ret = fmt.Sprintf("hlist.Cons[A%d, %s]", j, ret)
	}
	return ret
}

func ReversConsType(start, until int) string {
	ret := "Nil"
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
	f := &bytes.Buffer{}
	endBracket := ""
	for j := start; j <= until; j++ {
		fmt.Fprintf(f, "fp.Func1[A%d, ", j)
		endBracket = endBracket + "]"
	}
	fmt.Fprintf(f, "%s%s", rtype, endBracket)

	return f.String()
}

type Range struct {
	writer *writer
	start  int
	end    int
}

var defaultFunc = map[string]any{
	"TypeArgs":          FuncTypeArgs,
	"DeclArgs":          FuncDeclArgs,
	"CallArgs":          FuncCallArgs,
	"DeclTypeClassArgs": FuncDeclTypeClassArgs,
	"TupleType": func(n int) string {
		return fmt.Sprintf("Tuple%d[%s]", n, FuncTypeArgs(1, n))
	},
	"dec": func(n int) int {
		return n - 1
	},
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
				panic(err)
			}
		}
	} else {
		panic(err)
	}
}
