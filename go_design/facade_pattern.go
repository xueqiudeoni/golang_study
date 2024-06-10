package main

import "fmt"

type subA struct {
}

func (a *subA) methodA() {
	fmt.Println("methodA")
}

type subB struct{}

func (b *subB) methodB() {
	fmt.Println("methodB")
}

type subC struct{}

func (c *subC) methodC() {
	fmt.Println("methodC")
}

type facade struct {
	a *subA
	b *subB
	c *subC
}

func (f *facade) fmethod1() {
	f.a.methodA()
	f.b.methodB()
}
func (f *facade) fmethod2() {
	f.a.methodA()
	f.c.methodC()
}
func main() {
	sa := new(subA)
	sa.methodA()
	sb := new(subB)
	sc := new(subC)
	f := &facade{
		a: sa,
		b: sb,
		c: sc,
	}
	f.fmethod1()
	f.fmethod2()
}
