package testmeta_test

import (
	"fmt"
	"go/types"
	"os"
	"testing"

	"github.com/csgura/fp/metafp"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/should"
	"golang.org/x/tools/go/packages"
)

func TestCheckConstraint(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	should.BeNil(t, err)

	fpPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp")

	should.BeTrue(t, fpPkg.IsDefined())

	named := pkgs[0].Types.Scope().Lookup("Named")
	should.NotBeNil(t, named)
	tp := metafp.GetTypeInfo(named.Type())
	fmt.Printf("tp = %s\n", tp)

	rnt := fpPkg.Get().Scope().Lookup("RuntimeNamed")
	should.NotBeNil(t, rnt)

	rntp := metafp.GetTypeInfo(rnt.Type()).Instantiate(seq.Of(metafp.BasicType(types.Int).PtrType()))
	fmt.Printf("tp = %s\n", rntp)

	rtype := tp.ResultType()
	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(rntp.Get()))
	fmt.Printf("err = %s\n", cr.Error)
	should.BeTrue(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	should.BeTrue(t, cr.ParamMapping.Get("A").IsDefined()) // int
	should.BeTrue(t, cr.ParamMapping.Get("V").IsDefined()) // fp.RuntimeNamed[*int]

}

func TestCheckConstraintHashInt(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	should.BeNil(t, err)

	hashPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp/hash")

	should.BeTrue(t, hashPkg.IsDefined())

	nhash := hashPkg.Get().Scope().Lookup("Number")
	should.NotBeNil(t, nhash)
	tp := metafp.GetTypeInfo(nhash.Type())
	rtype := tp.ResultType()

	rntp := metafp.BasicType(types.Int)

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(rntp))
	fmt.Printf("err = %s\n", cr.Error)
	should.BeTrue(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	should.BeTrue(t, cr.ParamMapping.Get("T").IsDefined()) // int

}

func TestCheckConstraintShowHCons(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	should.BeNil(t, err)

	hashPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp/test/internal/show")

	should.BeTrue(t, hashPkg.IsDefined())

	nhash := hashPkg.Get().Scope().Lookup("TupleHCons")
	should.NotBeNil(t, nhash)
	tp := metafp.GetTypeInfo(nhash.Type())
	rtype := tp.ResultType().TypeArgs.Head().Get()

	strtp := metafp.BasicType(types.String)

	hlisttp := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("Intlist").Type())

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(strtp, hlisttp))
	fmt.Printf("err = %s\n", cr.Error)
	should.BeTrue(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	should.BeTrue(t, cr.ParamMapping.Get("H").IsDefined()) // string
	should.BeTrue(t, cr.ParamMapping.Get("T").IsDefined()) // hlist.Cons(int, hlist.Nil)

}

func TestCheckConstraintDecoder(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	should.BeNil(t, err)

	hashPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp/test/internal/js")

	should.BeTrue(t, hashPkg.IsDefined())

	fntype := hashPkg.Get().Scope().Lookup("DecoderNamed")
	should.NotBeNil(t, fntype)
	tp := metafp.GetTypeInfo(fntype.Type())
	rtype := tp.ResultType()

	argtype := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("StringNamed").Type())

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(argtype))
	fmt.Printf("err = %s\n", cr.Error)
	should.BeTrue(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	should.BeTrue(t, cr.ParamMapping.Get("T").IsDefined())
	should.BeTrue(t, cr.ParamMapping.Get("A").IsDefined()) // string

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
	should.BeNil(t, err)

	hashPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp/test/internal/js")

	should.BeTrue(t, hashPkg.IsDefined())

	fntype := hashPkg.Get().Scope().Lookup("EncoderNamed")
	should.NotBeNil(t, fntype)
	tp := metafp.GetTypeInfo(fntype.Type())
	rtype := tp.ResultType()

	argtype := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("carOpt").Type())

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, seq.Of(argtype))
	fmt.Printf("err = %s\n", cr.Error)
	should.BeTrue(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	should.BeTrue(t, cr.ParamMapping.Get("T").IsDefined())
	should.BeTrue(t, cr.ParamMapping.Get("A").IsDefined()) // string
	should.Equal(t, cr.ParamMapping.Get("A").Get().String(), "fp.Option[T any](int)")
	should.Equal(t, cr.ParamMapping.Get("T").Get().String(), "testmeta.NamedOptOfCar[T comparable](int)")

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
	should.BeNil(t, err)

	fpPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp")

	should.BeTrue(t, fpPkg.IsDefined())

	fpNamed := metafp.GetTypeInfo(fpPkg.Get().Scope().Lookup("NamedField").Type())

	tnamed := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("carOpt").Type())

	cr := tnamed.HasMethod(metafp.ConstraintCheckResult{Ok: true}, fpNamed.TypeParam, fpNamed.Method.Get("Value").Get())
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	should.BeTrue(t, cr.Ok)
	should.BeTrue(t, cr.ParamMapping.Get("T").IsDefined())

}

func TestCheckConstraintNgapSplit(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	should.BeNil(t, err)

	hashPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp/test/internal/ngap")

	should.BeTrue(t, hashPkg.IsDefined())

	tp4fn := hashPkg.Get().Scope().Lookup("Tuple4")
	should.NotBeNil(t, tp4fn)
	tp := metafp.GetTypeInfo(tp4fn.Type())
	rtype := tp.ResultType().TypeArgs.Head().Get()

	tp4 := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("tp4").Type())

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, tp.TypeParam, rtype, tp4.TypeArgs)
	fmt.Printf("err = %s\n", cr.Error)
	should.BeTrue(t, cr.Ok)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	should.BeTrue(t, cr.ParamMapping.Get("A1").IsDefined()) // string
	should.BeTrue(t, cr.ParamMapping.Get("A2").IsDefined()) // hlist.Cons(int, hlist.Nil)
	should.BeTrue(t, cr.ParamMapping.Get("A3").IsDefined()) // hlist.Cons(int, hlist.Nil)

}

func TestCheckConstraintTypeParam(t *testing.T) {

	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	should.BeNil(t, err)

	fpPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp/test/internal/js")

	should.BeTrue(t, fpPkg.IsDefined())

	ft := metafp.GetTypeInfo(fpPkg.Get().Scope().Lookup("EncoderNamed").Type())

	tnamed := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("NamedTypeParam").Type())

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, ft.TypeParam, ft.ResultType(), seq.Of(tnamed))
	fmt.Printf("err = %s\n", cr.Error)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	should.BeTrue(t, cr.Ok)
	should.BeTrue(t, cr.ParamMapping.Get("T").IsDefined())

}

func TestCheckTypeAlias(t *testing.T) {

	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	should.BeNil(t, err)

	fpPkg := metafp.FindPackage(pkgs, "github.com/csgura/fp/hlist")

	should.BeTrue(t, fpPkg.IsDefined())

	ft := metafp.GetTypeInfo(fpPkg.Get().Scope().Lookup("Nil").Type())

	tnamed := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("ShowNil").Type())

	cr := metafp.ConstraintCheck(metafp.ConstraintCheckResult{Ok: true}, ft.TypeParam, tnamed, seq.Of(ft))
	fmt.Printf("err = %s\n", cr.Error)
	fmt.Printf("mapping = %s\n", cr.ParamMapping)
	should.BeTrue(t, cr.Ok)

}
