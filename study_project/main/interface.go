package main

import "fmt"

// 接口嵌套
type Teacher interface {
	GetName() string
}
type Student interface {
	SetName(string)
}
type Human interface {
	Teacher
	Student
}
type Prof struct {
	Name string
}

func (a *Prof) SetName(name string) {
	a.Name = name
}
func (a Prof) GetName() string {
	return a.Name
}
func main() {
	var a = &Prof{
		"liming",
	}
	fmt.Println(a)
	var b Human = a
	b.SetName("sam")
	fmt.Println(b.GetName())

	//空接口
	//1、map实现：
	var userInfo = make(map[string]interface{})
	userInfo["name"] = "andy"
	userInfo["age"] = 10
	userInfo["hobby"] = []string{"sing", "dance"}
	//2、slice实现：
	var slice = make([]interface{}, 4, 4)
	slice[0] = "slice"
	slice[1] = true
	fmt.Println(slice)
	fmt.Println(userInfo["name"])
	//断言：判断接口中值的类型,由于空接口无法通过索引获取数组中的内容，因此常用于获取空接口创建的数组值
	//类型断言:interfacename.(type),返回2个参数：一个是interface转化为具体type的值，一个是bool类型，true表示断言成功，false表示断言失败
	hobbyValue, ok := userInfo["hobby"].([]string)
	if ok {
		fmt.Println(hobbyValue[0])
	}
}
