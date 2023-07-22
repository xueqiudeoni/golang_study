package main

import (
	"fmt"
	"reflect"
)
func test(x interface{}){
	t:=reflect.TypeOf(x)
	fmt.Println(t)
	v:=reflect.ValueOf(x)
	fmt.Println(v)
	//种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）
	fmt.Println(t.Name())
	fmt.Println(t.Kind())
}
type Teacher struct{
	name string
	age int 
	school string
}
func main(){
	var te Teacher=Teacher{"zs",20,"scuniwersity"}
	var a *int
	a=new(int)
	*a=20
	test(te)
	test(a)
	fmt.Println("------------------")
	test("aaa")
}