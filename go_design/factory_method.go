package main

import "fmt"

type Fruite interface {
	Show()
}
type abstractFactory interface {
	createFruite() Fruite
}
type apple struct {
}

func (a *apple) Show() {
	fmt.Println("apple show")
}

type banana struct{}

func (b *banana) Show() {
	fmt.Println("banana show")
}

type appleFactory struct{}

func (a appleFactory) createFruite() Fruite {
	var apple = new(apple)
	return apple
}

type bananaFactory struct{}

func (b bananaFactory) createFruite() Fruite {
	banana := new(banana)
	return banana
}
func main() {
	var fruitfactory abstractFactory
	fruitfactory = appleFactory{}
	app := fruitfactory.createFruite()
	app.Show()
	fruitfactory = bananaFactory{}
	fruitfactory.createFruite().Show()
}
