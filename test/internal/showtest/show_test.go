package showtest_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/show"
	"github.com/csgura/fp/test/internal/showtest"
)

func TestShow(t *testing.T) {
	v := showtest.Person{Name: "gura", Age: 29}

	assert.Equal(showtest.ShowPerson.Show(v), `showtest.Person{Name:"gura",Age:29}`)

	c := showtest.Collection{
		Index: map[string]showtest.Person{
			"gura":  v,
			"other": {Name: "other", Age: 30},
		},
		List:        []showtest.Person{v, {Name: "list", Age: 30}},
		Description: as.Ptr("example"),
		Set:         mutable.SetOf(1, 2, 3),
		Option: option.Some(showtest.Person{
			Name: "opt",
			Age:  12,
		}),
		NoMap: map[string]showtest.NoDerive{
			"hello": {
				Hello: "world",
			},
			"hi": {
				Hello: "there",
			},
		},
		StringSeq: fp.Seq[string]{"1"},
	}
	fmt.Println(showtest.ShowCollection.ShowIndent(c, show.Pretty))

	d := showtest.HasTuple{
		Entry: as.Tuple2("hello", 10),
		HList: hlist.Concat("hello", hlist.Concat(1, hlist.Nil{})),
	}

	fmt.Println("d = ", showtest.ShowHasTuple.ShowIndent(d, show.Space))

	// untyped struct 에 private field 있는 경우, 다른 패키지에서 호출 불가능
	// showtest.UntypedStructFunc(struct {
	// 	Level   int
	// 	Stage   string
	// 	privacy string
	// }{Level: 1, Stage: "hello", privacy: "p"})

	e := showtest.EmbeddedStructMutable{
		Hello: "world",
		World: struct {
			Level int
			Stage string
		}{Level: 1, Stage: "hello"},
	}.AsImmutable()

	fmt.Println("e = ", showtest.ShowEmbeddedStruct.ShowIndent(e, show.Space))
}
