package main

import (
	"fmt"
	"sync"
)

//三要点：一个类只能有一个实例
//		必须自行创建这个实例
//		必须自行向整个系统提供这个实例

// 类非公有化，外界不能通过类创建对象
// 但要有一个指针可以指向这个唯一的变量，且指针永远不能改变方向
// 若全部为私有化，那么外部模块永远无法访问这个类和对象，故需对外提供一个方法来获取这个唯一实例对象
type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
func (s *singleton) Something() {
	fmt.Println("单例对象的某方法")
}
func main() {
	s := GetInstance()
	s.Something()
}
