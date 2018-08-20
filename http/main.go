package main

import (
	"fmt"
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("")))

	http.HandleFunc("/index", sayHelloWorld)
	http.ListenAndServe(":8899", nil)
}
