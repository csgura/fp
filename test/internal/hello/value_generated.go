// Code generated by gombok, DO NOT EDIT.
package hello

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

func (r World) AsLabelled() fp.Labelled2[NameIsMessage[string], NameIsTimestamp[time.Time]] {
	return as.Labelled2(NameIsMessage[string]{r.message}, NameIsTimestamp[time.Time]{r.timestamp})
}

func (r WorldBuilder) FromLabelled(t fp.Labelled2[NameIsMessage[string], NameIsTimestamp[time.Time]]) WorldBuilder {
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

func (r HasOption) AsLabelled() fp.Labelled4[NameIsMessage[string], NameIsAddr[fp.Option[string]], NameIsPhone[[]string], NameIsEmptySeq[[]int]] {
	return as.Labelled4(NameIsMessage[string]{r.message}, NameIsAddr[fp.Option[string]]{r.addr}, NameIsPhone[[]string]{r.phone}, NameIsEmptySeq[[]int]{r.emptySeq})
}

func (r HasOptionBuilder) FromLabelled(t fp.Labelled4[NameIsMessage[string], NameIsAddr[fp.Option[string]], NameIsPhone[[]string], NameIsEmptySeq[[]int]]) HasOptionBuilder {
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

type NameIsAddr[T any] fp.Tuple1[T]

func (r NameIsAddr[T]) Name() string {
	return "addr"
}
func (r NameIsAddr[T]) Value() T {
	return r.I1
}
func (r NameIsAddr[T]) WithValue(v T) NameIsAddr[T] {
	r.I1 = v
	return r
}

type NameIsEmptySeq[T any] fp.Tuple1[T]

func (r NameIsEmptySeq[T]) Name() string {
	return "emptySeq"
}
func (r NameIsEmptySeq[T]) Value() T {
	return r.I1
}
func (r NameIsEmptySeq[T]) WithValue(v T) NameIsEmptySeq[T] {
	r.I1 = v
	return r
}

type NameIsMessage[T any] fp.Tuple1[T]

func (r NameIsMessage[T]) Name() string {
	return "message"
}
func (r NameIsMessage[T]) Value() T {
	return r.I1
}
func (r NameIsMessage[T]) WithValue(v T) NameIsMessage[T] {
	r.I1 = v
	return r
}

type NameIsPhone[T any] fp.Tuple1[T]

func (r NameIsPhone[T]) Name() string {
	return "phone"
}
func (r NameIsPhone[T]) Value() T {
	return r.I1
}
func (r NameIsPhone[T]) WithValue(v T) NameIsPhone[T] {
	r.I1 = v
	return r
}

type NameIsTimestamp[T any] fp.Tuple1[T]

func (r NameIsTimestamp[T]) Name() string {
	return "timestamp"
}
func (r NameIsTimestamp[T]) Value() T {
	return r.I1
}
func (r NameIsTimestamp[T]) WithValue(v T) NameIsTimestamp[T] {
	r.I1 = v
	return r
}
