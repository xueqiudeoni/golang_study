package main

import (
	"fmt"
	"sync"
)

// Map 类型针对两种常见用例进行了优化：
//
//	 1.当给定 key 的条目只写入一次但读取多次时，如在只会增长的缓存中。（读多写少）
//		2. 当多个 goroutine 读取、写入和覆盖不相交的键集的条目。（不同 goroutine 操作不同的 key）
//
// 在这两种情况下，与 Go map 与单独的 Mutex 或 RWMutex 配对相比，使用 sync.Map 可以显著减少锁竞争（很多时候只需要原子操作就可以）
func main() {
	var m sync.Map
	m.Store("a", "b")
	m.Store("c", "d")
	m.Load("a")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	m.Delete("c")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

}
