/*
 * @Author: haoxuebing
 * @Date: 2018-11-29 17:06:14
 * @Last Modified by:   haoxuebing
 * @Last Modified time: 2018-11-29 17:06:14
 * @Description: 对于结构体的方法，指针类型和非指针类型 Go语言会自动识别
 */
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func (p *Person) SetValue(n, s string, a int) {
	p.Name = n
	p.Age = a
	p.Sex = s
}

func main() {
	//结构体是指针变量的方法集
	p := &Person{"yoyo", 18, "女"}
	p.SetInfoPointer()
	p.SetInfoValue()      //内部做转化，先把指针p，转化成*p再调用
	(*p).SetInfoPointer() //把（*p）转化成p后再调用

	//普通结构体的方法集
	m := Person{"Mike", 19, "男"}
	(&m).SetInfoPointer()
	m.SetInfoPointer()
	m.SetInfoValue()

	var s Person
	s.SetValue("花花", "未知", 14)
	fmt.Println("s=", s)

}
