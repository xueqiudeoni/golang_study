package main

import "fmt"

type phone struct {
	v5
}

func (p *phone) charge() {
	fmt.Println("charge to phone")
	p.v5.useV5()
}
func NewPhone(v52 v5) *phone {
	return &phone{v52}
}

type v5 interface {
	useV5()
}
type v220 struct {
}

func (v220) useV220() {
	fmt.Println("v220")
}

type adapter struct {
	v220 *v220
}

func (a *adapter) useV5() {
	a.v220.useV220()
}
func NewAdapter(v220 *v220) *adapter {
	return &adapter{v220: v220}
}
func main() {
	p := NewPhone(NewAdapter(new(v220)))
	p.charge()
}
