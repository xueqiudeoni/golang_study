package main

import "fmt"
func sayHello(){
	for i := 0; i < 5; i++ {
		fmt.Println("hello golang")
	}
}
func errorTest(){
	defer func(){
		err:=recover()
		if err!=nil{
			fmt.Println("errTest发生错误")
		}
	}()
	num1:=10
	num2:=0
	result:=num1/num2
	fmt.Println(result)
}
func main(){
	//解决协程中的panic
	go sayHello()
	go errorTest()
}