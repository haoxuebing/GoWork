package main

import (
	"fmt"
	// "github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	// "strings"
)

func main() {
	myurl := "http://test.bestbing.cn/" //"http://www.xicidaili.com/wn/"
	request, _ := http.NewRequest("GET", myurl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Connection", "keep-alive")

	p, _ := url.Parse("http://221.210.120.153:54402")

	// p := &url.URL{
	// 	Host: "219.238.186.188:8118",
	// }

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(p),
		},
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("err1:", err)
		return
	}
	// defer response.Body.Close()

	buf := make([]byte, 1024)
	n, _ := response.Body.Read(buf)

	// if err2 != nil {
	// 	fmt.Println("err2:", err2)
	// 	return
	// }

	fmt.Println(string(buf[:n]))
}
