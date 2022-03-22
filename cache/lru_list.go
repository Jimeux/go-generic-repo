package cache

type node[K, V any] struct {
	key      K
	value    V
	previous *node[K, V]
	next     *node[K, V]
}

func newNode[K, V any](key K, value V) *node[K, V] {
	return &node[K, V]{
		key:   key,
		value: value,
	}
}

// list is a doubly-linked list with fixed dummy nodes for the
// head and tail. These allow inserts and LRU evictions, as well
// as arbitrary removals, to be completed in constant time.
//
//       LRU             MRU
//        ↓ 			  ↓    (size==5)
// head ⇄ N ⇄ N ⇄ N ⇄ N ⇄ N ⇄ tail
type list[K, V any] struct {
	head *node[K, V] // head.next points to the LRU node
	tail *node[K, V] // tail.previous points to the MRU node
	size int
}

// newList creates a new list with head and tail pointing to each other.
func newList[K, V any]() *list[K, V] {
	head := &node[K, V]{}
	tail := &node[K, V]{}
	head.next = tail
	tail.previous = head
	return &list[K, V]{
		head: head,
		tail: tail,
	}
}

// Size returns the current number of nodes excluding head and tail.
func (l *list[K, V]) Size() int {
	return l.size
}

// Add inserts n at the end of the list (before tail).
func (l *list[K, V]) Add(n *node[K, V]) {
	prev, next := l.tail.previous, l.tail
	prev.next = n
	next.previous = n
	n.next, n.previous = next, prev
	l.size++
}

// Remove deletes n from the list by removing any reference to it.
func (l *list[K, V]) Remove(n *node[K, V]) {
	prev, next := n.previous, n.next
	prev.next, next.previous = next, prev
	l.size--
}

// Evict removes the LRU node (next to head) from the list and returns it.
func (l *list[K, V]) Evict() *node[K, V] {
	lru := l.head.next
	l.Remove(lru)
	return lru
}
