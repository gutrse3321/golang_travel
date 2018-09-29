package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main()  {
	// 设置访问的路由
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	// 设置监听的端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Listening serve: ", err)
	}
}
func sayhelloName(writer http.ResponseWriter, request *http.Request) {
	// 解析参数，默认为不解析
	request.ParseForm()

	// 这些信息是输出到服务端的打印信息
	// http://localhost:9090/?url_long=123&fuck=2223
	// 输出一个映射，shell: map[fuck:[2223] url_long:[123]]
	fmt.Println(request.Form)
	fmt.Println("path: ", request.URL.Path)
	fmt.Println("scheme: ", request.URL.Scheme)
	fmt.Println("request.Form[\"url_long\"]:", request.Form["url_long"])
	fmt.Print("\n")

	for k, v := range  request.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	// 输出到页面的
	fmt.Fprintf(writer, "hello Tomo")
	fmt.Print("\n")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		// GET请求页面
		t, _ := template.ParseFiles("web/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		//fmt.Println("username: ", r.Form["username"])
		//fmt.Println("password: ", r.Form["password"])
		//转义
		//输出到服务端
		fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
		// 输出到页面的
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}
