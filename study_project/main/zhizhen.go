package main

import "fmt"
func main(){
	var arr [3]string
	arr=[3]string{"1,2,3"}
	var arrP *[3]string //数组指针
	arrP=&arr
	fmt.Println(arr,arrP)
	//指针数组
	var arrpp [3]*string
	var str1 string="str1"
	var str2 string="str2"
	var str3 string="str3"
	arrpp=[3]*string{&str1,&str2,&str3}
	*arrpp[1]="555"
	fmt.Println(str2)
	fmt.Println("-------------------------------------")
	var a map[int]string
	a=make(map[int]string,10)
	a[2020]="smith"
	a[2021]="tom"
	a[2022]="wendy"
	var aP *map[int]string
	aP=&a
	fmt.Println(a,aP)
	
	type Student struct{
		Name string
		Age int
		Sex bool
	}
	var s1 Student=Student{"zyr",18,true}
	var sp *Student
	sp=&s1
	fmt.Println(s1,sp,*sp)
}
