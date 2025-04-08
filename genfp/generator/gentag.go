package generator

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/csgura/fp/genfp"
	"golang.org/x/tools/go/packages"
)

type Annotation struct {
	name   string
	params map[string]string
}

func parseKeyValue(s string) tuple2[string, string] {
	s = strings.TrimSpace(s)
	idx := strings.Index(s, "=")
	if idx > 0 && len(s) > idx+1 {
		return tuple2[string, string]{strings.TrimSpace(s[:idx]), strings.TrimSpace(s[idx+1:])}
	}
	return tuple2[string, string]{s, "true"}
}

func parseAnnotation(s string) tuple2[string, Annotation] {
	pstart := strings.Index(s, "(")
	if pstart > 0 {
		pend := strings.LastIndex(s, ")")
		if pend > pstart {
			name := strings.TrimSpace(s[:pstart])
			params := s[pstart+1 : pend]

			itr := strings.Split(params, ",")
			p := seqToGoMap(seqMap(itr, parseKeyValue))
			return asTuple(name, Annotation{
				name:   name,
				params: p,
			})
		}

	}
	name := strings.TrimSpace(s)
	return asTuple(name, Annotation{
		name: name,
	})

}

func seqZip[A, B any](s1 []A, s2 []B) []tuple2[A, B] {
	minSize := min(len(s1), len(s2))

	ret := make([]tuple2[A, B], minSize)
	for i := 0; i < minSize; i++ {
		ret[i] = asTuple(s1[i], s2[i])
	}
	return ret
}

type Map[K comparable, V any] map[K]V

func (r Map[K, V]) Contains(k K) bool {
	_, ok := r[k]
	return ok
}

func extractTag(comment string) Map[string, Annotation] {
	list := strings.Split(comment, "\n")
	list = seqMap(list, strings.TrimSpace)
	list = seqFilter(list, func(v string) bool {
		return strings.HasPrefix(v, "@")
	})
	ret := seqToGoMap(seqMap(list, parseAnnotation))
	return ret
}

func seqExists[T any](r []T, p func(v T) bool) bool {
	for _, v := range r {
		if p(v) {
			return true
		}
	}
	return false
}

func FindGenerateFromUntil(p []*packages.Package, tags ...string) map[string][]GenerateFromUntil {
	ret := map[string][]GenerateFromUntil{}
	genseq := FindTaggedCompositeVariable(p, "GenerateFromUntil", tags...)
	for _, cl := range genseq {
		gfu, err := ParseGenerateFromUntil(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s\n", err)
		} else {
			s := ret[gfu.File]
			s = append(s, gfu)
			ret[gfu.File] = s
		}
	}

	return ret
}

func FindGenerateFromList(p []*packages.Package, tags ...string) map[string][]GenerateFromList {
	ret := map[string][]GenerateFromList{}
	genseq := FindTaggedCompositeVariable(p, "GenerateFromList", tags...)
	for _, cl := range genseq {
		gfu, err := ParseGenerateFromList(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s\n", err)
		} else {
			s := ret[gfu.File]
			s = append(s, gfu)
			ret[gfu.File] = s
		}
	}

	return ret
}

func FindGenerateFromStructs(p []*packages.Package, tags ...string) map[string][]GenerateFromStructs {
	ret := map[string][]GenerateFromStructs{}
	genseq := FindTaggedCompositeVariable(p, "GenerateFromStructs", tags...)
	for _, cl := range genseq {
		gfu, err := ParseGenerateFromStructs(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s\n", err)
		} else {
			s := ret[gfu.File]
			s = append(s, gfu)
			ret[gfu.File] = s
		}
	}

	return ret
}

func FindGenerateFromInterfaces(p []*packages.Package, tags ...string) map[string][]GenerateFromStructs {
	ret := map[string][]GenerateFromStructs{}
	genseq := FindTaggedCompositeVariable(p, "GenerateFromInterfaces", tags...)
	for _, cl := range genseq {
		gfu, err := ParseGenerateFromInterfaces(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s\n", err)
		} else {
			s := ret[gfu.File]
			s = append(s, gfu)
			ret[gfu.File] = s
		}
	}

	return ret
}

func FindGenerateAdaptor(p []*packages.Package, tags ...string) map[string][]GenerateAdaptorDirective {
	ret := map[string][]GenerateAdaptorDirective{}
	genseq := FindTaggedCompositeVariable(p, "GenerateAdaptor", tags...)
	for _, cl := range genseq {
		gfu, err := ParseGenerateAdaptor(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s\n", err)
		} else {
			s := ret[gfu.File]
			s = append(s, gfu)
			ret[gfu.File] = s
		}
	}

	return ret
}

func FindGenerateMonadFunctions(p []*packages.Package, tags ...string) map[string][]GenerateMonadFunctionsDirective {
	ret := map[string][]GenerateMonadFunctionsDirective{}
	genseq := FindTaggedCompositeVariable(p, "GenerateMonadFunctions", tags...)
	for _, cl := range genseq {
		gfu, err := ParseGenerateMonadFunctions(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s\n", err)
		} else {
			s := ret[gfu.File]
			s = append(s, gfu)
			ret[gfu.File] = s
		}
	}

	return ret
}

func FindGenerateTraverseFunctions(p []*packages.Package, tags ...string) map[string][]GenerateMonadFunctionsDirective {
	ret := map[string][]GenerateMonadFunctionsDirective{}
	genseq := FindTaggedCompositeVariable(p, "GenerateTraverseFunctions", tags...)
	for _, cl := range genseq {
		gfu, err := ParseGenerateMonadFunctions(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s\n", err)
		} else {
			s := ret[gfu.File]
			s = append(s, gfu)
			ret[gfu.File] = s
		}
	}

	return ret
}

func FindGenerateMonadTransfomers(p []*packages.Package, tags ...string) map[string][]GenerateMonadTransformerDirective {
	ret := map[string][]GenerateMonadTransformerDirective{}
	genseq := FindTaggedCompositeVariable(p, "GenerateMonadTransformer", tags...)
	for _, cl := range genseq {
		gfu, err := ParseGenerateMonadTransformer(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s\n", err)
		} else {
			s := ret[gfu.File]
			s = append(s, gfu)
			ret[gfu.File] = s
		}
	}

	return ret
}

func FindGenerateApplicatives(p []*packages.Package, tags ...string) map[string][]GenerateApplicative {
	ret := map[string][]GenerateApplicative{}
	genseq := FindTaggedCompositeVariable(p, "GenerateApplicative", tags...)
	for _, cl := range genseq {
		gfu, err := ParseGenerateApplicative(cl)
		if err != nil {
			fmt.Printf("invalid generate directive : %s\n", err)
		} else {
			s := ret[gfu.File]
			s = append(s, gfu)
			ret[gfu.File] = s
		}
	}

	return ret
}

func checkType(pk *packages.Package, typeExpr ast.Expr) *types.Named {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	types.CheckExpr(pk.Fset, pk.Types, typeExpr.End(), typeExpr, info)

	ti := info.Types[typeExpr]
	if named, ok := ti.Type.(*types.Named); ok {
		return named
	}
	return nil
}

func checkFuncType(pk *packages.Package, typeExpr ast.Expr) *types.Signature {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	types.CheckExpr(pk.Fset, pk.Types, typeExpr.End(), typeExpr, info)

	ti := info.Types[typeExpr]
	if named, ok := ti.Type.(*types.Signature); ok {
		return named
	}
	return nil
}

type TaggedLit struct {
	Package *packages.Package
	Type    *types.Named
	Lit     *ast.CompositeLit
}

func (r TaggedLit) WorkingPackage() genfp.WorkingPackage {
	return genfp.NewWorkingPackage(r.Package.Types, r.Package.Fset, r.Package.Syntax)
}

type FuncOrGen struct {
	fn *ast.FuncDecl
	gn *ast.GenDecl
}

func taggedFromFuncDecl(pk *packages.Package, typeName string, gd *ast.FuncDecl, tagSeq []string) []TaggedLit {
	gdDoc := gd.Doc
	comment := func() string {
		if gdDoc != nil {
			return gdDoc.Text()
		}

		return ""
	}()

	if comment != "" && seqExists(tagSeq, func(tag string) bool { return strings.Contains(comment, tag) }) {
		tags := extractTag(comment)

		if !seqExists(tagSeq, tags.Contains) {
			return nil
		}

		if gd.Type.Results.NumFields() == 1 {

			sig := checkFuncType(pk, gd.Type)
			if sig != nil {
				if sig.Results().Len() == 1 {
					if named, ok := sig.Results().At(0).Type().(*types.Named); ok {
						if named.Obj().Name() == typeName {
							if lastStmt, ok := seqLast(gd.Body.List); ok {
								if retStmt, ok := lastStmt.(*ast.ReturnStmt); ok && len(retStmt.Results) == 1 {
									if cl, ok := retStmt.Results[0].(*ast.CompositeLit); ok {
										return []TaggedLit{{pk, named, cl}}
									}
								}
							}

						}
					}
				}
			}
		}
	}

	return nil
}

func taggedFromGenDecl(pk *packages.Package, typeName string, gd *ast.GenDecl, tagSeq []string) []TaggedLit {
	gdDoc := gd.Doc

	return seqFlatMap(gd.Specs, func(v ast.Spec) []TaggedLit {
		if ts, ok := v.(*ast.ValueSpec); ok {
			comment := func() string {
				if ts.Doc != nil {
					return ts.Doc.Text()
				}
				if gdDoc != nil {
					return gdDoc.Text()
				}

				return ""
			}()

			if comment != "" && seqExists(tagSeq, func(tag string) bool { return strings.Contains(comment, tag) }) {
				return seqFlatMap(seqZip(ts.Names, ts.Values), func(v tuple2[*ast.Ident, ast.Expr]) []TaggedLit {

					tags := extractTag(comment)

					if !seqExists(tagSeq, tags.Contains) {
						return nil
					}

					if cl, ok := v.I2.(*ast.CompositeLit); ok {
						named := checkType(pk, cl.Type)
						if named != nil {
							if named.Obj().Name() == typeName {
								return []TaggedLit{{pk, named, cl}}
							}
						}
					}
					return nil
				})
			}
		}
		return nil
	})
}

func FindTaggedCompositeVariable(p []*packages.Package, typeName string, tags ...string) []TaggedLit {
	tagSeq := tags
	return seqFlatMap(p, func(pk *packages.Package) []TaggedLit {
		s2 := seqFlatMap(pk.Syntax, func(v *ast.File) []ast.Decl {

			return v.Decls
		})

		s3 := seqFlatMap(s2, func(v ast.Decl) []FuncOrGen {
			switch r := v.(type) {
			case *ast.GenDecl:
				return []FuncOrGen{{gn: r}}
			case *ast.FuncDecl:
				return []FuncOrGen{{fn: r}}
			}
			return []FuncOrGen{}
		})

		return seqFlatMap(s3, func(gd FuncOrGen) []TaggedLit {
			if gd.fn != nil {
				return taggedFromFuncDecl(pk, typeName, gd.fn, tagSeq)
			} else if gd.gn != nil {

				return taggedFromGenDecl(pk, typeName, gd.gn, tagSeq)
			}
			return nil
		})

	})
}

type TaggedVar struct {
	Package genfp.WorkingPackage
	Name    string
	Type    *types.Named
}

func FindTaggedNotInitalizedVariable(p []*packages.Package, tags ...string) []TaggedVar {
	tagSeq := tags
	return seqFlatMap(p, func(pk *packages.Package) []TaggedVar {
		s2 := seqFlatMap(pk.Syntax, func(v *ast.File) []ast.Decl {

			return v.Decls
		})

		s3 := seqFlatMap(s2, func(v ast.Decl) []*ast.GenDecl {
			switch r := v.(type) {
			case *ast.GenDecl:
				return []*ast.GenDecl{r}
			}
			return []*ast.GenDecl{}
		})

		return seqFlatMap(s3, func(gd *ast.GenDecl) []TaggedVar {

			return varFromGenDecl(pk, gd, tagSeq)

		})

	})
}

func varFromGenDecl(pk *packages.Package, gd *ast.GenDecl, tagSeq []string) []TaggedVar {
	gdDoc := gd.Doc

	return seqFlatMap(gd.Specs, func(v ast.Spec) []TaggedVar {
		if ts, ok := v.(*ast.ValueSpec); ok && ts.Type != nil {
			comment := func() string {
				if ts.Doc != nil {
					return ts.Doc.Text()
				}
				if gdDoc != nil {
					return gdDoc.Text()
				}

				return ""
			}()

			tpe := checkType(pk, ts.Type)

			if comment != "" && tpe != nil && seqExists(tagSeq, func(tag string) bool { return strings.Contains(comment, tag) }) {
				return seqFlatMap(ts.Names, func(v *ast.Ident) []TaggedVar {

					tags := extractTag(comment)

					if !seqExists(tagSeq, tags.Contains) {
						return nil
					}

					return []TaggedVar{
						{
							Package: genfp.NewWorkingPackage(pk.Types, pk.Fset, pk.Syntax),
							Name:    v.Name,
							Type:    tpe},
					}
				})
			}
		}
		return nil
	})
}
