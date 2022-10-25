package value_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
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
	b := value.PersonMutable{
		Name: "Hello",
		Age:  10,
	}.AsImmutable()

	assert.True(value.EqPerson.Eqv(a, b))
	assert.False(value.EqPerson.Eqv(a, b.WithAge(20)))

}

func TestHash(t *testing.T) {
	key := fp.New(value.Key.Builder).A(10).B(13).C([]byte("hello")).Build()

	fmt.Println("hash = ", key.Hash())
}

func TestMonoid(t *testing.T) {
	p1 := fp.New(value.Point.Builder).X(10).Y(12).Z(as.Tuple(1, 2)).Build()
	p2 := fp.New(value.Point.Builder).X(5).Y(4).Z(as.Tuple(2, 3)).Build()

	p3 := value.MonoidPoint.Combine(p1, p2)
	assert.Equal(p3.X(), 15)
	assert.Equal(p3.Y(), 16)
	assert.Equal(p3.Z().I1, 3)
	assert.Equal(p3.Z().I2, 5)

}
