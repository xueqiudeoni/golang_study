package main
import "fmt"
//方法：作用在指定的数据类型上，和指定的数据类型绑定，so自定义类型都可以有方法，而不仅是struct
	//方法的声明：
	// type A struct{
	// 	Num int
	// }
	// func (a A) test(){   //(a A)与test()绑定,结构体对象传入test方法中是值传递，和函数参数传递一致
	// 	fmt.Println(a.Num)
	// }
	// 调用：
	// var a A
	// a.test()
	type Person struct{
		Name string
	}
	//值传递
	// func (p Person) test(){
	// 	p.Name="lulu"
	// 	fmt.Println(p.Name)
	// }
	//址传递
	func (p *Person) test(){
		(*p).Name="andy"
		fmt.Println(p.Name)
	}
func main(){
	var p Person
	p.Name="same"
	//值传递
	// p.test()
	// fmt.Println(p.Name)
	//址传递
	(&p).test()
	fmt.Println(p.Name)
}