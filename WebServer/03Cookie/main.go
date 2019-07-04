package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	fmt.Println("http://localhost:8080")
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/getCookie", getCookie)

	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	//设置cookie
	cookie := http.Cookie{
		Name:  "hello",
		Value: "qqwweerr",
	}
	http.SetCookie(w, &cookie)

	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		fmt.Fprintln(w, cookie.Name+":"+cookie.Value)
	}
}

// func login(w http.ResponseWriter, r *http.Request) {
// 	sess := globalSessions.SessionStart(w, r)
// }
