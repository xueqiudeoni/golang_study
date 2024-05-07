package main

import "fmt"
func main(){
	//map基本语法：var map变量名 map[keytype]valuetype
	//key通常为int、string类型，value通常为数字（整数、浮点数）、string、map、结构体，不可为slice、map、function
	var a map[int]string
	//map需通过make进行初始化，才会分配内存
	a=make(map[int]string,10)//map可存放10个键值对,make函数的第二个参数可省略不写，默认创建1个空间
	//将键值对存入map
	a[2020200]="张三"
	a[2020201]="smith"
	//增加
	a[2020202]="liming"
	//修改
	a[2020200]="ruby"
	fmt.Println(a)
	b:=make(map[int]string)
	b[2020300]="tom"
	b[2020301]="sam"
	c:=map[int]string{
		2020100:"zhangsan",
		2020101:"lisi",
	}
	fmt.Println(b)
	fmt.Println(c)
	//删除
	delete(c,2020100)
	fmt.Println(c)
	//清空需将所有key删除，或map=make(...),make一个新的，让原来的失效被回收
	//查找
	value,flag :=b[2020308]
	fmt.Println(value)
	fmt.Println(flag)
	//获取长度
	fmt.Println(len(a))
	//遍历
	for k,v:=range a{
		fmt.Printf("key:%v,value=%v\t",k,v)
	}
	fmt.Println()
	d:=make(map[string]map[int]string)
	d["class1"]=make(map[int]string,3)
	d["class1"][200]="alice"
	d["class1"][201]="bob"
	d["class1"][202]="sam"
	d["class2"]=make(map[int]string,2)
	d["class2"][100]="susan"
	d["class2"][101]="ady"
	d["class3"]=make(map[int]string,2)
	d["class3"][100]="karry"
	d["class3"][101]="chew"
	for k,v:=range d{
		fmt.Println(k)
		for k1,v1:=range v{
			fmt.Printf("num:%d,name=%v\n",k1,v1)
		}
	}
}