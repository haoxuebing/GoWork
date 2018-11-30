package main

import (
	"fmt"
)

type IHuman interface {
	SayHi()
}

type Person struct {
	Name string
}

type Student struct {
	Name string
}

type Teacher struct {
	Name string
}
type WaiXinRen struct {
	Name string
}

func (p *Person) SayHi() {
	fmt.Println("This is ", p.Name)
}

func (s *Student) SayHi() {
	fmt.Println("This is ", s.Name)
}

func (t *Teacher) SayHi() {
	fmt.Println("This is ", t.Name)
}

func (t *WaiXinRen) SayHello() {
	fmt.Println("This is ", t.Name)
}

func WhoSayhi(i IHuman) {
	i.SayHi()
}

func main() {
	s := &Student{"学生"}
	t := &Teacher{"教师"}
	p := Person{"人类"}
	// w := WaiXinRen{"外星人"}
	WhoSayhi(s)
	WhoSayhi(t)
	WhoSayhi(&p)
	// WhoSayhi(&w)  //WaiXinRen 没有实现SayHi 故报错

}

func main01() {
	var i IHuman

	s := &Student{"学生"}
	i = s
	i.SayHi()

	t := &Teacher{"教师"}
	i = t
	i.SayHi()

	p := &Person{"人类"}
	i = p
	i.SayHi()

}
