// Code generated by gombok, DO NOT EDIT.
package value

import (
	"encoding/json"
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hlist"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/test/internal/hello"
	"net/http"
	"os"
	"reflect"
	"sync/atomic"
)

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

type HelloBuilder Hello

type HelloMutable struct {
	World string `json:"world"`
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

type AllKindTypesBuilder AllKindTypes

type AllKindTypesMutable struct {
	Embed Embed
	Hi    fp.Option[int]
	Tpe   reflect.Type
	Arr   []os.File
	M     map[string]int
	A     any
	P     *int
	L     Local
	T     fp.Try[fp.Option[Local]]
	M2    map[string]atomic.Bool
	Mm    fp.Map[string, int]
	Intf  fp.Future[int]
	Ch    chan fp.Try[fp.Either[int, string]]
	Fn3   fp.Func1[int, fp.Try[string]]
	Fn    func(a string) fp.Try[int]
	Fn2   func(fp.Try[string]) (result int, err error)
	Arr2  [2]int
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

func (r AllKindTypes) WithHi(v fp.Option[int]) AllKindTypes {
	r.hi = v
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

func (r AllKindTypes) Tpe() reflect.Type {
	return r.tpe
}

func (r AllKindTypes) WithTpe(v reflect.Type) AllKindTypes {
	r.tpe = v
	return r
}

func (r AllKindTypesBuilder) Tpe(v reflect.Type) AllKindTypesBuilder {
	r.tpe = v
	return r
}

func (r AllKindTypes) Arr() []os.File {
	return r.arr
}

func (r AllKindTypes) WithArr(v []os.File) AllKindTypes {
	r.arr = v
	return r
}

func (r AllKindTypesBuilder) Arr(v []os.File) AllKindTypesBuilder {
	r.arr = v
	return r
}

func (r AllKindTypes) M() map[string]int {
	return r.m
}

func (r AllKindTypes) WithM(v map[string]int) AllKindTypes {
	r.m = v
	return r
}

func (r AllKindTypesBuilder) M(v map[string]int) AllKindTypesBuilder {
	r.m = v
	return r
}

func (r AllKindTypes) A() any {
	return r.a
}

func (r AllKindTypes) WithA(v any) AllKindTypes {
	r.a = v
	return r
}

func (r AllKindTypesBuilder) A(v any) AllKindTypesBuilder {
	r.a = v
	return r
}

func (r AllKindTypes) P() *int {
	return r.p
}

func (r AllKindTypes) WithP(v *int) AllKindTypes {
	r.p = v
	return r
}

func (r AllKindTypesBuilder) P(v *int) AllKindTypesBuilder {
	r.p = v
	return r
}

func (r AllKindTypes) L() Local {
	return r.l
}

func (r AllKindTypes) WithL(v Local) AllKindTypes {
	r.l = v
	return r
}

func (r AllKindTypesBuilder) L(v Local) AllKindTypesBuilder {
	r.l = v
	return r
}

func (r AllKindTypes) T() fp.Try[fp.Option[Local]] {
	return r.t
}

func (r AllKindTypes) WithT(v fp.Try[fp.Option[Local]]) AllKindTypes {
	r.t = v
	return r
}

func (r AllKindTypesBuilder) T(v fp.Try[fp.Option[Local]]) AllKindTypesBuilder {
	r.t = v
	return r
}

func (r AllKindTypes) M2() map[string]atomic.Bool {
	return r.m2
}

func (r AllKindTypes) WithM2(v map[string]atomic.Bool) AllKindTypes {
	r.m2 = v
	return r
}

func (r AllKindTypesBuilder) M2(v map[string]atomic.Bool) AllKindTypesBuilder {
	r.m2 = v
	return r
}

func (r AllKindTypes) Mm() fp.Map[string, int] {
	return r.mm
}

func (r AllKindTypes) WithMm(v fp.Map[string, int]) AllKindTypes {
	r.mm = v
	return r
}

func (r AllKindTypesBuilder) Mm(v fp.Map[string, int]) AllKindTypesBuilder {
	r.mm = v
	return r
}

func (r AllKindTypes) Intf() fp.Future[int] {
	return r.intf
}

func (r AllKindTypes) WithIntf(v fp.Future[int]) AllKindTypes {
	r.intf = v
	return r
}

func (r AllKindTypesBuilder) Intf(v fp.Future[int]) AllKindTypesBuilder {
	r.intf = v
	return r
}

func (r AllKindTypes) Ch() chan fp.Try[fp.Either[int, string]] {
	return r.ch
}

func (r AllKindTypes) WithCh(v chan fp.Try[fp.Either[int, string]]) AllKindTypes {
	r.ch = v
	return r
}

func (r AllKindTypesBuilder) Ch(v chan fp.Try[fp.Either[int, string]]) AllKindTypesBuilder {
	r.ch = v
	return r
}

func (r AllKindTypes) Fn3() fp.Func1[int, fp.Try[string]] {
	return r.fn3
}

func (r AllKindTypes) WithFn3(v fp.Func1[int, fp.Try[string]]) AllKindTypes {
	r.fn3 = v
	return r
}

func (r AllKindTypesBuilder) Fn3(v fp.Func1[int, fp.Try[string]]) AllKindTypesBuilder {
	r.fn3 = v
	return r
}

func (r AllKindTypes) Fn() func(a string) fp.Try[int] {
	return r.fn
}

func (r AllKindTypes) WithFn(v func(a string) fp.Try[int]) AllKindTypes {
	r.fn = v
	return r
}

func (r AllKindTypesBuilder) Fn(v func(a string) fp.Try[int]) AllKindTypesBuilder {
	r.fn = v
	return r
}

func (r AllKindTypes) Fn2() func(fp.Try[string]) (result int, err error) {
	return r.fn2
}

func (r AllKindTypes) WithFn2(v func(fp.Try[string]) (result int, err error)) AllKindTypes {
	r.fn2 = v
	return r
}

func (r AllKindTypesBuilder) Fn2(v func(fp.Try[string]) (result int, err error)) AllKindTypesBuilder {
	r.fn2 = v
	return r
}

func (r AllKindTypes) Arr2() [2]int {
	return r.arr2
}

func (r AllKindTypes) WithArr2(v [2]int) AllKindTypes {
	r.arr2 = v
	return r
}

func (r AllKindTypesBuilder) Arr2(v [2]int) AllKindTypesBuilder {
	r.arr2 = v
	return r
}

func (r AllKindTypes) String() string {
	return fmt.Sprintf("AllKindTypes(hi=%v, tpe=%v, arr=%v, m=%v, a=%v, p=%v, l=%v, t=%v, m2=%v, mm=%v, intf=%v, ch=%v, fn3=%v, arr2=%v)", r.hi, r.tpe, r.arr, r.m, r.a, r.p, r.l, r.t, r.m2, r.mm, r.intf, r.ch, r.fn3, r.arr2)
}

func (r AllKindTypes) AsTuple() fp.Tuple16[fp.Option[int], reflect.Type, []os.File, map[string]int, any, *int, Local, fp.Try[fp.Option[Local]], map[string]atomic.Bool, fp.Map[string, int], fp.Future[int], chan fp.Try[fp.Either[int, string]], fp.Func1[int, fp.Try[string]], func(a string) fp.Try[int], func(fp.Try[string]) (result int, err error), [2]int] {
	return as.Tuple16(r.hi, r.tpe, r.arr, r.m, r.a, r.p, r.l, r.t, r.m2, r.mm, r.intf, r.ch, r.fn3, r.fn, r.fn2, r.arr2)
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
		Fn3:  r.fn3,
		Fn:   r.fn,
		Fn2:  r.fn2,
		Arr2: r.arr2,
	}
}

func (r AllKindTypesMutable) AsImmutable() AllKindTypes {
	return AllKindTypes{
		Embed: r.Embed,
		hi:    r.Hi,
		tpe:   r.Tpe,
		arr:   r.Arr,
		m:     r.M,
		a:     r.A,
		p:     r.P,
		l:     r.L,
		t:     r.T,
		m2:    r.M2,
		mm:    r.Mm,
		intf:  r.Intf,
		ch:    r.Ch,
		fn3:   r.Fn3,
		fn:    r.Fn,
		fn2:   r.Fn2,
		arr2:  r.Arr2,
	}
}

func (r AllKindTypesBuilder) FromTuple(t fp.Tuple16[fp.Option[int], reflect.Type, []os.File, map[string]int, any, *int, Local, fp.Try[fp.Option[Local]], map[string]atomic.Bool, fp.Map[string, int], fp.Future[int], chan fp.Try[fp.Either[int, string]], fp.Func1[int, fp.Try[string]], func(a string) fp.Try[int], func(fp.Try[string]) (result int, err error), [2]int]) AllKindTypesBuilder {
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
	r.fn3 = t.I13
	r.fn = t.I14
	r.fn2 = t.I15
	r.arr2 = t.I16
	return r
}

func (r AllKindTypes) AsMap() map[string]any {
	return map[string]any{
		"hi":   r.hi,
		"tpe":  r.tpe,
		"arr":  r.arr,
		"m":    r.m,
		"a":    r.a,
		"p":    r.p,
		"l":    r.l,
		"t":    r.t,
		"m2":   r.m2,
		"mm":   r.mm,
		"intf": r.intf,
		"ch":   r.ch,
		"fn3":  r.fn3,
		"fn":   r.fn,
		"fn2":  r.fn2,
		"arr2": r.arr2,
	}
}

func (r AllKindTypesBuilder) FromMap(m map[string]any) AllKindTypesBuilder {

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

	if v, ok := m["intf"].(fp.Future[int]); ok {
		r.intf = v
	}

	if v, ok := m["ch"].(chan fp.Try[fp.Either[int, string]]); ok {
		r.ch = v
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

func (r PersonBuilder) SomePhone(v string) PersonBuilder {
	r.phone = option.Some(v)
	return r
}

func (r PersonBuilder) NonePhone() PersonBuilder {
	r.phone = option.None[string]()
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

func (r Person) Blob() []byte {
	return r.blob
}

func (r Person) WithBlob(v []byte) Person {
	r.blob = v
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
		name:       r.Name,
		age:        r.Age,
		height:     r.Height,
		phone:      r.Phone,
		addr:       r.Addr,
		list:       r.List,
		seq:        r.Seq,
		blob:       r.Blob,
		_notExport: r._notExport,
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

func (r Person) AsMap() map[string]any {
	return map[string]any{
		"name":   r.name,
		"age":    r.age,
		"height": r.height,
		"phone":  r.phone,
		"addr":   r.addr,
		"list":   r.list,
		"seq":    r.seq,
		"blob":   r.blob,
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

func (r Wallet) WithOwner(v Person) Wallet {
	r.owner = v
	return r
}

func (r WalletBuilder) Owner(v Person) WalletBuilder {
	r.owner = v
	return r
}

func (r Wallet) Amount() int64 {
	return r.amount
}

func (r Wallet) WithAmount(v int64) Wallet {
	r.amount = v
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

func (r Wallet) AsMap() map[string]any {
	return map[string]any{
		"owner":  r.owner,
		"amount": r.amount,
	}
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

type EntryBuilder[A interface{ String() string }, B any] Entry[A, B]

type EntryMutable[A interface{ String() string }, B any] struct {
	Name  string
	Value A
	Tuple fp.Tuple2[A, B]
}

func (r EntryBuilder[A, B]) Build() Entry[A, B] {
	return Entry[A, B](r)
}

func (r Entry[A, B]) Builder() EntryBuilder[A, B] {
	return EntryBuilder[A, B](r)
}

func (r Entry[A, B]) Name() string {
	return r.name
}

func (r Entry[A, B]) WithName(v string) Entry[A, B] {
	r.name = v
	return r
}

func (r EntryBuilder[A, B]) Name(v string) EntryBuilder[A, B] {
	r.name = v
	return r
}

func (r Entry[A, B]) Value() A {
	return r.value
}

func (r Entry[A, B]) WithValue(v A) Entry[A, B] {
	r.value = v
	return r
}

func (r EntryBuilder[A, B]) Value(v A) EntryBuilder[A, B] {
	r.value = v
	return r
}

func (r Entry[A, B]) Tuple() fp.Tuple2[A, B] {
	return r.tuple
}

func (r Entry[A, B]) WithTuple(v fp.Tuple2[A, B]) Entry[A, B] {
	r.tuple = v
	return r
}

func (r EntryBuilder[A, B]) Tuple(v fp.Tuple2[A, B]) EntryBuilder[A, B] {
	r.tuple = v
	return r
}

func (r Entry[A, B]) String() string {
	return fmt.Sprintf("Entry(name=%v, value=%v, tuple=%v)", r.name, r.value, r.tuple)
}

func (r Entry[A, B]) AsTuple() fp.Tuple3[string, A, fp.Tuple2[A, B]] {
	return as.Tuple3(r.name, r.value, r.tuple)
}

func (r Entry[A, B]) AsMutable() EntryMutable[A, B] {
	return EntryMutable[A, B]{
		Name:  r.name,
		Value: r.value,
		Tuple: r.tuple,
	}
}

func (r EntryMutable[A, B]) AsImmutable() Entry[A, B] {
	return Entry[A, B]{
		name:  r.Name,
		value: r.Value,
		tuple: r.Tuple,
	}
}

func (r EntryBuilder[A, B]) FromTuple(t fp.Tuple3[string, A, fp.Tuple2[A, B]]) EntryBuilder[A, B] {
	r.name = t.I1
	r.value = t.I2
	r.tuple = t.I3
	return r
}

func (r Entry[A, B]) AsMap() map[string]any {
	return map[string]any{
		"name":  r.name,
		"value": r.value,
		"tuple": r.tuple,
	}
}

func (r EntryBuilder[A, B]) FromMap(m map[string]any) EntryBuilder[A, B] {

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

func (r Key) WithA(v int) Key {
	r.a = v
	return r
}

func (r KeyBuilder) A(v int) KeyBuilder {
	r.a = v
	return r
}

func (r Key) B() float32 {
	return r.b
}

func (r Key) WithB(v float32) Key {
	r.b = v
	return r
}

func (r KeyBuilder) B(v float32) KeyBuilder {
	r.b = v
	return r
}

func (r Key) C() []byte {
	return r.c
}

func (r Key) WithC(v []byte) Key {
	r.c = v
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

func (r Key) AsMap() map[string]any {
	return map[string]any{
		"a": r.a,
		"b": r.b,
		"c": r.c,
	}
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

func (r Point) WithX(v int) Point {
	r.x = v
	return r
}

func (r PointBuilder) X(v int) PointBuilder {
	r.x = v
	return r
}

func (r Point) Y() int {
	return r.y
}

func (r Point) WithY(v int) Point {
	r.y = v
	return r
}

func (r PointBuilder) Y(v int) PointBuilder {
	r.y = v
	return r
}

func (r Point) Z() fp.Tuple2[int, int] {
	return r.z
}

func (r Point) WithZ(v fp.Tuple2[int, int]) Point {
	r.z = v
	return r
}

func (r PointBuilder) Z(v fp.Tuple2[int, int]) PointBuilder {
	r.z = v
	return r
}

func (r Point) AsTuple() fp.Tuple3[int, int, fp.Tuple2[int, int]] {
	return as.Tuple3(r.x, r.y, r.z)
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

func (r Point) AsMap() map[string]any {
	return map[string]any{
		"x": r.x,
		"y": r.y,
		"z": r.z,
	}
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
	Hello    hello.World `json:"hello"`
	Language string      `json:"language"`
}

func (r GreetingBuilder) Build() Greeting {
	return Greeting(r)
}

func (r Greeting) Builder() GreetingBuilder {
	return GreetingBuilder(r)
}

func (r Greeting) Hello() hello.World {
	return r.hello
}

func (r Greeting) WithHello(v hello.World) Greeting {
	r.hello = v
	return r
}

func (r GreetingBuilder) Hello(v hello.World) GreetingBuilder {
	r.hello = v
	return r
}

func (r Greeting) Language() string {
	return r.language
}

func (r Greeting) WithLanguage(v string) Greeting {
	r.language = v
	return r
}

func (r GreetingBuilder) Language(v string) GreetingBuilder {
	r.language = v
	return r
}

func (r Greeting) String() string {
	return fmt.Sprintf("Greeting(hello=%v, language=%v)", r.hello, r.language)
}

func (r Greeting) AsTuple() fp.Tuple2[hello.World, string] {
	return as.Tuple2(r.hello, r.language)
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

func (r GreetingBuilder) FromTuple(t fp.Tuple2[hello.World, string]) GreetingBuilder {
	r.hello = t.I1
	r.language = t.I2
	return r
}

func (r Greeting) AsMap() map[string]any {
	return map[string]any{
		"hello":    r.hello,
		"language": r.language,
	}
}

func (r GreetingBuilder) FromMap(m map[string]any) GreetingBuilder {

	if v, ok := m["hello"].(hello.World); ok {
		r.hello = v
	}

	if v, ok := m["language"].(string); ok {
		r.language = v
	}

	return r
}

func (r Greeting) AsLabelled() fp.Labelled2[NameIsHello[hello.World], NameIsLanguage[string]] {
	return as.Labelled2(NameIsHello[hello.World]{r.hello}, NameIsLanguage[string]{r.language})
}

func (r GreetingBuilder) FromLabelled(t fp.Labelled2[NameIsHello[hello.World], NameIsLanguage[string]]) GreetingBuilder {
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

func (r Three) WithOne(v int) Three {
	r.one = v
	return r
}

func (r ThreeBuilder) One(v int) ThreeBuilder {
	r.one = v
	return r
}

func (r Three) Two() string {
	return r.two
}

func (r Three) WithTwo(v string) Three {
	r.two = v
	return r
}

func (r ThreeBuilder) Two(v string) ThreeBuilder {
	r.two = v
	return r
}

func (r Three) Three() float64 {
	return r.three
}

func (r Three) WithThree(v float64) Three {
	r.three = v
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

func (r Three) AsMap() map[string]any {
	return map[string]any{
		"one":   r.one,
		"two":   r.two,
		"three": r.three,
	}
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

func (r Three) AsLabelled() fp.Labelled3[NameIsOne[int], NameIsTwo[string], NameIsThree[float64]] {
	return as.Labelled3(NameIsOne[int]{r.one}, NameIsTwo[string]{r.two}, NameIsThree[float64]{r.three})
}

func (r ThreeBuilder) FromLabelled(t fp.Labelled3[NameIsOne[int], NameIsTwo[string], NameIsThree[float64]]) ThreeBuilder {
	r.one = t.I1.Value()
	r.two = t.I2.Value()
	r.three = t.I3.Value()
	return r
}

type NameIsHello[T any] fp.Tuple1[T]

func (r NameIsHello[T]) Name() string {
	return "hello"
}
func (r NameIsHello[T]) Value() T {
	return r.I1
}
func (r NameIsHello[T]) WithValue(v T) NameIsHello[T] {
	r.I1 = v
	return r
}

type NameIsLanguage[T any] fp.Tuple1[T]

func (r NameIsLanguage[T]) Name() string {
	return "language"
}
func (r NameIsLanguage[T]) Value() T {
	return r.I1
}
func (r NameIsLanguage[T]) WithValue(v T) NameIsLanguage[T] {
	r.I1 = v
	return r
}

type NameIsOne[T any] fp.Tuple1[T]

func (r NameIsOne[T]) Name() string {
	return "one"
}
func (r NameIsOne[T]) Value() T {
	return r.I1
}
func (r NameIsOne[T]) WithValue(v T) NameIsOne[T] {
	r.I1 = v
	return r
}

type NameIsThree[T any] fp.Tuple1[T]

func (r NameIsThree[T]) Name() string {
	return "three"
}
func (r NameIsThree[T]) Value() T {
	return r.I1
}
func (r NameIsThree[T]) WithValue(v T) NameIsThree[T] {
	r.I1 = v
	return r
}

type NameIsTwo[T any] fp.Tuple1[T]

func (r NameIsTwo[T]) Name() string {
	return "two"
}
func (r NameIsTwo[T]) Value() T {
	return r.I1
}
func (r NameIsTwo[T]) WithValue(v T) NameIsTwo[T] {
	r.I1 = v
	return r
}
