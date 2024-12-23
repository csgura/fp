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

func FindFpPackage(pk *types.Package) fp.Option[*types.Package] {
	ret := as.Seq(pk.Imports()).Find(func(v *types.Package) bool {
		return v.Path() == "github.com/csgura/fp"
	})
	if ret.IsDefined() {
		return ret
	}

	for _, p := range pk.Imports() {
		ret := FindFpPackage(p)
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
		return FindFpPackage(v.Types)
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
