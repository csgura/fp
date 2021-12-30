package hash_test

import (
	"fmt"
	"testing"

	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hash"
	"github.com/csgura/fp/internal/assert"
)

func TestHash(t *testing.T) {
	tup := as.Tuple3("hello", "world", 1)

	hasher := hash.Tuple3(hash.String, hash.String, hash.Number[int]())

	fmt.Println(hasher.Hash(tup))

	hl := tup.ToHList()

	hlHasher := hash.HCons(hash.String, hash.HCons(hash.String, hash.HCons(hash.Number[int](), hash.HNil)))

	fmt.Println(hlHasher.Hash(hl))

	assert.Equal(hasher.Hash(tup), hlHasher.Hash(hl))
}
