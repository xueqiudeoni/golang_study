package main

import (
	"fmt"
	"sync"
	"time"
)

// 	sync.Cond是一个条件变量，协调多个goroutine之间的同步，当条件满足的时候，去通知那些因为条件不满足被阻塞的 goroutine 继续执行。
//	sync.Cond 的接口比较简单，只有 Wait, Signal 和 Broadcast 三个方法。
//		Wait 方法用来阻塞当前 goroutine，直到条件满足。调用 Wait 方法之前，需要先调用 L.Lock 方法加锁。
//		Signal 方法用来唤醒 notifyList 中的第一个 goroutine。
//		Broadcast 方法用来唤醒 notifyList 中的所有 goroutine。
//	sync.Cond 的实现也比较简单，它的核心就是 notifyList，它是一个链表，用来保存所有因为条件不满足而被阻塞的 goroutine。
//	用关闭 channel 的方式也可以实现类似的广播功能，但是有个问题是 channel 不能被重复关闭，所以这种方式无法被多次使用。也就是说使用这种方式无法多次广播。
//	使用 channel 发送通知的方式也是可以的，但是这样实现起来就复杂很多了，就更容易出错了。
//	sync.Cond 中使用 copyChecker 来检查 sync.Cond 是否被复制，如果被复制了，就会 panic。需要注意的是，这里的复制是指调用了 Wait，Signal 或 Broadcast 方法之后，sync.Cond 被复制了。在调用这几个方法之前进行复制是没有影响的。

var done bool
var data string

func read(cond *sync.Cond) {
	fmt.Println("read")
	cond.L.Lock()
	for !done {
		cond.Wait()
		fmt.Println("read wait...")
	}
	cond.L.Unlock()
	fmt.Println("read done")
	fmt.Println(done, data)
}
func write(cond *sync.Cond) {
	time.Sleep(time.Second)
	cond.L.Lock()
	time.Sleep(time.Second)
	data = "hello sync cond"
	done = true
	fmt.Println("write done")
	cond.L.Unlock()
	cond.Signal()
	//cond.Broadcast()
}
func main() {
	c := sync.NewCond(&sync.Mutex{})
	go read(c)
	go read(c)
	go write(c)
	time.Sleep(time.Second * 10)
}
