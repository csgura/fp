package immutable_test

import (
	"testing"

	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/immutable"
	"github.com/csgura/fp/ord"
	"github.com/csgura/fp/slice"
)

type String struct {
	value string
	hash  uint32
}

var HashString = hash.New(eq.New(func(a, b *String) bool {
	return a.value == b.value
}), func(t *String) uint32 {
	return t.hash
})

func NewString(v string) *String {
	return &String{
		value: v,
		hash:  hash.String.Hash(v),
	}
}

func BenchmarkFpHashedMap8(b *testing.B) {

	nf_service_op := NewString("nf_service_op")
	nf_type := NewString("nf_type")
	api_root := NewString("api_root")
	version := NewString("version")
	method := NewString("method")
	service_name := NewString("service_name")
	path := NewString("path")
	description := NewString("description")
	world := NewString("world")

	obj := immutable.Map(HashString,
		as.Tuple(nf_service_op, "Update"),
		as.Tuple(nf_type, "nf_type"),
		as.Tuple(api_root, "api_root"),
		as.Tuple(version, "version"),
		as.Tuple(method, "method"),
		as.Tuple(service_name, "service_name"),
		as.Tuple(path, "path"),
		as.Tuple(description, "description"),
		as.Tuple(world, "world"),
	)

	for b.Loop() {
		obj.Get(nf_service_op)
		obj.Get(nf_type)
		obj.Get(api_root)
		obj.Get(version)
		obj.Get(method)
		obj.Get(service_name)
	}
}

func BenchmarkSortedMap(b *testing.B) {

	obj := slice.ToSortedMap(slice.Of(
		as.Tuple("nf_service_op", "Update"),
		as.Tuple("nf_type", "nf_type"),
		as.Tuple("api_root", "api_root"),
		as.Tuple("version", "version"),
		as.Tuple("method", "method"),
		as.Tuple("service_name", "service_name"),
		as.Tuple("path", "path"),
		as.Tuple("description", "description"),
		as.Tuple("world", "world"),
	), ord.Given[string]())

	for b.Loop() {
		obj.Get("nf_service_op")
		obj.Get("nf_type")
		obj.Get("api_root")
		obj.Get("version")
		obj.Get("method")
		obj.Get("service_name")
	}
}
