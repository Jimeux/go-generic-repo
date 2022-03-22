package cache

import (
	"sync"
)

// Cache is a simple generic LRU cache.
type Cache[K comparable, V any] struct {
	mu       sync.RWMutex
	cache    map[K]*node[K, V]
	lru      *list[K, V]
	capacity int
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		cache:    make(map[K]*node[K, V], capacity),
		lru:      newList[K, V](),
		capacity: capacity,
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	item, ok := c.cache[key]
	if ok {
		c.lru.Remove(item) // remove to refresh LRU position
		item.value = value // update value
	} else {
		item = newNode(key, value)
	}
	c.lru.Add(item)     // update LRU position
	c.cache[key] = item // overwrite
	if c.lru.Size() > c.capacity {
		// evict LRU element
		evicted := c.lru.Evict()
		delete(c.cache, evicted.key)
	}
	c.mu.Unlock()
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	item, ok := c.cache[key]
	c.mu.RUnlock()
	if !ok {
		var val V
		return val, false
	}
	return item.value, true
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	item, ok := c.cache[key]
	if ok {
		delete(c.cache, key)
		c.lru.Remove(item)
	}
	c.mu.Unlock()
}

func (c *Cache[K, V]) Flush() {
	c.mu.Lock()
	c.cache = make(map[K]*node[K, V], c.capacity)
	c.lru = newList[K, V]()
	c.mu.Unlock()
}
