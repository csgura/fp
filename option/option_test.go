package option_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
)

func Sum(a string, b int) string {
	return fmt.Sprintf("%s:%d", a, b)
}

func TestSome(t *testing.T) {
	opt := option.Some("10")
	opt.Foreach(func(v string) {
		println(v)
	})

	opt2 := option.FlatMap(opt, func(v string) fp.Option[int] {
		return option.Some(20)
	})

	opt2.Foreach(fp.Println[int])

	var ptr *string = nil
	ptrOpt := option.Of(ptr)
	println(ptrOpt.IsDefined())

	println("option.Ptr(ptr)", option.Ptr(ptr).IsDefined())
	option.Ptr(as.Ptr("hello")).Foreach(fp.Println[string])

	var close io.Writer = nil
	intfOpt := option.Of(close)
	println(intfOpt.IsDefined())

	close = &bytes.Buffer{}
	intfOpt = option.Of(close)
	println(intfOpt.IsDefined())

	var buf *bytes.Buffer
	bufOpt := option.Of(buf)
	println(bufOpt.IsDefined())

	intOpt := option.Some(10)
	strOpt := option.Map(intOpt, strconv.Itoa)
	fmt.Println(strOpt)

	intNone := option.None[int]()
	strOpt = option.Map(intNone, strconv.Itoa)
	fmt.Println(strOpt)

	intFunctor := option.Map[int, string]
	curried.Func2(intFunctor)(intOpt)(strconv.Itoa)

	optFn := curried.Flip(as.Curried2(intFunctor))(strconv.Itoa)
	fmt.Println(optFn(option.Some(42)))

}

func TestCompileError(t *testing.T) {
	res := option.Applicative2(func(a int, b int) int {
		println(a, b)
		return 10
	}).
		Ap(1).
		Ap(20)
	fmt.Println(res)

}

func TestJson(t *testing.T) {

	type Hello struct {
		Hello string         `json:"hello"`
		World fp.Option[int] `json:"world"`
	}

	strNull := `{
		"hello":"world",
		"world" : null
	}`

	str := `{
		"hello":"world",
		"world" : 20
	}`

	strNone := `{
		"hello":"world"
	}`

	var h Hello
	json.Unmarshal([]byte(strNull), &h)
	assert.True(h.World.IsEmpty())

	b, err := json.Marshal(h)
	assert.IsNil(err)
	assert.Equal(string(b), `{"hello":"world","world":null}`)

	json.Unmarshal([]byte(strNone), &h)

	assert.True(h.World.IsEmpty())
	b, err = json.Marshal(h)
	assert.IsNil(err)
	assert.Equal(string(b), `{"hello":"world","world":null}`)

	json.Unmarshal([]byte(str), &h)
	assert.True(h.World.IsDefined())

	b, err = json.Marshal(h)
	assert.IsNil(err)
	assert.Equal(string(b), `{"hello":"world","world":20}`)

}

func TestIterator(t *testing.T) {
	itr := option.Iterator(option.Some(1))

	count := 0
	for itr.HasNext() {
		count++
		v := itr.Next()
		fmt.Println("v =", v)

		if count > 2 {
			panic("invalid iterator")
		}
	}

}

func TestPtr(t *testing.T) {

	opt := option.Some(10)
	ptr := opt.Ptr()

	assert.True(ptr != nil)

	*ptr = 20

	assert.Equal(*ptr, 20)
	assert.Equal(opt.Get(), 10)

}

type World struct {
	name    string
	address *string
	age     int
}

func (r World) Name() string {
	return r.name
}

func (r World) Age() int {
	return r.age
}

func (r World) Address() *string {
	return r.address
}

type Hello struct {
	world World
}

func (r Hello) World() World {
	return r.world
}

func TestFilter(t *testing.T) {

	opt := option.Some(Hello{
		world: World{
			name:    "gura",
			age:     17,
			address: as.Ptr("seoul"),
		},
	})

	res := opt.Exists(fp.TestWith(Hello.World)(
		fp.TestWith(World.Name)(eq.GivenValue("gura")).
			And(fp.TestWith(World.Age)(eq.GivenValue(17))).
			And(fp.TestWith(World.Address)(eq.NotNilAnd(eq.GivenValue("seoul")))),
	))
	assert.True(res)

	res = opt.Exists(fp.TestWith(Hello.World)(
		eq.GivenFieldValue(World.Name, "gura").
			And(eq.GivenFieldValue(World.Age, 17)).
			And(eq.FieldNotNilAnd(World.Address, eq.GivenValue("suji"))),
	))
	assert.False(res)

	res = opt.Exists(fp.TestWith(Hello.World)(func(w World) bool {
		return true
	}))
	assert.True(res)
}
