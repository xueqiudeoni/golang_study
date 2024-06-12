package main

import (
	"fmt"
	"time"
)

//
//type observer interface {
//	onteachercoming()
//}
//type notifier interface {
//	addobserver(observer)
//	delobserver(observer)
//	notify()
//}
//type stu1 struct {
//	badthing string
//}
//
//func (s *stu1) dobadthing() {
//	fmt.Println("s1 is doing", s.badthing)
//}
//func (s *stu1) onteachercoming() {
//	fmt.Println("stu1 stop ", s.badthing)
//}
//
//type stu2 struct {
//	badthing string
//}
//
//func (s *stu2) dobadthing() {
//	fmt.Println("s2 is doing", s.badthing)
//}
//func (s *stu2) onteachercoming() {
//	fmt.Println("stu2 stop ", s.badthing)
//}
//
//type monitor struct {
//	obsList []observer
//}
//
//func (m *monitor) addobserver(ob observer) {
//	m.obsList = append(m.obsList, ob)
//}
//func (m *monitor) delobserver(ob observer) {
//	for k, o := range m.obsList {
//		if ob == o {
//			m.obsList = append(m.obsList[:k], m.obsList[k+1:]...)
//			break
//		}
//	}
//}
//func (m *monitor) notify() {
//	for _, o := range m.obsList {
//		o.onteachercoming()
//	}
//}
//func main() {
//	s1 := &stu1{badthing: " play game"}
//	s2 := &stu2{badthing: " watch tv"}
//	m := &monitor{}
//	s1.dobadthing()
//	s2.dobadthing()
//	m.addobserver(s1)
//	m.addobserver(s2)
//	m.notify()
//}

//		百晓生
//[丐帮]				[明教]
//洪七公				张无忌
//黄蓉					韦一笑
//乔峰					金毛狮王

const (
	PGaiBang  string = "丐帮"
	PMingJiao string = "明教"
)

type observer interface {
	onfriendBeFight(event *event)
	getName() string
	getParty() string
	title() string
}
type event struct {
	noti    notifier
	one     observer
	another observer
	msg     string
}
type notifier interface {
	Notify(event *event)
	AddObserver(observer observer)
	DelObserver(observer observer)
}
type hero struct {
	name  string
	party string
}

func (h *hero) onfriendBeFight(event *event) {
	if h.name == event.one.getName() || h.name == event.another.getName() {
		return
	}
	if h.party == event.one.getParty() {
		fmt.Println(h.title(), "拍手叫好!")
		return
	}
	if h.party == event.another.getParty() {
		fmt.Println(h.title(), "得知消息，发起报仇")
		h.fight(event.one, event.noti)
		return
	}
}
func (h *hero) title() string {
	return fmt.Sprintf("[" + h.party + h.name + "]")
}
func (h *hero) fight(another observer, baixiaosheng notifier) {
	msg := fmt.Sprintf("%s把%s打了", h.title(), another.title())
	event := &event{
		noti:    baixiaosheng,
		one:     h,
		another: another,
		msg:     msg,
	}
	baixiaosheng.Notify(event)
}
func (h *hero) getName() string {
	return h.name
}
func (h hero) getParty() string {
	return h.party
}

type baixiao struct {
	heroList []observer
}

func (b *baixiao) Notify(event *event) {
	fmt.Println("【世界消息】百晓生广播消息：", event.msg)
	for _, h := range b.heroList {
		//fmt.Println(string(h.getName()))
		h.onfriendBeFight(event)
		time.Sleep(2 * time.Second)
	}
}
func (b *baixiao) AddObserver(observer observer) {
	b.heroList = append(b.heroList, observer)
}
func (b *baixiao) DelObserver(observer2 observer) {
	for i, o := range b.heroList {
		if o == observer2 {
			b.heroList = append(b.heroList[:i], b.heroList[i+1:]...)
			break
		}
	}
}
func main() {
	hr := &hero{
		name:  "黄蓉",
		party: PGaiBang,
	}
	zwj := &hero{
		name:  "张无忌",
		party: PMingJiao,
	}
	hqg := &hero{
		name:  "洪七公",
		party: PGaiBang,
	}
	qf := &hero{
		name:  "乔峰",
		party: PGaiBang,
	}
	mj := &hero{
		name:  "灭绝师太",
		party: PMingJiao,
	}
	jmsw := &hero{
		name:  "金毛狮王",
		party: PMingJiao,
	}
	baixiaos := &baixiao{}
	baixiaos.AddObserver(hr)
	baixiaos.AddObserver(qf)
	baixiaos.AddObserver(hqg)
	baixiaos.AddObserver(mj)
	baixiaos.AddObserver(jmsw)
	baixiaos.AddObserver(zwj)
	fmt.Println("武林一片平静......")
	//for i, o := range baixiaos.heroList {
	//	fmt.Println(i, o)
	//}
	time.Sleep(time.Second * 5)
	hr.fight(zwj, baixiaos)
}
