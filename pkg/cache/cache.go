package cache

import (
	"container/list"
	"lrucache/pkg/interfaces"
	"sync"
)

type item struct {
	key   string
	value string
}

// LRU ...
type LRU struct {
	sync.RWMutex
	capacity int
	elements map[string]*list.Element
	queue    *list.List
}

// NewLRUCache ...
func NewLRUCache(n int) interfaces.LRUCache {
	return &LRU{
		capacity: n,
		elements: make(map[string]*list.Element),
		queue:    list.New(),
	}
}

// Add ...
func (c *LRU) Add(key, value string) bool {
	defer c.Unlock()
	c.Lock()
	if _, ok := c.elements[key]; ok {
		return false
	}

	if c.queue.Len() == c.capacity {
		elem := c.queue.Back()
		itemRemove := c.queue.Remove(elem).(*item)
		delete(c.elements, itemRemove.key)
	}

	it := &item{key: key, value: value}
	elem := c.queue.PushFront(it)
	c.elements[key] = elem
	return true
}

// Get ...
func (c *LRU) Get(key string) (value string, ok bool) {
	defer c.RUnlock()
	c.RLock()
	el, ok := c.elements[key]
	if !ok {
		return
	}
	c.queue.MoveToFront(el)
	return el.Value.(*item).value, true
}

// Remove ...
func (c *LRU) Remove(key string) (ok bool) {
	defer c.Unlock()
	c.Lock()
	elem, ok := c.elements[key]
	if !ok {
		return ok
	}
	itemRemove := c.queue.Remove(elem).(*item)
	delete(c.elements, itemRemove.key)
	return true
}
