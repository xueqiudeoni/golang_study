package main

import (
	"fmt"
	"sync"
	"runtime"
)
//定义协程计数器
var wg sync.WaitGroup
//互斥锁（共享资源的访问控制）：sync.Mutex,Mutex有2个方法：Lock()、Unlock()
var mutex sync.Mutex
func test(){
	mutex.Lock()//加锁
	for i:=0;i<10;i++{
		fmt.Println("test hello golang",i)
	}
	wg.Done()
	mutex.Unlock()//解锁
}
func test2(){
	for i:=0;i<10;i++{
		fmt.Println("test2 hello golang",i)
	}
	wg.Done() 
}
func main(){
	wg.Add(1)
	go test()
	wg.Add(1)
	go test2()
	for i:=0;i<10;i++{
		fmt.Println("main hello golang",i)
	}
	//等待所有协程执行完毕
	wg.Wait()
	fmt.Println("主线程结束")
	//获取CPU个数
	numCpu:=runtime.NumCPU()
	fmt.Println(numCpu)
	//runtime.GOMAXPROCS设置允许的CPU个数,其返回值为先前的设置，若<1则不会更改当前设置
	runtime.GOMAXPROCS(10)
	nowCpu:=runtime.GOMAXPROCS(10)
	fmt.Println(nowCpu)
}
