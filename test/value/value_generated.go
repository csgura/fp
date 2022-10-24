package value

import (
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"os"
	"reflect"
	"sync/atomic"
)

type NotIgnoredBuilder NotIgnored

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

func (r NotIgnoredBuilder) FromTuple(t fp.Tuple1[int]) NotIgnoredBuilder {
	r.ig = t.I1
	return r
}

func (r NotIgnored) AsMap() map[string]any {
	return map[string]any{
		"ig": r.ig,
	}
}

func (r NotIgnoredBuilder) FromMap(m map[string]any) NotIgnoredBuilder {

	if v, ok := m["ig"].(int); ok {
		r.ig = v
	}

	return r
}

func (r NotIgnored) AsLabelled() fp.Tuple1[fp.Tuple2[string, int]] {
	return as.Tuple1(as.Tuple2("ig", r.ig))
}

func (r NotIgnoredBuilder) FromLabelled(t fp.Tuple1[fp.Tuple2[string, int]]) NotIgnoredBuilder {
	r.ig = t.I1.I2
	return r
}

type HelloBuilder Hello

func (r HelloBuilder) Build() Hello {
	return Hello(r)
}

func (r Hello) Builder() HelloBuilder {
	return HelloBuilder(r)
}

func (r Hello) World() string {
	return r.world
}

func (r Hello) WithWorld(v string) Hello {
	r.world = v
	return r
}

func (r HelloBuilder) World(v string) HelloBuilder {
	r.world = v
	return r
}

func (r Hello) Hi() int {
	return r.hi
}

func (r Hello) WithHi(v int) Hello {
	r.hi = v
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

func (r HelloBuilder) FromTuple(t fp.Tuple2[string, int]) HelloBuilder {
	r.world = t.I1
	r.hi = t.I2
	return r
}

func (r Hello) AsMap() map[string]any {
	return map[string]any{
		"world": r.world,
		"hi":    r.hi,
	}
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

func (r Hello) AsLabelled() fp.Tuple2[fp.Tuple2[string, string], fp.Tuple2[string, int]] {
	return as.Tuple2(as.Tuple2("world", r.world), as.Tuple2("hi", r.hi))
}

func (r HelloBuilder) FromLabelled(t fp.Tuple2[fp.Tuple2[string, string], fp.Tuple2[string, int]]) HelloBuilder {
	r.world = t.I1.I2
	r.hi = t.I2.I2
	return r
}

type MyMyBuilder MyMy

func (r MyMyBuilder) Build() MyMy {
	return MyMy(r)
}

func (r MyMy) Builder() MyMyBuilder {
	return MyMyBuilder(r)
}

func (r MyMy) Hi() fp.Option[int] {
	return r.hi
}

func (r MyMy) WithHi(v fp.Option[int]) MyMy {
	r.hi = v
	return r
}

func (r MyMyBuilder) Hi(v fp.Option[int]) MyMyBuilder {
	r.hi = v
	return r
}

func (r MyMy) Tpe() reflect.Type {
	return r.tpe
}

func (r MyMy) WithTpe(v reflect.Type) MyMy {
	r.tpe = v
	return r
}

func (r MyMyBuilder) Tpe(v reflect.Type) MyMyBuilder {
	r.tpe = v
	return r
}

func (r MyMy) Arr() []os.File {
	return r.arr
}

func (r MyMy) WithArr(v []os.File) MyMy {
	r.arr = v
	return r
}

func (r MyMyBuilder) Arr(v []os.File) MyMyBuilder {
	r.arr = v
	return r
}

func (r MyMy) M() map[string]int {
	return r.m
}

func (r MyMy) WithM(v map[string]int) MyMy {
	r.m = v
	return r
}

func (r MyMyBuilder) M(v map[string]int) MyMyBuilder {
	r.m = v
	return r
}

func (r MyMy) A() any {
	return r.a
}

func (r MyMy) WithA(v any) MyMy {
	r.a = v
	return r
}

func (r MyMyBuilder) A(v any) MyMyBuilder {
	r.a = v
	return r
}

func (r MyMy) P() *int {
	return r.p
}

func (r MyMy) WithP(v *int) MyMy {
	r.p = v
	return r
}

func (r MyMyBuilder) P(v *int) MyMyBuilder {
	r.p = v
	return r
}

func (r MyMy) L() Local {
	return r.l
}

func (r MyMy) WithL(v Local) MyMy {
	r.l = v
	return r
}

func (r MyMyBuilder) L(v Local) MyMyBuilder {
	r.l = v
	return r
}

func (r MyMy) T() fp.Try[fp.Option[Local]] {
	return r.t
}

func (r MyMy) WithT(v fp.Try[fp.Option[Local]]) MyMy {
	r.t = v
	return r
}

func (r MyMyBuilder) T(v fp.Try[fp.Option[Local]]) MyMyBuilder {
	r.t = v
	return r
}

func (r MyMy) M2() map[string]atomic.Bool {
	return r.m2
}

func (r MyMy) WithM2(v map[string]atomic.Bool) MyMy {
	r.m2 = v
	return r
}

func (r MyMyBuilder) M2(v map[string]atomic.Bool) MyMyBuilder {
	r.m2 = v
	return r
}

func (r MyMy) Mm() fp.Map[string, int] {
	return r.mm
}

func (r MyMy) WithMm(v fp.Map[string, int]) MyMy {
	r.mm = v
	return r
}

func (r MyMyBuilder) Mm(v fp.Map[string, int]) MyMyBuilder {
	r.mm = v
	return r
}

func (r MyMy) String() string {
	return fmt.Sprintf("MyMy(hi=%v, tpe=%v, arr=%v, m=%v, a=%v, p=%v, l=%v, t=%v, m2=%v, mm=%v)", r.hi, r.tpe, r.arr, r.m, r.a, r.p, r.l, r.t, r.m2, r.mm)
}

func (r MyMy) AsTuple() fp.Tuple10[fp.Option[int], reflect.Type, []os.File, map[string]int, any, *int, Local, fp.Try[fp.Option[Local]], map[string]atomic.Bool, fp.Map[string, int]] {
	return as.Tuple10(r.hi, r.tpe, r.arr, r.m, r.a, r.p, r.l, r.t, r.m2, r.mm)
}

func (r MyMyBuilder) FromTuple(t fp.Tuple10[fp.Option[int], reflect.Type, []os.File, map[string]int, any, *int, Local, fp.Try[fp.Option[Local]], map[string]atomic.Bool, fp.Map[string, int]]) MyMyBuilder {
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
	return r
}

func (r MyMy) AsMap() map[string]any {
	return map[string]any{
		"hi":  r.hi,
		"tpe": r.tpe,
		"arr": r.arr,
		"m":   r.m,
		"a":   r.a,
		"p":   r.p,
		"l":   r.l,
		"t":   r.t,
		"m2":  r.m2,
		"mm":  r.mm,
	}
}

func (r MyMyBuilder) FromMap(m map[string]any) MyMyBuilder {

	if v, ok := m["hi"].(fp.Option[int]); ok {
		r.hi = v
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

	return r
}

func (r MyMy) AsLabelled() fp.Tuple10[fp.Tuple2[string, fp.Option[int]], fp.Tuple2[string, reflect.Type], fp.Tuple2[string, []os.File], fp.Tuple2[string, map[string]int], fp.Tuple2[string, any], fp.Tuple2[string, *int], fp.Tuple2[string, Local], fp.Tuple2[string, fp.Try[fp.Option[Local]]], fp.Tuple2[string, map[string]atomic.Bool], fp.Tuple2[string, fp.Map[string, int]]] {
	return as.Tuple10(as.Tuple2("hi", r.hi), as.Tuple2("tpe", r.tpe), as.Tuple2("arr", r.arr), as.Tuple2("m", r.m), as.Tuple2("a", r.a), as.Tuple2("p", r.p), as.Tuple2("l", r.l), as.Tuple2("t", r.t), as.Tuple2("m2", r.m2), as.Tuple2("mm", r.mm))
}

func (r MyMyBuilder) FromLabelled(t fp.Tuple10[fp.Tuple2[string, fp.Option[int]], fp.Tuple2[string, reflect.Type], fp.Tuple2[string, []os.File], fp.Tuple2[string, map[string]int], fp.Tuple2[string, any], fp.Tuple2[string, *int], fp.Tuple2[string, Local], fp.Tuple2[string, fp.Try[fp.Option[Local]]], fp.Tuple2[string, map[string]atomic.Bool], fp.Tuple2[string, fp.Map[string, int]]]) MyMyBuilder {
	r.hi = t.I1.I2
	r.tpe = t.I2.I2
	r.arr = t.I3.I2
	r.m = t.I4.I2
	r.a = t.I5.I2
	r.p = t.I6.I2
	r.l = t.I7.I2
	r.t = t.I8.I2
	r.m2 = t.I9.I2
	r.mm = t.I10.I2
	return r
}

type PersonBuilder Person

func (r PersonBuilder) Build() Person {
	return Person(r)
}

func (r Person) Builder() PersonBuilder {
	return PersonBuilder(r)
}

func (r Person) Name() string {
	return r.name
}

func (r Person) WithName(v string) Person {
	r.name = v
	return r
}

func (r PersonBuilder) Name(v string) PersonBuilder {
	r.name = v
	return r
}

func (r Person) Age() int {
	return r.age
}

func (r Person) WithAge(v int) Person {
	r.age = v
	return r
}

func (r PersonBuilder) Age(v int) PersonBuilder {
	r.age = v
	return r
}

func (r Person) Height() float64 {
	return r.height
}

func (r Person) WithHeight(v float64) Person {
	r.height = v
	return r
}

func (r PersonBuilder) Height(v float64) PersonBuilder {
	r.height = v
	return r
}

func (r Person) Phone() fp.Option[string] {
	return r.phone
}

func (r Person) WithPhone(v fp.Option[string]) Person {
	r.phone = v
	return r
}

func (r PersonBuilder) Phone(v fp.Option[string]) PersonBuilder {
	r.phone = v
	return r
}

func (r Person) Addr() []string {
	return r.addr
}

func (r Person) WithAddr(v []string) Person {
	r.addr = v
	return r
}

func (r PersonBuilder) Addr(v []string) PersonBuilder {
	r.addr = v
	return r
}

func (r Person) List() hlist.Cons[string, hlist.Cons[int, hlist.Nil]] {
	return r.list
}

func (r Person) WithList(v hlist.Cons[string, hlist.Cons[int, hlist.Nil]]) Person {
	r.list = v
	return r
}

func (r PersonBuilder) List(v hlist.Cons[string, hlist.Cons[int, hlist.Nil]]) PersonBuilder {
	r.list = v
	return r
}

func (r Person) Seq() fp.Seq[float64] {
	return r.seq
}

func (r Person) WithSeq(v fp.Seq[float64]) Person {
	r.seq = v
	return r
}

func (r PersonBuilder) Seq(v fp.Seq[float64]) PersonBuilder {
	r.seq = v
	return r
}

func (r Person) String() string {
	return fmt.Sprintf("Person(name=%v, age=%v, height=%v, phone=%v, addr=%v, list=%v, seq=%v)", r.name, r.age, r.height, r.phone, r.addr, r.list, r.seq)
}

func (r Person) AsTuple() fp.Tuple7[string, int, float64, fp.Option[string], []string, hlist.Cons[string, hlist.Cons[int, hlist.Nil]], fp.Seq[float64]] {
	return as.Tuple7(r.name, r.age, r.height, r.phone, r.addr, r.list, r.seq)
}

func (r PersonBuilder) FromTuple(t fp.Tuple7[string, int, float64, fp.Option[string], []string, hlist.Cons[string, hlist.Cons[int, hlist.Nil]], fp.Seq[float64]]) PersonBuilder {
	r.name = t.I1
	r.age = t.I2
	r.height = t.I3
	r.phone = t.I4
	r.addr = t.I5
	r.list = t.I6
	r.seq = t.I7
	return r
}

func (r Person) AsMap() map[string]any {
	return map[string]any{
		"name":   r.name,
		"age":    r.age,
		"height": r.height,
		"phone":  r.phone,
		"addr":   r.addr,
		"list":   r.list,
		"seq":    r.seq,
	}
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

	return r
}

func (r Person) AsLabelled() fp.Tuple7[fp.Tuple2[string, string], fp.Tuple2[string, int], fp.Tuple2[string, float64], fp.Tuple2[string, fp.Option[string]], fp.Tuple2[string, []string], fp.Tuple2[string, hlist.Cons[string, hlist.Cons[int, hlist.Nil]]], fp.Tuple2[string, fp.Seq[float64]]] {
	return as.Tuple7(as.Tuple2("name", r.name), as.Tuple2("age", r.age), as.Tuple2("height", r.height), as.Tuple2("phone", r.phone), as.Tuple2("addr", r.addr), as.Tuple2("list", r.list), as.Tuple2("seq", r.seq))
}

func (r PersonBuilder) FromLabelled(t fp.Tuple7[fp.Tuple2[string, string], fp.Tuple2[string, int], fp.Tuple2[string, float64], fp.Tuple2[string, fp.Option[string]], fp.Tuple2[string, []string], fp.Tuple2[string, hlist.Cons[string, hlist.Cons[int, hlist.Nil]]], fp.Tuple2[string, fp.Seq[float64]]]) PersonBuilder {
	r.name = t.I1.I2
	r.age = t.I2.I2
	r.height = t.I3.I2
	r.phone = t.I4.I2
	r.addr = t.I5.I2
	r.list = t.I6.I2
	r.seq = t.I7.I2
	return r
}
