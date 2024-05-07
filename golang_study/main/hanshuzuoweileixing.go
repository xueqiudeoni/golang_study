package main
import "fmt"
func main(){
	a:=test
	fmt.Printf("a的数据类型是:%T,test的数据类型是:%T",a,test)
	test1(10,5,test)
}
func test(num *int){
	fmt.Println(num)
}
func test1(num int,age int ,testFun func(*int)){
	fmt.Println(num,age,testFun)
}