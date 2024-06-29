package lfu

// 最少使用
import (
	"cache"
	"container/heap"
)

type LFUCache struct {
	maxBytes  int
	onEvicted func(key string, value interface{})
	usedBytes int
	queue     *queue
	cache     map[string]*entry
}

func NewLFUCache(maxBytes int, onEvicted func(key string, value interface{})) cache.Cache {
	return &LFUCache{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		queue:     new(queue),
		cache:     make(map[string]*entry),
	}
}

func (lfu *LFUCache) Set(key string, value interface{}) {
	if e, ok := lfu.cache[key]; ok {
		lfu.usedBytes = lfu.usedBytes - cache.CalcLen(e.value) + cache.CalcLen(value)
		lfu.queue.update(e, value, e.weight+1)
		return
	}
	en := &entry{key: key, value: value, weight: 1}
	heap.Push(lfu.queue, en)
	lfu.cache[key] = en

	lfu.usedBytes += en.Len()
	if lfu.maxBytes > 0 && lfu.usedBytes > lfu.maxBytes {
		lfu.removeElement(heap.Pop(lfu.queue))
	}
}
func (lfu *LFUCache) removeElement(e interface{}) {
	if e == nil {
		return
	}
	en := e.(*entry)
	delete(lfu.cache, en.key)
	lfu.usedBytes -= en.Len()

	if lfu.onEvicted != nil {
		lfu.onEvicted(en.key, en.value)
	}
}

func (lfu *LFUCache) Get(key string) interface{} {
	if e, ok := lfu.cache[key]; ok {
		lfu.queue.update(e, e.value, e.weight+1)
		return e.value
	}
	return nil
}
func (lfu *LFUCache) Del(key string) {
	if e, ok := lfu.cache[key]; ok {
		heap.Remove(lfu.queue, e.index)
		lfu.removeElement(e)
	}
}
func (lfu *LFUCache) DelOldest() {
	if lfu.queue.Len() == 0 {
		return
	}
	lfu.removeElement(heap.Pop(lfu.queue))
}
func (lfu *LFUCache) Len() int {
	return lfu.queue.Len()
}
