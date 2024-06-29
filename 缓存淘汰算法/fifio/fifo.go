package fifo

import (
	"cache"
	"container/list"
)

type fifo struct {
	maxByte   int
	onEvicted func(key string, value interface{})
	usedBytes int
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
func New(maxBytes int, onEvicted func(key string, value interface{})) cache.Cache {
	return &fifo{
		maxByte:   maxBytes,
		onEvicted: onEvicted,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
	}
}
func (fifo *fifo) Set(key string, value interface{}) {
	if e, ok := fifo.cache[key]; ok {
		fifo.ll.MoveToBack(e)
		en := e.Value.(*entry)
		fifo.usedBytes = fifo.usedBytes - cache.CalcLen(en.value) + cache.CalcLen(value)
		en.value = value
		return
	}
	en := &entry{key, value}
	e := fifo.ll.PushBack(en)
	fifo.cache[key] = e
	fifo.usedBytes += en.Len()
	if fifo.maxByte > 0 && fifo.usedBytes > fifo.maxByte {
		fifo.DelOldest()
	}
}
func (fifo *fifo) Get(key string) interface{} {
	if e, ok := fifo.cache[key]; ok {
		return e.Value.(*entry).value
	}
	return nil
}
func (fifo *fifo) Del(key string) {
	if e, ok := fifo.cache[key]; ok {
		fifo.removeElement(e)
	}
}
func (fifo *fifo) DelOldest() {
	fifo.removeElement(fifo.ll.Front())
}
func (fifo *fifo) removeElement(e *list.Element) {
	if e == nil {
		return
	}
	fifo.ll.Remove(e)
	en := e.Value.(*entry)
	fifo.usedBytes -= en.Len()
	delete(fifo.cache, en.key)
	if fifo.onEvicted != nil {
		fifo.onEvicted(en.key, en.value)
	}
}
func (f *fifo) Len() int {
	return f.ll.Len()
}
