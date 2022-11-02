package docexample_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/test/internal/docexample"
)

func TestJson(t *testing.T) {

	addr := docexample.AddressMutable{
		Country: "Korea",
		City:    "Seoul",
		Street:  "Teheran",
	}.AsImmutable()

	b, err := json.Marshal(addr)

	assert.Success(err)

	var target *docexample.Address
	err = json.Unmarshal(b, &target)
	assert.Success(err)

	fmt.Println(target)
}

func TestShow(t *testing.T) {

	addr := docexample.AddressMutable{
		Country: "Korea",
		City:    "Seoul",
		Street:  "Teheran",
	}.AsImmutable()

	fmt.Println(docexample.ShowAddress.Show(addr))
}
