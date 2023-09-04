// Code generated by gombok, DO NOT EDIT.
package docexample

import (
	"encoding/json"
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/option"
	"net/http"
)

type PersonMutable struct {
	Name string
	Age  int
}

func (r Person) Name() string {
	return r.name
}

func (r Person) Age() int {
	return r.age
}

func (r Person) WithName(v string) Person {
	r.name = v
	return r
}

func (r Person) WithAge(v int) Person {
	r.age = v
	return r
}

type PersonBuilder Person

func (r PersonBuilder) Build() Person {
	return Person(r)
}

func (r Person) Builder() PersonBuilder {
	return PersonBuilder(r)
}

func (r PersonBuilder) Name(v string) PersonBuilder {
	r.name = v
	return r
}

func (r PersonBuilder) Age(v int) PersonBuilder {
	r.age = v
	return r
}

func (r PersonBuilder) FromTuple(t fp.Tuple2[string, int]) PersonBuilder {
	r.name = t.I1
	r.age = t.I2
	return r
}

func (r PersonBuilder) Apply(name string, age int) PersonBuilder {
	r.name = name
	r.age = age
	return r
}

func (r PersonBuilder) FromMap(m map[string]any) PersonBuilder {

	if v, ok := m["name"].(string); ok {
		r.name = v
	}

	if v, ok := m["age"].(int); ok {
		r.age = v
	}

	return r
}

func (r Person) String() string {
	return fmt.Sprintf("Person(name=%v, age=%v)", r.name, r.age)
}

func (r Person) AsTuple() fp.Tuple2[string, int] {
	return as.Tuple2(r.name, r.age)
}

func (r Person) Unapply() (string, int) {
	return r.name, r.age
}

func (r Person) AsMutable() PersonMutable {
	return PersonMutable{
		Name: r.name,
		Age:  r.age,
	}
}

func (r PersonMutable) AsImmutable() Person {
	return Person{
		name: r.Name,
		age:  r.Age,
	}
}

func (r Person) AsMap() map[string]any {
	m := map[string]any{}
	m["name"] = r.name
	m["age"] = r.age
	return m
}

type AddressMutable struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
}

func (r Address) Country() string {
	return r.country
}

func (r Address) City() string {
	return r.city
}

func (r Address) Street() string {
	return r.street
}

func (r Address) WithCountry(v string) Address {
	r.country = v
	return r
}

func (r Address) WithCity(v string) Address {
	r.city = v
	return r
}

func (r Address) WithStreet(v string) Address {
	r.street = v
	return r
}

type AddressBuilder Address

func (r AddressBuilder) Build() Address {
	return Address(r)
}

func (r Address) Builder() AddressBuilder {
	return AddressBuilder(r)
}

func (r AddressBuilder) Country(v string) AddressBuilder {
	r.country = v
	return r
}

func (r AddressBuilder) City(v string) AddressBuilder {
	r.city = v
	return r
}

func (r AddressBuilder) Street(v string) AddressBuilder {
	r.street = v
	return r
}

func (r AddressBuilder) FromTuple(t fp.Tuple3[string, string, string]) AddressBuilder {
	r.country = t.I1
	r.city = t.I2
	r.street = t.I3
	return r
}

func (r AddressBuilder) Apply(country string, city string, street string) AddressBuilder {
	r.country = country
	r.city = city
	r.street = street
	return r
}

func (r AddressBuilder) FromMap(m map[string]any) AddressBuilder {

	if v, ok := m["country"].(string); ok {
		r.country = v
	}

	if v, ok := m["city"].(string); ok {
		r.city = v
	}

	if v, ok := m["street"].(string); ok {
		r.street = v
	}

	return r
}

func (r Address) String() string {
	return fmt.Sprintf("Address(country=%v, city=%v, street=%v)", r.country, r.city, r.street)
}

func (r Address) AsTuple() fp.Tuple3[string, string, string] {
	return as.Tuple3(r.country, r.city, r.street)
}

func (r Address) Unapply() (string, string, string) {
	return r.country, r.city, r.street
}

func (r Address) AsMutable() AddressMutable {
	return AddressMutable{
		Country: r.country,
		City:    r.city,
		Street:  r.street,
	}
}

func (r AddressMutable) AsImmutable() Address {
	return Address{
		country: r.Country,
		city:    r.City,
		street:  r.Street,
	}
}

func (r Address) AsMap() map[string]any {
	m := map[string]any{}
	m["country"] = r.country
	m["city"] = r.city
	m["street"] = r.street
	return m
}

func (r Address) MarshalJSON() ([]byte, error) {
	m := r.AsMutable()
	return json.Marshal(m)
}

func (r *Address) UnmarshalJSON(b []byte) error {
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

type CarMutable struct {
	Company string
	Model   string
	Year    int
}

func (r Car) Company() string {
	return r.company
}

func (r Car) Model() string {
	return r.model
}

func (r Car) Year() int {
	return r.year
}

func (r Car) WithCompany(v string) Car {
	r.company = v
	return r
}

func (r Car) WithModel(v string) Car {
	r.model = v
	return r
}

func (r Car) WithYear(v int) Car {
	r.year = v
	return r
}

type CarBuilder Car

func (r CarBuilder) Build() Car {
	return Car(r)
}

func (r Car) Builder() CarBuilder {
	return CarBuilder(r)
}

func (r CarBuilder) Company(v string) CarBuilder {
	r.company = v
	return r
}

func (r CarBuilder) Model(v string) CarBuilder {
	r.model = v
	return r
}

func (r CarBuilder) Year(v int) CarBuilder {
	r.year = v
	return r
}

func (r CarBuilder) FromTuple(t fp.Tuple3[string, string, int]) CarBuilder {
	r.company = t.I1
	r.model = t.I2
	r.year = t.I3
	return r
}

func (r CarBuilder) Apply(company string, model string, year int) CarBuilder {
	r.company = company
	r.model = model
	r.year = year
	return r
}

func (r CarBuilder) FromMap(m map[string]any) CarBuilder {

	if v, ok := m["company"].(string); ok {
		r.company = v
	}

	if v, ok := m["model"].(string); ok {
		r.model = v
	}

	if v, ok := m["year"].(int); ok {
		r.year = v
	}

	return r
}

func (r CarBuilder) FromLabelled(t fp.Labelled3[NamedCompany[string], NamedModel[string], NamedYear[int]]) CarBuilder {
	r.company = t.I1.Value()
	r.model = t.I2.Value()
	r.year = t.I3.Value()
	return r
}

func (r Car) String() string {
	return fmt.Sprintf("Car(company=%v, model=%v, year=%v)", r.company, r.model, r.year)
}

func (r Car) AsTuple() fp.Tuple3[string, string, int] {
	return as.Tuple3(r.company, r.model, r.year)
}

func (r Car) Unapply() (string, string, int) {
	return r.company, r.model, r.year
}

func (r Car) AsMutable() CarMutable {
	return CarMutable{
		Company: r.company,
		Model:   r.model,
		Year:    r.year,
	}
}

func (r CarMutable) AsImmutable() Car {
	return Car{
		company: r.Company,
		model:   r.Model,
		year:    r.Year,
	}
}

func (r Car) AsMap() map[string]any {
	m := map[string]any{}
	m["company"] = r.company
	m["model"] = r.model
	m["year"] = r.year
	return m
}

func (r Car) AsLabelled() fp.Labelled3[NamedCompany[string], NamedModel[string], NamedYear[int]] {
	return as.Labelled3(NamedCompany[string]{r.company}, NamedModel[string]{r.model}, NamedYear[int]{r.year})
}

type EntryMutable[A comparable, B any] struct {
	Key   A
	Value B
}

func (r Entry[A, B]) Key() A {
	return r.key
}

func (r Entry[A, B]) Value() B {
	return r.value
}

func (r Entry[A, B]) WithKey(v A) Entry[A, B] {
	r.key = v
	return r
}

func (r Entry[A, B]) WithValue(v B) Entry[A, B] {
	r.value = v
	return r
}

type EntryBuilder[A comparable, B any] Entry[A, B]

func (r EntryBuilder[A, B]) Build() Entry[A, B] {
	return Entry[A, B](r)
}

func (r Entry[A, B]) Builder() EntryBuilder[A, B] {
	return EntryBuilder[A, B](r)
}

func (r EntryBuilder[A, B]) Key(v A) EntryBuilder[A, B] {
	r.key = v
	return r
}

func (r EntryBuilder[A, B]) Value(v B) EntryBuilder[A, B] {
	r.value = v
	return r
}

func (r EntryBuilder[A, B]) FromTuple(t fp.Tuple2[A, B]) EntryBuilder[A, B] {
	r.key = t.I1
	r.value = t.I2
	return r
}

func (r EntryBuilder[A, B]) Apply(key A, value B) EntryBuilder[A, B] {
	r.key = key
	r.value = value
	return r
}

func (r EntryBuilder[A, B]) FromMap(m map[string]any) EntryBuilder[A, B] {

	if v, ok := m["key"].(A); ok {
		r.key = v
	}

	if v, ok := m["value"].(B); ok {
		r.value = v
	}

	return r
}

func (r Entry[A, B]) String() string {
	return fmt.Sprintf("Entry(key=%v, value=%v)", r.key, r.value)
}

func (r Entry[A, B]) AsTuple() fp.Tuple2[A, B] {
	return as.Tuple2(r.key, r.value)
}

func (r Entry[A, B]) Unapply() (A, B) {
	return r.key, r.value
}

func (r Entry[A, B]) AsMutable() EntryMutable[A, B] {
	return EntryMutable[A, B]{
		Key:   r.key,
		Value: r.value,
	}
}

func (r EntryMutable[A, B]) AsImmutable() Entry[A, B] {
	return Entry[A, B]{
		key:   r.Key,
		value: r.Value,
	}
}

func (r Entry[A, B]) AsMap() map[string]any {
	m := map[string]any{}
	m["key"] = r.key
	m["value"] = r.value
	return m
}

type CarsOwnedMutable struct {
	Owner Person
	Cars  fp.Seq[Car]
}

func (r CarsOwned) Owner() Person {
	return r.owner
}

func (r CarsOwned) Cars() fp.Seq[Car] {
	return r.cars
}

func (r CarsOwned) WithOwner(v Person) CarsOwned {
	r.owner = v
	return r
}

func (r CarsOwned) WithCars(v fp.Seq[Car]) CarsOwned {
	r.cars = v
	return r
}

type CarsOwnedBuilder CarsOwned

func (r CarsOwnedBuilder) Build() CarsOwned {
	return CarsOwned(r)
}

func (r CarsOwned) Builder() CarsOwnedBuilder {
	return CarsOwnedBuilder(r)
}

func (r CarsOwnedBuilder) Owner(v Person) CarsOwnedBuilder {
	r.owner = v
	return r
}

func (r CarsOwnedBuilder) Cars(v fp.Seq[Car]) CarsOwnedBuilder {
	r.cars = v
	return r
}

func (r CarsOwnedBuilder) FromTuple(t fp.Tuple2[Person, fp.Seq[Car]]) CarsOwnedBuilder {
	r.owner = t.I1
	r.cars = t.I2
	return r
}

func (r CarsOwnedBuilder) Apply(owner Person, cars fp.Seq[Car]) CarsOwnedBuilder {
	r.owner = owner
	r.cars = cars
	return r
}

func (r CarsOwnedBuilder) FromMap(m map[string]any) CarsOwnedBuilder {

	if v, ok := m["owner"].(Person); ok {
		r.owner = v
	}

	if v, ok := m["cars"].(fp.Seq[Car]); ok {
		r.cars = v
	}

	return r
}

func (r CarsOwned) String() string {
	return fmt.Sprintf("CarsOwned(owner=%v, cars=%v)", r.owner, r.cars)
}

func (r CarsOwned) AsTuple() fp.Tuple2[Person, fp.Seq[Car]] {
	return as.Tuple2(r.owner, r.cars)
}

func (r CarsOwned) Unapply() (Person, fp.Seq[Car]) {
	return r.owner, r.cars
}

func (r CarsOwned) AsMutable() CarsOwnedMutable {
	return CarsOwnedMutable{
		Owner: r.owner,
		Cars:  r.cars,
	}
}

func (r CarsOwnedMutable) AsImmutable() CarsOwned {
	return CarsOwned{
		owner: r.Owner,
		cars:  r.Cars,
	}
}

func (r CarsOwned) AsMap() map[string]any {
	m := map[string]any{}
	m["owner"] = r.owner
	m["cars"] = r.cars
	return m
}

type UserMutable struct {
	Name   string
	Email  fp.Option[string]
	Active bool
}

func (r User) Name() string {
	return r.name
}

func (r User) Email() fp.Option[string] {
	return r.email
}

func (r User) Active() bool {
	return r.active
}

func (r User) WithName(v string) User {
	r.name = v
	return r
}

func (r User) WithEmail(v fp.Option[string]) User {
	r.email = v
	return r
}

func (r User) WithSomeEmail(v string) User {
	r.email = option.Some(v)
	return r
}

func (r User) WithNoneEmail() User {
	r.email = option.None[string]()
	return r
}

func (r User) WithActive(v bool) User {
	r.active = v
	return r
}

type UserBuilder User

func (r UserBuilder) Build() User {
	return User(r)
}

func (r User) Builder() UserBuilder {
	return UserBuilder(r)
}

func (r UserBuilder) Name(v string) UserBuilder {
	r.name = v
	return r
}

func (r UserBuilder) Email(v fp.Option[string]) UserBuilder {
	r.email = v
	return r
}

func (r UserBuilder) SomeEmail(v string) UserBuilder {
	r.email = option.Some(v)
	return r
}

func (r UserBuilder) NoneEmail() UserBuilder {
	r.email = option.None[string]()
	return r
}

func (r UserBuilder) Active(v bool) UserBuilder {
	r.active = v
	return r
}

func (r UserBuilder) FromTuple(t fp.Tuple3[string, fp.Option[string], bool]) UserBuilder {
	r.name = t.I1
	r.email = t.I2
	r.active = t.I3
	return r
}

func (r UserBuilder) Apply(name string, email fp.Option[string], active bool) UserBuilder {
	r.name = name
	r.email = email
	r.active = active
	return r
}

func (r UserBuilder) FromMap(m map[string]any) UserBuilder {

	if v, ok := m["name"].(string); ok {
		r.name = v
	}

	if v, ok := m["email"].(fp.Option[string]); ok {
		r.email = v
	} else if v, ok := m["email"].(string); ok {
		r.email = option.Some(v)
	}

	if v, ok := m["active"].(bool); ok {
		r.active = v
	}

	return r
}

func (r User) String() string {
	return fmt.Sprintf("User(name=%v, email=%v, active=%v)", r.name, r.email, r.active)
}

func (r User) AsTuple() fp.Tuple3[string, fp.Option[string], bool] {
	return as.Tuple3(r.name, r.email, r.active)
}

func (r User) Unapply() (string, fp.Option[string], bool) {
	return r.name, r.email, r.active
}

func (r User) AsMutable() UserMutable {
	return UserMutable{
		Name:   r.name,
		Email:  r.email,
		Active: r.active,
	}
}

func (r UserMutable) AsImmutable() User {
	return User{
		name:   r.Name,
		email:  r.Email,
		active: r.Active,
	}
}

func (r User) AsMap() map[string]any {
	m := map[string]any{}
	m["name"] = r.name
	if r.email.IsDefined() {
		m["email"] = r.email.Get()
	}
	m["active"] = r.active
	return m
}

func (r MapEntry[K, V]) Deref() fp.Tuple2[K, V] {
	return fp.Tuple2[K, V](r)
}

func IntoMapEntry[K any, V any](v fp.Tuple2[K, V]) MapEntry[K, V] {
	return MapEntry[K, V](v)
}

func (r MapEntry[K, V]) Head() K {
	return fp.Tuple2[K, V](r).Head()
}

func (r MapEntry[K, V]) Init() K {
	return fp.Tuple2[K, V](r).Init()
}

func (r MapEntry[K, V]) Last() V {
	return fp.Tuple2[K, V](r).Last()
}

func (r MapEntry[K, V]) String() string {
	return fp.Tuple2[K, V](r).String()
}

func (r MapEntry[K, V]) Tail() V {
	return fp.Tuple2[K, V](r).Tail()
}

func (r MapEntry[K, V]) Unapply() (K, V) {
	return fp.Tuple2[K, V](r).Unapply()
}

func (r OptionalInt) Deref() fp.Option[int] {
	return fp.Option[int](r)
}

func IntoOptionalInt(v fp.Option[int]) OptionalInt {
	return OptionalInt(v)
}

func (r OptionalInt) Exists(p func(v int) bool) bool {
	return fp.Option[int](r).Exists(p)
}

func (r OptionalInt) Filter(p func(v int) bool) fp.Option[int] {
	return fp.Option[int](r).Filter(p)
}

func (r OptionalInt) FilterNot(p func(v int) bool) fp.Option[int] {
	return fp.Option[int](r).FilterNot(p)
}

func (r OptionalInt) FlatMap(mf func(int) fp.Option[int]) fp.Option[int] {
	return fp.Option[int](r).FlatMap(mf)
}

func (r OptionalInt) ForAll(p func(v int) bool) bool {
	return fp.Option[int](r).ForAll(p)
}

func (r OptionalInt) Foreach(f func(v int)) {
	fp.Option[int](r).Foreach(f)
}

func (r OptionalInt) Get() int {
	return fp.Option[int](r).Get()
}

func (r OptionalInt) IsDefined() bool {
	return fp.Option[int](r).IsDefined()
}

func (r OptionalInt) IsEmpty() bool {
	return fp.Option[int](r).IsEmpty()
}

func (r OptionalInt) Map(mf func(int) int) fp.Option[int] {
	return fp.Option[int](r).Map(mf)
}

func (r OptionalInt) MarshalJSON() ([]byte, error) {
	return fp.Option[int](r).MarshalJSON()
}

func (r OptionalInt) Or(f func() fp.Option[int]) fp.Option[int] {
	return fp.Option[int](r).Or(f)
}

func (r OptionalInt) OrElse(t int) int {
	return fp.Option[int](r).OrElse(t)
}

func (r OptionalInt) OrElseGet(f func() int) int {
	return fp.Option[int](r).OrElseGet(f)
}

func (r OptionalInt) OrOption(v fp.Option[int]) fp.Option[int] {
	return fp.Option[int](r).OrOption(v)
}

func (r OptionalInt) OrPtr(v *int) fp.Option[int] {
	return fp.Option[int](r).OrPtr(v)
}

func (r OptionalInt) OrZero() int {
	return fp.Option[int](r).OrZero()
}

func (r OptionalInt) Ptr() *int {
	return fp.Option[int](r).Ptr()
}

func (r OptionalInt) Recover(f func() int) fp.Option[int] {
	return fp.Option[int](r).Recover(f)
}

func (r OptionalInt) String() string {
	return fp.Option[int](r).String()
}

func (r OptionalInt) ToSeq() []int {
	return fp.Option[int](r).ToSeq()
}

func (r OptionalInt) Unapply() (int, bool) {
	return fp.Option[int](r).Unapply()
}

func (r *OptionalInt) UnmarshalJSON(b []byte) error {
	return (*fp.Option[int])(r).UnmarshalJSON(b)
}

func (r OptionalStringer[T]) Deref() fp.Option[T] {
	return fp.Option[T](r)
}

func IntoOptionalStringer[T fmt.Stringer](v fp.Option[T]) OptionalStringer[T] {
	return OptionalStringer[T](v)
}

func (r OptionalStringer[T]) Exists(p func(v T) bool) bool {
	return fp.Option[T](r).Exists(p)
}

func (r OptionalStringer[T]) Filter(p func(v T) bool) fp.Option[T] {
	return fp.Option[T](r).Filter(p)
}

func (r OptionalStringer[T]) FilterNot(p func(v T) bool) fp.Option[T] {
	return fp.Option[T](r).FilterNot(p)
}

func (r OptionalStringer[T]) FlatMap(mf func(T) fp.Option[T]) fp.Option[T] {
	return fp.Option[T](r).FlatMap(mf)
}

func (r OptionalStringer[T]) ForAll(p func(v T) bool) bool {
	return fp.Option[T](r).ForAll(p)
}

func (r OptionalStringer[T]) Foreach(f func(v T)) {
	fp.Option[T](r).Foreach(f)
}

func (r OptionalStringer[T]) Get() T {
	return fp.Option[T](r).Get()
}

func (r OptionalStringer[T]) IsDefined() bool {
	return fp.Option[T](r).IsDefined()
}

func (r OptionalStringer[T]) IsEmpty() bool {
	return fp.Option[T](r).IsEmpty()
}

func (r OptionalStringer[T]) Map(mf func(T) T) fp.Option[T] {
	return fp.Option[T](r).Map(mf)
}

func (r OptionalStringer[T]) MarshalJSON() ([]byte, error) {
	return fp.Option[T](r).MarshalJSON()
}

func (r OptionalStringer[T]) Or(f func() fp.Option[T]) fp.Option[T] {
	return fp.Option[T](r).Or(f)
}

func (r OptionalStringer[T]) OrElse(t T) T {
	return fp.Option[T](r).OrElse(t)
}

func (r OptionalStringer[T]) OrElseGet(f func() T) T {
	return fp.Option[T](r).OrElseGet(f)
}

func (r OptionalStringer[T]) OrOption(v fp.Option[T]) fp.Option[T] {
	return fp.Option[T](r).OrOption(v)
}

func (r OptionalStringer[T]) OrPtr(v *T) fp.Option[T] {
	return fp.Option[T](r).OrPtr(v)
}

func (r OptionalStringer[T]) OrZero() T {
	return fp.Option[T](r).OrZero()
}

func (r OptionalStringer[T]) Ptr() *T {
	return fp.Option[T](r).Ptr()
}

func (r OptionalStringer[T]) Recover(f func() T) fp.Option[T] {
	return fp.Option[T](r).Recover(f)
}

func (r OptionalStringer[T]) String() string {
	return fp.Option[T](r).String()
}

func (r OptionalStringer[T]) ToSeq() []T {
	return fp.Option[T](r).ToSeq()
}

func (r OptionalStringer[T]) Unapply() (T, bool) {
	return fp.Option[T](r).Unapply()
}

func (r *OptionalStringer[T]) UnmarshalJSON(b []byte) error {
	return (*fp.Option[T])(r).UnmarshalJSON(b)
}

type NamedCompany[T any] fp.Tuple1[T]

func (r NamedCompany[T]) Name() string {
	return "company"
}
func (r NamedCompany[T]) Value() T {
	return r.I1
}
func (r NamedCompany[T]) WithValue(v T) NamedCompany[T] {
	r.I1 = v
	return r
}

type NamedModel[T any] fp.Tuple1[T]

func (r NamedModel[T]) Name() string {
	return "model"
}
func (r NamedModel[T]) Value() T {
	return r.I1
}
func (r NamedModel[T]) WithValue(v T) NamedModel[T] {
	r.I1 = v
	return r
}

type NamedYear[T any] fp.Tuple1[T]

func (r NamedYear[T]) Name() string {
	return "year"
}
func (r NamedYear[T]) Value() T {
	return r.I1
}
func (r NamedYear[T]) WithValue(v T) NamedYear[T] {
	r.I1 = v
	return r
}
