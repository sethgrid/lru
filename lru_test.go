package lru

import "testing"

func TestCapacity(t *testing.T) {
	size := 4
	cache := New(size)

	counter := 0
	cur := cache.ll.Front()
	for {
		if cur == nil {
			break
		}

		if v, ok := cur.Value.(int); ok && v == 0 {
			cur.Value = 1
			counter++
			cur = cur.Next()
			continue
		}
		break
	}
	if counter != size {
		t.Errorf("got %d, want %d for internal linked list size", counter, size)
	}
}

func TestEvict(t *testing.T) {
	size := 3
	cache := New(size)

	cache.Add("one", 1)

	if !isPresentAndEquals(cache, "one", 1) {
		t.Error("want key one to eq 1")
	}

	// add three more values, should evict "one"
	cache.Add("two", 2)
	cache.Add("three", 3)
	cache.Add("four", 4)

	if isPresent(cache, "one") {
		t.Error("want one to be evicted")
	}

	// "two"" is the next to get evicted. Do a Get to bring it back to the front.
	// the following Add should not evict "two", but should evict "three", which is now oldest.
	if !isPresentAndEquals(cache, "two", 2) {
		t.Error("want key two to eq 2")
	}

	cache.Add("five", 5)

	if isPresent(cache, "three") {
		t.Error("want three to be evicted")
	}

	if !isPresentAndEquals(cache, "two", 2) {
		t.Error("want key two to eq 2")
	}

	// finally, make sure that a key we never added is not present
	if isPresent(cache, "six") {
		t.Error("want six to never have existed")
	}
}

func isNotPresent(cache *LRU, key string, value int) bool {
	_, ok := cache.Get(key)
	return !ok
}

func isPresent(cache *LRU, key string) bool {
	_, ok := cache.Get(key)
	return ok
}

func isPresentAndEquals(cache *LRU, key string, value int) bool {
	v, ok := cache.Get(key)
	if !ok {
		return false
	}
	if v != value {
		return false
	}
	return true
}
