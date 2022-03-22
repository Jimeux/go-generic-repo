package cache

import (
	"strconv"
	"testing"
)

func TestCache(t *testing.T) {
	type Tester struct {
		A int
		B string
	}

	t.Run("int cache", func(t *testing.T) {
		testCache(t, 5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	})
	t.Run("string cache", func(t *testing.T) {
		testCache(t, 5, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})
	})
	t.Run("struct cache", func(t *testing.T) {
		testers := make([]Tester, 10)
		for i := range testers {
			testers[i] = Tester{1, "A"}
		}
		testCache(t, 5, testers)
	})
	t.Run("struct pointer cache", func(t *testing.T) {
		testers := make([]*Tester, 10)
		for i := range testers {
			testers[i] = &Tester{1, "A"}
		}
		testCache(t, 5, testers)
	})
}

func testCache[T comparable](t *testing.T, capacity int, values []T) {
	t.Helper()

	cache := NewCache[string, T](capacity)
	for i := 0; i < len(values); i++ {
		key := "key" + strconv.Itoa(i)
		cache.Set(key, values[i])
		result, ok := cache.Get(key)
		if !ok {
			t.Fatalf("expected cache to be set for key%d", i)
		}
		if result != values[i] {
			t.Fatalf("expected %v to be equal to %v", result, values[i])
		}

		if i > capacity {
			_, ok = cache.Get("key" + strconv.Itoa(i-capacity))
			if ok {
				t.Fatalf("expected LRU key%d to be evicted", i-capacity)
			}
		}
	}
}
