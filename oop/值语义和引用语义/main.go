package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) SetValue(n, s string, a int) {
	p.Age = a
	p.Name = n
	p.Sex = s

	fmt.Println("SetValue：", p)
	fmt.Printf("SetValue：&p=%p\n", &p)
}

func (p *Person) SetPoint(n, s string, a int) {
	p.Age = a
	p.Name = n
	p.Sex = s

	fmt.Println("SetPoint：", p)
	fmt.Printf("SetPoint：&p=%p\n", p)
}

func main() {
	var p1 Person
	p1.SetValue("nn", "ss", 15)
	fmt.Printf("p1=%v, &p1=%p\n", p1, &p1)

	var p2 Person
	(&p2).SetPoint("nnn", "sss", 155)
	fmt.Printf("p2=%v,&p2=%p\n", p2, &p2)
}
