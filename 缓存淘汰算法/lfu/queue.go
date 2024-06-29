package lfu

import (
	"cache"
	"container/heap"
	"sort"
)

type entry struct {
	key    string
	value  interface{}
	weight int //entry在queue中的权重，访问次数越多，权重越大
	index  int //entry在heap中的索引
}

func (e *entry) Len() int {
	return cache.CalcLen(e.value)
}

type queue []*entry

type Interface interface {
	sort.Interface
	Push(v interface{})
	Pop() interface{}
}

func (q queue) Len() int {
	return len(q)
}
func (q queue) Less(i, j int) bool {
	return q[i].weight < q[j].weight
}
func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index, q[j].index = i, j
}
func (q *queue) Push(v interface{}) {
	n := len(*q)
	en := v.(*entry)
	en.index = n
	*q = append(*q, en)
}
func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	v := old[n-1]
	old[n-1] = nil //防止内存泄漏
	v.index = -1   //安全
	*q = old[0 : n-1]
	return v
}
func (q *queue) update(en *entry, value interface{}, weight int) {
	en.value = value
	en.weight = weight
	heap.Fix(q, en.index)
}
