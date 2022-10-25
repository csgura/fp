package value_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/test/value"
)

func TestBuilder(t *testing.T) {
	a := fp.New(value.Hello.Builder).World("world").Hi(0).Build()
	fmt.Println(a)
	fmt.Println(a.WithWorld("No").World())
}

func TestEq(t *testing.T) {
	a := fp.New(value.Person.Builder).Name("Hello").Age(10).Build()
	b := fp.New(value.Person.Builder).Name("Hello").Age(10).Build()

	assert.True(value.EqPerson.Eqv(a, b))
	assert.False(value.EqPerson.Eqv(a, b.WithAge(20)))

}

func TestHash(t *testing.T) {
	key := fp.New(value.Key.Builder).A(10).B(13).C([]byte("hello")).Build()

	fmt.Println("hash = ", key.Hash())
}
