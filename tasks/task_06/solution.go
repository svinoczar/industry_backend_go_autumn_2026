package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRUCache[K comparable, V any] struct {
	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element, max(0, capacity)),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	elem, exists := c.items[key]
	if !exists {
		return value, false
	}

	c.ll.MoveToFront(elem)

	return elem.Value.(entry[K, V]).value, true
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}

	if elem, exists := c.items[key]; exists {
		item := elem.Value.(entry[K, V])
		item.value = value
		elem.Value = item
		return
	}

	if c.ll.Len() >= c.capacity {
		oldest := c.ll.Back()
		if oldest != nil {
			item := oldest.Value.(entry[K, V])
			delete(c.items, item.key)
			c.ll.Remove(oldest)
		}
	}

	elem := c.ll.PushFront(entry[K, V]{
		key:   key,
		value: value,
	})
	c.items[key] = elem
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
