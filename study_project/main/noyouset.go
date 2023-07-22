package main
import (
	"fmt"
	"errors"
)
func main(){
	err:=test()
	if err!=nil{
		fmt.Println("自定义错误：",err)
		panic(err)//使用panic程序将不再往下运行
	}
	fmt.Println("除法操作执行成功")
	fmt.Println("exec下面的操作")
}
func test() (err error){
	//自定义错误：需调用errors包下的New函数，函数返回error类型
	num1:=10
	num2:=0
	if num2==0{
		return errors.New("除数不能为0")
	}else{
		result:=num1/num2
		fmt.Println(result)
		return nil
	}
	
}