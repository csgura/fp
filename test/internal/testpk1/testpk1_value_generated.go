// Code generated by gombok, DO NOT EDIT.
package testpk1

import (
	"encoding/json"
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
	"net/http"
	"time"
)

type WorldBuilder World

type WorldMutable struct {
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

func (r WorldBuilder) Build() World {
	return World(r)
}

func (r World) Builder() WorldBuilder {
	return WorldBuilder(r)
}

func (r World) Message() string {
	return r.message
}

func (r World) WithMessage(v string) World {
	r.message = v
	return r
}

func (r WorldBuilder) Message(v string) WorldBuilder {
	r.message = v
	return r
}

func (r World) Timestamp() time.Time {
	return r.timestamp
}

func (r World) WithTimestamp(v time.Time) World {
	r.timestamp = v
	return r
}

func (r WorldBuilder) Timestamp(v time.Time) WorldBuilder {
	r.timestamp = v
	return r
}

func (r World) String() string {
	return fmt.Sprintf("World(message=%v, timestamp=%v)", r.message, r.timestamp)
}

func (r World) AsTuple() fp.Tuple2[string, time.Time] {
	return as.Tuple2(r.message, r.timestamp)
}

func (r World) AsMutable() WorldMutable {
	return WorldMutable{
		Message:   r.message,
		Timestamp: r.timestamp,
	}
}

func (r WorldMutable) AsImmutable() World {
	return World{
		message:   r.Message,
		timestamp: r.Timestamp,
	}
}

func (r WorldBuilder) FromTuple(t fp.Tuple2[string, time.Time]) WorldBuilder {
	r.message = t.I1
	r.timestamp = t.I2
	return r
}

func (r World) AsMap() map[string]any {
	return map[string]any{
		"message":   r.message,
		"timestamp": r.timestamp,
	}
}

func (r WorldBuilder) FromMap(m map[string]any) WorldBuilder {

	if v, ok := m["message"].(string); ok {
		r.message = v
	}

	if v, ok := m["timestamp"].(time.Time); ok {
		r.timestamp = v
	}

	return r
}

func (r World) AsLabelled() fp.Labelled2[NamedMessage[string], NamedTimestamp[time.Time]] {
	return as.Labelled2(NamedMessage[string]{r.message}, NamedTimestamp[time.Time]{r.timestamp})
}

func (r WorldBuilder) FromLabelled(t fp.Labelled2[NamedMessage[string], NamedTimestamp[time.Time]]) WorldBuilder {
	r.message = t.I1.Value()
	r.timestamp = t.I2.Value()
	return r
}

func (r World) MarshalJSON() ([]byte, error) {
	m := r.AsMutable()
	return json.Marshal(m)
}

func (r *World) UnmarshalJSON(b []byte) error {
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

type HasOptionBuilder HasOption

type HasOptionMutable struct {
	Message  string
	Addr     fp.Option[string]
	Phone    []string
	EmptySeq []int
}

func (r HasOptionBuilder) Build() HasOption {
	return HasOption(r)
}

func (r HasOption) Builder() HasOptionBuilder {
	return HasOptionBuilder(r)
}

func (r HasOption) Message() string {
	return r.message
}

func (r HasOption) WithMessage(v string) HasOption {
	r.message = v
	return r
}

func (r HasOptionBuilder) Message(v string) HasOptionBuilder {
	r.message = v
	return r
}

func (r HasOption) Addr() fp.Option[string] {
	return r.addr
}

func (r HasOption) WithAddr(v fp.Option[string]) HasOption {
	r.addr = v
	return r
}

func (r HasOptionBuilder) Addr(v fp.Option[string]) HasOptionBuilder {
	r.addr = v
	return r
}

func (r HasOption) WithSomeAddr(v string) HasOption {
	r.addr = option.Some(v)
	return r
}

func (r HasOption) WithNoneAddr() HasOption {
	r.addr = option.None[string]()
	return r
}

func (r HasOptionBuilder) SomeAddr(v string) HasOptionBuilder {
	r.addr = option.Some(v)
	return r
}

func (r HasOptionBuilder) NoneAddr() HasOptionBuilder {
	r.addr = option.None[string]()
	return r
}

func (r HasOption) Phone() []string {
	return r.phone
}

func (r HasOption) WithPhone(v []string) HasOption {
	r.phone = v
	return r
}

func (r HasOptionBuilder) Phone(v []string) HasOptionBuilder {
	r.phone = v
	return r
}

func (r HasOption) EmptySeq() []int {
	return r.emptySeq
}

func (r HasOption) WithEmptySeq(v []int) HasOption {
	r.emptySeq = v
	return r
}

func (r HasOptionBuilder) EmptySeq(v []int) HasOptionBuilder {
	r.emptySeq = v
	return r
}

func (r HasOption) String() string {
	return fmt.Sprintf("HasOption(message=%v, addr=%v, phone=%v, emptySeq=%v)", r.message, r.addr, r.phone, r.emptySeq)
}

func (r HasOption) AsTuple() fp.Tuple4[string, fp.Option[string], []string, []int] {
	return as.Tuple4(r.message, r.addr, r.phone, r.emptySeq)
}

func (r HasOption) AsMutable() HasOptionMutable {
	return HasOptionMutable{
		Message:  r.message,
		Addr:     r.addr,
		Phone:    r.phone,
		EmptySeq: r.emptySeq,
	}
}

func (r HasOptionMutable) AsImmutable() HasOption {
	return HasOption{
		message:  r.Message,
		addr:     r.Addr,
		phone:    r.Phone,
		emptySeq: r.EmptySeq,
	}
}

func (r HasOptionBuilder) FromTuple(t fp.Tuple4[string, fp.Option[string], []string, []int]) HasOptionBuilder {
	r.message = t.I1
	r.addr = t.I2
	r.phone = t.I3
	r.emptySeq = t.I4
	return r
}

func (r HasOption) AsMap() map[string]any {
	return map[string]any{
		"message":  r.message,
		"addr":     r.addr,
		"phone":    r.phone,
		"emptySeq": r.emptySeq,
	}
}

func (r HasOptionBuilder) FromMap(m map[string]any) HasOptionBuilder {

	if v, ok := m["message"].(string); ok {
		r.message = v
	}

	if v, ok := m["addr"].(fp.Option[string]); ok {
		r.addr = v
	}

	if v, ok := m["phone"].([]string); ok {
		r.phone = v
	}

	if v, ok := m["emptySeq"].([]int); ok {
		r.emptySeq = v
	}

	return r
}

func (r HasOption) AsLabelled() fp.Labelled4[NamedMessage[string], NamedAddr[fp.Option[string]], NamedPhone[[]string], NamedEmptySeq[[]int]] {
	return as.Labelled4(NamedMessage[string]{r.message}, NamedAddr[fp.Option[string]]{r.addr}, NamedPhone[[]string]{r.phone}, NamedEmptySeq[[]int]{r.emptySeq})
}

func (r HasOptionBuilder) FromLabelled(t fp.Labelled4[NamedMessage[string], NamedAddr[fp.Option[string]], NamedPhone[[]string], NamedEmptySeq[[]int]]) HasOptionBuilder {
	r.message = t.I1.Value()
	r.addr = t.I2.Value()
	r.phone = t.I3.Value()
	r.emptySeq = t.I4.Value()
	return r
}

type CustomValueMutable struct {
	A string
	B int
}

func (r CustomValueBuilder) Build() CustomValue {
	return CustomValue(r)
}

func (r CustomValue) Builder() CustomValueBuilder {
	return CustomValueBuilder(r)
}

func (r CustomValue) WithA(v string) CustomValue {
	r.a = v
	return r
}

func (r CustomValueBuilder) A(v string) CustomValueBuilder {
	r.a = v
	return r
}

func (r CustomValue) B() int {
	return r.b
}

func (r CustomValue) String() string {
	return fmt.Sprintf("CustomValue(a=%v, b=%v)", r.a, r.b)
}

func (r CustomValue) AsTuple() fp.Tuple2[string, int] {
	return as.Tuple2(r.a, r.b)
}

func (r CustomValue) AsMutable() CustomValueMutable {
	return CustomValueMutable{
		A: r.a,
		B: r.b,
	}
}

func (r CustomValueMutable) AsImmutable() CustomValue {
	return CustomValue{
		a: r.A,
		b: r.B,
	}
}

func (r CustomValueBuilder) FromTuple(t fp.Tuple2[string, int]) CustomValueBuilder {
	r.a = t.I1
	r.b = t.I2
	return r
}

func (r CustomValue) AsMap() map[string]any {
	return map[string]any{
		"a": r.a,
		"b": r.b,
	}
}

func (r CustomValueBuilder) FromMap(m map[string]any) CustomValueBuilder {

	if v, ok := m["a"].(string); ok {
		r.a = v
	}

	if v, ok := m["b"].(int); ok {
		r.b = v
	}

	return r
}

type AliasedStructBuilder AliasedStruct

type AliasedStructMutable struct {
	Message   string
	Timestamp time.Time
}

func (r AliasedStructBuilder) Build() AliasedStruct {
	return AliasedStruct(r)
}

func (r AliasedStruct) Builder() AliasedStructBuilder {
	return AliasedStructBuilder(r)
}

func (r AliasedStruct) Message() string {
	return r.message
}

func (r AliasedStruct) WithMessage(v string) AliasedStruct {
	r.message = v
	return r
}

func (r AliasedStructBuilder) Message(v string) AliasedStructBuilder {
	r.message = v
	return r
}

func (r AliasedStruct) Timestamp() time.Time {
	return r.timestamp
}

func (r AliasedStruct) WithTimestamp(v time.Time) AliasedStruct {
	r.timestamp = v
	return r
}

func (r AliasedStructBuilder) Timestamp(v time.Time) AliasedStructBuilder {
	r.timestamp = v
	return r
}

func (r AliasedStruct) String() string {
	return fmt.Sprintf("AliasedStruct(message=%v, timestamp=%v)", r.message, r.timestamp)
}

func (r AliasedStruct) AsTuple() fp.Tuple2[string, time.Time] {
	return as.Tuple2(r.message, r.timestamp)
}

func (r AliasedStruct) AsMutable() AliasedStructMutable {
	return AliasedStructMutable{
		Message:   r.message,
		Timestamp: r.timestamp,
	}
}

func (r AliasedStructMutable) AsImmutable() AliasedStruct {
	return AliasedStruct{
		message:   r.Message,
		timestamp: r.Timestamp,
	}
}

func (r AliasedStructBuilder) FromTuple(t fp.Tuple2[string, time.Time]) AliasedStructBuilder {
	r.message = t.I1
	r.timestamp = t.I2
	return r
}

func (r AliasedStruct) AsMap() map[string]any {
	return map[string]any{
		"message":   r.message,
		"timestamp": r.timestamp,
	}
}

func (r AliasedStructBuilder) FromMap(m map[string]any) AliasedStructBuilder {

	if v, ok := m["message"].(string); ok {
		r.message = v
	}

	if v, ok := m["timestamp"].(time.Time); ok {
		r.timestamp = v
	}

	return r
}

type HListInsideHListBuilder HListInsideHList

type HListInsideHListMutable struct {
	Tp    fp.Tuple2[string, int]
	Value string
	Hello World
}

func (r HListInsideHListBuilder) Build() HListInsideHList {
	return HListInsideHList(r)
}

func (r HListInsideHList) Builder() HListInsideHListBuilder {
	return HListInsideHListBuilder(r)
}

func (r HListInsideHList) Tp() fp.Tuple2[string, int] {
	return r.tp
}

func (r HListInsideHList) WithTp(v fp.Tuple2[string, int]) HListInsideHList {
	r.tp = v
	return r
}

func (r HListInsideHListBuilder) Tp(v fp.Tuple2[string, int]) HListInsideHListBuilder {
	r.tp = v
	return r
}

func (r HListInsideHList) Value() string {
	return r.value
}

func (r HListInsideHList) WithValue(v string) HListInsideHList {
	r.value = v
	return r
}

func (r HListInsideHListBuilder) Value(v string) HListInsideHListBuilder {
	r.value = v
	return r
}

func (r HListInsideHList) Hello() World {
	return r.hello
}

func (r HListInsideHList) WithHello(v World) HListInsideHList {
	r.hello = v
	return r
}

func (r HListInsideHListBuilder) Hello(v World) HListInsideHListBuilder {
	r.hello = v
	return r
}

func (r HListInsideHList) String() string {
	return fmt.Sprintf("HListInsideHList(tp=%v, value=%v, hello=%v)", r.tp, r.value, r.hello)
}

func (r HListInsideHList) AsTuple() fp.Tuple3[fp.Tuple2[string, int], string, World] {
	return as.Tuple3(r.tp, r.value, r.hello)
}

func (r HListInsideHList) AsMutable() HListInsideHListMutable {
	return HListInsideHListMutable{
		Tp:    r.tp,
		Value: r.value,
		Hello: r.hello,
	}
}

func (r HListInsideHListMutable) AsImmutable() HListInsideHList {
	return HListInsideHList{
		tp:    r.Tp,
		value: r.Value,
		hello: r.Hello,
	}
}

func (r HListInsideHListBuilder) FromTuple(t fp.Tuple3[fp.Tuple2[string, int], string, World]) HListInsideHListBuilder {
	r.tp = t.I1
	r.value = t.I2
	r.hello = t.I3
	return r
}

func (r HListInsideHList) AsMap() map[string]any {
	return map[string]any{
		"tp":    r.tp,
		"value": r.value,
		"hello": r.hello,
	}
}

func (r HListInsideHListBuilder) FromMap(m map[string]any) HListInsideHListBuilder {

	if v, ok := m["tp"].(fp.Tuple2[string, int]); ok {
		r.tp = v
	}

	if v, ok := m["value"].(string); ok {
		r.value = v
	}

	if v, ok := m["hello"].(World); ok {
		r.hello = v
	}

	return r
}

type WrapperBuilder[T any] Wrapper[T]

type WrapperMutable[T any] struct {
	Unwrap T
}

func (r WrapperBuilder[T]) Build() Wrapper[T] {
	return Wrapper[T](r)
}

func (r Wrapper[T]) Builder() WrapperBuilder[T] {
	return WrapperBuilder[T](r)
}

func (r Wrapper[T]) Unwrap() T {
	return r.unwrap
}

func (r Wrapper[T]) WithUnwrap(v T) Wrapper[T] {
	r.unwrap = v
	return r
}

func (r WrapperBuilder[T]) Unwrap(v T) WrapperBuilder[T] {
	r.unwrap = v
	return r
}

func (r Wrapper[T]) String() string {
	return fmt.Sprintf("Wrapper(unwrap=%v)", r.unwrap)
}

func (r Wrapper[T]) AsTuple() fp.Tuple1[T] {
	return as.Tuple1(r.unwrap)
}

func (r Wrapper[T]) AsMutable() WrapperMutable[T] {
	return WrapperMutable[T]{
		Unwrap: r.unwrap,
	}
}

func (r WrapperMutable[T]) AsImmutable() Wrapper[T] {
	return Wrapper[T]{
		unwrap: r.Unwrap,
	}
}

func (r WrapperBuilder[T]) FromTuple(t fp.Tuple1[T]) WrapperBuilder[T] {
	r.unwrap = t.I1
	return r
}

func (r Wrapper[T]) AsMap() map[string]any {
	return map[string]any{
		"unwrap": r.unwrap,
	}
}

func (r WrapperBuilder[T]) FromMap(m map[string]any) WrapperBuilder[T] {

	if v, ok := m["unwrap"].(T); ok {
		r.unwrap = v
	}

	return r
}

type TestOrderedEqBuilder TestOrderedEq

type TestOrderedEqMutable struct {
	List  fp.Seq[int]
	Tlist fp.Seq[fp.Tuple2[int, int]]
}

func (r TestOrderedEqBuilder) Build() TestOrderedEq {
	return TestOrderedEq(r)
}

func (r TestOrderedEq) Builder() TestOrderedEqBuilder {
	return TestOrderedEqBuilder(r)
}

func (r TestOrderedEq) List() fp.Seq[int] {
	return r.list
}

func (r TestOrderedEq) WithList(v fp.Seq[int]) TestOrderedEq {
	r.list = v
	return r
}

func (r TestOrderedEqBuilder) List(v fp.Seq[int]) TestOrderedEqBuilder {
	r.list = v
	return r
}

func (r TestOrderedEq) Tlist() fp.Seq[fp.Tuple2[int, int]] {
	return r.tlist
}

func (r TestOrderedEq) WithTlist(v fp.Seq[fp.Tuple2[int, int]]) TestOrderedEq {
	r.tlist = v
	return r
}

func (r TestOrderedEqBuilder) Tlist(v fp.Seq[fp.Tuple2[int, int]]) TestOrderedEqBuilder {
	r.tlist = v
	return r
}

func (r TestOrderedEq) String() string {
	return fmt.Sprintf("TestOrderedEq(list=%v, tlist=%v)", r.list, r.tlist)
}

func (r TestOrderedEq) AsTuple() fp.Tuple2[fp.Seq[int], fp.Seq[fp.Tuple2[int, int]]] {
	return as.Tuple2(r.list, r.tlist)
}

func (r TestOrderedEq) AsMutable() TestOrderedEqMutable {
	return TestOrderedEqMutable{
		List:  r.list,
		Tlist: r.tlist,
	}
}

func (r TestOrderedEqMutable) AsImmutable() TestOrderedEq {
	return TestOrderedEq{
		list:  r.List,
		tlist: r.Tlist,
	}
}

func (r TestOrderedEqBuilder) FromTuple(t fp.Tuple2[fp.Seq[int], fp.Seq[fp.Tuple2[int, int]]]) TestOrderedEqBuilder {
	r.list = t.I1
	r.tlist = t.I2
	return r
}

func (r TestOrderedEq) AsMap() map[string]any {
	return map[string]any{
		"list":  r.list,
		"tlist": r.tlist,
	}
}

func (r TestOrderedEqBuilder) FromMap(m map[string]any) TestOrderedEqBuilder {

	if v, ok := m["list"].(fp.Seq[int]); ok {
		r.list = v
	}

	if v, ok := m["tlist"].(fp.Seq[fp.Tuple2[int, int]]); ok {
		r.tlist = v
	}

	return r
}

type MapEqBuilder MapEq

type MapEqMutable struct {
	M  map[string]World
	M2 fp.Map[string, World]
}

func (r MapEqBuilder) Build() MapEq {
	return MapEq(r)
}

func (r MapEq) Builder() MapEqBuilder {
	return MapEqBuilder(r)
}

func (r MapEq) M() map[string]World {
	return r.m
}

func (r MapEq) WithM(v map[string]World) MapEq {
	r.m = v
	return r
}

func (r MapEqBuilder) M(v map[string]World) MapEqBuilder {
	r.m = v
	return r
}

func (r MapEq) M2() fp.Map[string, World] {
	return r.m2
}

func (r MapEq) WithM2(v fp.Map[string, World]) MapEq {
	r.m2 = v
	return r
}

func (r MapEqBuilder) M2(v fp.Map[string, World]) MapEqBuilder {
	r.m2 = v
	return r
}

func (r MapEq) String() string {
	return fmt.Sprintf("MapEq(m=%v, m2=%v)", r.m, r.m2)
}

func (r MapEq) AsTuple() fp.Tuple2[map[string]World, fp.Map[string, World]] {
	return as.Tuple2(r.m, r.m2)
}

func (r MapEq) AsMutable() MapEqMutable {
	return MapEqMutable{
		M:  r.m,
		M2: r.m2,
	}
}

func (r MapEqMutable) AsImmutable() MapEq {
	return MapEq{
		m:  r.M,
		m2: r.M2,
	}
}

func (r MapEqBuilder) FromTuple(t fp.Tuple2[map[string]World, fp.Map[string, World]]) MapEqBuilder {
	r.m = t.I1
	r.m2 = t.I2
	return r
}

func (r MapEq) AsMap() map[string]any {
	return map[string]any{
		"m":  r.m,
		"m2": r.m2,
	}
}

func (r MapEqBuilder) FromMap(m map[string]any) MapEqBuilder {

	if v, ok := m["m"].(map[string]World); ok {
		r.m = v
	}

	if v, ok := m["m2"].(fp.Map[string, World]); ok {
		r.m2 = v
	}

	return r
}

type SeqMonoidBuilder SeqMonoid

type SeqMonoidMutable struct {
	V  string
	S  fp.Seq[string]
	M  map[string]int
	M2 fp.Map[string, World]
}

func (r SeqMonoidBuilder) Build() SeqMonoid {
	return SeqMonoid(r)
}

func (r SeqMonoid) Builder() SeqMonoidBuilder {
	return SeqMonoidBuilder(r)
}

func (r SeqMonoid) V() string {
	return r.v
}

func (r SeqMonoid) WithV(v string) SeqMonoid {
	r.v = v
	return r
}

func (r SeqMonoidBuilder) V(v string) SeqMonoidBuilder {
	r.v = v
	return r
}

func (r SeqMonoid) S() fp.Seq[string] {
	return r.s
}

func (r SeqMonoid) WithS(v fp.Seq[string]) SeqMonoid {
	r.s = v
	return r
}

func (r SeqMonoidBuilder) S(v fp.Seq[string]) SeqMonoidBuilder {
	r.s = v
	return r
}

func (r SeqMonoid) M() map[string]int {
	return r.m
}

func (r SeqMonoid) WithM(v map[string]int) SeqMonoid {
	r.m = v
	return r
}

func (r SeqMonoidBuilder) M(v map[string]int) SeqMonoidBuilder {
	r.m = v
	return r
}

func (r SeqMonoid) M2() fp.Map[string, World] {
	return r.m2
}

func (r SeqMonoid) WithM2(v fp.Map[string, World]) SeqMonoid {
	r.m2 = v
	return r
}

func (r SeqMonoidBuilder) M2(v fp.Map[string, World]) SeqMonoidBuilder {
	r.m2 = v
	return r
}

func (r SeqMonoid) String() string {
	return fmt.Sprintf("SeqMonoid(v=%v, s=%v, m=%v, m2=%v)", r.v, r.s, r.m, r.m2)
}

func (r SeqMonoid) AsTuple() fp.Tuple4[string, fp.Seq[string], map[string]int, fp.Map[string, World]] {
	return as.Tuple4(r.v, r.s, r.m, r.m2)
}

func (r SeqMonoid) AsMutable() SeqMonoidMutable {
	return SeqMonoidMutable{
		V:  r.v,
		S:  r.s,
		M:  r.m,
		M2: r.m2,
	}
}

func (r SeqMonoidMutable) AsImmutable() SeqMonoid {
	return SeqMonoid{
		v:  r.V,
		s:  r.S,
		m:  r.M,
		m2: r.M2,
	}
}

func (r SeqMonoidBuilder) FromTuple(t fp.Tuple4[string, fp.Seq[string], map[string]int, fp.Map[string, World]]) SeqMonoidBuilder {
	r.v = t.I1
	r.s = t.I2
	r.m = t.I3
	r.m2 = t.I4
	return r
}

func (r SeqMonoid) AsMap() map[string]any {
	return map[string]any{
		"v":  r.v,
		"s":  r.s,
		"m":  r.m,
		"m2": r.m2,
	}
}

func (r SeqMonoidBuilder) FromMap(m map[string]any) SeqMonoidBuilder {

	if v, ok := m["v"].(string); ok {
		r.v = v
	}

	if v, ok := m["s"].(fp.Seq[string]); ok {
		r.s = v
	}

	if v, ok := m["m"].(map[string]int); ok {
		r.m = v
	}

	if v, ok := m["m2"].(fp.Map[string, World]); ok {
		r.m2 = v
	}

	return r
}

type MapEqParamBuilder[K any, V any] MapEqParam[K, V]

type MapEqParamMutable[K any, V any] struct {
	M fp.Map[K, V]
}

func (r MapEqParamBuilder[K, V]) Build() MapEqParam[K, V] {
	return MapEqParam[K, V](r)
}

func (r MapEqParam[K, V]) Builder() MapEqParamBuilder[K, V] {
	return MapEqParamBuilder[K, V](r)
}

func (r MapEqParam[K, V]) M() fp.Map[K, V] {
	return r.m
}

func (r MapEqParam[K, V]) WithM(v fp.Map[K, V]) MapEqParam[K, V] {
	r.m = v
	return r
}

func (r MapEqParamBuilder[K, V]) M(v fp.Map[K, V]) MapEqParamBuilder[K, V] {
	r.m = v
	return r
}

func (r MapEqParam[K, V]) String() string {
	return fmt.Sprintf("MapEqParam(m=%v)", r.m)
}

func (r MapEqParam[K, V]) AsTuple() fp.Tuple1[fp.Map[K, V]] {
	return as.Tuple1(r.m)
}

func (r MapEqParam[K, V]) AsMutable() MapEqParamMutable[K, V] {
	return MapEqParamMutable[K, V]{
		M: r.m,
	}
}

func (r MapEqParamMutable[K, V]) AsImmutable() MapEqParam[K, V] {
	return MapEqParam[K, V]{
		m: r.M,
	}
}

func (r MapEqParamBuilder[K, V]) FromTuple(t fp.Tuple1[fp.Map[K, V]]) MapEqParamBuilder[K, V] {
	r.m = t.I1
	return r
}

func (r MapEqParam[K, V]) AsMap() map[string]any {
	return map[string]any{
		"m": r.m,
	}
}

func (r MapEqParamBuilder[K, V]) FromMap(m map[string]any) MapEqParamBuilder[K, V] {

	if v, ok := m["m"].(fp.Map[K, V]); ok {
		r.m = v
	}

	return r
}

type NotUsedProblemBuilder NotUsedProblem

type NotUsedProblemMutable struct {
	M MapEqParam[string, int]
}

func (r NotUsedProblemBuilder) Build() NotUsedProblem {
	return NotUsedProblem(r)
}

func (r NotUsedProblem) Builder() NotUsedProblemBuilder {
	return NotUsedProblemBuilder(r)
}

func (r NotUsedProblem) M() MapEqParam[string, int] {
	return r.m
}

func (r NotUsedProblem) WithM(v MapEqParam[string, int]) NotUsedProblem {
	r.m = v
	return r
}

func (r NotUsedProblemBuilder) M(v MapEqParam[string, int]) NotUsedProblemBuilder {
	r.m = v
	return r
}

func (r NotUsedProblem) String() string {
	return fmt.Sprintf("NotUsedProblem(m=%v)", r.m)
}

func (r NotUsedProblem) AsTuple() fp.Tuple1[MapEqParam[string, int]] {
	return as.Tuple1(r.m)
}

func (r NotUsedProblem) AsMutable() NotUsedProblemMutable {
	return NotUsedProblemMutable{
		M: r.m,
	}
}

func (r NotUsedProblemMutable) AsImmutable() NotUsedProblem {
	return NotUsedProblem{
		m: r.M,
	}
}

func (r NotUsedProblemBuilder) FromTuple(t fp.Tuple1[MapEqParam[string, int]]) NotUsedProblemBuilder {
	r.m = t.I1
	return r
}

func (r NotUsedProblem) AsMap() map[string]any {
	return map[string]any{
		"m": r.m,
	}
}

func (r NotUsedProblemBuilder) FromMap(m map[string]any) NotUsedProblemBuilder {

	if v, ok := m["m"].(MapEqParam[string, int]); ok {
		r.m = v
	}

	return r
}

type NodeBuilder Node

type NodeMutable struct {
	Value string
	Left  *Node
	Right *Node
}

func (r NodeBuilder) Build() Node {
	return Node(r)
}

func (r Node) Builder() NodeBuilder {
	return NodeBuilder(r)
}

func (r Node) Value() string {
	return r.value
}

func (r Node) WithValue(v string) Node {
	r.value = v
	return r
}

func (r NodeBuilder) Value(v string) NodeBuilder {
	r.value = v
	return r
}

func (r Node) Left() *Node {
	return r.left
}

func (r Node) WithLeft(v *Node) Node {
	r.left = v
	return r
}

func (r NodeBuilder) Left(v *Node) NodeBuilder {
	r.left = v
	return r
}

func (r Node) Right() *Node {
	return r.right
}

func (r Node) WithRight(v *Node) Node {
	r.right = v
	return r
}

func (r NodeBuilder) Right(v *Node) NodeBuilder {
	r.right = v
	return r
}

func (r Node) String() string {
	return fmt.Sprintf("Node(value=%v, left=%v, right=%v)", r.value, r.left, r.right)
}

func (r Node) AsTuple() fp.Tuple3[string, *Node, *Node] {
	return as.Tuple3(r.value, r.left, r.right)
}

func (r Node) AsMutable() NodeMutable {
	return NodeMutable{
		Value: r.value,
		Left:  r.left,
		Right: r.right,
	}
}

func (r NodeMutable) AsImmutable() Node {
	return Node{
		value: r.Value,
		left:  r.Left,
		right: r.Right,
	}
}

func (r NodeBuilder) FromTuple(t fp.Tuple3[string, *Node, *Node]) NodeBuilder {
	r.value = t.I1
	r.left = t.I2
	r.right = t.I3
	return r
}

func (r Node) AsMap() map[string]any {
	return map[string]any{
		"value": r.value,
		"left":  r.left,
		"right": r.right,
	}
}

func (r NodeBuilder) FromMap(m map[string]any) NodeBuilder {

	if v, ok := m["value"].(string); ok {
		r.value = v
	}

	if v, ok := m["left"].(*Node); ok {
		r.left = v
	}

	if v, ok := m["right"].(*Node); ok {
		r.right = v
	}

	return r
}

type NamedAddr[T any] fp.Tuple1[T]

func (r NamedAddr[T]) Name() string {
	return "addr"
}
func (r NamedAddr[T]) Value() T {
	return r.I1
}
func (r NamedAddr[T]) WithValue(v T) NamedAddr[T] {
	r.I1 = v
	return r
}

type NamedEmptySeq[T any] fp.Tuple1[T]

func (r NamedEmptySeq[T]) Name() string {
	return "emptySeq"
}
func (r NamedEmptySeq[T]) Value() T {
	return r.I1
}
func (r NamedEmptySeq[T]) WithValue(v T) NamedEmptySeq[T] {
	r.I1 = v
	return r
}

type NamedMessage[T any] fp.Tuple1[T]

func (r NamedMessage[T]) Name() string {
	return "message"
}
func (r NamedMessage[T]) Value() T {
	return r.I1
}
func (r NamedMessage[T]) WithValue(v T) NamedMessage[T] {
	r.I1 = v
	return r
}

type NamedPhone[T any] fp.Tuple1[T]

func (r NamedPhone[T]) Name() string {
	return "phone"
}
func (r NamedPhone[T]) Value() T {
	return r.I1
}
func (r NamedPhone[T]) WithValue(v T) NamedPhone[T] {
	r.I1 = v
	return r
}

type NamedTimestamp[T any] fp.Tuple1[T]

func (r NamedTimestamp[T]) Name() string {
	return "timestamp"
}
func (r NamedTimestamp[T]) Value() T {
	return r.I1
}
func (r NamedTimestamp[T]) WithValue(v T) NamedTimestamp[T] {
	r.I1 = v
	return r
}
