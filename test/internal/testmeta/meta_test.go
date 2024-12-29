package testmeta_test

import (
	"fmt"
	"go/types"
	"os"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/fptest"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"golang.org/x/tools/go/packages"
)

func FindPackage(pk *types.Package, path string) fp.Option[*types.Package] {
	ret := as.Seq(pk.Imports()).Find(func(v *types.Package) bool {
		return v.Path() == path
	})
	if ret.IsDefined() {
		return ret
	}

	for _, p := range pk.Imports() {
		ret := FindPackage(p, path)
		if ret.IsDefined() {
			return ret
		}
	}

	return option.None[*types.Package]()
}

func TestCheckConstraint(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	fptest.IsNil(t, err)

	fpPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp")
	}).NextOption()

	fptest.True(t, fpPkg.IsDefined())

	named := pkgs[0].Types.Scope().Lookup("Named")
	fptest.NotNil(t, named)
	tp := metafp.GetTypeInfo(named.Type())
	fmt.Printf("tp = %s\n", tp)

	rnt := fpPkg.Get().Scope().Lookup("RuntimeNamed")
	fptest.NotNil(t, rnt)

	rntp := metafp.GetTypeInfo(rnt.Type()).Instantiate(seq.Of(metafp.BasicType(types.Int).PtrType()))
	fmt.Printf("tp = %s\n", rntp)

	rtype := tp.ResultType()
	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(rntp.Get()))
	fmt.Printf("err = %s\n", cr.Error)
	fptest.True(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	fptest.True(t, cr.ParamMapping.Get("A").IsDefined()) // int
	fptest.True(t, cr.ParamMapping.Get("V").IsDefined()) // fp.RuntimeNamed[*int]

}

func TestCheckConstraintHashInt(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	fptest.IsNil(t, err)

	hashPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp/hash")
	}).NextOption()

	fptest.True(t, hashPkg.IsDefined())

	nhash := hashPkg.Get().Scope().Lookup("Number")
	fptest.NotNil(t, nhash)
	tp := metafp.GetTypeInfo(nhash.Type())
	rtype := tp.ResultType()

	rntp := metafp.BasicType(types.Int)

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(rntp))
	fmt.Printf("err = %s\n", cr.Error)
	fptest.True(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	fptest.True(t, cr.ParamMapping.Get("T").IsDefined()) // int

}

func TestCheckConstraintShowHCons(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	fptest.IsNil(t, err)

	hashPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp/test/internal/show")
	}).NextOption()

	fptest.True(t, hashPkg.IsDefined())

	nhash := hashPkg.Get().Scope().Lookup("TupleHCons")
	fptest.NotNil(t, nhash)
	tp := metafp.GetTypeInfo(nhash.Type())
	rtype := tp.ResultType().TypeArgs.Head().Get()

	strtp := metafp.BasicType(types.String)

	hlisttp := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("Intlist").Type())

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(strtp, hlisttp))
	fmt.Printf("err = %s\n", cr.Error)
	fptest.True(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	fptest.True(t, cr.ParamMapping.Get("H").IsDefined()) // string
	fptest.True(t, cr.ParamMapping.Get("T").IsDefined()) // hlist.Cons(int, hlist.Nil)

}

func TestCheckConstraintDecoder(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	fptest.IsNil(t, err)

	hashPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp/test/internal/js")
	}).NextOption()

	fptest.True(t, hashPkg.IsDefined())

	fntype := hashPkg.Get().Scope().Lookup("DecoderNamed")
	fptest.NotNil(t, fntype)
	tp := metafp.GetTypeInfo(fntype.Type())
	rtype := tp.ResultType()

	argtype := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("StringNamed").Type())

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(argtype))
	fmt.Printf("err = %s\n", cr.Error)
	fptest.True(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	fptest.True(t, cr.ParamMapping.Get("T").IsDefined())
	fptest.True(t, cr.ParamMapping.Get("A").IsDefined()) // string

}

func TestCheckConstraintEncoderOption(t *testing.T) {

	/*
		param = [T github.com/csgura/fp.NamedField[A] A any], genericType =js.Encoder[T any](T), typeArgs = [testmeta.NamedOptOfCar[T comparable](T)]
		param = [T github.com/csgura/fp.NamedField[A] A any], genericType =fp.NamedField[T any](A), typeArgs = [T]
		sig = func[A any]() (ret github.com/csgura/fp.NamedField[A]), paramCons = [A], paramIns = [T]
		sig = func[T github.com/csgura/fp.NamedField[T]]() (ret github.com/csgura/fp/test/internal/js.Encoder[T]), paramCons = [T], paramIns = [github.com/csgura/fp/test/internal/testmeta.NamedOptOfCar[T comparable]]
	*/

	/*
		param = [T github.com/csgura/fp.NamedField[A] A any], genericType =js.Encoder[T any](T), typeArgs = [gendebug.NamedOptOfCar[T comparable](fp.Option[T any](T))]
		param = [T github.com/csgura/fp.NamedField[A] A any], genericType =fp.NamedField[T any](A), typeArgs = [fp.Option[T any](T)]
		sig = func[A any]() (ret github.com/csgura/fp.NamedField[A]), paramCons = [A], paramIns = [github.com/csgura/fp.Option[T]]
		sig = func[T github.com/csgura/fp.NamedField[github.com/csgura/fp.Option[T]]]() (ret github.com/csgura/fp/test/internal/js.Encoder[T]), paramCons = [T], paramIns = [github.com/csgura/fp/test/internal
	*/
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	fptest.IsNil(t, err)

	hashPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp/test/internal/js")
	}).NextOption()

	fptest.True(t, hashPkg.IsDefined())

	fntype := hashPkg.Get().Scope().Lookup("EncoderNamed")
	fptest.NotNil(t, fntype)
	tp := metafp.GetTypeInfo(fntype.Type())
	rtype := tp.ResultType()

	argtype := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("carOpt").Type())

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(argtype))
	fmt.Printf("err = %s\n", cr.Error)
	fptest.True(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	fptest.True(t, cr.ParamMapping.Get("T").IsDefined())
	fptest.True(t, cr.ParamMapping.Get("A").IsDefined()) // string
	fptest.Equal(t, cr.ParamMapping.Get("A").Get().String(), "fp.Option[T any](int)")
	fptest.Equal(t, cr.ParamMapping.Get("T").Get().String(), "testmeta.NamedOptOfCar[T comparable](int)")

}

func TestCheckConstraintInterface(t *testing.T) {

	/*
		param = [T github.com/csgura/fp.NamedField[A] A any], genericType =js.Encoder[T any](T), typeArgs = [testmeta.NamedOptOfCar[T comparable](T)]
		param = [T github.com/csgura/fp.NamedField[A] A any], genericType =fp.NamedField[T any](A), typeArgs = [T]
		sig = func[A any]() (ret github.com/csgura/fp.NamedField[A]), paramCons = [A], paramIns = [T]
		sig = func[T github.com/csgura/fp.NamedField[T]]() (ret github.com/csgura/fp/test/internal/js.Encoder[T]), paramCons = [T], paramIns = [github.com/csgura/fp/test/internal/testmeta.NamedOptOfCar[T comparable]]
	*/

	/*
		param = [T github.com/csgura/fp.NamedField[A] A any], genericType =js.Encoder[T any](T), typeArgs = [gendebug.NamedOptOfCar[T comparable](fp.Option[T any](T))]
		param = [T github.com/csgura/fp.NamedField[A] A any], genericType =fp.NamedField[T any](A), typeArgs = [fp.Option[T any](T)]
		sig = func[A any]() (ret github.com/csgura/fp.NamedField[A]), paramCons = [A], paramIns = [github.com/csgura/fp.Option[T]]
		sig = func[T github.com/csgura/fp.NamedField[github.com/csgura/fp.Option[T]]]() (ret github.com/csgura/fp/test/internal/js.Encoder[T]), paramCons = [T], paramIns = [github.com/csgura/fp/test/internal
	*/
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	fptest.IsNil(t, err)

	fpPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp")
	}).NextOption()

	fptest.True(t, fpPkg.IsDefined())

	fpNamed := metafp.GetTypeInfo(fpPkg.Get().Scope().Lookup("NamedField").Type())

	tnamed := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("carOpt").Type())

	cr := tnamed.HasMethod(metafp.ConstraintCheckResult{Ok: true}, fpNamed.TypeParam, fpNamed.Method.Get("Value").Get())
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	fptest.True(t, cr.Ok)
	fptest.True(t, cr.ParamMapping.Get("T").IsDefined())

}
