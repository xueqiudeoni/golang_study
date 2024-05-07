package main

import "fmt"
func main(){
	//new分配内存，new函数的实参是一个类型而不是具体数值，其返回值是对应类型的指针，如*int
	num:=new(int)
	fmt.Printf("num:  type:%T,值:%v,add:%p,指针指向的值:%v",num,num,&num,*num)
}
