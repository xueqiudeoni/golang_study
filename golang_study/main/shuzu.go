package main

import "fmt"
func main(){
	//键盘输入：1、ScanIn
	var age int
	fmt.Scanln(&age)
	fmt.Println("年龄：",age)
	//2、Scanf
	var isVIP bool
	var score int
	fmt.Println("请输入成绩,是否VIP,使用空格进行分割")
	fmt.Scanf("%d %t",&score,&isVIP)
	fmt.Printf("成绩：%v,是否VIP:%t",score,isVIP)
	fmt.Println("-----------------------------------------------")
	var scores [6]int //数组定义：var 数组名 [数组长度]数组类型
	fmt.Println(len(scores))
	for i:= 0;i<len(scores);i++{
		fmt.Printf("请输入第%d个学生的成绩",i+1)
		fmt.Scanln(&scores[i])
	}
	//数组遍历：1、for循环
	for i:=0;i<len(scores);i++{
		fmt.Printf("第%d个学生成绩为:%v\n",i+1,scores[i])
	}
	//2、for-range循环
	for key,value:=range scores{
		fmt.Printf("第%d个学生成绩为:%v\n",key+1,value)
	}
	fmt.Println("-----------------------------------------------")
	//数组初始化,4种：
	var arr1 [3]int=[3]int{1,3,5}
	var arr2=[3]int{34,5,6}
	var arr3=[...]int{3,5,6}
	var arr4=[...]int{0:6,2:5,1:9}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	test(&arr1)
	fmt.Println(arr1)
	二维数组的定义
	var arr [2][3]int16
	//数组初始化
	var arr [2][3]int16=[2][3]int16{{1,4,7},{2,3,4}}
	fmt.Println(arr)
	//二维数组遍历：1、for
	for i:=0;i<len(arr);i++{
		 for j:=0;j<len(arr[i]);j++{
			fmt.Print(arr[i][j],"\t")
		 }
		 fmt.Println()
	}
	//2、for range
	for key,value:=range arr{
		for k,v:=range value{
			fmt.Printf("arr[%v][%v]=%v\t",key,k,v)
		}
		fmt.Println()
	}
}
func test(arr *[3]int) *[3]int{
	(*arr)[0]=0
	return arr
}