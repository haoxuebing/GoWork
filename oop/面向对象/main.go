package main

import "fmt"

type Person interface {
	SayHello()
	SayHi()
}

func SayHello() {
	fmt.Println("Father say hello")
}

func SayHi() {
	fmt.Println("Father say hi")
}

type Teacher struct {
	Person
}

func main() {

}
