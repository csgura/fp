package try_test

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/curried"
	"github.com/csgura/fp/try"
)

func print[T any](v T) {
	fmt.Println(v)
}

func TestTry(t *testing.T) {
	v := try.Success(10)

	v.Foreach(print[int])

	f2 := try.Success(try.Success(20))
	v = try.Flatten(f2)
	v.Foreach(print[int])

	e := try.Failure[string](fmt.Errorf("bad request"))
	e.Failed().Foreach(print[error])

	e.Recover(func(err error) string {
		return "recover"
	}).Foreach(print[string])

	e.RecoverWith(func(err error) fp.Try[string] {
		return try.Success("recoverWith")
	}).Foreach(print[string])

	v.ToOption().Foreach(print[int])

	// fp.Try[*url.URL]
	var u fp.Try[*url.URL] = try.Func1(url.Parse)("http://[abc")
	fmt.Println(u)

	var p fp.Try[string] = try.Map(u, (*url.URL).Port)

	var intPort fp.Try[int] = try.Flatten(try.Map(p, try.Func1(strconv.Atoi)))
	fmt.Println(intPort)

	try.FlatMap(p, try.Func1(strconv.Atoi)).Foreach(fp.Println[int])

}

func TestFlatMap(t *testing.T) {

	// fp.Try[*url.URL]
	var u fp.Try[*url.URL] = try.Func1(url.Parse)("http://localhost:8080/abcd")
	fmt.Println(u)

	var p fp.Try[string] = try.Map(u, (*url.URL).Port)

	var intPort fp.Try[int] = try.FlatMap(p, try.Func1(strconv.Atoi))
	fmt.Println(intPort)

}

func TestApplicative(t *testing.T) {

	cf := curried.Concat[*url.URL](curried.Concat[string](fp.Id[int]))

	var intPort fp.Try[int] = try.Applicative3(curried.Revert3(cf)).
		ApTry(try.Func1(url.Parse)("http://localhost:8080/abcd")).
		Map((*url.URL).Port).
		FlatMap(try.Func1(strconv.Atoi))
	fmt.Println(intPort)

}

func TestCompose(t *testing.T) {

	var intPort fp.Try[int] = try.Compose(
		try.Func1(url.Parse),
		fp.Compose((*url.URL).Port, try.Func1(strconv.Atoi)),
	)("http://localhost:8080/abcd")

	fmt.Println(intPort)

}

func ParsePort() (int, error) {
	u, err := url.Parse("http://localhost:8080/abcd")
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(u.Port())
}

func ParsePortFn() fp.Try[int] {
	return try.Compose(
		try.Func1(url.Parse),
		fp.Compose((*url.URL).Port, try.Func1(strconv.Atoi)),
	)("http://localhost:8080/abcd")
}

func TestProcessAp(t *testing.T) {
	tstr := try.Success("25380")
	killResult := try.Applicative3(fp.Nop2[string, int](try.Unit1((*os.Process).Kill))).
		ApTry(tstr).
		FlatMap(try.Func1(strconv.Atoi))
	fmt.Println(killResult)

}
