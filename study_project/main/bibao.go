package main
import "fmt"
//闭包：一个函数与其相应的引用环境组合的整体
//匿名函数+引用的参数/变量
func GetSum() func(int) int{
	var sum int =0
	return func (num int) int{
		sum=sum+num
		return sum
	}
}
func main(){
 	f:=GetSum()
 	fmt.Println(f(1))//1
	fmt.Println(f(2))//3
	fmt.Println(f(3))//6
}

