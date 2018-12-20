package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	getPage()
}

func getPage() {
	boss := "https://www.zhipin.com/job_detail/?query=node.js&scity=101010100&industry=&position="
	req, _ := http.NewRequest("GET", boss, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")
	cli := &http.Client{}
	res, _ := cli.Do(req)

	dom, _ := goquery.NewDocumentFromResponse(res)
	f, _ := os.OpenFile("node.txt", os.O_RDWR|os.O_CREATE, 0)

	defer f.Close()
	dom.Find(".job-list ul li").Each(func(i int, ctx *goquery.Selection) {

		title := ctx.Find(".job-title").Text()
		company := ctx.Find(".company-text .name a").Text()
		salery := ctx.Find(".info-primary .red").Text()
		experience := ctx.Find(".info-primary p").Text()
		dtl, _ := ctx.Find(".info-primary .name a").Attr("href")

		dtlReq, _ := http.NewRequest("GET", "https://www.zhipin.com/"+dtl, nil)
		dtlReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")

		dtlCli := &http.Client{}
		dtlRes, _ := dtlCli.Do(dtlReq)
		dtlDom, _ := goquery.NewDocumentFromResponse(dtlRes)

		jd := dtlDom.Find(".job-sec .text").Text()

		jd = strings.Replace(jd, "\n", "", -1)
		jd = strings.Replace(jd, "\t", "", -1)
		jd = strings.Replace(jd, " ", "", -1)

		fmt.Println(title + ":" + company + ":" + salery + ":" + experience + ":职位描述：" + jd)

		f.WriteString(title + ":" + company + ":" + salery + ":" + experience + "\n") //"\n职位描述：" + jd +

	})

	//todo
	//换IP 换userAgent  爬下一页

	// buf := make([]byte, 1024)
	// tmp := ""
	// for {
	// 	n, _ := res.Body.Read(buf)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	tmp += string(buf[:n])
	// }
	// fmt.Println(tmp)

	fmt.Println("END")

}
