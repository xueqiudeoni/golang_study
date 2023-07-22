package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main()  {
	//logger，可以通过调用函数Print系列(Print|Printf|Println）、Fatal系列（Fatal|Fatalf|Fatalln）、和Panic系列（Panic|Panicf|Panicln）来使用，比自行创建一个logger对象更容易使用。
	//log输出到文件
	f,err:=os.OpenFile("./test.txt",os.O_APPEND|os.O_WRONLY,0644)
	if err!=nil {
		fmt.Println("open file failed:",err)
		return
	}
	log.SetOutput(f)
	for{
		log.Println("this is test log")
		time.Sleep(10*time.Second)
	}
	//log级别：debug、trace、info、warning、error、fatal
}
