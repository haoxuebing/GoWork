package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("server start")

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/index", logIndex)
	http.HandleFunc("/login", loginFun)
	http.HandleFunc("/getParams", getParams)

	http.HandleFunc("/upload", uploadFile)

	http.ListenAndServe(":9090", nil)

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func logIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
}

func loginFun(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	user := r.Form["username"]
	pwd := r.Form["password"]
	fmt.Println("username:", user)
	fmt.Println("password:", pwd)

	fmt.Fprintf(w, user[0]+"\n"+pwd[0])

}

func getParams(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //可获取GET和POST请求的参数，所以最好不要用该变量来获取GET参数
	for k, v := range r.Form {
		fmt.Printf("key:%s,val:%s\n", k, strings.Join(v, ""))
	}

	//最好使用如下方式获取GET参数
	queryForm, _ := url.ParseQuery(r.URL.RawQuery)
	fmt.Println(queryForm["id"][0])

}

// 接收上传文件
func uploadFile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func postFile(filename string, targetURL string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetURL, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil
}
