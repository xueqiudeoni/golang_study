package lru

import (
	"cache"
	"container/list"
)

type LRUCache struct {
	maxBytes  int
	usedBytes int
	onEvicted func(key string, value interface{})
	ll        *list.List
	cache     map[string]*list.Element
}
type entry struct {
	key   string
	value interface{}
}

func (e *entry) Len() int {
	return cache.CalcLen(e.value)
}
func NewLruCache(maxBytes int, onEvicted func(key string, value interface{})) cache.Cache {
	return &LRUCache{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
	}
}
func (c *LRUCache) Set(k string, v interface{}) {
	if e, ok := c.cache[k]; ok {
		c.ll.MoveToBack(e)
		en := e.Value.(*entry)
		c.usedBytes = c.usedBytes - cache.CalcLen(en.value) + cache.CalcLen(v)
		en.value = v
		return
	}
	n := &entry{k, v}
	e := c.ll.PushBack(n)
	c.cache[k] = e
	c.usedBytes += cache.CalcLen(v)
	if c.maxBytes > 0 && c.usedBytes > c.maxBytes {
		c.DelOldest()
	}
}
func (c *LRUCache) Get(key string) interface{} {
	if e, ok := c.cache[key]; ok {
		c.ll.MoveToBack(e)
		return e.Value.(*entry).value
	}
	return nil
}
func (c *LRUCache) Del(k string) {
	if e, ok := c.cache[k]; ok {
		c.ll.Remove(e)
	}
}
func (c *LRUCache) DelOldest() {
	c.removeElement(c.ll.Front())
}
func (c *LRUCache) removeElement(e *list.Element) {
	if e == nil {
		return
	}
	c.ll.Remove(c.ll.Front())
	en := e.Value.(*entry)
	c.usedBytes -= en.Len()
	delete(c.cache, en.key)
	if c.onEvicted != nil {
		c.onEvicted(en.key, en.value)
	}
}
func (c *LRUCache) Len() int {
	return c.ll.Len()
}
