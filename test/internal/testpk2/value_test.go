package testpk2_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/test/internal/testpk1"
	"github.com/csgura/fp/test/internal/testpk2"
	"github.com/csgura/fp/try"
)

func TestString(t *testing.T) {
	a := fp.New(testpk2.AllKindTypes.Builder).
		A("world").
		T(try.Success(option.None[testpk2.Local]())).
		Fn3(func(a1 int) fp.Try[string] {
			return try.Success("success")
		}).
		Build()
	fmt.Println(a)
}

func TestBuilder(t *testing.T) {
	a := fp.New(testpk2.Hello.Builder).World("world").Hi(0).Build()
	fmt.Println(a)
	fmt.Println(a.WithWorld("No").World())
}

func TestEq(t *testing.T) {
	a := fp.New(testpk2.Person.Builder).Name("Hello").Age(10).Build()
	b := testpk2.PersonMutable{
		Name: "Hello",
		Age:  10,
	}.AsImmutable()

	assert.True(testpk2.EqPerson.Eqv(a, b))
	assert.False(testpk2.EqPerson.Eqv(a, b.WithAge(20)))

}

func TestHash(t *testing.T) {
	key := fp.New(testpk2.Key.Builder).A(10).B(13).C([]byte("hello")).Build()

	fmt.Println("hash = ", key.Hash())
}

func TestMonoid(t *testing.T) {
	p1 := fp.New(testpk2.Point.Builder).X(10).Y(12).Z(as.Tuple(1, 2)).Build()
	p2 := fp.New(testpk2.Point.Builder).X(5).Y(4).Z(as.Tuple(2, 3)).Build()

	p3 := testpk2.MonoidPoint.Combine(p1, p2)
	assert.Equal(p3.X(), 15)
	assert.Equal(p3.Y(), 16)
	assert.Equal(p3.Z().I1, 3)
	assert.Equal(p3.Z().I2, 5)

}

func TestJson(t *testing.T) {
	g := testpk2.GreetingMutable{
		Hello: testpk1.WorldMutable{
			Message:   "hello",
			Timestamp: time.Now(),
		}.AsImmutable(),
		Language: "En",
	}.AsImmutable()

	res := testpk2.EncoderGreeting.Encode(g).Get()
	fmt.Println(res)

	parsedG := testpk2.DecoderGreeting.Decode(res)
	parsedG.Failed().Foreach(func(v error) {
		fmt.Printf("parse error : %s\n", v)
	})
	assert.True(parsedG.IsSuccess())
	assert.Equal(parsedG.Get().Hello().Message(), "hello")
	assert.Equal(parsedG.Get().Language(), "En")

	var rev testpk2.Greeting
	err := json.Unmarshal([]byte(res), &rev)
	assert.Success(err)
	assert.True(rev.Language() == g.Language())

	res2, err := json.Marshal(rev)
	assert.Success(err)
	fmt.Println(string(res2))
	assert.True(res == string(res2))

	t3 := testpk2.ThreeMutable{
		One:   1,
		Two:   "2",
		Three: 3,
	}.AsImmutable()

	res = testpk2.EncoderThree.Encode(t3).Get()
	fmt.Println(res)

	parsedT3 := testpk2.DecoderThree.Decode(res)
	assert.True(parsedT3.IsSuccess())
	assert.Equal(parsedT3.Get().One(), 1)
	assert.Equal(parsedT3.Get().Two(), "2")

}

func TestRead(t *testing.T) {
	t3 := testpk2.ThreeMutable{
		One:   12,
		Two:   "hello world",
		Three: 13.5,
	}.AsImmutable()

	str := testpk2.ShowThree.Show(t3)
	fmt.Println(str)

	res := testpk2.ReadThree.Read(str)
	res.Failed().Foreach(fp.Println[error])
	assert.True(res.IsSuccess())
	assert.Equal(res.Get().Three(), 13.5)
	assert.Equal(res.Get().Two(), "hello world")

}
