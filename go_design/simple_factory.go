package main

import "fmt"

type Fruit interface {
	show()
}
type Apple struct {
}

func (apple *Apple) show() {
	fmt.Println("apple")
}

type Banana struct{}

func (banana *Banana) show() {
	fmt.Println("banana")
}

type Factory struct{}

func (factory *Factory) CreateFruit(kind string) (fruit Fruit) {
	if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "apple" {
		fruit = new(Apple)
	}
	return fruit
}
func main() {
	factory := &Factory{}
	apple := factory.CreateFruit("apple")
	apple.show()
	banana := factory.CreateFruit("banana")
	banana.show()
}
