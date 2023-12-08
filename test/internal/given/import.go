package given

import (
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/seq"
)

//lint:file-ignore U1000 test code

// @fp.ImportGiven
var _ eq.Derives[fp.Eq[any]]

var t1 = seq.Of(as.Tuple2("hello", 10))

var i1 int = 10
