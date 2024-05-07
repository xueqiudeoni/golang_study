package main
import "fmt"

func main()  {
	test(1,2,3,3,44,78)
}
func test(args...int){//agrs...int 可传入多个int类型的数据
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}