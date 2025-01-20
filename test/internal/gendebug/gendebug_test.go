package gendebug_test

import (
	"testing"

	"github.com/csgura/fp/mshow"
	"github.com/csgura/fp/test/internal/gendebug"
)

func TestShow(t *testing.T) {
	mshow.FullShow(gendebug.ShowInitiatingMessageValue())
}
