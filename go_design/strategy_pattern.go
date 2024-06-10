package main

import (
	"fmt"
	"math"
)

//type weaponStrategy interface {
//	useweapon()
//}
//type AK47 struct {
//}
//
//func (ak47 *AK47) useweapon() {
//	fmt.Println("AK47")
//}
//
//type Knife struct {
//}
//
//func (k Knife) useweapon() {
//	fmt.Println("knife")
//}
//
//type hero struct {
//	weaponStrategy
//}
//
//func (h *hero) fight() {
//	fmt.Println("hero useweapon")
//	h.weaponStrategy.useweapon()
//}
//func (h *hero) setweapon(strategy weaponStrategy) {
//	h.weaponStrategy = strategy
//}
//func main() {
//	hero := new(hero)
//	hero.setweapon(new(AK47))
//	hero.fight()
//	hero.weaponStrategy = new(Knife)
//	hero.fight()
//}

type sellstrategy interface {
	discount(price float64) float64
}
type eightstrategy struct {
}

func (es *eightstrategy) discount(price float64) float64 {
	fmt.Println("8折")
	return price * 0.8
}

type fanxianstrategy struct{}

func (es *fanxianstrategy) discount(price float64) float64 {
	fmt.Println("满200返100现")
	if price >= 200 {
		tmp := math.Floor(price / 200)
		return price - 100*tmp
	}
	return price
}

type human struct {
	sellstrategy
}
type product struct {
	price float64
	kind  string
}

func (h *human) choosestrategy(sellstrategy2 sellstrategy) {
	h.sellstrategy = sellstrategy2
}
func (h *human) paygoods(product2 *product) float64 {
	return h.sellstrategy.discount(product2.price)
}
func main() {
	man := new(human)
	man.choosestrategy(new(fanxianstrategy))
	pro := &product{
		price: 2000,
		kind:  "ipad",
	}
	pr := man.paygoods(pro)
	fmt.Println("买", pro.kind, "花了", pr)
}
