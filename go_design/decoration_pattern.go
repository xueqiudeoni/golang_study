package main

import "fmt"

type Phone interface {
	PhoneShow()
}
type huawei struct {
}

func (h huawei) PhoneShow() {
	fmt.Println("huawei show")
}

type xiaomi struct{}

func (xm xiaomi) PhoneShow() {
	fmt.Println("xiaomi show")
}

type decrate struct {
	phone Phone
}

func (d decrate) PhoneShow() {
	fmt.Println("decrate show")
}

type MoDecrate struct {
	decrate
}

func (m MoDecrate) PhoneShow() {
	m.phone.PhoneShow()
	fmt.Println("mo show")
}
func NewMoDecrate(phone Phone) Phone {
	return MoDecrate{decrate{phone}}
}

type KeDecrate struct {
	decrate
}

func (k KeDecrate) PhoneShow() {
	k.phone.PhoneShow()
	fmt.Println("k show")
}
func NewKeDecrate(phone Phone) Phone {
	return KeDecrate{decrate{phone}}
}
func main() {
	var hwei Phone
	hwei = new(huawei)
	hwei.PhoneShow()
	fmt.Println("---------------------------")
	mod := NewMoDecrate(hwei)
	mod.PhoneShow()
	fmt.Println("---------------------------")
	ked := NewKeDecrate(hwei)
	ked.PhoneShow()
	fmt.Println("---------------------------")
	kemod := NewMoDecrate(ked)
	kemod.PhoneShow()
}
