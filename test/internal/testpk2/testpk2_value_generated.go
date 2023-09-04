// Code generated by gombok, DO NOT EDIT.
package testpk2

import (
	"encoding/json"
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/test/internal/testpk1"
	"io"
	"net/http"
	"os"
	"reflect"
	"sync/atomic"
)

type HelloBuilder Hello

type HelloMutable struct {
	World string `json:"world,omitempty"`
	Hi    int    `bson:"hi" json:"merong"`
}

func (r HelloBuilder) Build() Hello {
	return Hello(r)
}

func (r Hello) Builder() HelloBuilder {
	return HelloBuilder(r)
}

func (r Hello) World() string {
	return r.world
}

func (r Hello) Hi() int {
	return r.hi
}

func (r Hello) WithWorld(v string) Hello {
	r.world = v
	return r
}

func (r Hello) WithHi(v int) Hello {
	r.hi = v
	return r
}

func (r HelloBuilder) World(v string) HelloBuilder {
	r.world = v
	return r
}

func (r HelloBuilder) Hi(v int) HelloBuilder {
	r.hi = v
	return r
}

func (r Hello) String() string {
	return fmt.Sprintf("Hello(world=%v, hi=%v)", r.world, r.hi)
}

func (r Hello) AsTuple() fp.Tuple2[string, int] {
	return as.Tuple2(r.world, r.hi)
}

func (r Hello) Unapply() (string, int) {
	return r.world, r.hi
}

func (r Hello) AsMutable() HelloMutable {
	return HelloMutable{
		World: r.world,
		Hi:    r.hi,
	}
}

func (r HelloMutable) AsImmutable() Hello {
	return Hello{
		world: r.World,
		hi:    r.Hi,
	}
}

func (r HelloBuilder) FromTuple(t fp.Tuple2[string, int]) HelloBuilder {
	r.world = t.I1
	r.hi = t.I2
	return r
}

func (r HelloBuilder) Apply(world string, hi int) HelloBuilder {
	r.world = world
	r.hi = hi
	return r
}

func (r Hello) AsMap() map[string]any {
	m := map[string]any{}
	m["world"] = r.world
	m["hi"] = r.hi
	return m
}

func (r HelloBuilder) FromMap(m map[string]any) HelloBuilder {

	if v, ok := m["world"].(string); ok {
		r.world = v
	}

	if v, ok := m["hi"].(int); ok {
		r.hi = v
	}

	return r
}

type AllKindTypesBuilder AllKindTypes

type AllKindTypesMutable struct {
	Embed
	Hi   fp.Option[int]
	Tpe  reflect.Type
	Arr  []os.File
	M    map[string]int
	A    any
	P    *int
	L    Local
	T    fp.Try[fp.Option[Local]]
	M2   map[string]atomic.Bool
	Mm   fp.Map[string, int]
	Intf fp.Future[int]
	Ch   chan fp.Try[fp.Either[int, string]]
	Ch2  chan<- int
	Ch3  <-chan int
	Fn3  fp.Func1[int, fp.Try[string]]
	Fn   func(a string) fp.Try[int]
	Fn2  func(fp.Try[string]) (result int, err error)
	Arr2 [2]int
	St   struct {
		Embed
		A int
		B fp.Option[string]
	}
	I2 interface {
		io.Closer
		Hello() fp.Try[int]
	}
}

func (r AllKindTypesBuilder) Build() AllKindTypes {
	return AllKindTypes(r)
}

func (r AllKindTypes) Builder() AllKindTypesBuilder {
	return AllKindTypesBuilder(r)
}

func (r AllKindTypes) Hi() fp.Option[int] {
	return r.hi
}

func (r AllKindTypes) Tpe() reflect.Type {
	return r.tpe
}

func (r AllKindTypes) Arr() []os.File {
	return r.arr
}

func (r AllKindTypes) M() map[string]int {
	return r.m
}

func (r AllKindTypes) A() any {
	return r.a
}

func (r AllKindTypes) P() *int {
	return r.p
}

func (r AllKindTypes) L() Local {
	return r.l
}

func (r AllKindTypes) T() fp.Try[fp.Option[Local]] {
	return r.t
}

func (r AllKindTypes) M2() map[string]atomic.Bool {
	return r.m2
}

func (r AllKindTypes) Mm() fp.Map[string, int] {
	return r.mm
}

func (r AllKindTypes) Intf() fp.Future[int] {
	return r.intf
}

func (r AllKindTypes) Ch() chan fp.Try[fp.Either[int, string]] {
	return r.ch
}

func (r AllKindTypes) Ch2() chan<- int {
	return r.ch2
}

func (r AllKindTypes) Ch3() <-chan int {
	return r.ch3
}

func (r AllKindTypes) Fn3() fp.Func1[int, fp.Try[string]] {
	return r.fn3
}

func (r AllKindTypes) Fn() func(a string) fp.Try[int] {
	return r.fn
}

func (r AllKindTypes) Fn2() func(fp.Try[string]) (result int, err error) {
	return r.fn2
}

func (r AllKindTypes) Arr2() [2]int {
	return r.arr2
}

func (r AllKindTypes) St() struct {
	Embed
	A int
	B fp.Option[string]
} {
	return r.st
}

func (r AllKindTypes) I2() interface {
	io.Closer
	Hello() fp.Try[int]
} {
	return r.i2
}

func (r AllKindTypes) WithHi(v fp.Option[int]) AllKindTypes {
	r.hi = v
	return r
}

func (r AllKindTypes) WithSomeHi(v int) AllKindTypes {
	r.hi = option.Some(v)
	return r
}

func (r AllKindTypes) WithNoneHi() AllKindTypes {
	r.hi = option.None[int]()
	return r
}

func (r AllKindTypes) WithTpe(v reflect.Type) AllKindTypes {
	r.tpe = v
	return r
}

func (r AllKindTypes) WithArr(v []os.File) AllKindTypes {
	r.arr = v
	return r
}

func (r AllKindTypes) WithM(v map[string]int) AllKindTypes {
	r.m = v
	return r
}

func (r AllKindTypes) WithA(v any) AllKindTypes {
	r.a = v
	return r
}

func (r AllKindTypes) WithP(v *int) AllKindTypes {
	r.p = v
	return r
}

func (r AllKindTypes) WithL(v Local) AllKindTypes {
	r.l = v
	return r
}

func (r AllKindTypes) WithT(v fp.Try[fp.Option[Local]]) AllKindTypes {
	r.t = v
	return r
}

func (r AllKindTypes) WithM2(v map[string]atomic.Bool) AllKindTypes {
	r.m2 = v
	return r
}

func (r AllKindTypes) WithMm(v fp.Map[string, int]) AllKindTypes {
	r.mm = v
	return r
}

func (r AllKindTypes) WithIntf(v fp.Future[int]) AllKindTypes {
	r.intf = v
	return r
}

func (r AllKindTypes) WithCh(v chan fp.Try[fp.Either[int, string]]) AllKindTypes {
	r.ch = v
	return r
}

func (r AllKindTypes) WithCh2(v chan<- int) AllKindTypes {
	r.ch2 = v
	return r
}

func (r AllKindTypes) WithCh3(v <-chan int) AllKindTypes {
	r.ch3 = v
	return r
}

func (r AllKindTypes) WithFn3(v fp.Func1[int, fp.Try[string]]) AllKindTypes {
	r.fn3 = v
	return r
}

func (r AllKindTypes) WithFn(v func(a string) fp.Try[int]) AllKindTypes {
	r.fn = v
	return r
}

func (r AllKindTypes) WithFn2(v func(fp.Try[string]) (result int, err error)) AllKindTypes {
	r.fn2 = v
	return r
}

func (r AllKindTypes) WithArr2(v [2]int) AllKindTypes {
	r.arr2 = v
	return r
}

func (r AllKindTypes) WithSt(v struct {
	Embed
	A int
	B fp.Option[string]
}) AllKindTypes {
	r.st = v
	return r
}

func (r AllKindTypes) WithI2(v interface {
	io.Closer
	Hello() fp.Try[int]
}) AllKindTypes {
	r.i2 = v
	return r
}

func (r AllKindTypesBuilder) Hi(v fp.Option[int]) AllKindTypesBuilder {
	r.hi = v
	return r
}

func (r AllKindTypesBuilder) SomeHi(v int) AllKindTypesBuilder {
	r.hi = option.Some(v)
	return r
}

func (r AllKindTypesBuilder) NoneHi() AllKindTypesBuilder {
	r.hi = option.None[int]()
	return r
}

func (r AllKindTypesBuilder) Tpe(v reflect.Type) AllKindTypesBuilder {
	r.tpe = v
	return r
}

func (r AllKindTypesBuilder) Arr(v []os.File) AllKindTypesBuilder {
	r.arr = v
	return r
}

func (r AllKindTypesBuilder) M(v map[string]int) AllKindTypesBuilder {
	r.m = v
	return r
}

func (r AllKindTypesBuilder) A(v any) AllKindTypesBuilder {
	r.a = v
	return r
}

func (r AllKindTypesBuilder) P(v *int) AllKindTypesBuilder {
	r.p = v
	return r
}

func (r AllKindTypesBuilder) L(v Local) AllKindTypesBuilder {
	r.l = v
	return r
}

func (r AllKindTypesBuilder) T(v fp.Try[fp.Option[Local]]) AllKindTypesBuilder {
	r.t = v
	return r
}

func (r AllKindTypesBuilder) M2(v map[string]atomic.Bool) AllKindTypesBuilder {
	r.m2 = v
	return r
}

func (r AllKindTypesBuilder) Mm(v fp.Map[string, int]) AllKindTypesBuilder {
	r.mm = v
	return r
}

func (r AllKindTypesBuilder) Intf(v fp.Future[int]) AllKindTypesBuilder {
	r.intf = v
	return r
}

func (r AllKindTypesBuilder) Ch(v chan fp.Try[fp.Either[int, string]]) AllKindTypesBuilder {
	r.ch = v
	return r
}

func (r AllKindTypesBuilder) Ch2(v chan<- int) AllKindTypesBuilder {
	r.ch2 = v
	return r
}

func (r AllKindTypesBuilder) Ch3(v <-chan int) AllKindTypesBuilder {
	r.ch3 = v
	return r
}

func (r AllKindTypesBuilder) Fn3(v fp.Func1[int, fp.Try[string]]) AllKindTypesBuilder {
	r.fn3 = v
	return r
}

func (r AllKindTypesBuilder) Fn(v func(a string) fp.Try[int]) AllKindTypesBuilder {
	r.fn = v
	return r
}

func (r AllKindTypesBuilder) Fn2(v func(fp.Try[string]) (result int, err error)) AllKindTypesBuilder {
	r.fn2 = v
	return r
}

func (r AllKindTypesBuilder) Arr2(v [2]int) AllKindTypesBuilder {
	r.arr2 = v
	return r
}

func (r AllKindTypesBuilder) St(v struct {
	Embed
	A int
	B fp.Option[string]
}) AllKindTypesBuilder {
	r.st = v
	return r
}

func (r AllKindTypesBuilder) I2(v interface {
	io.Closer
	Hello() fp.Try[int]
}) AllKindTypesBuilder {
	r.i2 = v
	return r
}

func (r AllKindTypes) String() string {
	return fmt.Sprintf("AllKindTypes(hi=%v, tpe=%v, arr=%v, m=%v, a=%v, p=%v, l=%v, t=%v, m2=%v, mm=%v, intf=%v, ch=%v, ch2=%v, ch3=%v, fn3=%v, arr2=%v, st=%v, i2=%v)", r.hi, r.tpe, r.arr, r.m, r.a, r.p, r.l, r.t, r.m2, r.mm, r.intf, r.ch, r.ch2, r.ch3, r.fn3, r.arr2, r.st, r.i2)
}

func (r AllKindTypes) AsTuple() fp.Tuple20[fp.Option[int], reflect.Type, []os.File, map[string]int, any, *int, Local, fp.Try[fp.Option[Local]], map[string]atomic.Bool, fp.Map[string, int], fp.Future[int], chan fp.Try[fp.Either[int, string]], chan<- int, <-chan int, fp.Func1[int, fp.Try[string]], func(a string) fp.Try[int], func(fp.Try[string]) (result int, err error), [2]int, struct {
	Embed
	A int
	B fp.Option[string]
}, interface {
	io.Closer
	Hello() fp.Try[int]
}] {
	return as.Tuple20(r.hi, r.tpe, r.arr, r.m, r.a, r.p, r.l, r.t, r.m2, r.mm, r.intf, r.ch, r.ch2, r.ch3, r.fn3, r.fn, r.fn2, r.arr2, r.st, r.i2)
}

func (r AllKindTypes) Unapply() (fp.Option[int], reflect.Type, []os.File, map[string]int, any, *int, Local, fp.Try[fp.Option[Local]], map[string]atomic.Bool, fp.Map[string, int], fp.Future[int], chan fp.Try[fp.Either[int, string]], chan<- int, <-chan int, fp.Func1[int, fp.Try[string]], func(a string) fp.Try[int], func(fp.Try[string]) (result int, err error), [2]int, struct {
	Embed
	A int
	B fp.Option[string]
}, interface {
	io.Closer
	Hello() fp.Try[int]
}) {
	return r.hi, r.tpe, r.arr, r.m, r.a, r.p, r.l, r.t, r.m2, r.mm, r.intf, r.ch, r.ch2, r.ch3, r.fn3, r.fn, r.fn2, r.arr2, r.st, r.i2
}

func (r AllKindTypes) AsMutable() AllKindTypesMutable {
	return AllKindTypesMutable{
		Hi:   r.hi,
		Tpe:  r.tpe,
		Arr:  r.arr,
		M:    r.m,
		A:    r.a,
		P:    r.p,
		L:    r.l,
		T:    r.t,
		M2:   r.m2,
		Mm:   r.mm,
		Intf: r.intf,
		Ch:   r.ch,
		Ch2:  r.ch2,
		Ch3:  r.ch3,
		Fn3:  r.fn3,
		Fn:   r.fn,
		Fn2:  r.fn2,
		Arr2: r.arr2,
		St:   r.st,
		I2:   r.i2,
	}
}

func (r AllKindTypesMutable) AsImmutable() AllKindTypes {
	return AllKindTypes{
		hi:   r.Hi,
		tpe:  r.Tpe,
		arr:  r.Arr,
		m:    r.M,
		a:    r.A,
		p:    r.P,
		l:    r.L,
		t:    r.T,
		m2:   r.M2,
		mm:   r.Mm,
		intf: r.Intf,
		ch:   r.Ch,
		ch2:  r.Ch2,
		ch3:  r.Ch3,
		fn3:  r.Fn3,
		fn:   r.Fn,
		fn2:  r.Fn2,
		arr2: r.Arr2,
		st:   r.St,
		i2:   r.I2,
	}
}

func (r AllKindTypesBuilder) FromTuple(t fp.Tuple20[fp.Option[int], reflect.Type, []os.File, map[string]int, any, *int, Local, fp.Try[fp.Option[Local]], map[string]atomic.Bool, fp.Map[string, int], fp.Future[int], chan fp.Try[fp.Either[int, string]], chan<- int, <-chan int, fp.Func1[int, fp.Try[string]], func(a string) fp.Try[int], func(fp.Try[string]) (result int, err error), [2]int, struct {
	Embed
	A int
	B fp.Option[string]
}, interface {
	io.Closer
	Hello() fp.Try[int]
}]) AllKindTypesBuilder {
	r.hi = t.I1
	r.tpe = t.I2
	r.arr = t.I3
	r.m = t.I4
	r.a = t.I5
	r.p = t.I6
	r.l = t.I7
	r.t = t.I8
	r.m2 = t.I9
	r.mm = t.I10
	r.intf = t.I11
	r.ch = t.I12
	r.ch2 = t.I13
	r.ch3 = t.I14
	r.fn3 = t.I15
	r.fn = t.I16
	r.fn2 = t.I17
	r.arr2 = t.I18
	r.st = t.I19
	r.i2 = t.I20
	return r
}

func (r AllKindTypesBuilder) Apply(hi fp.Option[int], tpe reflect.Type, arr []os.File, m map[string]int, a any, p *int, l Local, t fp.Try[fp.Option[Local]], m2 map[string]atomic.Bool, mm fp.Map[string, int], intf fp.Future[int], ch chan fp.Try[fp.Either[int, string]], ch2 chan<- int, ch3 <-chan int, fn3 fp.Func1[int, fp.Try[string]], fn func(a string) fp.Try[int], fn2 func(fp.Try[string]) (result int, err error), arr2 [2]int, st struct {
	Embed
	A int
	B fp.Option[string]
}, i2 interface {
	io.Closer
	Hello() fp.Try[int]
}) AllKindTypesBuilder {
	r.hi = hi
	r.tpe = tpe
	r.arr = arr
	r.m = m
	r.a = a
	r.p = p
	r.l = l
	r.t = t
	r.m2 = m2
	r.mm = mm
	r.intf = intf
	r.ch = ch
	r.ch2 = ch2
	r.ch3 = ch3
	r.fn3 = fn3
	r.fn = fn
	r.fn2 = fn2
	r.arr2 = arr2
	r.st = st
	r.i2 = i2
	return r
}

func (r AllKindTypes) AsMap() map[string]any {
	m := map[string]any{}
	if r.hi.IsDefined() {
		m["hi"] = r.hi.Get()
	}
	m["tpe"] = r.tpe
	m["arr"] = r.arr
	m["m"] = r.m
	m["a"] = r.a
	m["p"] = r.p
	m["l"] = r.l
	m["t"] = r.t
	m["m2"] = r.m2
	m["mm"] = r.mm
	m["intf"] = r.intf
	m["ch"] = r.ch
	m["ch2"] = r.ch2
	m["ch3"] = r.ch3
	m["fn3"] = r.fn3
	m["fn"] = r.fn
	m["fn2"] = r.fn2
	m["arr2"] = r.arr2
	m["st"] = r.st
	m["i2"] = r.i2
	return m
}

func (r AllKindTypesBuilder) FromMap(m map[string]any) AllKindTypesBuilder {

	if v, ok := m["hi"].(fp.Option[int]); ok {
		r.hi = v
	} else if v, ok := m["hi"].(int); ok {
		r.hi = option.Some(v)
	}

	if v, ok := m["tpe"].(reflect.Type); ok {
		r.tpe = v
	}

	if v, ok := m["arr"].([]os.File); ok {
		r.arr = v
	}

	if v, ok := m["m"].(map[string]int); ok {
		r.m = v
	}

	if v, ok := m["a"].(any); ok {
		r.a = v
	}

	if v, ok := m["p"].(*int); ok {
		r.p = v
	}

	if v, ok := m["l"].(Local); ok {
		r.l = v
	}

	if v, ok := m["t"].(fp.Try[fp.Option[Local]]); ok {
		r.t = v
	}

	if v, ok := m["m2"].(map[string]atomic.Bool); ok {
		r.m2 = v
	}

	if v, ok := m["mm"].(fp.Map[string, int]); ok {
		r.mm = v
	}

	if v, ok := m["intf"].(fp.Future[int]); ok {
		r.intf = v
	}

	if v, ok := m["ch"].(chan fp.Try[fp.Either[int, string]]); ok {
		r.ch = v
	}

	if v, ok := m["ch2"].(chan<- int); ok {
		r.ch2 = v
	}

	if v, ok := m["ch3"].(<-chan int); ok {
		r.ch3 = v
	}

	if v, ok := m["fn3"].(fp.Func1[int, fp.Try[string]]); ok {
		r.fn3 = v
	}

	if v, ok := m["fn"].(func(a string) fp.Try[int]); ok {
		r.fn = v
	}

	if v, ok := m["fn2"].(func(fp.Try[string]) (result int, err error)); ok {
		r.fn2 = v
	}

	if v, ok := m["arr2"].([2]int); ok {
		r.arr2 = v
	}

	if v, ok := m["st"].(struct {
		Embed
		A int
		B fp.Option[string]
	}); ok {
		r.st = v
	}

	if v, ok := m["i2"].(interface {
		io.Closer
		Hello() fp.Try[int]
	}); ok {
		r.i2 = v
	}

	return r
}

type PersonBuilder Person

type PersonMutable struct {
	Name       string
	Age        int
	Height     float64
	Phone      fp.Option[string]
	Addr       []string
	List       hlist.Cons[string, hlist.Cons[int, hlist.Nil]]
	Seq        fp.Seq[float64]
	Blob       []byte
	_notExport string
}

func (r PersonBuilder) Build() Person {
	return Person(r)
}

func (r Person) Builder() PersonBuilder {
	return PersonBuilder(r)
}

func (r Person) Name() string {
	return r.name
}

func (r Person) Age() int {
	return r.age
}

func (r Person) Height() float64 {
	return r.height
}

func (r Person) Phone() fp.Option[string] {
	return r.phone
}

func (r Person) Addr() []string {
	return r.addr
}

func (r Person) List() hlist.Cons[string, hlist.Cons[int, hlist.Nil]] {
	return r.list
}

func (r Person) Seq() fp.Seq[float64] {
	return r.seq
}

func (r Person) Blob() []byte {
	return r.blob
}

func (r Person) WithName(v string) Person {
	r.name = v
	return r
}

func (r Person) WithAge(v int) Person {
	r.age = v
	return r
}

func (r Person) WithHeight(v float64) Person {
	r.height = v
	return r
}

func (r Person) WithPhone(v fp.Option[string]) Person {
	r.phone = v
	return r
}

func (r Person) WithSomePhone(v string) Person {
	r.phone = option.Some(v)
	return r
}

func (r Person) WithNonePhone() Person {
	r.phone = option.None[string]()
	return r
}

func (r Person) WithAddr(v []string) Person {
	r.addr = v
	return r
}

func (r Person) WithList(v hlist.Cons[string, hlist.Cons[int, hlist.Nil]]) Person {
	r.list = v
	return r
}

func (r Person) WithSeq(v fp.Seq[float64]) Person {
	r.seq = v
	return r
}

func (r Person) WithBlob(v []byte) Person {
	r.blob = v
	return r
}

func (r PersonBuilder) Name(v string) PersonBuilder {
	r.name = v
	return r
}

func (r PersonBuilder) Age(v int) PersonBuilder {
	r.age = v
	return r
}

func (r PersonBuilder) Height(v float64) PersonBuilder {
	r.height = v
	return r
}

func (r PersonBuilder) Phone(v fp.Option[string]) PersonBuilder {
	r.phone = v
	return r
}

func (r PersonBuilder) SomePhone(v string) PersonBuilder {
	r.phone = option.Some(v)
	return r
}

func (r PersonBuilder) NonePhone() PersonBuilder {
	r.phone = option.None[string]()
	return r
}

func (r PersonBuilder) Addr(v []string) PersonBuilder {
	r.addr = v
	return r
}

func (r PersonBuilder) List(v hlist.Cons[string, hlist.Cons[int, hlist.Nil]]) PersonBuilder {
	r.list = v
	return r
}

func (r PersonBuilder) Seq(v fp.Seq[float64]) PersonBuilder {
	r.seq = v
	return r
}

func (r PersonBuilder) Blob(v []byte) PersonBuilder {
	r.blob = v
	return r
}

func (r Person) String() string {
	return fmt.Sprintf("Person(name=%v, age=%v, height=%v, phone=%v, addr=%v, list=%v, seq=%v, blob=%v)", r.name, r.age, r.height, r.phone, r.addr, r.list, r.seq, r.blob)
}

func (r Person) AsTuple() fp.Tuple8[string, int, float64, fp.Option[string], []string, hlist.Cons[string, hlist.Cons[int, hlist.Nil]], fp.Seq[float64], []byte] {
	return as.Tuple8(r.name, r.age, r.height, r.phone, r.addr, r.list, r.seq, r.blob)
}

func (r Person) Unapply() (string, int, float64, fp.Option[string], []string, hlist.Cons[string, hlist.Cons[int, hlist.Nil]], fp.Seq[float64], []byte) {
	return r.name, r.age, r.height, r.phone, r.addr, r.list, r.seq, r.blob
}

func (r Person) AsMutable() PersonMutable {
	return PersonMutable{
		Name:   r.name,
		Age:    r.age,
		Height: r.height,
		Phone:  r.phone,
		Addr:   r.addr,
		List:   r.list,
		Seq:    r.seq,
		Blob:   r.blob,
	}
}

func (r PersonMutable) AsImmutable() Person {
	return Person{
		name:   r.Name,
		age:    r.Age,
		height: r.Height,
		phone:  r.Phone,
		addr:   r.Addr,
		list:   r.List,
		seq:    r.Seq,
		blob:   r.Blob,
	}
}

func (r PersonBuilder) FromTuple(t fp.Tuple8[string, int, float64, fp.Option[string], []string, hlist.Cons[string, hlist.Cons[int, hlist.Nil]], fp.Seq[float64], []byte]) PersonBuilder {
	r.name = t.I1
	r.age = t.I2
	r.height = t.I3
	r.phone = t.I4
	r.addr = t.I5
	r.list = t.I6
	r.seq = t.I7
	r.blob = t.I8
	return r
}

func (r PersonBuilder) Apply(name string, age int, height float64, phone fp.Option[string], addr []string, list hlist.Cons[string, hlist.Cons[int, hlist.Nil]], seq fp.Seq[float64], blob []byte) PersonBuilder {
	r.name = name
	r.age = age
	r.height = height
	r.phone = phone
	r.addr = addr
	r.list = list
	r.seq = seq
	r.blob = blob
	return r
}

func (r Person) AsMap() map[string]any {
	m := map[string]any{}
	m["name"] = r.name
	m["age"] = r.age
	m["height"] = r.height
	if r.phone.IsDefined() {
		m["phone"] = r.phone.Get()
	}
	m["addr"] = r.addr
	m["list"] = r.list
	m["seq"] = r.seq
	m["blob"] = r.blob
	return m
}

func (r PersonBuilder) FromMap(m map[string]any) PersonBuilder {

	if v, ok := m["name"].(string); ok {
		r.name = v
	}

	if v, ok := m["age"].(int); ok {
		r.age = v
	}

	if v, ok := m["height"].(float64); ok {
		r.height = v
	}

	if v, ok := m["phone"].(fp.Option[string]); ok {
		r.phone = v
	} else if v, ok := m["phone"].(string); ok {
		r.phone = option.Some(v)
	}

	if v, ok := m["addr"].([]string); ok {
		r.addr = v
	}

	if v, ok := m["list"].(hlist.Cons[string, hlist.Cons[int, hlist.Nil]]); ok {
		r.list = v
	}

	if v, ok := m["seq"].(fp.Seq[float64]); ok {
		r.seq = v
	}

	if v, ok := m["blob"].([]byte); ok {
		r.blob = v
	}

	return r
}

type WalletBuilder Wallet

type WalletMutable struct {
	Owner  Person
	Amount int64
}

func (r WalletBuilder) Build() Wallet {
	return Wallet(r)
}

func (r Wallet) Builder() WalletBuilder {
	return WalletBuilder(r)
}

func (r Wallet) Owner() Person {
	return r.owner
}

func (r Wallet) Amount() int64 {
	return r.amount
}

func (r Wallet) WithOwner(v Person) Wallet {
	r.owner = v
	return r
}

func (r Wallet) WithAmount(v int64) Wallet {
	r.amount = v
	return r
}

func (r WalletBuilder) Owner(v Person) WalletBuilder {
	r.owner = v
	return r
}

func (r WalletBuilder) Amount(v int64) WalletBuilder {
	r.amount = v
	return r
}

func (r Wallet) String() string {
	return fmt.Sprintf("Wallet(owner=%v, amount=%v)", r.owner, r.amount)
}

func (r Wallet) AsTuple() fp.Tuple2[Person, int64] {
	return as.Tuple2(r.owner, r.amount)
}

func (r Wallet) Unapply() (Person, int64) {
	return r.owner, r.amount
}

func (r Wallet) AsMutable() WalletMutable {
	return WalletMutable{
		Owner:  r.owner,
		Amount: r.amount,
	}
}

func (r WalletMutable) AsImmutable() Wallet {
	return Wallet{
		owner:  r.Owner,
		amount: r.Amount,
	}
}

func (r WalletBuilder) FromTuple(t fp.Tuple2[Person, int64]) WalletBuilder {
	r.owner = t.I1
	r.amount = t.I2
	return r
}

func (r WalletBuilder) Apply(owner Person, amount int64) WalletBuilder {
	r.owner = owner
	r.amount = amount
	return r
}

func (r Wallet) AsMap() map[string]any {
	m := map[string]any{}
	m["owner"] = r.owner
	m["amount"] = r.amount
	return m
}

func (r WalletBuilder) FromMap(m map[string]any) WalletBuilder {

	if v, ok := m["owner"].(Person); ok {
		r.owner = v
	}

	if v, ok := m["amount"].(int64); ok {
		r.amount = v
	}

	return r
}

type EntryBuilder[A comparable, B any, C fmt.Stringer, D interface {
	Hello() string
}] Entry[A, B, C, D]

type EntryMutable[A comparable, B any, C fmt.Stringer, D interface {
	Hello() string
}] struct {
	Name  string
	Value A
	Tuple fp.Tuple2[A, B]
}

func (r EntryBuilder[A, B, C, D]) Build() Entry[A, B, C, D] {
	return Entry[A, B, C, D](r)
}

func (r Entry[A, B, C, D]) Builder() EntryBuilder[A, B, C, D] {
	return EntryBuilder[A, B, C, D](r)
}

func (r Entry[A, B, C, D]) Name() string {
	return r.name
}

func (r Entry[A, B, C, D]) Value() A {
	return r.value
}

func (r Entry[A, B, C, D]) Tuple() fp.Tuple2[A, B] {
	return r.tuple
}

func (r Entry[A, B, C, D]) WithName(v string) Entry[A, B, C, D] {
	r.name = v
	return r
}

func (r Entry[A, B, C, D]) WithValue(v A) Entry[A, B, C, D] {
	r.value = v
	return r
}

func (r Entry[A, B, C, D]) WithTuple(v fp.Tuple2[A, B]) Entry[A, B, C, D] {
	r.tuple = v
	return r
}

func (r EntryBuilder[A, B, C, D]) Name(v string) EntryBuilder[A, B, C, D] {
	r.name = v
	return r
}

func (r EntryBuilder[A, B, C, D]) Value(v A) EntryBuilder[A, B, C, D] {
	r.value = v
	return r
}

func (r EntryBuilder[A, B, C, D]) Tuple(v fp.Tuple2[A, B]) EntryBuilder[A, B, C, D] {
	r.tuple = v
	return r
}

func (r Entry[A, B, C, D]) String() string {
	return fmt.Sprintf("Entry(name=%v, value=%v, tuple=%v)", r.name, r.value, r.tuple)
}

func (r Entry[A, B, C, D]) AsTuple() fp.Tuple3[string, A, fp.Tuple2[A, B]] {
	return as.Tuple3(r.name, r.value, r.tuple)
}

func (r Entry[A, B, C, D]) Unapply() (string, A, fp.Tuple2[A, B]) {
	return r.name, r.value, r.tuple
}

func (r Entry[A, B, C, D]) AsMutable() EntryMutable[A, B, C, D] {
	return EntryMutable[A, B, C, D]{
		Name:  r.name,
		Value: r.value,
		Tuple: r.tuple,
	}
}

func (r EntryMutable[A, B, C, D]) AsImmutable() Entry[A, B, C, D] {
	return Entry[A, B, C, D]{
		name:  r.Name,
		value: r.Value,
		tuple: r.Tuple,
	}
}

func (r EntryBuilder[A, B, C, D]) FromTuple(t fp.Tuple3[string, A, fp.Tuple2[A, B]]) EntryBuilder[A, B, C, D] {
	r.name = t.I1
	r.value = t.I2
	r.tuple = t.I3
	return r
}

func (r EntryBuilder[A, B, C, D]) Apply(name string, value A, tuple fp.Tuple2[A, B]) EntryBuilder[A, B, C, D] {
	r.name = name
	r.value = value
	r.tuple = tuple
	return r
}

func (r Entry[A, B, C, D]) AsMap() map[string]any {
	m := map[string]any{}
	m["name"] = r.name
	m["value"] = r.value
	m["tuple"] = r.tuple
	return m
}

func (r EntryBuilder[A, B, C, D]) FromMap(m map[string]any) EntryBuilder[A, B, C, D] {

	if v, ok := m["name"].(string); ok {
		r.name = v
	}

	if v, ok := m["value"].(A); ok {
		r.value = v
	}

	if v, ok := m["tuple"].(fp.Tuple2[A, B]); ok {
		r.tuple = v
	}

	return r
}

type KeyBuilder Key

type KeyMutable struct {
	A int
	B float32
	C []byte
}

func (r KeyBuilder) Build() Key {
	return Key(r)
}

func (r Key) Builder() KeyBuilder {
	return KeyBuilder(r)
}

func (r Key) A() int {
	return r.a
}

func (r Key) B() float32 {
	return r.b
}

func (r Key) C() []byte {
	return r.c
}

func (r Key) WithA(v int) Key {
	r.a = v
	return r
}

func (r Key) WithB(v float32) Key {
	r.b = v
	return r
}

func (r Key) WithC(v []byte) Key {
	r.c = v
	return r
}

func (r KeyBuilder) A(v int) KeyBuilder {
	r.a = v
	return r
}

func (r KeyBuilder) B(v float32) KeyBuilder {
	r.b = v
	return r
}

func (r KeyBuilder) C(v []byte) KeyBuilder {
	r.c = v
	return r
}

func (r Key) String() string {
	return fmt.Sprintf("Key(a=%v, b=%v, c=%v)", r.a, r.b, r.c)
}

func (r Key) AsTuple() fp.Tuple3[int, float32, []byte] {
	return as.Tuple3(r.a, r.b, r.c)
}

func (r Key) Unapply() (int, float32, []byte) {
	return r.a, r.b, r.c
}

func (r Key) AsMutable() KeyMutable {
	return KeyMutable{
		A: r.a,
		B: r.b,
		C: r.c,
	}
}

func (r KeyMutable) AsImmutable() Key {
	return Key{
		a: r.A,
		b: r.B,
		c: r.C,
	}
}

func (r KeyBuilder) FromTuple(t fp.Tuple3[int, float32, []byte]) KeyBuilder {
	r.a = t.I1
	r.b = t.I2
	r.c = t.I3
	return r
}

func (r KeyBuilder) Apply(a int, b float32, c []byte) KeyBuilder {
	r.a = a
	r.b = b
	r.c = c
	return r
}

func (r Key) AsMap() map[string]any {
	m := map[string]any{}
	m["a"] = r.a
	m["b"] = r.b
	m["c"] = r.c
	return m
}

func (r KeyBuilder) FromMap(m map[string]any) KeyBuilder {

	if v, ok := m["a"].(int); ok {
		r.a = v
	}

	if v, ok := m["b"].(float32); ok {
		r.b = v
	}

	if v, ok := m["c"].([]byte); ok {
		r.c = v
	}

	return r
}

type PointBuilder Point

type PointMutable struct {
	X int
	Y int
	Z fp.Tuple2[int, int]
}

func (r PointBuilder) Build() Point {
	return Point(r)
}

func (r Point) Builder() PointBuilder {
	return PointBuilder(r)
}

func (r Point) X() int {
	return r.x
}

func (r Point) Y() int {
	return r.y
}

func (r Point) Z() fp.Tuple2[int, int] {
	return r.z
}

func (r Point) WithX(v int) Point {
	r.x = v
	return r
}

func (r Point) WithY(v int) Point {
	r.y = v
	return r
}

func (r Point) WithZ(v fp.Tuple2[int, int]) Point {
	r.z = v
	return r
}

func (r PointBuilder) X(v int) PointBuilder {
	r.x = v
	return r
}

func (r PointBuilder) Y(v int) PointBuilder {
	r.y = v
	return r
}

func (r PointBuilder) Z(v fp.Tuple2[int, int]) PointBuilder {
	r.z = v
	return r
}

func (r Point) AsTuple() fp.Tuple3[int, int, fp.Tuple2[int, int]] {
	return as.Tuple3(r.x, r.y, r.z)
}

func (r Point) Unapply() (int, int, fp.Tuple2[int, int]) {
	return r.x, r.y, r.z
}

func (r Point) AsMutable() PointMutable {
	return PointMutable{
		X: r.x,
		Y: r.y,
		Z: r.z,
	}
}

func (r PointMutable) AsImmutable() Point {
	return Point{
		x: r.X,
		y: r.Y,
		z: r.Z,
	}
}

func (r PointBuilder) FromTuple(t fp.Tuple3[int, int, fp.Tuple2[int, int]]) PointBuilder {
	r.x = t.I1
	r.y = t.I2
	r.z = t.I3
	return r
}

func (r PointBuilder) Apply(x int, y int, z fp.Tuple2[int, int]) PointBuilder {
	r.x = x
	r.y = y
	r.z = z
	return r
}

func (r Point) AsMap() map[string]any {
	m := map[string]any{}
	m["x"] = r.x
	m["y"] = r.y
	m["z"] = r.z
	return m
}

func (r PointBuilder) FromMap(m map[string]any) PointBuilder {

	if v, ok := m["x"].(int); ok {
		r.x = v
	}

	if v, ok := m["y"].(int); ok {
		r.y = v
	}

	if v, ok := m["z"].(fp.Tuple2[int, int]); ok {
		r.z = v
	}

	return r
}

type GreetingBuilder Greeting

type GreetingMutable struct {
	Hello    testpk1.World `json:"hello"`
	Language string        `json:"language,omitempty"`
}

func (r GreetingBuilder) Build() Greeting {
	return Greeting(r)
}

func (r Greeting) Builder() GreetingBuilder {
	return GreetingBuilder(r)
}

func (r Greeting) Hello() testpk1.World {
	return r.hello
}

func (r Greeting) Language() string {
	return r.language
}

func (r Greeting) WithHello(v testpk1.World) Greeting {
	r.hello = v
	return r
}

func (r Greeting) WithLanguage(v string) Greeting {
	r.language = v
	return r
}

func (r GreetingBuilder) Hello(v testpk1.World) GreetingBuilder {
	r.hello = v
	return r
}

func (r GreetingBuilder) Language(v string) GreetingBuilder {
	r.language = v
	return r
}

func (r Greeting) String() string {
	return fmt.Sprintf("Greeting(hello=%v, language=%v)", r.hello, r.language)
}

func (r Greeting) AsTuple() fp.Tuple2[testpk1.World, string] {
	return as.Tuple2(r.hello, r.language)
}

func (r Greeting) Unapply() (testpk1.World, string) {
	return r.hello, r.language
}

func (r Greeting) AsMutable() GreetingMutable {
	return GreetingMutable{
		Hello:    r.hello,
		Language: r.language,
	}
}

func (r GreetingMutable) AsImmutable() Greeting {
	return Greeting{
		hello:    r.Hello,
		language: r.Language,
	}
}

func (r GreetingBuilder) FromTuple(t fp.Tuple2[testpk1.World, string]) GreetingBuilder {
	r.hello = t.I1
	r.language = t.I2
	return r
}

func (r GreetingBuilder) Apply(hello testpk1.World, language string) GreetingBuilder {
	r.hello = hello
	r.language = language
	return r
}

func (r Greeting) AsMap() map[string]any {
	m := map[string]any{}
	m["hello"] = r.hello
	m["language"] = r.language
	return m
}

func (r GreetingBuilder) FromMap(m map[string]any) GreetingBuilder {

	if v, ok := m["hello"].(testpk1.World); ok {
		r.hello = v
	}

	if v, ok := m["language"].(string); ok {
		r.language = v
	}

	return r
}

func (r Greeting) AsLabelled() fp.Labelled2[NamedHello[testpk1.World], NamedLanguage[string]] {
	return as.Labelled2(NamedHello[testpk1.World]{r.hello}, NamedLanguage[string]{r.language})
}

func (r GreetingBuilder) FromLabelled(t fp.Labelled2[NamedHello[testpk1.World], NamedLanguage[string]]) GreetingBuilder {
	r.hello = t.I1.Value()
	r.language = t.I2.Value()
	return r
}

func (r Greeting) MarshalJSON() ([]byte, error) {
	m := r.AsMutable()
	return json.Marshal(m)
}

func (r *Greeting) UnmarshalJSON(b []byte) error {
	if r == nil {
		return fp.Error(http.StatusBadRequest, "target ptr is nil")
	}
	m := r.AsMutable()
	err := json.Unmarshal(b, &m)
	if err == nil {
		*r = m.AsImmutable()
	}
	return err
}

type ThreeBuilder Three

type ThreeMutable struct {
	One   int
	Two   string
	Three float64
}

func (r ThreeBuilder) Build() Three {
	return Three(r)
}

func (r Three) Builder() ThreeBuilder {
	return ThreeBuilder(r)
}

func (r Three) One() int {
	return r.one
}

func (r Three) Two() string {
	return r.two
}

func (r Three) Three() float64 {
	return r.three
}

func (r Three) WithOne(v int) Three {
	r.one = v
	return r
}

func (r Three) WithTwo(v string) Three {
	r.two = v
	return r
}

func (r Three) WithThree(v float64) Three {
	r.three = v
	return r
}

func (r ThreeBuilder) One(v int) ThreeBuilder {
	r.one = v
	return r
}

func (r ThreeBuilder) Two(v string) ThreeBuilder {
	r.two = v
	return r
}

func (r ThreeBuilder) Three(v float64) ThreeBuilder {
	r.three = v
	return r
}

func (r Three) String() string {
	return fmt.Sprintf("Three(one=%v, two=%v, three=%v)", r.one, r.two, r.three)
}

func (r Three) AsTuple() fp.Tuple3[int, string, float64] {
	return as.Tuple3(r.one, r.two, r.three)
}

func (r Three) Unapply() (int, string, float64) {
	return r.one, r.two, r.three
}

func (r Three) AsMutable() ThreeMutable {
	return ThreeMutable{
		One:   r.one,
		Two:   r.two,
		Three: r.three,
	}
}

func (r ThreeMutable) AsImmutable() Three {
	return Three{
		one:   r.One,
		two:   r.Two,
		three: r.Three,
	}
}

func (r ThreeBuilder) FromTuple(t fp.Tuple3[int, string, float64]) ThreeBuilder {
	r.one = t.I1
	r.two = t.I2
	r.three = t.I3
	return r
}

func (r ThreeBuilder) Apply(one int, two string, three float64) ThreeBuilder {
	r.one = one
	r.two = two
	r.three = three
	return r
}

func (r Three) AsMap() map[string]any {
	m := map[string]any{}
	m["one"] = r.one
	m["two"] = r.two
	m["three"] = r.three
	return m
}

func (r ThreeBuilder) FromMap(m map[string]any) ThreeBuilder {

	if v, ok := m["one"].(int); ok {
		r.one = v
	}

	if v, ok := m["two"].(string); ok {
		r.two = v
	}

	if v, ok := m["three"].(float64); ok {
		r.three = v
	}

	return r
}

func (r Three) AsLabelled() fp.Labelled3[NamedOne[int], NamedTwo[string], NamedThree[float64]] {
	return as.Labelled3(NamedOne[int]{r.one}, NamedTwo[string]{r.two}, NamedThree[float64]{r.three})
}

func (r ThreeBuilder) FromLabelled(t fp.Labelled3[NamedOne[int], NamedTwo[string], NamedThree[float64]]) ThreeBuilder {
	r.one = t.I1.Value()
	r.two = t.I2.Value()
	r.three = t.I3.Value()
	return r
}

type TreeBuilder Tree

type TreeMutable struct {
	Root testpk1.Node
}

func (r TreeBuilder) Build() Tree {
	return Tree(r)
}

func (r Tree) Builder() TreeBuilder {
	return TreeBuilder(r)
}

func (r Tree) Root() testpk1.Node {
	return r.root
}

func (r Tree) WithRoot(v testpk1.Node) Tree {
	r.root = v
	return r
}

func (r TreeBuilder) Root(v testpk1.Node) TreeBuilder {
	r.root = v
	return r
}

func (r Tree) String() string {
	return fmt.Sprintf("Tree(root=%v)", r.root)
}

func (r Tree) AsTuple() fp.Tuple1[testpk1.Node] {
	return as.Tuple1(r.root)
}

func (r Tree) Unapply() testpk1.Node {
	return r.root
}

func (r Tree) AsMutable() TreeMutable {
	return TreeMutable{
		Root: r.root,
	}
}

func (r TreeMutable) AsImmutable() Tree {
	return Tree{
		root: r.Root,
	}
}

func (r TreeBuilder) FromTuple(t fp.Tuple1[testpk1.Node]) TreeBuilder {
	r.root = t.I1
	return r
}

func (r TreeBuilder) Apply(root testpk1.Node) TreeBuilder {
	r.root = root
	return r
}

func (r Tree) AsMap() map[string]any {
	m := map[string]any{}
	m["root"] = r.root
	return m
}

func (r TreeBuilder) FromMap(m map[string]any) TreeBuilder {

	if v, ok := m["root"].(testpk1.Node); ok {
		r.root = v
	}

	return r
}

func (r AliasedStruct) GetPubField() string {
	return r.PubField
}

func (r AliasedStruct) WithPubField(v string) AliasedStruct {
	r.PubField = v
	return r
}

func (r AliasedStruct) WithDupGetter(v string) AliasedStruct {
	r.DupGetter = v
	return r
}

func (r AliasedStruct) Deref() testpk1.DefinedOtherPackage {
	return testpk1.DefinedOtherPackage(r)
}

func IntoAliasedStruct(v testpk1.DefinedOtherPackage) AliasedStruct {
	return AliasedStruct(v)
}

func (r AliasedStruct) AsMap() map[string]any {
	return testpk1.DefinedOtherPackage(r).AsMap()
}

func (r AliasedStruct) AsMutable() testpk1.DefinedOtherPackageMutable {
	return testpk1.DefinedOtherPackage(r).AsMutable()
}

func (r AliasedStruct) AsTuple() fp.Tuple3[string, string, string] {
	return testpk1.DefinedOtherPackage(r).AsTuple()
}

func (r AliasedStruct) Builder() testpk1.DefinedOtherPackageBuilder {
	return testpk1.DefinedOtherPackage(r).Builder()
}

func (r *AliasedStruct) GetDupGetter() string {
	return (*testpk1.DefinedOtherPackage)(r).GetDupGetter()
}

func (r AliasedStruct) PrivField() string {
	return testpk1.DefinedOtherPackage(r).PrivField()
}

func (r *AliasedStruct) PtrRecv() {
	(*testpk1.DefinedOtherPackage)(r).PtrRecv()
}

func (r *AliasedStruct) PtrRecvRet() string {
	return (*testpk1.DefinedOtherPackage)(r).PtrRecvRet()
}

func (r AliasedStruct) Unapply() (string, string, string) {
	return testpk1.DefinedOtherPackage(r).Unapply()
}

func (r AliasedStruct) WithPrivField(v string) testpk1.DefinedOtherPackage {
	return testpk1.DefinedOtherPackage(r).WithPrivField(v)
}

func (r GetterOverride) GetPubField() string {
	return r.PubField
}

func (r GetterOverride) GetDupGetter() string {
	return r.DupGetter
}

func (r GetterOverride) Deref() testpk1.DefinedOtherPackage {
	return testpk1.DefinedOtherPackage(r)
}

func IntoGetterOverride(v testpk1.DefinedOtherPackage) GetterOverride {
	return GetterOverride(v)
}

func (r GetterOverride) AsMap() map[string]any {
	return testpk1.DefinedOtherPackage(r).AsMap()
}

func (r GetterOverride) AsMutable() testpk1.DefinedOtherPackageMutable {
	return testpk1.DefinedOtherPackage(r).AsMutable()
}

func (r GetterOverride) AsTuple() fp.Tuple3[string, string, string] {
	return testpk1.DefinedOtherPackage(r).AsTuple()
}

func (r GetterOverride) Builder() testpk1.DefinedOtherPackageBuilder {
	return testpk1.DefinedOtherPackage(r).Builder()
}

func (r GetterOverride) PrivField() string {
	return testpk1.DefinedOtherPackage(r).PrivField()
}

func (r *GetterOverride) PtrRecv() {
	(*testpk1.DefinedOtherPackage)(r).PtrRecv()
}

func (r *GetterOverride) PtrRecvRet() string {
	return (*testpk1.DefinedOtherPackage)(r).PtrRecvRet()
}

func (r GetterOverride) String() string {
	return testpk1.DefinedOtherPackage(r).String()
}

func (r GetterOverride) Unapply() (string, string, string) {
	return testpk1.DefinedOtherPackage(r).Unapply()
}

func (r GetterOverride) WithPrivField(v string) testpk1.DefinedOtherPackage {
	return testpk1.DefinedOtherPackage(r).WithPrivField(v)
}

type NotIgnoredBuilder NotIgnored

type NotIgnoredMutable struct {
	Ig int
}

func (r NotIgnoredBuilder) Build() NotIgnored {
	return NotIgnored(r)
}

func (r NotIgnored) Builder() NotIgnoredBuilder {
	return NotIgnoredBuilder(r)
}

func (r NotIgnored) Ig() int {
	return r.ig
}

func (r NotIgnored) WithIg(v int) NotIgnored {
	r.ig = v
	return r
}

func (r NotIgnoredBuilder) Ig(v int) NotIgnoredBuilder {
	r.ig = v
	return r
}

func (r NotIgnored) String() string {
	return fmt.Sprintf("NotIgnored(ig=%v)", r.ig)
}

func (r NotIgnored) AsTuple() fp.Tuple1[int] {
	return as.Tuple1(r.ig)
}

func (r NotIgnored) Unapply() int {
	return r.ig
}

func (r NotIgnored) AsMutable() NotIgnoredMutable {
	return NotIgnoredMutable{
		Ig: r.ig,
	}
}

func (r NotIgnoredMutable) AsImmutable() NotIgnored {
	return NotIgnored{
		ig: r.Ig,
	}
}

func (r NotIgnoredBuilder) FromTuple(t fp.Tuple1[int]) NotIgnoredBuilder {
	r.ig = t.I1
	return r
}

func (r NotIgnoredBuilder) Apply(ig int) NotIgnoredBuilder {
	r.ig = ig
	return r
}

func (r NotIgnored) AsMap() map[string]any {
	m := map[string]any{}
	m["ig"] = r.ig
	return m
}

func (r NotIgnoredBuilder) FromMap(m map[string]any) NotIgnoredBuilder {

	if v, ok := m["ig"].(int); ok {
		r.ig = v
	}

	return r
}

type NamedHello[T any] fp.Tuple1[T]

func (r NamedHello[T]) Name() string {
	return "hello"
}
func (r NamedHello[T]) Value() T {
	return r.I1
}
func (r NamedHello[T]) WithValue(v T) NamedHello[T] {
	r.I1 = v
	return r
}

type NamedLanguage[T any] fp.Tuple1[T]

func (r NamedLanguage[T]) Name() string {
	return "language"
}
func (r NamedLanguage[T]) Value() T {
	return r.I1
}
func (r NamedLanguage[T]) WithValue(v T) NamedLanguage[T] {
	r.I1 = v
	return r
}

type NamedOne[T any] fp.Tuple1[T]

func (r NamedOne[T]) Name() string {
	return "one"
}
func (r NamedOne[T]) Value() T {
	return r.I1
}
func (r NamedOne[T]) WithValue(v T) NamedOne[T] {
	r.I1 = v
	return r
}

type NamedThree[T any] fp.Tuple1[T]

func (r NamedThree[T]) Name() string {
	return "three"
}
func (r NamedThree[T]) Value() T {
	return r.I1
}
func (r NamedThree[T]) WithValue(v T) NamedThree[T] {
	r.I1 = v
	return r
}

type NamedTwo[T any] fp.Tuple1[T]

func (r NamedTwo[T]) Name() string {
	return "two"
}
func (r NamedTwo[T]) Value() T {
	return r.I1
}
func (r NamedTwo[T]) WithValue(v T) NamedTwo[T] {
	r.I1 = v
	return r
}
