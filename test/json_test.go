package main_test

import (
	"encoding/json"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
)

type Hello struct {
	Hello fp.Option[string] `json:"hello"`
}

func TestUnmarshal(t *testing.T) {

	str := "{}"
	unit := fp.Unit{}
	err := json.Unmarshal([]byte(str), &unit)
	assert.IsNil(err)

	hello := Hello{}

	err = json.Unmarshal([]byte(str), &hello)
	assert.IsNil(err)
	assert.False(hello.Hello.IsDefined())

	err = json.Unmarshal([]byte(`{
		"hello" : "world"
	}`), &hello)
	assert.IsNil(err)
	assert.True(hello.Hello.IsDefined())
	assert.Equal(hello.Hello.Get(), "world")

	err = json.Unmarshal([]byte(`{
		"hello" : null
	}`), &hello)
	assert.IsNil(err)
	assert.False(hello.Hello.IsDefined())

}

func TestMarshal(t *testing.T) {

	hello := Hello{}
	res, err := json.Marshal(hello)
	assert.IsNil(err)
	assert.Equal(string(res), `{"hello":null}`)

	hello = Hello{Hello: option.Some("world")}
	res, err = json.Marshal(hello)
	assert.IsNil(err)
	assert.Equal(string(res), `{"hello":"world"}`)
}
