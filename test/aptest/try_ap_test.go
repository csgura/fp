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
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/try"
)

func TestApplicative(t *testing.T) {

	cf := fp.Id3[*url.URL, string, int]

	var intPort fp.Try[int] = try.Chain3(cf).
		ApTry(try.Func1(url.Parse)("http://localhost:8080/abcd")).
		Map((*url.URL).Port).
		FlatMap(try.Func1(strconv.Atoi))
	assert.True(intPort.IsSuccess())
	assert.Equal(intPort.Get(), 8080)
}

func TestProcessAp(t *testing.T) {
	tstr := try.Success("25380")
	killResult := try.Chain4(fp.Id4[string, int, *os.Process, fp.Unit]).
		ApTry(tstr).
		FlatMap(try.Func1(strconv.Atoi)).
		FlatMap(try.Func1(os.FindProcess)).
		FlatMap(try.Unit1((*os.Process).Kill))

	fmt.Println(killResult)
	assert.True(killResult.IsFailure())

}
