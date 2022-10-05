package either_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/either"
)

func TestEither(t *testing.T) {
	l := either.Left[int, float64](10)
	l.Left().Foreach(fp.Println[int])

	s := l.Swap()
	s.Foreach(fp.Println[int])

}
