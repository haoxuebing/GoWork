package main

import (
	"fmt"
)

//可以为不同的结构体添加同名的方法

// type po *int
//结构体为指针类型，则不能为结构体添加方法
// func (tmp po) test() {

// }

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (tmp Person) PrintInfo() {
	fmt.Println("tmp=", tmp)
}

func (tmp *Person) SetPserson(n, s string, a int) {
	tmp.Age = a
	tmp.Name = n
	tmp.Sex = s
}

func main() {
	p := Person{"mike", 12, "男"}
	p.PrintInfo()

	var p2 Person
	(&p2).SetPserson("yoyo", "女", 18)
	p2.PrintInfo()
}
