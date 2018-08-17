package main

import (
	"fmt"
	"reflect"
)

type Books struct {
	Title,
	Content,
	Author string
}

func main() {

	book1 := new(Books)
	book1.Title = "this is book1 title"
	book1.Content = "this is book1 content"
	book1.Author = "this is book1 author"

	book2 := Books{"this is book2 title", "this is book2 content", "this is book2 author"}

	fmt.Println("book1:", book1, ",Type:", reflect.TypeOf(book1))
	fmt.Println("book2:", book2, ",Type:", reflect.TypeOf(book2))
}
