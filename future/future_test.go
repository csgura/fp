package future_test

import (
	"testing"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/future"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/promise"
	"github.com/csgura/fp/seq"
)

func TestFuture(t *testing.T) {

	p := promise.New[fp.Option[string]]()

	p.Success(option.Some("hello"))
	fp.Println(p.Future())

	s := seq.Of(1, 2, 3, 4)
	seqFuture := seq.Map(s, func(v int) fp.Future[int] {
		return future.Apply(func() int {
			time.Sleep(100 * time.Millisecond)
			return v * v
		})
	})

	futureSeq := future.Sequence(seqFuture)
	fp.Println(futureSeq)

	ch := make(chan fp.Try[fp.Seq[int]])

	futureSeq.OnComplete(func(result fp.Try[fp.Seq[int]]) {
		ch <- result
	})

	<-ch
	fp.Println(futureSeq)

}
