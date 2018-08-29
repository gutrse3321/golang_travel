package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	// 地址: 127.0.0.1:8080/welcome
	app.Handle("GET", "/welcome", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	// 输出字符串
	// 127.0.0.1:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	// 输出json
	// 127.0.0.1:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris"})
	})
	// 监听8080端口
	// 文件配置，切换dev和pro
	app.Run(iris.Addr(":8080"), iris.WithConfiguration((iris.YAML("./config/config.yml"))))
}
