package main

import "fmt"

type doctor struct{}

func (d *doctor) treatEye() {
	fmt.Println(" doctor treat eye")
}
func (d *doctor) treatNose() {
	fmt.Println(" doctor treat nose")
}

type cmd interface {
	treat()
}
type commandtreatEye struct {
	doctor doctor
}

func (c *commandtreatEye) treat() {
	c.doctor.treatEye()
}

type commandtreatNose struct {
	doctor doctor
}

func (c *commandtreatNose) treat() {
	c.doctor.treatNose()
}

type nurse struct {
	cmdList []cmd
}

func (n *nurse) notify() {
	if n.cmdList == nil {
		return
	}
	for _, cmd := range n.cmdList {
		cmd.treat()
	}
}
func main() {
	doc := doctor{}
	cmdeye := commandtreatEye{doctor: doc}
	cmdnose := commandtreatNose{doctor: doc}
	nurse := new(nurse)
	nurse.cmdList = append(nurse.cmdList, &cmdnose)
	nurse.cmdList = append(nurse.cmdList, &cmdeye)
	nurse.notify()
}

//type cooker struct {}
//func (c *cooker) cookChicken() {
//	fmt.Println("cookChicken")
//}
//func (c *cooker) cookYang()  {
//	fmt.Println("cook yang")
//}
//
//type cmdcook interface {
//	cook()
//}
//type cmdcookchicken struct {
//	cooker
//}
//func (c *cmdcookchicken)cook()  {
//	c.cooker.cookChicken()
//}
//
//type cmdcookyang struct {
//	cooker
//}
//func (c *cmdcookyang)cook()  {
//	c.cooker.cookYang()
//}
//
//type waiter struct {
//	cmdList []cmdcook
//}
//func (w *waiter) notify() {
//	if w.cmdList == nil {return}
//	for _, cmd := range w.cmdList {
//		cmd.cook()
//	}
//}
