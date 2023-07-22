package main
import "fmt"
func main(){
	num:=20
	fmt.Println(&num)
	test(&num)
	fmt.Println(num)
}
func test(num *int){
	*num=30
}

