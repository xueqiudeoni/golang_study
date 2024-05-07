package main

import "fmt"
func main(){
	var intarr [6]int=[6]int{1,2,3,4,5,6}
	//切片构建在数组之上
	slice:=intarr[1:3]
	fmt.Println("intarr:",intarr)
	fmt.Println("slice:",slice)
	//切片元素个数
	fmt.Println("slice的个数:",len(slice))
	//容量可动态变化
	fmt.Println("slice 的容量:",cap(slice))
	//定义切片：make函数的三个参数：切片类型 切片长度 切片容量
	//make是底层创建一个数组，对外部可见，所以不可直接操作此数组，需通过slice区见解访问各个元素
	slice1:=make([]int,4,20)
	fmt.Println(slice1)
	fmt.Println("切片长度:",len(slice1))
	fmt.Println("切片容量:",cap(slice1))
	slice1[0]=99
	fmt.Println(slice1)
	//切片遍历：1、for
	for i:=0;i<len(slice);i++{
		fmt.Printf("slice[%v]=%v",i+1,slice[i])
	}
	//2、for range
	for key,value:=range slice{
		fmt.Printf("slice[%v]=%v",key+1,value)
	}
	fmt.Println()
	//切片可动态增长
	//slice的追加，底层数组对数组进行扩容：创建一个新数组，将老数组中的元素复制到新数组，新数组追加元素，底层数组指向新数组
	slice2:=append(slice,99,100)
	fmt.Println(slice2)
	fmt.Println(slice)
	//底层数组不能维护，需通过切片间接操作
	slice=append(slice,80,90)
	fmt.Println(slice)
	//通过append函数将切片追加给切片，函数中...必须写
	slice3:=[]int{22,33}
	slice=append(slice,slice3...)
	fmt.Println(slice)
	fmt.Println("-----------------------------------------------------------------")
	//切片的拷贝
	var a []int=[]int{1,2,3,4,5,6}
	var b []int=make([]int,10)
	copy(b,a)
	fmt.Println(b)
}