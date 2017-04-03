package v2

import (
	"sync"

	"github.com/sethgrid/lru/v2/linkedlist"
)

type Cache struct {
	mu         *sync.Mutex
	size       int
	linkedList *linkedlist.LinkedList
	m          map[string]*linkedlist.Element
}

func New(size int) *Cache {
	return &Cache{
		mu:         &sync.Mutex{},
		size:       size,
		linkedList: linkedlist.New(),
		m:          make(map[string]*linkedlist.Element),
	}
}

func (c *Cache) Set(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if e, ok := c.m[key]; ok {
		// update the cache by deleting the former record
		delete(c.m, e.Key())
		e.Delete()
	}
	e := c.linkedList.First().InsertNext(key, value)
	c.m[key] = e

	if len(c.m) > c.size {
		e := c.linkedList.Last()
		delete(c.m, e.Key())
		e.Delete()
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.m[key]
	if !ok {
		return "", false
	}
	v := e.Value()

	// push onto front of ll
	c.linkedList.First().InsertNext(key, v)

	// remove from where ever else it is in the list
	e.Delete()

	return v, true
}
