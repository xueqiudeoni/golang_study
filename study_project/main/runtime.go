package main

import (
	"fmt"
	"runtime"
)

func main() {
	//a, _ := os.Stat("./trace.out")
	//fmt.Println(a)

	go show("java")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched() // 暂停当前，运行goroutine
		fmt.Println("golang")
	}
}
func show(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}
