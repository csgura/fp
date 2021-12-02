package either_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/either"
)

func TestEither(t *testing.T) {
	l := either.Left[string, int]("error")
	l.Left().Foreach(fp.Println[string])

	s := l.Swap()
	s.Foreach(fp.Println[string])

}
