package main

import (
	"fmt"
)

func My[T any](i T) T { //泛型
	return i
}

type Mytest[A any] struct { //泛型结构体
	Name A
}
type myType interface {
	getValue() string
}

func test_1[T myType](t T) {
	fmt.Println(t.getValue())
}

type my struct {
	Name string
}

func (m my) getValue() string {
	return m.Name
}
func main() {
	fmt.Println(My[string]("123") + "456")
	fmt.Println(My[int](123) + 456)
	m := Mytest[int]{
		Name: 123,
	}
	fmt.Println(m)
	a := my{Name: "swag"}
	test_1(a)
}
