package showtest_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/mutable"
	"github.com/csgura/fp/test/internal/showtest"
)

func TestShow(t *testing.T) {
	v := showtest.Person{Name: "gura", Age: 29}

	assert.Equal(showtest.ShowPerson.Show(v), `showtest.Person(Name:"gura",Age:29)`)

	c := showtest.Collection{
		Index: map[string]showtest.Person{
			"gura":  v,
			"other": {Name: "other", Age: 30},
		},
		List:        []showtest.Person{v, {Name: "list", Age: 30}},
		Description: as.Ptr("example"),
		Set:         mutable.SetOf(1, 2, 3),
	}
	fmt.Println(showtest.ShowCollection.Show(c))
}
