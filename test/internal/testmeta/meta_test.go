package testmeta_test

import (
	"fmt"
	"go/types"
	"os"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/assert"
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
	assert.IsNil(err)

	fpPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp")
	}).NextOption()

	assert.True(fpPkg.IsDefined())

	named := pkgs[0].Types.Scope().Lookup("Named")
	assert.NotNil(named)
	tp := metafp.GetTypeInfo(named.Type())
	fmt.Printf("tp = %s\n", tp)

	rnt := fpPkg.Get().Scope().Lookup("RuntimeNamed")
	assert.NotNil(rnt)

	rntp := metafp.GetTypeInfo(rnt.Type()).Instantiate(seq.Of(metafp.BasicType(types.Int).PtrType()))
	fmt.Printf("tp = %s\n", rntp)

	rtype := tp.ResultType()
	cr := metafp.ConstraintCheck(tp.TypeParam, rtype, seq.Of(rntp.Get()))
	fmt.Printf("err = %s\n", cr.Error)
	assert.True(cr.Ok)
}

func TestCheckConstraintHashInt(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	assert.IsNil(err)

	hashPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp/hash")
	}).NextOption()

	assert.True(hashPkg.IsDefined())

	nhash := hashPkg.Get().Scope().Lookup("Number")
	assert.NotNil(nhash)
	tp := metafp.GetTypeInfo(nhash.Type())
	rtype := tp.ResultType()

	rntp := metafp.BasicType(types.Int)

	cr := metafp.ConstraintCheck(tp.TypeParam, rtype, seq.Of(rntp))
	fmt.Printf("err = %s\n", cr.Error)
	assert.True(cr.Ok)

}

func TestCheckConstraintShowHCons(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	assert.IsNil(err)

	hashPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp/test/internal/show")
	}).NextOption()

	assert.True(hashPkg.IsDefined())

	nhash := hashPkg.Get().Scope().Lookup("TupleHCons")
	assert.NotNil(nhash)
	tp := metafp.GetTypeInfo(nhash.Type())
	rtype := tp.ResultType().TypeArgs.Head().Get()

	strtp := metafp.BasicType(types.String)

	hlisttp := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("Intlist").Type())

	cr := metafp.ConstraintCheck(tp.TypeParam, rtype, seq.Of(strtp, hlisttp))
	fmt.Printf("err = %s\n", cr.Error)
	assert.True(cr.Ok)

}

func TestCheckConstraintDecoder(t *testing.T) {
	cwd, _ := os.Getwd()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, cwd)
	assert.IsNil(err)

	hashPkg := iterator.FilterMap(iterator.FromSeq(pkgs), func(v *packages.Package) fp.Option[*types.Package] {
		return FindPackage(v.Types, "github.com/csgura/fp/test/internal/js")
	}).NextOption()

	assert.True(hashPkg.IsDefined())

	fntype := hashPkg.Get().Scope().Lookup("DecoderNamed")
	assert.NotNil(fntype)
	tp := metafp.GetTypeInfo(fntype.Type())
	rtype := tp.ResultType()

	argtype := metafp.GetTypeInfo(pkgs[0].Types.Scope().Lookup("StringNamed").Type())

	cr := metafp.ConstraintCheck(tp.TypeParam, rtype, seq.Of(argtype))
	fmt.Printf("err = %s\n", cr.Error)
	assert.True(cr.Ok)

}
