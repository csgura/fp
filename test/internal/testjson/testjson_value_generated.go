// Code generated by gombok, DO NOT EDIT.
package testjson

import (
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
)

type RootBuilder Root

type RootMutable struct {
	A int
	B string
	C float64
	D bool
	E *int
	F []int
	G map[string]int
	H Child
}

func (r RootBuilder) Build() Root {
	return Root(r)
}

func (r Root) Builder() RootBuilder {
	return RootBuilder(r)
}

func (r Root) A() int {
	return r.a
}

func (r Root) B() string {
	return r.b
}

func (r Root) C() float64 {
	return r.c
}

func (r Root) D() bool {
	return r.d
}

func (r Root) E() *int {
	return r.e
}

func (r Root) F() []int {
	return r.f
}

func (r Root) G() map[string]int {
	return r.g
}

func (r Root) H() Child {
	return r.h
}

func (r Root) WithA(v int) Root {
	r.a = v
	return r
}

func (r RootBuilder) A(v int) RootBuilder {
	r.a = v
	return r
}

func (r Root) WithB(v string) Root {
	r.b = v
	return r
}

func (r RootBuilder) B(v string) RootBuilder {
	r.b = v
	return r
}

func (r Root) WithC(v float64) Root {
	r.c = v
	return r
}

func (r RootBuilder) C(v float64) RootBuilder {
	r.c = v
	return r
}

func (r Root) WithD(v bool) Root {
	r.d = v
	return r
}

func (r RootBuilder) D(v bool) RootBuilder {
	r.d = v
	return r
}

func (r Root) WithE(v *int) Root {
	r.e = v
	return r
}

func (r RootBuilder) E(v *int) RootBuilder {
	r.e = v
	return r
}

func (r Root) WithF(v []int) Root {
	r.f = v
	return r
}

func (r RootBuilder) F(v []int) RootBuilder {
	r.f = v
	return r
}

func (r Root) WithG(v map[string]int) Root {
	r.g = v
	return r
}

func (r RootBuilder) G(v map[string]int) RootBuilder {
	r.g = v
	return r
}

func (r Root) WithH(v Child) Root {
	r.h = v
	return r
}

func (r RootBuilder) H(v Child) RootBuilder {
	r.h = v
	return r
}

func (r Root) String() string {
	return fmt.Sprintf("Root(a=%v, b=%v, c=%v, d=%v, e=%v, f=%v, g=%v, h=%v)", r.a, r.b, r.c, r.d, r.e, r.f, r.g, r.h)
}

func (r Root) AsTuple() fp.Tuple8[int, string, float64, bool, *int, []int, map[string]int, Child] {
	return as.Tuple8(r.a, r.b, r.c, r.d, r.e, r.f, r.g, r.h)
}

func (r Root) Unapply() (int, string, float64, bool, *int, []int, map[string]int, Child) {
	return r.a, r.b, r.c, r.d, r.e, r.f, r.g, r.h
}

func (r Root) AsMutable() RootMutable {
	return RootMutable{
		A: r.a,
		B: r.b,
		C: r.c,
		D: r.d,
		E: r.e,
		F: r.f,
		G: r.g,
		H: r.h,
	}
}

func (r RootMutable) AsImmutable() Root {
	return Root{
		a: r.A,
		b: r.B,
		c: r.C,
		d: r.D,
		e: r.E,
		f: r.F,
		g: r.G,
		h: r.H,
	}
}

func (r RootBuilder) FromTuple(t fp.Tuple8[int, string, float64, bool, *int, []int, map[string]int, Child]) RootBuilder {
	r.a = t.I1
	r.b = t.I2
	r.c = t.I3
	r.d = t.I4
	r.e = t.I5
	r.f = t.I6
	r.g = t.I7
	r.h = t.I8
	return r
}

func (r RootBuilder) Apply(a int, b string, c float64, d bool, e *int, f []int, g map[string]int, h Child) RootBuilder {
	r.a = a
	r.b = b
	r.c = c
	r.d = d
	r.e = e
	r.f = f
	r.g = g
	r.h = h
	return r
}

func (r Root) AsMap() map[string]any {
	m := map[string]any{}
	m["a"] = r.a
	m["b"] = r.b
	m["c"] = r.c
	m["d"] = r.d
	m["e"] = r.e
	m["f"] = r.f
	m["g"] = r.g
	m["h"] = r.h
	return m
}

func (r RootBuilder) FromMap(m map[string]any) RootBuilder {

	if v, ok := m["a"].(int); ok {
		r.a = v
	}

	if v, ok := m["b"].(string); ok {
		r.b = v
	}

	if v, ok := m["c"].(float64); ok {
		r.c = v
	}

	if v, ok := m["d"].(bool); ok {
		r.d = v
	}

	if v, ok := m["e"].(*int); ok {
		r.e = v
	}

	if v, ok := m["f"].([]int); ok {
		r.f = v
	}

	if v, ok := m["g"].(map[string]int); ok {
		r.g = v
	}

	if v, ok := m["h"].(Child); ok {
		r.h = v
	}

	return r
}

func (r Root) AsLabelled() fp.Labelled8[NamedA[int], NamedB[string], NamedC[float64], NamedD[bool], NamedE[*int], NamedF[[]int], NamedG[map[string]int], NamedH[Child]] {
	return as.Labelled8(NamedA[int]{r.a}, NamedB[string]{r.b}, NamedC[float64]{r.c}, NamedD[bool]{r.d}, NamedE[*int]{r.e}, NamedF[[]int]{r.f}, NamedG[map[string]int]{r.g}, NamedH[Child]{r.h})
}

func (r RootBuilder) FromLabelled(t fp.Labelled8[NamedA[int], NamedB[string], NamedC[float64], NamedD[bool], NamedE[*int], NamedF[[]int], NamedG[map[string]int], NamedH[Child]]) RootBuilder {
	r.a = t.I1.Value()
	r.b = t.I2.Value()
	r.c = t.I3.Value()
	r.d = t.I4.Value()
	r.e = t.I5.Value()
	r.f = t.I6.Value()
	r.g = t.I7.Value()
	r.h = t.I8.Value()
	return r
}

type ChildBuilder Child

type ChildMutable struct {
	A map[string]any
	B any
}

func (r ChildBuilder) Build() Child {
	return Child(r)
}

func (r Child) Builder() ChildBuilder {
	return ChildBuilder(r)
}

func (r Child) A() map[string]any {
	return r.a
}

func (r Child) B() any {
	return r.b
}

func (r Child) WithA(v map[string]any) Child {
	r.a = v
	return r
}

func (r ChildBuilder) A(v map[string]any) ChildBuilder {
	r.a = v
	return r
}

func (r Child) WithB(v any) Child {
	r.b = v
	return r
}

func (r ChildBuilder) B(v any) ChildBuilder {
	r.b = v
	return r
}

func (r Child) String() string {
	return fmt.Sprintf("Child(a=%v, b=%v)", r.a, r.b)
}

func (r Child) AsTuple() fp.Tuple2[map[string]any, any] {
	return as.Tuple2(r.a, r.b)
}

func (r Child) Unapply() (map[string]any, any) {
	return r.a, r.b
}

func (r Child) AsMutable() ChildMutable {
	return ChildMutable{
		A: r.a,
		B: r.b,
	}
}

func (r ChildMutable) AsImmutable() Child {
	return Child{
		a: r.A,
		b: r.B,
	}
}

func (r ChildBuilder) FromTuple(t fp.Tuple2[map[string]any, any]) ChildBuilder {
	r.a = t.I1
	r.b = t.I2
	return r
}

func (r ChildBuilder) Apply(a map[string]any, b any) ChildBuilder {
	r.a = a
	r.b = b
	return r
}

func (r Child) AsMap() map[string]any {
	m := map[string]any{}
	m["a"] = r.a
	m["b"] = r.b
	return m
}

func (r ChildBuilder) FromMap(m map[string]any) ChildBuilder {

	if v, ok := m["a"].(map[string]any); ok {
		r.a = v
	}

	if v, ok := m["b"].(any); ok {
		r.b = v
	}

	return r
}

func (r Child) AsLabelled() fp.Labelled2[NamedA[map[string]any], NamedB[any]] {
	return as.Labelled2(NamedA[map[string]any]{r.a}, NamedB[any]{r.b})
}

func (r ChildBuilder) FromLabelled(t fp.Labelled2[NamedA[map[string]any], NamedB[any]]) ChildBuilder {
	r.a = t.I1.Value()
	r.b = t.I2.Value()
	return r
}

type NodeBuilder Node

type NodeMutable struct {
	Name  string
	Left  *Node
	Right *Node
}

func (r NodeBuilder) Build() Node {
	return Node(r)
}

func (r Node) Builder() NodeBuilder {
	return NodeBuilder(r)
}

func (r Node) Name() string {
	return r.name
}

func (r Node) Left() *Node {
	return r.left
}

func (r Node) Right() *Node {
	return r.right
}

func (r Node) WithName(v string) Node {
	r.name = v
	return r
}

func (r NodeBuilder) Name(v string) NodeBuilder {
	r.name = v
	return r
}

func (r Node) WithLeft(v *Node) Node {
	r.left = v
	return r
}

func (r NodeBuilder) Left(v *Node) NodeBuilder {
	r.left = v
	return r
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
	return fmt.Sprintf("Node(name=%v, left=%v, right=%v)", r.name, r.left, r.right)
}

func (r Node) AsTuple() fp.Tuple3[string, *Node, *Node] {
	return as.Tuple3(r.name, r.left, r.right)
}

func (r Node) Unapply() (string, *Node, *Node) {
	return r.name, r.left, r.right
}

func (r Node) AsMutable() NodeMutable {
	return NodeMutable{
		Name:  r.name,
		Left:  r.left,
		Right: r.right,
	}
}

func (r NodeMutable) AsImmutable() Node {
	return Node{
		name:  r.Name,
		left:  r.Left,
		right: r.Right,
	}
}

func (r NodeBuilder) FromTuple(t fp.Tuple3[string, *Node, *Node]) NodeBuilder {
	r.name = t.I1
	r.left = t.I2
	r.right = t.I3
	return r
}

func (r NodeBuilder) Apply(name string, left *Node, right *Node) NodeBuilder {
	r.name = name
	r.left = left
	r.right = right
	return r
}

func (r Node) AsMap() map[string]any {
	m := map[string]any{}
	m["name"] = r.name
	m["left"] = r.left
	m["right"] = r.right
	return m
}

func (r NodeBuilder) FromMap(m map[string]any) NodeBuilder {

	if v, ok := m["name"].(string); ok {
		r.name = v
	}

	if v, ok := m["left"].(*Node); ok {
		r.left = v
	}

	if v, ok := m["right"].(*Node); ok {
		r.right = v
	}

	return r
}

func (r Node) AsLabelled() fp.Labelled3[NamedName[string], NamedLeft[*Node], NamedRight[*Node]] {
	return as.Labelled3(NamedName[string]{r.name}, NamedLeft[*Node]{r.left}, NamedRight[*Node]{r.right})
}

func (r NodeBuilder) FromLabelled(t fp.Labelled3[NamedName[string], NamedLeft[*Node], NamedRight[*Node]]) NodeBuilder {
	r.name = t.I1.Value()
	r.left = t.I2.Value()
	r.right = t.I3.Value()
	return r
}

type TreeBuilder Tree

type TreeMutable struct {
	Root *Node
}

func (r TreeBuilder) Build() Tree {
	return Tree(r)
}

func (r Tree) Builder() TreeBuilder {
	return TreeBuilder(r)
}

func (r Tree) Root() *Node {
	return r.root
}

func (r Tree) WithRoot(v *Node) Tree {
	r.root = v
	return r
}

func (r TreeBuilder) Root(v *Node) TreeBuilder {
	r.root = v
	return r
}

func (r Tree) String() string {
	return fmt.Sprintf("Tree(root=%v)", r.root)
}

func (r Tree) AsTuple() fp.Tuple1[*Node] {
	return as.Tuple1(r.root)
}

func (r Tree) Unapply() *Node {
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

func (r TreeBuilder) FromTuple(t fp.Tuple1[*Node]) TreeBuilder {
	r.root = t.I1
	return r
}

func (r TreeBuilder) Apply(root *Node) TreeBuilder {
	r.root = root
	return r
}

func (r Tree) AsMap() map[string]any {
	m := map[string]any{}
	m["root"] = r.root
	return m
}

func (r TreeBuilder) FromMap(m map[string]any) TreeBuilder {

	if v, ok := m["root"].(*Node); ok {
		r.root = v
	}

	return r
}

func (r Tree) AsLabelled() fp.Labelled1[NamedRoot[*Node]] {
	return as.Labelled1(NamedRoot[*Node]{r.root})
}

func (r TreeBuilder) FromLabelled(t fp.Labelled1[NamedRoot[*Node]]) TreeBuilder {
	r.root = t.I1.Value()
	return r
}

type EntryBuilder[V any] Entry[V]

type EntryMutable[V any] struct {
	Name  string
	Value V
}

func (r EntryBuilder[V]) Build() Entry[V] {
	return Entry[V](r)
}

func (r Entry[V]) Builder() EntryBuilder[V] {
	return EntryBuilder[V](r)
}

func (r Entry[V]) Name() string {
	return r.name
}

func (r Entry[V]) Value() V {
	return r.value
}

func (r Entry[V]) WithName(v string) Entry[V] {
	r.name = v
	return r
}

func (r EntryBuilder[V]) Name(v string) EntryBuilder[V] {
	r.name = v
	return r
}

func (r Entry[V]) WithValue(v V) Entry[V] {
	r.value = v
	return r
}

func (r EntryBuilder[V]) Value(v V) EntryBuilder[V] {
	r.value = v
	return r
}

func (r Entry[V]) String() string {
	return fmt.Sprintf("Entry(name=%v, value=%v)", r.name, r.value)
}

func (r Entry[V]) AsTuple() fp.Tuple2[string, V] {
	return as.Tuple2(r.name, r.value)
}

func (r Entry[V]) Unapply() (string, V) {
	return r.name, r.value
}

func (r Entry[V]) AsMutable() EntryMutable[V] {
	return EntryMutable[V]{
		Name:  r.name,
		Value: r.value,
	}
}

func (r EntryMutable[V]) AsImmutable() Entry[V] {
	return Entry[V]{
		name:  r.Name,
		value: r.Value,
	}
}

func (r EntryBuilder[V]) FromTuple(t fp.Tuple2[string, V]) EntryBuilder[V] {
	r.name = t.I1
	r.value = t.I2
	return r
}

func (r EntryBuilder[V]) Apply(name string, value V) EntryBuilder[V] {
	r.name = name
	r.value = value
	return r
}

func (r Entry[V]) AsMap() map[string]any {
	m := map[string]any{}
	m["name"] = r.name
	m["value"] = r.value
	return m
}

func (r EntryBuilder[V]) FromMap(m map[string]any) EntryBuilder[V] {

	if v, ok := m["name"].(string); ok {
		r.name = v
	}

	if v, ok := m["value"].(V); ok {
		r.value = v
	}

	return r
}

func (r Entry[V]) AsLabelled() fp.Labelled2[NamedName[string], NamedValue[V]] {
	return as.Labelled2(NamedName[string]{r.name}, NamedValue[V]{r.value})
}

func (r EntryBuilder[V]) FromLabelled(t fp.Labelled2[NamedName[string], NamedValue[V]]) EntryBuilder[V] {
	r.name = t.I1.Value()
	r.value = t.I2.Value()
	return r
}

type NotUsedParamBuilder[K any, V any] NotUsedParam[K, V]

type NotUsedParamMutable[K any, V any] struct {
	Param string
	Value V
}

func (r NotUsedParamBuilder[K, V]) Build() NotUsedParam[K, V] {
	return NotUsedParam[K, V](r)
}

func (r NotUsedParam[K, V]) Builder() NotUsedParamBuilder[K, V] {
	return NotUsedParamBuilder[K, V](r)
}

func (r NotUsedParam[K, V]) Param() string {
	return r.param
}

func (r NotUsedParam[K, V]) Value() V {
	return r.value
}

func (r NotUsedParam[K, V]) WithParam(v string) NotUsedParam[K, V] {
	r.param = v
	return r
}

func (r NotUsedParamBuilder[K, V]) Param(v string) NotUsedParamBuilder[K, V] {
	r.param = v
	return r
}

func (r NotUsedParam[K, V]) WithValue(v V) NotUsedParam[K, V] {
	r.value = v
	return r
}

func (r NotUsedParamBuilder[K, V]) Value(v V) NotUsedParamBuilder[K, V] {
	r.value = v
	return r
}

func (r NotUsedParam[K, V]) String() string {
	return fmt.Sprintf("NotUsedParam(param=%v, value=%v)", r.param, r.value)
}

func (r NotUsedParam[K, V]) AsTuple() fp.Tuple2[string, V] {
	return as.Tuple2(r.param, r.value)
}

func (r NotUsedParam[K, V]) Unapply() (string, V) {
	return r.param, r.value
}

func (r NotUsedParam[K, V]) AsMutable() NotUsedParamMutable[K, V] {
	return NotUsedParamMutable[K, V]{
		Param: r.param,
		Value: r.value,
	}
}

func (r NotUsedParamMutable[K, V]) AsImmutable() NotUsedParam[K, V] {
	return NotUsedParam[K, V]{
		param: r.Param,
		value: r.Value,
	}
}

func (r NotUsedParamBuilder[K, V]) FromTuple(t fp.Tuple2[string, V]) NotUsedParamBuilder[K, V] {
	r.param = t.I1
	r.value = t.I2
	return r
}

func (r NotUsedParamBuilder[K, V]) Apply(param string, value V) NotUsedParamBuilder[K, V] {
	r.param = param
	r.value = value
	return r
}

func (r NotUsedParam[K, V]) AsMap() map[string]any {
	m := map[string]any{}
	m["param"] = r.param
	m["value"] = r.value
	return m
}

func (r NotUsedParamBuilder[K, V]) FromMap(m map[string]any) NotUsedParamBuilder[K, V] {

	if v, ok := m["param"].(string); ok {
		r.param = v
	}

	if v, ok := m["value"].(V); ok {
		r.value = v
	}

	return r
}

func (r NotUsedParam[K, V]) AsLabelled() fp.Labelled2[NamedParam[string], NamedValue[V]] {
	return as.Labelled2(NamedParam[string]{r.param}, NamedValue[V]{r.value})
}

func (r NotUsedParamBuilder[K, V]) FromLabelled(t fp.Labelled2[NamedParam[string], NamedValue[V]]) NotUsedParamBuilder[K, V] {
	r.param = t.I1.Value()
	r.value = t.I2.Value()
	return r
}

type MovieBuilder Movie

type MovieMutable struct {
	Name    string
	Casting Entry[string]
	NotUsed NotUsedParam[int, string]
}

func (r MovieBuilder) Build() Movie {
	return Movie(r)
}

func (r Movie) Builder() MovieBuilder {
	return MovieBuilder(r)
}

func (r Movie) Name() string {
	return r.name
}

func (r Movie) Casting() Entry[string] {
	return r.casting
}

func (r Movie) NotUsed() NotUsedParam[int, string] {
	return r.notUsed
}

func (r Movie) WithName(v string) Movie {
	r.name = v
	return r
}

func (r MovieBuilder) Name(v string) MovieBuilder {
	r.name = v
	return r
}

func (r Movie) WithCasting(v Entry[string]) Movie {
	r.casting = v
	return r
}

func (r MovieBuilder) Casting(v Entry[string]) MovieBuilder {
	r.casting = v
	return r
}

func (r Movie) WithNotUsed(v NotUsedParam[int, string]) Movie {
	r.notUsed = v
	return r
}

func (r MovieBuilder) NotUsed(v NotUsedParam[int, string]) MovieBuilder {
	r.notUsed = v
	return r
}

func (r Movie) String() string {
	return fmt.Sprintf("Movie(name=%v, casting=%v, notUsed=%v)", r.name, r.casting, r.notUsed)
}

func (r Movie) AsTuple() fp.Tuple3[string, Entry[string], NotUsedParam[int, string]] {
	return as.Tuple3(r.name, r.casting, r.notUsed)
}

func (r Movie) Unapply() (string, Entry[string], NotUsedParam[int, string]) {
	return r.name, r.casting, r.notUsed
}

func (r Movie) AsMutable() MovieMutable {
	return MovieMutable{
		Name:    r.name,
		Casting: r.casting,
		NotUsed: r.notUsed,
	}
}

func (r MovieMutable) AsImmutable() Movie {
	return Movie{
		name:    r.Name,
		casting: r.Casting,
		notUsed: r.NotUsed,
	}
}

func (r MovieBuilder) FromTuple(t fp.Tuple3[string, Entry[string], NotUsedParam[int, string]]) MovieBuilder {
	r.name = t.I1
	r.casting = t.I2
	r.notUsed = t.I3
	return r
}

func (r MovieBuilder) Apply(name string, casting Entry[string], notUsed NotUsedParam[int, string]) MovieBuilder {
	r.name = name
	r.casting = casting
	r.notUsed = notUsed
	return r
}

func (r Movie) AsMap() map[string]any {
	m := map[string]any{}
	m["name"] = r.name
	m["casting"] = r.casting
	m["notUsed"] = r.notUsed
	return m
}

func (r MovieBuilder) FromMap(m map[string]any) MovieBuilder {

	if v, ok := m["name"].(string); ok {
		r.name = v
	}

	if v, ok := m["casting"].(Entry[string]); ok {
		r.casting = v
	}

	if v, ok := m["notUsed"].(NotUsedParam[int, string]); ok {
		r.notUsed = v
	}

	return r
}

func (r Movie) AsLabelled() fp.Labelled3[NamedName[string], NamedCasting[Entry[string]], NamedNotUsed[NotUsedParam[int, string]]] {
	return as.Labelled3(NamedName[string]{r.name}, NamedCasting[Entry[string]]{r.casting}, NamedNotUsed[NotUsedParam[int, string]]{r.notUsed})
}

func (r MovieBuilder) FromLabelled(t fp.Labelled3[NamedName[string], NamedCasting[Entry[string]], NamedNotUsed[NotUsedParam[int, string]]]) MovieBuilder {
	r.name = t.I1.Value()
	r.casting = t.I2.Value()
	r.notUsed = t.I3.Value()
	return r
}

type NoPrivateBuilder NoPrivate

type NoPrivateMutable struct {
	Root string
}

func (r NoPrivateBuilder) Build() NoPrivate {
	return NoPrivate(r)
}

func (r NoPrivate) Builder() NoPrivateBuilder {
	return NoPrivateBuilder(r)
}

func (r NoPrivate) String() string {
	return fmt.Sprintf("NoPrivate(Root=%v)", r.Root)
}

func (r NoPrivate) AsTuple() fp.Tuple1[string] {
	return as.Tuple1(r.Root)
}

func (r NoPrivate) Unapply() string {
	return r.Root
}

func (r NoPrivate) AsMutable() NoPrivateMutable {
	return NoPrivateMutable{
		Root: r.Root,
	}
}

func (r NoPrivateMutable) AsImmutable() NoPrivate {
	return NoPrivate{
		Root: r.Root,
	}
}

func (r NoPrivateBuilder) FromTuple(t fp.Tuple1[string]) NoPrivateBuilder {
	r.Root = t.I1
	return r
}

func (r NoPrivateBuilder) Apply(Root string) NoPrivateBuilder {
	r.Root = Root
	return r
}

func (r NoPrivate) AsMap() map[string]any {
	m := map[string]any{}
	m["Root"] = r.Root
	return m
}

func (r NoPrivateBuilder) FromMap(m map[string]any) NoPrivateBuilder {

	if v, ok := m["Root"].(string); ok {
		r.Root = v
	}

	return r
}

func (r NoPrivate) AsLabelled() fp.Labelled1[PubNamedRoot[string]] {
	return as.Labelled1(PubNamedRoot[string]{r.Root})
}

func (r NoPrivateBuilder) FromLabelled(t fp.Labelled1[PubNamedRoot[string]]) NoPrivateBuilder {
	r.Root = t.I1.Value()
	return r
}

type PubNamedRoot[T any] fp.Tuple1[T]

func (r PubNamedRoot[T]) Name() string {
	return "Root"
}
func (r PubNamedRoot[T]) Value() T {
	return r.I1
}
func (r PubNamedRoot[T]) WithValue(v T) PubNamedRoot[T] {
	r.I1 = v
	return r
}

type NamedA[T any] fp.Tuple1[T]

func (r NamedA[T]) Name() string {
	return "a"
}
func (r NamedA[T]) Value() T {
	return r.I1
}
func (r NamedA[T]) WithValue(v T) NamedA[T] {
	r.I1 = v
	return r
}

type NamedB[T any] fp.Tuple1[T]

func (r NamedB[T]) Name() string {
	return "b"
}
func (r NamedB[T]) Value() T {
	return r.I1
}
func (r NamedB[T]) WithValue(v T) NamedB[T] {
	r.I1 = v
	return r
}

type NamedC[T any] fp.Tuple1[T]

func (r NamedC[T]) Name() string {
	return "c"
}
func (r NamedC[T]) Value() T {
	return r.I1
}
func (r NamedC[T]) WithValue(v T) NamedC[T] {
	r.I1 = v
	return r
}

type NamedCasting[T any] fp.Tuple1[T]

func (r NamedCasting[T]) Name() string {
	return "casting"
}
func (r NamedCasting[T]) Value() T {
	return r.I1
}
func (r NamedCasting[T]) WithValue(v T) NamedCasting[T] {
	r.I1 = v
	return r
}

type NamedD[T any] fp.Tuple1[T]

func (r NamedD[T]) Name() string {
	return "d"
}
func (r NamedD[T]) Value() T {
	return r.I1
}
func (r NamedD[T]) WithValue(v T) NamedD[T] {
	r.I1 = v
	return r
}

type NamedE[T any] fp.Tuple1[T]

func (r NamedE[T]) Name() string {
	return "e"
}
func (r NamedE[T]) Value() T {
	return r.I1
}
func (r NamedE[T]) WithValue(v T) NamedE[T] {
	r.I1 = v
	return r
}

type NamedF[T any] fp.Tuple1[T]

func (r NamedF[T]) Name() string {
	return "f"
}
func (r NamedF[T]) Value() T {
	return r.I1
}
func (r NamedF[T]) WithValue(v T) NamedF[T] {
	r.I1 = v
	return r
}

type NamedG[T any] fp.Tuple1[T]

func (r NamedG[T]) Name() string {
	return "g"
}
func (r NamedG[T]) Value() T {
	return r.I1
}
func (r NamedG[T]) WithValue(v T) NamedG[T] {
	r.I1 = v
	return r
}

type NamedH[T any] fp.Tuple1[T]

func (r NamedH[T]) Name() string {
	return "h"
}
func (r NamedH[T]) Value() T {
	return r.I1
}
func (r NamedH[T]) WithValue(v T) NamedH[T] {
	r.I1 = v
	return r
}

type NamedLeft[T any] fp.Tuple1[T]

func (r NamedLeft[T]) Name() string {
	return "left"
}
func (r NamedLeft[T]) Value() T {
	return r.I1
}
func (r NamedLeft[T]) WithValue(v T) NamedLeft[T] {
	r.I1 = v
	return r
}

type NamedName[T any] fp.Tuple1[T]

func (r NamedName[T]) Name() string {
	return "name"
}
func (r NamedName[T]) Value() T {
	return r.I1
}
func (r NamedName[T]) WithValue(v T) NamedName[T] {
	r.I1 = v
	return r
}

type NamedNotUsed[T any] fp.Tuple1[T]

func (r NamedNotUsed[T]) Name() string {
	return "notUsed"
}
func (r NamedNotUsed[T]) Value() T {
	return r.I1
}
func (r NamedNotUsed[T]) WithValue(v T) NamedNotUsed[T] {
	r.I1 = v
	return r
}

type NamedParam[T any] fp.Tuple1[T]

func (r NamedParam[T]) Name() string {
	return "param"
}
func (r NamedParam[T]) Value() T {
	return r.I1
}
func (r NamedParam[T]) WithValue(v T) NamedParam[T] {
	r.I1 = v
	return r
}

type NamedRight[T any] fp.Tuple1[T]

func (r NamedRight[T]) Name() string {
	return "right"
}
func (r NamedRight[T]) Value() T {
	return r.I1
}
func (r NamedRight[T]) WithValue(v T) NamedRight[T] {
	r.I1 = v
	return r
}

type NamedRoot[T any] fp.Tuple1[T]

func (r NamedRoot[T]) Name() string {
	return "root"
}
func (r NamedRoot[T]) Value() T {
	return r.I1
}
func (r NamedRoot[T]) WithValue(v T) NamedRoot[T] {
	r.I1 = v
	return r
}

type NamedValue[T any] fp.Tuple1[T]

func (r NamedValue[T]) Name() string {
	return "value"
}
func (r NamedValue[T]) Value() T {
	return r.I1
}
func (r NamedValue[T]) WithValue(v T) NamedValue[T] {
	r.I1 = v
	return r
}
