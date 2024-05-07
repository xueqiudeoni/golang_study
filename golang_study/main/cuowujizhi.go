package main
import "fmt"
func main(){
	test()
	fmt.Println("除法操作执行成功")
	fmt.Println("exec下面的操作")

}
func test(){
	//defer+匿名函数的调用
	defer func(){
		err:=recover()//recover内置捕获错误函数，返回值nil时表示没有捕获到错误
		if err!=nil{
			fmt.Println("错误已捕获")
			fmt.Println("err:",err)
		}
	}()
	num1:=10
	num2:=0
	result:=num1/num2
	fmt.Println(result)
}