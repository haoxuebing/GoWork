package main

import (
	"fmt"
)

type Tom struct {
	X, Y int
}

type IBase interface {
	Read()
	Write()
}

func (t *Tom) Read() {
	fmt.Println("This is Tom Read", t.X)
}

func (t *Tom) Write() {
	fmt.Println("This is Tom Write", t.Y)
}

type ITom interface {
	Read()
}
