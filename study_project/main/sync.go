package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var lock sync.Mutex
	lock.Lock()
	lock.Unlock() //防止多个goroutine同一时刻操作统一资源
	var rwlock sync.RWMutex
	rwlock.RLock() //获取读锁，后续goroutine能读不能写
	rwlock.RUnlock()
	rwlock.Lock() //获取写锁，后续goroutine不管读、写都得等待获取锁
	rwlock.Unlock()
	//var once sync.Once//某些函数只需执行一次
	//once.Do()
	var syncMap sync.Map //不需初始化的并发安全map
	//原生map
	syncMap.Store(11, "def") //存储
	//syncMap.LoadOrStore()加载value值或存储
	syncMap.Delete(11)
	//syncMap.Range(fun()bool)顺序调用map中的值
	//原子操作，go内置针对内置数据类型的并发安全操作
	var i int32 = 10
	atomic.AddInt32(&i, 1)

}
