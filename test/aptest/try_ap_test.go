//go:build ap
// +build ap

package main_test

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/try"
)

func TestApplicative(t *testing.T) {

	cf := curried.Concat[*url.URL](curried.Concat[string](fp.Id[int]))

	var intPort fp.Try[int] = try.Applicative3(curried.Revert3(cf)).
		ApTry(try.Func1(url.Parse)("http://localhost:8080/abcd")).
		Map((*url.URL).Port).
		FlatMap(try.Func1(strconv.Atoi))
	assert.True(intPort.IsSuccess())
	assert.Equal(intPort.Get(), 8080)
}

func TestProcessAp(t *testing.T) {
	tstr := try.Success("25380")
	killResult := try.Flatten(
		try.Applicative3(fp.Nop2[string, int](try.Unit1((*os.Process).Kill))).
			ApTry(tstr).
			FlatMap(try.Func1(strconv.Atoi)).
			FlatMap(try.Func1(os.FindProcess)),
	)
	fmt.Println(killResult)
	assert.True(killResult.IsFailure())

}
