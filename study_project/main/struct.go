package main

import "fmt"
//定义结构体
type Teacher struct{
	name string
	age int 
	school string
}
func main(){
	//创建结构体实例：法1：
	var ma Teacher
	ma.name="teacher_ma"
	ma.age=30
	ma.school="qh"
	fmt.Println(ma)
	fmt.Println(ma.age+10)
	//法2：
	var t Teacher=Teacher{"shanshan",30,"hljuniversity"}
	fmt.Println(t)
	//法3：
	var s *Teacher=new (Teacher)
	(*s).name="sam"
	(*s).age=20
	s.school="scuniversity"
	fmt.Println(s)
	//法4:
	var a *Teacher=&Teacher{"rady",44,"shkks"}
	fmt.Println(a)
	//结构体之间的转换
	type Professor struct{
		Age int
	}
	type Stu struct{
		Age int
	}
	var s1 Stu=Stu{10}
	var p1 Professor=Professor{30}
	p1=Professor(s1)
	fmt.Println(s1)
	fmt.Println(p1)

}