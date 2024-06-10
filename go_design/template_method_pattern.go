package main

import "fmt"

type beverage interface {
	boilWater()
	brew()
	pourInCup()
	addThings()
	wantAddThing() bool
}
type template struct {
	b beverage
}

func (t *template) makeBeverage() {
	if t.b == nil {
		return
	}
	t.b.boilWater()
	t.b.brew()
	t.b.pourInCup()
	if t.b.wantAddThing() == true {
		t.b.addThings()
	}
}

type makeCoffeeBeverage struct {
	template
}

func (m *makeCoffeeBeverage) boilWater() {
	fmt.Println("coffee boil water")
}
func (m *makeCoffeeBeverage) brew() {
	fmt.Println("coffee brew")
}
func (m *makeCoffeeBeverage) pourInCup() {
	fmt.Println("coffee pour in cup")
}
func (m *makeCoffeeBeverage) addThings() {
	fmt.Println("coffee add things")
}
func (m *makeCoffeeBeverage) wantAddThing() bool {
	return false
}
func newmakecoffee() *makeCoffeeBeverage {
	makecoffee := new(makeCoffeeBeverage)
	// b为beverage，是makeCoffeeBeverage的接口，需给接口赋值，指定具体的子类对象
	//触发b全部接口的多态特性
	makecoffee.b = makecoffee
	return makecoffee
}
func main() {
	coffee := newmakecoffee()
	coffee.makeBeverage()
}
