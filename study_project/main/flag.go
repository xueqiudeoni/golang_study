package main

import (
	"flag"
	"fmt"
	"time"
)

func main()  {
	//flag.Type():flag.Type(flag名，默认值，帮助信息)
	//name:=flag.String("name","zs","姓名")
	//flag.TypeVar():flag.TypeVar(Type指针，flag名，默认值，帮助信息)
	var name string
	flag.StringVar(&name,"name","zs","姓名")
	var age int
	flag.IntVar(&age,"age",18,"年龄")
	var delay time.Duration
	flag.DurationVar(&delay,"d",0,"延迟时间间隔")
	flag.Parse()//解析命令行参数
	fmt.Println(name,age,delay)
	fmt.Println(flag.Arg())//返回命令行参数后的其他参数
	fmt.Println(flag.NArg())//返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag())//返回使用命令行参数个数
}
