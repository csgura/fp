package genfp

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

func seqFlatMap[T, U any](opt []T, fn func(v T) []U) []U {
	ret := make([]U, 0, len(opt))

	for _, v := range opt {
		ret = append(ret, fn(v)...)
	}

	return ret
}

func seqMap[T, U any](opt []T, fn func(v T) U) []U {
	ret := make([]U, len(opt))

	for i, v := range opt {
		ret[i] = fn(v)
	}

	return ret
}

func seqFilter[T any](r []T, p func(v T) bool) []T {
	ret := make([]T, 0, len(r))
	for _, v := range r {
		if p(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

type tuple2[A, B any] struct {
	I1 A
	I2 B
}

func (r tuple2[A, B]) Unapply() (A, B) {
	return r.I1, r.I2
}

func seqToGoMap[K comparable, V any](s []tuple2[K, V]) map[K]V {
	ret := map[K]V{}
	for _, e := range s {
		k, v := e.Unapply()
		ret[k] = v
	}
	return ret
}

func asTuple[A, B any](a A, b B) tuple2[A, B] {
	return tuple2[A, B]{a, b}
}

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
		gfu, err := ParseGenerateFromUntil(cl.Lit)
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

func checkType(pk *packages.Package, typeExpr ast.Expr, pos token.Pos) *types.Named {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	types.CheckExpr(pk.Fset, pk.Types, pos, typeExpr, info)

	ti := info.Types[typeExpr]
	if named, ok := ti.Type.(*types.Named); ok {
		return named
	}
	return nil
}

type TaggedLit struct {
	Package *packages.Package
	Type    *types.Named
	Lit     *ast.CompositeLit
}

func FindTaggedCompositeVariable(p []*packages.Package, typeName string, tags ...string) []TaggedLit {
	tagSeq := tags
	return seqFlatMap(p, func(pk *packages.Package) []TaggedLit {
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

		return seqFlatMap(s3, func(gd *ast.GenDecl) []TaggedLit {
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
								named := checkType(pk, cl.Type, v.I2.Pos())
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
		})
	})
}
