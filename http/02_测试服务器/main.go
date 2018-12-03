package main

import (
	"fmt"
	"net/http"
)

func myHandle(w http.ResponseWriter, r *http.Request) {

	fmt.Println("r.Method=", r.Method)
	fmt.Println("r.URL=", r.URL)

	fmt.Fprint(w, "hello world")
}

func main() {

	http.HandleFunc("/go", myHandle)
	http.ListenAndServe("127.0.0.1:8000", nil)

}
