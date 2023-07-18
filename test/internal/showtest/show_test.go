package showtest_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/option"
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
		StringSeq: fp.Seq[string]{"1"},
	}
	fmt.Println(showtest.ShowCollection.ShowIndent(c, fp.ShowOption{
		Indent:    "  ",
		OmitEmpty: false,
	}))

	d := showtest.HasTuple{
		Entry: as.Tuple2("hello", 10),
	}

	fmt.Println("d = ", showtest.ShowHasTuple.Show(d))

	e := showtest.EmbeddedStructMutable{
		Hello: "world",
		World: struct {
			Level int
			Stage string
		}{Level: 10, Stage: "first"},
	}.AsImmutable()

	fmt.Println("e = ", showtest.ShowEmbeddedStruct.Show(e))
}
