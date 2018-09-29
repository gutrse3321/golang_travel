package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main()  {
	// 设置访问的路由
	http.HandleFunc("/", sayhelloName)
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
