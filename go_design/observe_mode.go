package main

import "fmt"

type observer interface {
	onteachercoming()
}
type notifier interface {
	addobserver(observer)
	delobserver(observer)
	notify()
}
type stu1 struct {
	badthing string
}

func (s *stu1) dobadthing() {
	fmt.Println("s1 is doing", s.badthing)
}
func (s *stu1) onteachercoming() {
	fmt.Println("stu1 stop ", s.badthing)
}

type stu2 struct {
	badthing string
}

func (s *stu2) dobadthing() {
	fmt.Println("s2 is doing", s.badthing)
}
func (s *stu2) onteachercoming() {
	fmt.Println("stu2 stop ", s.badthing)
}

type monitor struct {
	obsList []observer
}

func (m *monitor) addobserver(ob observer) {
	m.obsList = append(m.obsList, ob)
}
func (m *monitor) delobserver(ob observer) {
	for k, o := range m.obsList {
		if ob == o {
			m.obsList = append(m.obsList[:k], m.obsList[k+1:]...)
			break
		}
	}
}
func (m *monitor) notify() {
	for _, o := range m.obsList {
		o.onteachercoming()
	}
}
func main() {
	s1 := &stu1{badthing: " play game"}
	s2 := &stu2{badthing: " watch tv"}
	m := &monitor{}
	s1.dobadthing()
	s2.dobadthing()
	m.addobserver(s1)
	m.addobserver(s2)
	m.notify()
}
