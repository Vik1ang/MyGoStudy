package main

import "github.com/kataras/iris/v12"

func main() {
	// 1. 创建app结构体对象
	app := iris.New()

	// 2. 端口监听
	err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		return
	}
}
