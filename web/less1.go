package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
	// 设置访问的路由
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
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
		// Unix将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）。
		crutime := time.Now().Unix()

		// 返回一个新的使用MD5校验的hash.Hash接口。
		h := md5.New()
		// WriteString函数将字符串s的内容写入w中。如果h已经实现了WriteString方法，函数会直接调用该方法。
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("web/login.gtpl")
		t.Execute(w, token)
	} else {
		// 请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// 验证token合法性
			fmt.Println("token: ", token)
		} else {
			// 不存在token，报错
			fmt.Println("token: null")
		}

		//fmt.Println("username: ", r.Form["username"])
		//fmt.Println("password: ", r.Form["password"])
		//转义，预防跨站脚本
		//func HTMLEscape(w io.Writer, b []byte) 把 b 进行转义之后写到 w
		//func HTMLEscapeString(s string) string 转义 s 之后返回结果字符串
		//func HTMLEscaper(args ...interface{}) string 支持多个参数一起转义，返回结果字符串
		//输出到服务端
		fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))

		// 验证是否是伪造的下拉菜单值
		slice := []string{"apple", "pear", "banana"}
		v := r.Form.Get("fruit")
		for index, item := range slice {
			if item == v {
				fmt.Println("fruit ", index, ": ", v)
			} else {
				fmt.Println("fruit ", index, ": null")
			}
		}
		// 输出到页面的
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("web/upload.gtpl")
		t.Execute(w, token)
	} else {
		// 上传的文件存储在 maxMemory 大小的内存里面，如果文件大小超过了 maxMemory，那么剩下的部分将存储在系统的临时文件中
		r.ParseMultipartForm(32 << 20)
		// 获取上传的文件句柄
		file, handler, err := r.FormFile("upFile")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)

		f, err := os.OpenFile("web/test" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		// 存储文件
		io.Copy(f, file)
	}
}