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

	a := new([]int)
	fmt.Println(a)
	//输出&[]，ａ本身是一个地址
	b := make([]int, 1)
	fmt.Println(b)
	//输出[0]，ｂ本身是一个slice对象，其内容默认为0

	book1 := new(Books)
	book1.Title = "this is book1 title"
	book1.Content = "this is book1 content"
	book1.Author = "this is book1 author"
	book2 := Books{"this is book2 title", "this is book2 content", "this is book2 author"}
	fmt.Println("book1:", book1, ",Type:", reflect.TypeOf(book1)) //book1: &{this is book1 title this is book1 content this is book1 author} ,Type: *main.Books
	fmt.Println("book2:", book2, ",Type:", reflect.TypeOf(book2)) //book2: {this is book2 title this is book2 content this is book2 author} ,Type: main.Books
}
