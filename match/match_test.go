package match_test

import (
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/match"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/try"
)

func TestMatch(t *testing.T) {

	v := option.Some(11)

	r := match.Of(v,
		match.Case(match.SomeAnd(match.Equal(10)), fp.Const[int]("10")),
		match.Case(match.Some[int](), fp.Const[int]("some")),
		match.CaseNone[int](as.Supplier("none")),
	)

	fmt.Printf("r = %s\n", r)

	t2 := as.Tuple2(option.Some("hello"), option.None[int]())

	r = match.Of(t2,
		match.Case(match.Tuple2(match.Some[string](), match.Some[int]()), as.Tupled2(func(v1 string, v2 int) string {
			return "some,some"
		})),
		match.Case(match.Tuple2(match.Some[string](), match.None[int]()), func(t fp.Entry[fp.Unit]) string {
			return "some,none"

		}),
		match.CaseTuple2(match.None[string](), match.Some[int](), func(u fp.Unit, v int) string {
			return "none,some"
		}),
	)

	fmt.Printf("r = %s\n", r)

	s := seq.Of(1, 2)

	r = match.Of(s,
		match.CaseSeqConsAnd(match.Any[int](), match.SeqHead(match.SomeAnd(match.Equal(2))), func(h int, h2 int) string {
			return "int,2"
		}),
		match.CaseSeqCons(func(h int, tail fp.Seq[int]) string {
			return "head"
		}),
		match.CaseAny(func(v fp.Seq[int]) string {
			return "list"
		}),
	)

	fmt.Printf("r = %s\n", r)

	s = seq.Of[int]()

	r = match.Of(s,
		match.CaseSeqConsAnd(match.Any[int](), match.SeqHead(match.SomeAnd(match.Equal(2))), func(h int, h2 int) string {
			return "int,2"
		}),
		match.CaseSeqCons(func(h int, tail fp.Seq[int]) string {
			return "head"
		}),
		match.CaseSeqEmpty[int](func() string {
			return "empty"
		}),
	)

	fmt.Printf("r = %s\n", r)

	tt := try.Failure[int](fp.Error(404, "server error"))

	tt = tt.RecoverCase(match.Error(
		match.CaseErrorCode(500, fp.Const[error](10)),
		match.CaseErrorCode(404, fp.Const[error](20)),
	))

	fmt.Printf("tt = %s\n", tt)

	tt = try.Failure[int](&net.OpError{})

	tt = tt.RecoverCase(match.Error(
		match.CaseErrorType[*net.OpError](func(err *net.OpError) int {
			return 100
		}),
		match.CaseErrorCode(404, fp.Const[error](20)),
	))
	fmt.Printf("tt = %s\n", tt)

	tt = try.Failure[int](&net.OpError{})

	tt = tt.RecoverCase(match.Error(
		match.CaseErrorType[net.Error](func(err net.Error) int {
			return 1000
		}),
		match.CaseErrorCode(404, fp.Const[error](20)),
	))

	fmt.Printf("tt = %s\n", tt)

	tt = try.Failure[int](&net.OpError{
		Err: io.ErrUnexpectedEOF,
	})

	tt = tt.RecoverCase(match.Error(
		match.CaseErrorIs(io.ErrUnexpectedEOF, func(err error) int {
			return 10000
		}),
		match.CaseErrorCode(404, fp.Const[error](20)),
	))

	fmt.Printf("tt = %s\n", tt)

	tt.RecoverCaseWith(
		match.Error(
			match.CaseErrorCode(500, fp.Const[error](try.Success(10))),
		),
	)

}
