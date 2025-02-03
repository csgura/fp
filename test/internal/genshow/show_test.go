package genshow_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/test/internal/genshow"
)

func TestShow(t *testing.T) {

	v := genshow.Person{Name: "gura", Age: 29}

	assert.Equal(genshow.ShowPerson(v, fp.ShowOption{}), `genshow.Person{Name:"gura",Age:29}`)

}
