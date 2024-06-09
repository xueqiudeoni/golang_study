package main

import "fmt"

type Goods struct {
	Kind string
	Fact bool
}
type Shopping interface {
	Buy(goods *Goods)
}
type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("korea shopping")
}

type AmericShopping struct {
}

func (as *AmericShopping) Buy(goods *Goods) {
	fmt.Println("Americ shopping")
}

type Proxy struct {
	shopping Shopping
}

func (p *Proxy) distingush(goods *Goods) bool {
	fmt.Println("proxy distingushing ", goods.Kind)
	if goods.Fact == false {
		fmt.Println(goods.Kind, "is a fake good")
	}
	return goods.Fact
}
func (p *Proxy) check(goods *Goods) {
	fmt.Println("proxy checked ", goods.Kind)
}
func (p Proxy) Buy(goods *Goods) {
	if p.distingush(goods) == true {
		p.shopping.Buy(goods)
		p.check(goods)
	}
}
func NewProxy(shopping Shopping) *Proxy {
	return &Proxy{shopping}
}
func main() {
	g1 := &Goods{
		Kind: "korea mask",
		Fact: true,
	}
	g2 := &Goods{
		Kind: "degree",
		Fact: false,
	}
	var shopping Shopping
	shopping = new(KoreaShopping)

	proxy := NewProxy(shopping)
	proxy.Buy(g1)
	proxy.Buy(g2)
}
