package lru

import "container/list"

type Item struct {
	Key string
	Val int
}

type LRU struct {
	ll       *list.List
	capacity int
}

func New(capacity int) *LRU {
	ll := list.New()
	for i := 0; i < capacity; i++ {
		ll.PushBack(0)
	}
	return &LRU{ll: ll, capacity: capacity}
}

func (c *LRU) Add(key string, val int) {
	c.ll.PushFront(Item{Key: key, Val: val})
	for c.ll.Len() > c.capacity {
		c.ll.Remove(c.ll.Back())
	}
}

func (c *LRU) Get(key string) (int, bool) {
	e := c.ll.Front()
	for e != nil {
		if v, ok := e.Value.(Item); ok && v.Key == key {
			c.ll.Remove(e)
			c.ll.PushFront(e.Value)
			return e.Value.(Item).Val, true
		}
		e = e.Next()
	}
	return 0, false
}
