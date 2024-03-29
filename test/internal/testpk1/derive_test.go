package testpk1_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/eq"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/test/internal/testpk1"
)

func TestEncoderOption(t *testing.T) {
	str := testpk1.EncoderHasOption().Encode(testpk1.HasOptionMutable{
		Message: "testpk1",
		Addr:    option.None[string](),
		Phone:   []string{"1234", "2345"},
	}.AsImmutable())

	fmt.Println(str)
}

func TestShow(t *testing.T) {
	v := testpk1.WorldMutable{
		Message:   "testpk1",
		Timestamp: time.Now(),
	}.AsImmutable()

	fmt.Println(testpk1.ShowWorld().Show(v))
}

func TestHListInsideHList(t *testing.T) {
	v := testpk1.HListInsideHListMutable{
		Tp:    as.Tuple("10", 10),
		Value: "20",
		Hello: testpk1.WorldMutable{
			Message:   "message is sparta",
			Timestamp: time.Now(),
		}.AsImmutable(),
	}.AsImmutable()

	str := testpk1.ShowHListInsideHList().Show(v)

	fmt.Println(str)
	res := testpk1.ReadHListInsideHList().Read(str)

	res.Failed().Foreach(fp.Println[error])
	assert.True(res.IsSuccess())
	fmt.Println(res)
}

func TestNamedType(t *testing.T) {
	a := testpk1.MySeq[int]{1, 2, 3}
	assert.True(testpk1.EqMySeq(eq.Given[int]()).Eqv(a, a))
	d := testpk1.MonoidMySeq[int]().Combine(a, a)
	assert.True(len(d) == 6)
}

func TestRecursiveEq(t *testing.T) {

	a := testpk1.NodeMutable{
		Value: "root",
		Left: as.Ptr(testpk1.NodeMutable{
			Value: "left",
		}.AsImmutable()),
		Right: as.Ptr(testpk1.NodeMutable{
			Value: "right",
		}.AsImmutable()),
	}.AsImmutable()

	b := testpk1.NodeMutable{
		Value: "root",
		Left: as.Ptr(testpk1.NodeMutable{
			Value: "left",
		}.AsImmutable()),
		Right: as.Ptr(testpk1.NodeMutable{
			Value: "not right",
		}.AsImmutable()),
	}.AsImmutable()

	assert.False(testpk1.EqNode().Eqv(a, b))
}
