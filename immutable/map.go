package immutable

import (
	"fmt"
	"math/bits"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/hash"
)

//lint:file-ignore U1000 source copy from github.com/benbjohnson/immutable/blob/master/immutable.go

// https://github.com/benbjohnson/immutable/blob/master/immutable.go

// Size thresholds for each type of branch node.
const (
	maxArrayMapSize      = 8
	maxBitmapIndexedSize = 16
)

// Segment bit shifts within the map tree.
const (
	mapNodeBits = 5
	mapNodeSize = 1 << mapNodeBits
	mapNodeMask = mapNodeSize - 1
)

// hamt represents an immutable hash map implementation. The map uses a Hasher
// to generate hashes and check for equality of key values.
//
// It is implemented as an Hash Array Mapped Trie.
type hamt[K, V any] struct {
	size   int            // total number of key/value pairs
	root   mapNode[K, V]  // root node of trie
	hasher fp.Hashable[K] // hasher implementation
}

// NewMap returns a new instance of Map. If hasher is nil, a default hasher
// implementation will automatically be chosen based on the first key added.
// Default hasher implementations only exist for int, string, and byte slice types.
func MapBase[K, V any](hasher fp.Hashable[K], t ...fp.Tuple2[K, V]) fp.MapBase[K, V] {
	if len(t) > 0 {
		b := MapBuilder[K, V](hasher)

		for _, v := range t {
			b.Add(v.I1, v.I2)
		}
		return b.build()
	} else {
		return &hamt[K, V]{
			hasher: hasher,
		}
	}
}

func Map[K, V any](hasher fp.Hashable[K], t ...fp.Tuple2[K, V]) fp.Map[K, V] {
	return fp.MakeMap(MapBase(hasher, t...))
}

func MapGiven[K comparable, V any](t ...fp.Tuple2[K, V]) fp.Map[K, V] {
	return fp.MakeMap(MapBase(hash.Given[K](), t...))
}

// Len returns the number of elements in the map.
func (m *hamt[K, V]) Size() int {
	return m.size
}

// clone returns a shallow copy of m.
func (m *hamt[K, V]) clone() *hamt[K, V] {
	other := *m
	return &other
}

// Get returns the value for a given key and a flag indicating whether the
// key exists. This flag distinguishes a nil value set on a key versus a
// non-existent key in the map.
func (m *hamt[K, V]) Get(key K) fp.Option[V] {
	if m.root == nil {
		return fp.None[V]()
	}
	keyHash := m.hasher.Hash(key)
	ret := m.root.get(key, 0, keyHash, m.hasher)
	if ret == nil {
		return fp.None[V]()
	}
	return fp.Some(*ret)
}

// Set returns a map with the key set to the new value. A nil value is allowed.
//
// This function will return a new map even if the updated value is the same as
// the existing value because Map does not track value equality.
func (m *hamt[K, V]) Updated(key K, value V) fp.MapBase[K, V] {
	return m.set(key, value, false)
}

func (m *hamt[K, V]) set(key K, value V, mutable bool) *hamt[K, V] {
	// Set a hasher on the first value if one does not already exist.
	hasher := m.hasher
	// if hasher == nil {
	// 	hasher = NewHasher(key)
	// }

	// Generate copy if necessary.
	other := m
	if !mutable {
		other = m.clone()
	}
	other.hasher = hasher

	// If the map is empty, initialize with a simple array node.
	if m.root == nil {
		other.size = 1
		other.root = &mapArrayNode[K, V]{entries: []mapEntry[K, V]{{key: key, value: value}}}
		return other
	}

	// Otherwise copy the map and delegate insertion to the root.
	// Resized will return true if the key does not currently exist.
	var resized bool
	other.root = m.root.set(key, value, 0, hasher.Hash(key), hasher, mutable, &resized)
	if resized {
		other.size++
	}
	return other
}

// Delete returns a map with the given key removed.
// Removing a non-existent key will cause this method to return the same map.
func (m *hamt[K, V]) Removed(key ...K) fp.MapBase[K, V] {
	ret := m
	for _, k := range key {
		ret = ret.delete(k, false)
	}
	return ret
}

func (m *hamt[K, V]) delete(key K, mutable bool) *hamt[K, V] {
	// Return original map if no keys exist.
	if m.root == nil {
		return m
	}

	// If the delete did not change the node then return the original map.
	var resized bool
	newRoot := m.root.delete(key, 0, m.hasher.Hash(key), m.hasher, mutable, &resized)
	if !resized {
		return m
	}

	// Generate copy if necessary.
	other := m
	if !mutable {
		other = m.clone()
	}

	// Return copy of map with new root and decreased size.
	other.size = m.size - 1
	other.root = newRoot
	return other
}

// Iterator returns a new iterator for the map.
func (m *hamt[K, V]) Iterator() fp.Iterator[fp.Tuple2[K, V]] {
	itr := MapIterator(m)
	return itr
}

func iteratorMap[T, U any](opt fp.Iterator[T], fn func(v T) U) fp.Iterator[U] {
	return fp.MakeIterator(
		func() bool {
			return opt.HasNext()
		},
		func() U {
			return fn(opt.Next())
		},
	)
}

func (m *hamt[K, V]) String() string {
	return fmt.Sprintf("Map(%s)", iteratorMap(m.Iterator(), func(t fp.Tuple2[K, V]) any {
		return fmt.Sprintf("%v: %v", t.I1, t.I2)
	}).MakeString(","))
}

// MapBuilder represents an efficient builder for creating Maps.
type mapBuilder[K, V any] struct {
	m *hamt[K, V] // current state
}

// NewMapBuilder returns a new instance of MapBuilder.
func MapBuilder[K, V any](hasher fp.Hashable[K]) *mapBuilder[K, V] {
	return &mapBuilder[K, V]{m: &hamt[K, V]{
		hasher: hasher,
	}}
}

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func (b *mapBuilder[K, V]) build() fp.MapBase[K, V] {
	assert(b.m != nil, "immutable.SortedMapBuilder.Build(): duplicate call to fetch map")
	m := b.m
	b.m = nil
	return m
}

// Map returns the underlying map. Only call once.
// Builder is invalid after call. Will panic on second invocation.
func (b *mapBuilder[K, V]) Build() fp.Map[K, V] {
	return fp.MakeMap(b.build())
}

// Len returns the number of elements in the underlying map.
// func (b *mapBuilder[K, V]) Size() int {
// 	assert(b.m != nil, "immutable.MapBuilder: builder invalid after Map() invocation")
// 	return b.m.Size()
// }

// Get returns the value for the given key.
// func (b *mapBuilder[K, V]) Get(key K) fp.Option[V] {
// 	assert(b.m != nil, "immutable.MapBuilder: builder invalid after Map() invocation")
// 	return b.m.Get(key)
// }

// Set sets the value of the given key. See Map.Set() for additional details.
func (b *mapBuilder[K, V]) Add(key K, value V) *mapBuilder[K, V] {
	assert(b.m != nil, "immutable.MapBuilder: builder invalid after Build() invocation")
	b.m = b.m.set(key, value, true)
	return b
}

// // Delete removes the given key. See Map.Delete() for additional details.
// func (b *mapBuilder[K, V]) Delete(key K) {
// 	assert(b.m != nil, "immutable.MapBuilder: builder invalid after Map() invocation")
// 	b.m = b.m.delete(key, true)
// }

// // Iterator returns a new iterator for the underlying map.
// func (b *mapBuilder[K, V]) Iterator() fp.Iterator[fp.Tuple2[K, V]] {
// 	assert(b.m != nil, "immutable.MapBuilder: builder invalid after Map() invocation")
// 	return b.m.Iterator()
// }

// mapNode represents any node in the map tree.
type mapNode[K, V any] interface {
	get(key K, shift uint, keyHash uint32, h fp.Hashable[K]) fp.Ptr[V]
	set(key K, value V, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V]
	delete(key K, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V]
}

// var _ mapNode = (*mapArrayNode[K,V])(nil)
// var _ mapNode = (*mapBitmapIndexedNode[K,V])(nil)
// var _ mapNode = (*mapHashArrayNode)(nil)
// var _ mapNode = (*mapValueNode[K,V])(nil)
// var _ mapNode = (*mapHashCollisionNode[K,V])(nil)

// mapLeafNode represents a node that stores a single key hash at the leaf of the map tree.
type mapLeafNode[K, V any] interface {
	mapNode[K, V]
	keyHashValue() uint32
}

// var _ mapLeafNode = (*mapValueNode[K,V])(nil)
// var _ mapLeafNode = (*mapHashCollisionNode[K,V])(nil)

// mapArrayNode is a map node that stores key/value pairs in a slice.
// Entries are stored in insertion order. An array node expands into a bitmap
// indexed node once a given threshold size is crossed.
type mapArrayNode[K, V any] struct {
	entries []mapEntry[K, V]
}

// indexOf returns the entry index of the given key. Returns -1 if key not found.
func (n *mapArrayNode[K, V]) indexOf(key K, h fp.Hashable[K]) int {
	for i := range n.entries {
		if h.Eqv(n.entries[i].key, key) {
			return i
		}
	}
	return -1
}

// get returns the value for the given key.
func (n *mapArrayNode[K, V]) get(key K, shift uint, keyHash uint32, h fp.Hashable[K]) fp.Ptr[V] {
	i := n.indexOf(key, h)
	if i == -1 {
		return nil
	}
	return &n.entries[i].value
}

// set inserts or updates the value for a given key. If the key is inserted and
// the new size crosses the max size threshold, a bitmap indexed node is returned.
func (n *mapArrayNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	idx := n.indexOf(key, h)

	// Mark as resized if the key doesn't exist.
	if idx == -1 {
		*resized = true
	}

	// If we are adding and it crosses the max size threshold, expand the node.
	// We do this by continually setting the entries to a value node and expanding.
	if idx == -1 && len(n.entries) >= maxArrayMapSize {
		var node mapNode[K, V] = newMapValueNode(h.Hash(key), key, value)
		for _, entry := range n.entries {
			node = node.set(entry.key, entry.value, 0, h.Hash(entry.key), h, false, resized)
		}
		return node
	}

	// Update in-place if mutable.
	if mutable {
		if idx != -1 {
			n.entries[idx] = mapEntry[K, V]{key, value}
		} else {
			n.entries = append(n.entries, mapEntry[K, V]{key, value})
		}
		return n
	}

	// Update existing entry if a match is found.
	// Otherwise append to the end of the element list if it doesn't exist.
	var other mapArrayNode[K, V]
	if idx != -1 {
		other.entries = make([]mapEntry[K, V], len(n.entries))
		copy(other.entries, n.entries)
		other.entries[idx] = mapEntry[K, V]{key, value}
	} else {
		other.entries = make([]mapEntry[K, V], len(n.entries)+1)
		copy(other.entries, n.entries)
		other.entries[len(other.entries)-1] = mapEntry[K, V]{key, value}
	}
	return &other
}

// delete removes the given key from the node. Returns the same node if key does
// not exist. Returns a nil node when removing the last entry.
func (n *mapArrayNode[K, V]) delete(key K, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	idx := n.indexOf(key, h)

	// Return original node if key does not exist.
	if idx == -1 {
		return n
	}
	*resized = true

	// Return nil if this node will contain no nodes.
	if len(n.entries) == 1 {
		return nil
	}

	// Update in-place, if mutable.
	if mutable {
		copy(n.entries[idx:], n.entries[idx+1:])
		n.entries[len(n.entries)-1] = mapEntry[K, V]{}
		n.entries = n.entries[:len(n.entries)-1]
		return n
	}

	// Otherwise create a copy with the given entry removed.
	other := &mapArrayNode[K, V]{entries: make([]mapEntry[K, V], len(n.entries)-1)}
	copy(other.entries[:idx], n.entries[:idx])
	copy(other.entries[idx:], n.entries[idx+1:])
	return other
}

// mapBitmapIndexedNode represents a map branch node with a variable number of
// node slots and indexed using a bitmap. Indexes for the node slots are
// calculated by counting the number of set bits before the target bit using popcount.
type mapBitmapIndexedNode[K, V any] struct {
	bitmap uint32
	nodes  []mapNode[K, V]
}

// get returns the value for the given key.
func (n *mapBitmapIndexedNode[K, V]) get(key K, shift uint, keyHash uint32, h fp.Hashable[K]) fp.Ptr[V] {
	bit := uint32(1) << ((keyHash >> shift) & mapNodeMask)
	if (n.bitmap & bit) == 0 {
		return nil
	}
	child := n.nodes[bits.OnesCount32(n.bitmap&(bit-1))]
	return child.get(key, shift+mapNodeBits, keyHash, h)
}

// set inserts or updates the value for the given key. If a new key is inserted
// and the size crosses the max size threshold then a hash array node is returned.
func (n *mapBitmapIndexedNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	// Extract the index for the bit segment of the key hash.
	keyHashFrag := (keyHash >> shift) & mapNodeMask

	// Determine the bit based on the hash index.
	bit := uint32(1) << keyHashFrag
	exists := (n.bitmap & bit) != 0

	// Mark as resized if the key doesn't exist.
	if !exists {
		*resized = true
	}

	// Find index of node based on popcount of bits before it.
	idx := bits.OnesCount32(n.bitmap & (bit - 1))

	// If the node already exists, delegate set operation to it.
	// If the node doesn't exist then create a simple value leaf node.
	var newNode mapNode[K, V]
	if exists {
		newNode = n.nodes[idx].set(key, value, shift+mapNodeBits, keyHash, h, mutable, resized)
	} else {
		newNode = newMapValueNode(keyHash, key, value)
	}

	// Convert to a hash-array node once we exceed the max bitmap size.
	// Copy each node based on their bit position within the bitmap.
	if !exists && len(n.nodes) > maxBitmapIndexedSize {
		var other mapHashArrayNode[K, V]
		for i := uint(0); i < uint(len(other.nodes)); i++ {
			if n.bitmap&(uint32(1)<<i) != 0 {
				other.nodes[i] = n.nodes[other.count]
				other.count++
			}
		}
		other.nodes[keyHashFrag] = newNode
		other.count++
		return &other
	}

	// Update in-place if mutable.
	if mutable {
		if exists {
			n.nodes[idx] = newNode
		} else {
			n.bitmap |= bit
			n.nodes = append(n.nodes, nil)
			copy(n.nodes[idx+1:], n.nodes[idx:])
			n.nodes[idx] = newNode
		}
		return n
	}

	// If node exists at given slot then overwrite it with new node.
	// Otherwise expand the node list and insert new node into appropriate position.
	other := &mapBitmapIndexedNode[K, V]{bitmap: n.bitmap | bit}
	if exists {
		other.nodes = make([]mapNode[K, V], len(n.nodes))
		copy(other.nodes, n.nodes)
		other.nodes[idx] = newNode
	} else {
		other.nodes = make([]mapNode[K, V], len(n.nodes)+1)
		copy(other.nodes, n.nodes[:idx])
		other.nodes[idx] = newNode
		copy(other.nodes[idx+1:], n.nodes[idx:])
	}
	return other
}

// delete removes the key from the tree. If the key does not exist then the
// original node is returned. If removing the last child node then a nil is
// returned. Note that shrinking the node will not convert it to an array node.
func (n *mapBitmapIndexedNode[K, V]) delete(key K, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	bit := uint32(1) << ((keyHash >> shift) & mapNodeMask)

	// Return original node if key does not exist.
	if (n.bitmap & bit) == 0 {
		return n
	}

	// Find index of node based on popcount of bits before it.
	idx := bits.OnesCount32(n.bitmap & (bit - 1))

	// Delegate delete to child node.
	child := n.nodes[idx]
	newChild := child.delete(key, shift+mapNodeBits, keyHash, h, mutable, resized)

	// Return original node if key doesn't exist in child.
	if !*resized {
		return n
	}

	// Remove if returned child has been deleted.
	if newChild == nil {
		// If we won't have any children then return nil.
		if len(n.nodes) == 1 {
			return nil
		}

		// Update in-place if mutable.
		if mutable {
			n.bitmap ^= bit
			copy(n.nodes[idx:], n.nodes[idx+1:])
			n.nodes[len(n.nodes)-1] = nil
			n.nodes = n.nodes[:len(n.nodes)-1]
			return n
		}

		// Return copy with bit removed from bitmap and node removed from node list.
		other := &mapBitmapIndexedNode[K, V]{bitmap: n.bitmap ^ bit, nodes: make([]mapNode[K, V], len(n.nodes)-1)}
		copy(other.nodes[:idx], n.nodes[:idx])
		copy(other.nodes[idx:], n.nodes[idx+1:])
		return other
	}

	// Generate copy, if necessary.
	other := n
	if !mutable {
		other = &mapBitmapIndexedNode[K, V]{bitmap: n.bitmap, nodes: make([]mapNode[K, V], len(n.nodes))}
		copy(other.nodes, n.nodes)
	}

	// Update child.
	other.nodes[idx] = newChild
	return other
}

// mapHashArrayNode is a map branch node that stores nodes in a fixed length
// array. Child nodes are indexed by their index bit segment for the current depth.
type mapHashArrayNode[K, V any] struct {
	count uint                       // number of set nodes
	nodes [mapNodeSize]mapNode[K, V] // child node slots, may contain empties
}

// clone returns a shallow copy of n.
func (n *mapHashArrayNode[K, V]) clone() *mapHashArrayNode[K, V] {
	other := *n
	return &other
}

// get returns the value for the given key.
func (n *mapHashArrayNode[K, V]) get(key K, shift uint, keyHash uint32, h fp.Hashable[K]) fp.Ptr[V] {
	node := n.nodes[(keyHash>>shift)&mapNodeMask]
	if node == nil {
		return nil
	}
	return node.get(key, shift+mapNodeBits, keyHash, h)
}

// set returns a node with the value set for the given key.
func (n *mapHashArrayNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	idx := (keyHash >> shift) & mapNodeMask
	node := n.nodes[idx]

	// If node at index doesn't exist, create a simple value leaf node.
	// Otherwise delegate set to child node.
	var newNode mapNode[K, V]
	if node == nil {
		*resized = true
		newNode = newMapValueNode(keyHash, key, value)
	} else {
		newNode = node.set(key, value, shift+mapNodeBits, keyHash, h, mutable, resized)
	}

	// Generate copy, if necessary.
	other := n
	if !mutable {
		other = n.clone()
	}

	// Update child node (and update size, if new).
	if node == nil {
		other.count++
	}
	other.nodes[idx] = newNode
	return other
}

// delete returns a node with the given key removed. Returns the same node if
// the key does not exist. If node shrinks to within bitmap-indexed size then
// converts to a bitmap-indexed node.
func (n *mapHashArrayNode[K, V]) delete(key K, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	idx := (keyHash >> shift) & mapNodeMask
	node := n.nodes[idx]

	// Return original node if child is not found.
	if node == nil {
		return n
	}

	// Return original node if child is unchanged.
	newNode := node.delete(key, shift+mapNodeBits, keyHash, h, mutable, resized)
	if !*resized {
		return n
	}

	// If we remove a node and drop below a threshold, convert back to bitmap indexed node.
	if newNode == nil && n.count <= maxBitmapIndexedSize {
		other := &mapBitmapIndexedNode[K, V]{nodes: make([]mapNode[K, V], 0, n.count-1)}
		for i, child := range n.nodes {
			if child != nil && uint32(i) != idx {
				other.bitmap |= 1 << uint(i)
				other.nodes = append(other.nodes, child)
			}
		}
		return other
	}

	// Generate copy, if necessary.
	other := n
	if !mutable {
		other = n.clone()
	}

	// Return copy of node with child updated.
	other.nodes[idx] = newNode
	if newNode == nil {
		other.count--
	}
	return other
}

// mapValueNode represents a leaf node with a single key/value pair.
// A value node can be converted to a hash collision leaf node if a different
// key with the same keyHash is inserted.
type mapValueNode[K, V any] struct {
	keyHash uint32
	key     K
	value   V
}

// newMapValueNode returns a new instance of mapValueNode.
func newMapValueNode[K, V any](keyHash uint32, key K, value V) *mapValueNode[K, V] {
	return &mapValueNode[K, V]{
		keyHash: keyHash,
		key:     key,
		value:   value,
	}
}

// keyHashValue returns the key hash for this node.
func (n *mapValueNode[K, V]) keyHashValue() uint32 {
	return n.keyHash
}

// get returns the value for the given key.
func (n *mapValueNode[K, V]) get(key K, shift uint, keyHash uint32, h fp.Hashable[K]) fp.Ptr[V] {
	if !h.Eqv(n.key, key) {
		return nil
	}
	return &n.value
}

// set returns a new node with the new value set for the key. If the key equals
// the node's key then a new value node is returned. If key is not equal to the
// node's key but has the same hash then a hash collision node is returned.
// Otherwise the nodes are merged into a branch node.
func (n *mapValueNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	// If the keys match then return a new value node overwriting the value.
	if h.Eqv(n.key, key) {
		// Update in-place if mutable.
		if mutable {
			n.value = value
			return n
		}
		// Otherwise return a new copy.
		return newMapValueNode(n.keyHash, key, value)
	}

	*resized = true

	// Recursively merge nodes together if key hashes are different.
	if n.keyHash != keyHash {
		return mergeIntoNode(n, shift, keyHash, key, value)
	}

	// Merge into collision node if hash matches.
	return &mapHashCollisionNode[K, V]{keyHash: keyHash, entries: []mapEntry[K, V]{
		{key: n.key, value: n.value},
		{key: key, value: value},
	}}
}

// delete returns nil if the key matches the node's key. Otherwise returns the original node.
func (n *mapValueNode[K, V]) delete(key K, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	// Return original node if the keys do not match.
	if !h.Eqv(n.key, key) {
		return n
	}

	// Otherwise remove the node if keys do match.
	*resized = true
	return nil
}

// mapHashCollisionNode represents a leaf node that contains two or more key/value
// pairs with the same key hash. Single pairs for a hash are stored as value nodes.
type mapHashCollisionNode[K, V any] struct {
	keyHash uint32 // key hash for all entries
	entries []mapEntry[K, V]
}

// keyHashValue returns the key hash for all entries on the node.
func (n *mapHashCollisionNode[K, V]) keyHashValue() uint32 {
	return n.keyHash
}

// indexOf returns the index of the entry for the given key.
// Returns -1 if the key does not exist in the node.
func (n *mapHashCollisionNode[K, V]) indexOf(key K, h fp.Hashable[K]) int {
	for i := range n.entries {
		if h.Eqv(n.entries[i].key, key) {
			return i
		}
	}
	return -1
}

// get returns the value for the given key.
func (n *mapHashCollisionNode[K, V]) get(key K, shift uint, keyHash uint32, h fp.Hashable[K]) fp.Ptr[V] {
	for i := range n.entries {
		if h.Eqv(n.entries[i].key, key) {
			return &n.entries[i].value
		}
	}
	return nil
}

// set returns a copy of the node with key set to the given value.
func (n *mapHashCollisionNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	// Merge node with key/value pair if this is not a hash collision.
	if n.keyHash != keyHash {
		*resized = true
		return mergeIntoNode(n, shift, keyHash, key, value)
	}

	// Update in-place if mutable.
	if mutable {
		if idx := n.indexOf(key, h); idx == -1 {
			*resized = true
			n.entries = append(n.entries, mapEntry[K, V]{key, value})
		} else {
			n.entries[idx] = mapEntry[K, V]{key, value}
		}
		return n
	}

	// Append to end of node if key doesn't exist & mark resized.
	// Otherwise copy nodes and overwrite at matching key index.
	other := &mapHashCollisionNode[K, V]{keyHash: n.keyHash}
	if idx := n.indexOf(key, h); idx == -1 {
		*resized = true
		other.entries = make([]mapEntry[K, V], len(n.entries)+1)
		copy(other.entries, n.entries)
		other.entries[len(other.entries)-1] = mapEntry[K, V]{key, value}
	} else {
		other.entries = make([]mapEntry[K, V], len(n.entries))
		copy(other.entries, n.entries)
		other.entries[idx] = mapEntry[K, V]{key, value}
	}
	return other
}

// delete returns a node with the given key deleted. Returns the same node if
// the key does not exist. If removing the key would shrink the node to a single
// entry then a value node is returned.
func (n *mapHashCollisionNode[K, V]) delete(key K, shift uint, keyHash uint32, h fp.Hashable[K], mutable bool, resized *bool) mapNode[K, V] {
	idx := n.indexOf(key, h)

	// Return original node if key is not found.
	if idx == -1 {
		return n
	}

	// Mark as resized if key exists.
	*resized = true

	// Convert to value node if we move to one entry.
	if len(n.entries) == 2 {
		return &mapValueNode[K, V]{
			keyHash: n.keyHash,
			key:     n.entries[idx^1].key,
			value:   n.entries[idx^1].value,
		}
	}

	// Remove entry in-place if mutable.
	if mutable {
		copy(n.entries[idx:], n.entries[idx+1:])
		n.entries[len(n.entries)-1] = mapEntry[K, V]{}
		n.entries = n.entries[:len(n.entries)-1]
		return n
	}

	// Return copy without entry if immutable.
	other := &mapHashCollisionNode[K, V]{keyHash: n.keyHash, entries: make([]mapEntry[K, V], len(n.entries)-1)}
	copy(other.entries[:idx], n.entries[:idx])
	copy(other.entries[idx:], n.entries[idx+1:])
	return other
}

// mergeIntoNode merges a key/value pair into an existing node.
// Caller must verify that node's keyHash is not equal to keyHash.
func mergeIntoNode[K, V any](node mapLeafNode[K, V], shift uint, keyHash uint32, key K, value V) mapNode[K, V] {
	idx1 := (node.keyHashValue() >> shift) & mapNodeMask
	idx2 := (keyHash >> shift) & mapNodeMask

	// Recursively build branch nodes to combine the node and its key.
	other := &mapBitmapIndexedNode[K, V]{bitmap: (1 << idx1) | (1 << idx2)}
	if idx1 == idx2 {
		other.nodes = []mapNode[K, V]{mergeIntoNode(node, shift+mapNodeBits, keyHash, key, value)}
	} else {
		if newNode := newMapValueNode(keyHash, key, value); idx1 < idx2 {
			other.nodes = []mapNode[K, V]{node, newNode}
		} else {
			other.nodes = []mapNode[K, V]{newNode, node}
		}
	}
	return other
}

// mapEntry represents a single key/value pair.
type mapEntry[K, V any] struct {
	key   K
	value V
}

// MapIterator represents an iterator over a map's key/value pairs. Although
// map keys are not sorted, the iterator's order is deterministic.
func MapIterator[K, V any](m *hamt[K, V]) fp.Iterator[fp.Tuple2[K, V]] {
	// source map

	var stack [32]mapIteratorElem[K, V] // search stack
	var depth int                       // stack depth

	hasNext := func() bool {
		return depth != -1
	}

	first := func() {
		for ; ; depth++ {
			elem := &stack[depth]

			switch node := elem.node.(type) {
			case *mapBitmapIndexedNode[K, V]:
				elem.index = 0
				stack[depth+1].node = node.nodes[0]

			case *mapHashArrayNode[K, V]:
				for i := 0; i < len(node.nodes); i++ {
					if node.nodes[i] != nil { // find first node
						elem.index = i
						stack[depth+1].node = node.nodes[i]
						break
					}
				}

			default: // *mapArrayNode, mapLeafNode
				elem.index = 0
				return
			}
		}
	}

	moveStack := func() {
		for ; depth >= 0; depth-- {
			elem := &stack[depth]

			switch node := elem.node.(type) {
			case *mapArrayNode[K, V]:
				if elem.index < len(node.entries)-1 {
					elem.index++
					return
				}

			case *mapBitmapIndexedNode[K, V]:
				if elem.index < len(node.nodes)-1 {
					elem.index++
					stack[depth+1].node = node.nodes[elem.index]
					depth++
					first()
					return
				}

			case *mapHashArrayNode[K, V]:
				for i := elem.index + 1; i < len(node.nodes); i++ {
					if node.nodes[i] != nil {
						elem.index = i
						stack[depth+1].node = node.nodes[elem.index]
						depth++
						first()
						return
					}
				}

			case *mapValueNode[K, V]:
				continue // always the last value, traverse up

			case *mapHashCollisionNode[K, V]:
				if elem.index < len(node.entries)-1 {
					elem.index++
					return
				}
			}
		}
	}

	next := func() fp.Tuple2[K, V] {
		if !hasNext() {
			panic("next on empty")
		}

		var key K
		var value V
		// Retrieve current index & value. Current node is always a leaf.
		elem := &stack[depth]
		switch node := elem.node.(type) {
		case *mapArrayNode[K, V]:
			entry := &node.entries[elem.index]
			key, value = entry.key, entry.value
		case *mapValueNode[K, V]:
			key, value = node.key, node.value
		case *mapHashCollisionNode[K, V]:
			entry := &node.entries[elem.index]
			key, value = entry.key, entry.value
		}

		// Move up stack until we find a node that has remaining position ahead
		// and move that element forward by one.
		moveStack()
		return as.Tuple(key, value)
	}

	if m.root == nil {
		depth = -1
	} else {

		// Initialize the stack to the left most element.
		stack[0] = mapIteratorElem[K, V]{node: m.root}
		depth = 0
		first()
	}
	return fp.MakeIterator(hasNext, next)
}

// mapIteratorElem represents a node/index pair in the MapIterator stack.
type mapIteratorElem[K, V any] struct {
	node  mapNode[K, V]
	index int
}

type set[T any] struct {
	m fp.MapBase[T, bool]
}

func (r set[T]) Contains(v T) bool {
	return r.m.Get(v).IsDefined()
}
func (r set[T]) Size() int {
	return r.m.Size()
}
func (r set[T]) Iterator() fp.Iterator[T] {
	//return iterator.Map(r.m.Iterator(), fp.Tuple2[T, bool].Head)

	itr := r.m.Iterator()
	return fp.MakeIterator(
		func() bool {
			return itr.HasNext()
		},
		func() T {
			return itr.Next().I1
		},
	)
}
func (r set[T]) Incl(v T) fp.SetMinimal[T] {
	return set[T]{r.m.Updated(v, true)}
}
func (r set[T]) Excl(v T) fp.SetMinimal[T] {
	return set[T]{r.m.Removed(v)}
}

func (r set[T]) String() string {
	return fmt.Sprintf("Set(%s)", r.Iterator().MakeString(","))
}

func SetMinimal[T any](hasher fp.Hashable[T], v ...T) fp.SetMinimal[T] {

	tp := make(fp.Slice[fp.Tuple2[T, bool]], len(v))
	for i, v := range v {
		tp[i] = as.Tuple2(v, true)
	}

	return set[T]{MapBase(hasher, tp...)}
}

func Set[T any](hasher fp.Hashable[T], v ...T) fp.Set[T] {
	return fp.MakeSet(func() fp.SetMinimal[T] {
		return SetMinimal(hasher)
	}, SetMinimal(hasher, v...))
}

type setBuilder[V any] struct {
	m *hamt[V, bool]
}

func (r *setBuilder[V]) Add(v V) *setBuilder[V] {
	r.m.set(v, true, true)
	return r
}

func (r *setBuilder[V]) Build() fp.Set[V] {
	return fp.MakeSet(func() fp.SetMinimal[V] {
		return SetMinimal(r.m.hasher)
	}, set[V]{r.m})
}

func SetBuilder[V any](hasher fp.Hashable[V]) *setBuilder[V] {
	return &setBuilder[V]{m: &hamt[V, bool]{
		hasher: hasher,
	}}
}
